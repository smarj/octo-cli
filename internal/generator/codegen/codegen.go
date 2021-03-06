package codegen

import (
	"path/filepath"
	"sort"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/octo-cli/octo-cli/internal/generator/util"
	"github.com/spf13/afero"
)

func newRunMethodAdder(name, updateMethod string) codeGroupAdder {
	return func(group *jen.Group) {
		group.Id(updateMethod).Params(jen.Lit(name), jen.Id("c").Dot(util.ToArgName(name)))
	}
}

type codeGroupAdder func(group *jen.Group)

type runMethod struct {
	receiverName    string
	method          string
	urlPath         string
	codeGroupAdders []codeGroupAdder
}

type paramLocation int

//nolint:deadcode,varcheck
const (
	locInvalid paramLocation = iota
	locPath
	locQuery
	locBody
	locHeader
	locPreview
)

type structField struct {
	name          string
	fieldType     string
	tags          map[string]string
	paramLocation paramLocation
	paramOrder    int
}

func (s structField) required() bool {
	if s.tags == nil {
		return false
	}
	_, ok := s.tags["required"]
	return ok
}

func (s structField) less(o structField) bool {
	if s.paramLocation < o.paramLocation {
		return true
	}
	if s.paramLocation > o.paramLocation {
		return false
	}
	if s.paramOrder < o.paramOrder {
		return true
	}
	if s.paramOrder > o.paramOrder {
		return false
	}
	if s.required() != o.required() {
		return o.required()
	}
	return s.name < o.name
}

type structTmplHelper struct {
	name   string
	fields []structField
}

func (s *structTmplHelper) stmt(group *jen.Group) {
	group.Line().Type().Id(s.name).StructFunc(func(group *jen.Group) {
		for _, field := range s.fields {

			if strings.HasPrefix(field.fieldType, "internal.") {
				fieldType := strings.TrimPrefix(field.fieldType, "internal.")
				group.Id(field.name).Qual("github.com/octo-cli/octo-cli/internal", fieldType).Tag(field.tags)
				continue
			}
			group.Id(field.name).Id(field.fieldType).Tag(field.tags)
		}
	})
}

type cmdStructAndMethod struct {
	cmdStruct *structTmplHelper
	runMethod *runMethod
}

type svcTmpl struct {
	svcStruct           *structTmplHelper
	cmdStructAndMethods []cmdStructAndMethod
}

func (s svcTmpl) stmt(stmt *jen.Group) {
	if s.svcStruct.name != "" {
		s.svcStruct.stmt(stmt)
	}
	for _, csm := range s.cmdStructAndMethods {
		if csm.cmdStruct != nil {
			csm.cmdStruct.stmt(stmt)
		}
		if csm.runMethod != nil {
			r := csm.runMethod
			stmt.Line().Func().Params(jen.Id("c").Op("*").
				Id(r.receiverName)).
				Id("Run(isValueSetMap map[string]bool) error").
				BlockFunc(func(group *jen.Group) {
					group.Id("c.SetIsValueSetMap(isValueSetMap)")
					group.Id("c.SetURLPath").Params(jen.Lit(r.urlPath))
					for _, cga := range r.codeGroupAdders {
						if cga == nil {
							continue
						}
						cga(group)
					}
					group.Return(jen.Id("c.DoRequest").Params(jen.Lit(r.method)))
				})
		}
	}
}

type fileTmpl struct {
	cmdHelps       map[string]map[string]string
	flagHelps      map[string]map[string]map[string]string
	primaryStructs []structTmplHelper
	svcTmpls       []svcTmpl
}

func stringMapDict(mp map[string]string) jen.Dict {
	return jen.DictFunc(func(dict jen.Dict) {
		for k, v := range mp {
			dict[jen.Lit(k)] = jen.Lit(v)
		}
	})
}

func doubleStringMap(mp map[string]map[string]string) jen.Dict {
	return jen.DictFunc(func(dict jen.Dict) {
		for k, v := range mp {
			dict[jen.Lit(k)] = jen.Values(stringMapDict(v))
		}
	})
}

func (f fileTmpl) addFlagHelps(file *jen.File) {
	if len(f.flagHelps) == 0 {
		return
	}
	file.Var().Id("FlagHelps").Op("=").
		Map(jen.String()).
		Map(jen.String()).
		Map(jen.String()).String().Values(jen.DictFunc(func(dict jen.Dict) {
		for k, v := range f.flagHelps {
			dict[jen.Lit(k)] = jen.Values(doubleStringMap(v))
		}
	}))
}

func (f fileTmpl) addCmdHelps(file *jen.File) {
	if len(f.cmdHelps) == 0 {
		return
	}
	file.Var().Id("CmdHelps").Op("=").
		Map(jen.String()).
		Map(jen.String()).String().Values(
		doubleStringMap(f.cmdHelps),
	)
}

func (f fileTmpl) jenFile() *jen.File {
	file := jen.NewFile("generated")
	file.HeaderComment("Code generated by octo-cli/generator; DO NOT EDIT.")
	for _, primaryStruct := range f.primaryStructs {
		primaryStruct.stmt(file.Group)
	}
	for _, svcTmpl := range f.svcTmpls {
		svcTmpl.stmt(file.Group)
	}
	f.addCmdHelps(file)
	f.addFlagHelps(file)
	return file
}

func writeFiles(files map[string]fileTmpl, outputPath string, fs afero.Fs) error {
	for name, fileTmpl := range files {
		err := writeGoFile(name, fileTmpl, outputPath, fs)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeGoFile(filename string, fileTmpl fileTmpl, path string, fs afero.Fs) error {
	out, err := generateGoFile(fileTmpl)
	if err != nil {
		return err
	}
	fl := filepath.Join(path, filename)
	return afero.WriteFile(fs, fl, out, 0644)
}

func generateGoFile(fileTmpl fileTmpl) ([]byte, error) {
	for _, svcTmpl := range fileTmpl.svcTmpls {
		tmplSorting(&svcTmpl)
	}
	return []byte(fileTmpl.jenFile().GoString()), nil
}

func sortCmdStructFields(fields []structField) {
	if len(fields) == 0 {
		return
	}
	newFields := make([]structField, 0, len(fields))
	holdouts := make([]structField, 0, len(fields))
	for _, field := range fields {
		if field.name == "" {
			holdouts = append(holdouts, field)
			continue
		}
		newFields = append(newFields, field)
	}
	sort.Slice(newFields, func(i, j int) bool {
		return newFields[i].less(newFields[j])
	})
	sort.Slice(holdouts, func(i, j int) bool {
		return holdouts[i].fieldType < holdouts[j].fieldType
	})
	newFields = append(newFields, holdouts...)
	copy(fields, newFields)
}

func tmplSorting(svcTmpl *svcTmpl) {
	sort.Slice(svcTmpl.svcStruct.fields, func(i, j int) bool {
		return svcTmpl.svcStruct.fields[i].name < svcTmpl.svcStruct.fields[j].name
	})
	for _, csm := range svcTmpl.cmdStructAndMethods {
		sortCmdStructFields(csm.cmdStruct.fields)
	}
	sort.Slice(svcTmpl.cmdStructAndMethods, func(i, j int) bool {
		return svcTmpl.cmdStructAndMethods[i].cmdStruct.name < svcTmpl.cmdStructAndMethods[j].cmdStruct.name
	})
}
