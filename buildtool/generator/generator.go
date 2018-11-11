package generator

import (
	"bytes"
	"fmt"
	"go/format"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/fatih/structtag"
	"github.com/pkg/errors"
	"github.com/spf13/afero"
	"golang.org/x/tools/imports"
)

var paramTypes = map[string]string{
	"string":    "string",
	"integer":   "int64",
	"string[]":  "[]string",
	"integer[]": "[]int64",
	"boolean":   "bool",
}

func sortedRoutesMapKeys(mp map[string]Routes) []string {
	var keys []string
	for k := range mp {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return keys
}

var updateMethodMap = map[string]string{
	"url":     "UpdateURLPath",
	"body":    "UpdateBody",
	"query":   "UpdateURLQuery",
	"headers": "UpdateHeader",
	"preview": "UpdatePreview",
}

func dirFileMap(path string, fs afero.Fs) (map[string][]byte, error) {
	output := map[string][]byte{}
	fileInfos, err := afero.ReadDir(fs, path)
	if err != nil {
		return nil, err
	}
	for _, fileInfo := range fileInfos {
		fileInfo.Sys()
		if fileInfo.IsDir() {
			continue
		}
		fileBytes, err := afero.ReadFile(fs, filepath.Join(path, fileInfo.Name()))
		if err != nil {
			return nil, err
		}
		output[fileInfo.Name()] = fileBytes
	}
	return output, nil
}

func verify(routesPath, outputPath string) ([]string, error) {
	realFs := afero.NewOsFs()
	tmpFs := afero.NewMemMapFs()
	tempDir, err := afero.TempDir(tmpFs, "", "")
	if err != nil {
		return nil, err
	}
	Generate(routesPath, tempDir, tmpFs)

	wantFiles, err := dirFileMap(tempDir, tmpFs)
	if err != nil {
		return nil, err
	}
	gotFiles, err := dirFileMap(outputPath, realFs)
	if err != nil {
		return nil, err
	}
	diffFiles := map[string]interface{}{}
	for name, wantBytes := range wantFiles {
		gotBytes := gotFiles[name]
		if !bytes.Equal(gotBytes, wantBytes) {
			diffFiles[name] = nil
		}
	}
	for name := range gotFiles {
		if _, ok := wantFiles[name]; !ok {
			diffFiles[name] = nil
		}
	}

	var output []string

	for v := range diffFiles {
		if strings.HasSuffix(v, "_test.go") ||
			strings.HasSuffix(v, "artisanally_handcrafted_code.go") {
			continue
		}
		output = append(output, v)
	}
	return output, nil
}

func Generate(routesPath, outputPath string, fs afero.Fs) {
	if fs == nil {
		fs = afero.NewOsFs()
	}
	CLITmpl := StructTmplHelper{
		Name: "CLI",
	}
	routesMap, err := parseRoutesFile(routesPath)
	if err != nil {
		panic(err)
	}
	var svcTmpls []SvcTmpl
	for _, svcName := range sortedRoutesMapKeys(routesMap) {
		svc := routesMap[svcName]
		svcName := strings.Title(svcName)

		CLITmpl.Fields = append(CLITmpl.Fields, StructField{
			Name: svcName,
			Type: svcName + "Cmd",
			Tags: newTags(newTag("cmd", "")),
		})

		svcTmpl := SvcTmpl{
			SvcStruct: StructTmplHelper{
				Name: svcName + "Cmd",
			},
		}

		for _, route := range svc {
			if svcName == "Orgs" && route.IDName == "list" && route.Path == "/user/orgs" {
				route.IDName = "list-for-current-user"
			}
			if svcName == "Orgs" && route.IDName == "get-membership" && route.Path == "/orgs/:org/memberships/:username" {
				route.IDName = "get-membership-for-user"
			}

			structName := svcName + ToArgName(route.IDName) + "Cmd"
			tmplHelper := StructTmplHelper{
				Name:   structName,
				Fields: []StructField{{Type: "internal.BaseCmd"}},
			}

			runMethod := RunMethod{
				ReceiverName: structName,
				Method:       strings.ToUpper(route.Method),
				URLPath:      route.Path,
			}

			skipThisRoute := false
			for _, preview := range route.Previews {
				tags := newTags(newTag("name", preview.Name+"-preview"))
				if preview.Required {
					setTag(tags, newTag("required", ""))
				}
				setTag(tags, newTag("help", preview.Description))

				previewParamName := ToArgName(preview.Name)
				tmplHelper.Fields = append(tmplHelper.Fields,
					StructField{
						Name: previewParamName,
						Type: "bool",
						Tags: tags,
					},
				)
				runMethod.Params = append(runMethod.Params, RunMethodParam{
					Name:         preview.Name,
					UpdateMethod: updateMethodMap["preview"],
					ValueField:   previewParamName,
				})
			}
			for _, param := range route.Params {
				paramName := ToArgName(param.Name)
				paramType, ok := paramTypes[param.Type]
				if !ok {
					skipThisRoute = true
					break
				}
				tags := &structtag.Tags{}
				if param.Required {
					setTag(tags, &structtag.Tag{Key: "required"})
				}
				setTag(tags, &structtag.Tag{Key: "name", Name: param.Name})
				if param.Description != "" {
					setTag(tags, &structtag.Tag{Key: "help", Name: param.Description})
				}
				sf := StructField{
					Name: paramName,
					Type: paramType,
					Tags: tags,
				}
				tmplHelper.Fields = append(tmplHelper.Fields, sf)

				runMethod.Params = append(runMethod.Params, RunMethodParam{
					Name:         param.Name,
					UpdateMethod: updateMethodMap[param.Location],
					ValueField:   paramName,
				})
			}
			if skipThisRoute {
				continue
			}

			svcTmpl.CmdStructAndMethods = append(svcTmpl.CmdStructAndMethods, CmdStructAndMethod{
				CmdStruct: tmplHelper,
				RunMethod: runMethod,
			})
			helpText := route.Name
			if route.DocumentationURL != "" {
				helpText = fmt.Sprintf("%v - %v", route.Name, route.DocumentationURL)
			}
			svcTmpl.SvcStruct.Fields = append(svcTmpl.SvcStruct.Fields, StructField{
				Name: ToArgName(route.IDName),
				Type: structName,
				Tags: newTags(newTag("cmd", ""), newTag("help", helpText)),
			})

		}
		svcTmpls = append(svcTmpls, svcTmpl)
	}
	err = fs.MkdirAll(outputPath, 0755)
	if err != nil {
		panic(err)
	}
	files := map[string]FileTmpl{
		"cli.go": {
			PrimaryStructs: []StructTmplHelper{
				CLITmpl,
			},
		},
	}
	for _, svcTmpl := range svcTmpls {
		filename := strings.ToLower(svcTmpl.SvcStruct.Name) + ".go"
		files[filename] = FileTmpl{
			SvcTmpls: []SvcTmpl{svcTmpl},
		}
	}
	for filename, fileTmpl := range files {
		err = writeGoFile(filename, "main", fileTmpl, outputPath, fs)
		if err != nil {
			panic(err)
		}
	}
}

type RunMethodParam struct {
	Name         string
	ValueField   string
	UpdateMethod string
}

type RunMethod struct {
	ReceiverName string
	Method       string
	URLPath      string
	Params       []RunMethodParam
}

// StructField is one field in a StructTmplHelper
type StructField struct {
	Name string
	Type string
	Tags *structtag.Tags
}

type StructTmplHelper struct {
	Name   string
	Fields []StructField
}

type CmdStructAndMethod struct {
	CmdStruct StructTmplHelper
	RunMethod RunMethod
}

type SvcTmpl struct {
	SvcStruct           StructTmplHelper
	CmdStructAndMethods []CmdStructAndMethod
}

type FileTmpl struct {
	PrimaryStructs []StructTmplHelper
	SvcTmpls       []SvcTmpl
}

var tmpl = template.Must(template.New("").Parse(tmplt))

// language=GoTemplate
const tmplt = `
{{define "main"}}
// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import "github.com/octo-cli/octo-cli/internal"
{{range .PrimaryStructs}}
{{template "structtype" .}}
{{end}}
{{range .SvcTmpls}}{{template "svctmpl" .}}{{end}}
{{end}}

{{define "svctmpl"}}{{if .SvcStruct}}{{template "structtype" .SvcStruct}}{{end}}
{{range .CmdStructAndMethods}}
{{if .CmdStruct}}{{template "structtype" .CmdStruct}}{{end}}
{{if .RunMethod}}{{template "runmethod" .RunMethod}}{{end}}
{{end}}
{{end}}

{{define "structtype"}}type {{.Name}} struct { {{range .Fields}}
	{{.Name}} {{.Type}} {{if .Tags}}{{printf "%#q" .Tags}} {{end}}{{end}}
}{{end}}

{{define "runmethod"}}
func (c *{{.ReceiverName}}) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("{{.URLPath}}"){{range .Params}}{{template "runmethodparam" .}}{{end}}
	return c.DoRequest("{{.Method}}")
}{{end}}

{{define "runmethodparam"}}
	c.{{.UpdateMethod}}("{{.Name}}", c.{{.ValueField}}){{end}}
`

//newTag is a helper to create a new *structtag.Tag with fewer lines of code
func newTag(key, name string, options ...string) *structtag.Tag {
	return &structtag.Tag{
		Key:     key,
		Name:    name,
		Options: options,
	}
}

//newTags creates a new *structtag.Tags from a list of tags
//  it will panic if one of the tags has no key
func newTags(tag ...*structtag.Tag) *structtag.Tags {
	tags := &structtag.Tags{}
	for _, tag := range tag {
		err := tags.Set(tag)
		if err != nil {
			panic(err)
		}
	}
	return tags
}

func setTag(tags *structtag.Tags, tag ...*structtag.Tag) {
	for _, tg := range tag {
		err := tags.Set(tg)
		if err != nil {
			panic(err)
		}
	}
}

//writeGoFile executes the named template and does the equivalent of `go fmt` and `goimports` on the output
func writeGoFile(filename, templateName string, p interface{}, path string, fs afero.Fs) error {
	var buf bytes.Buffer
	err := tmpl.ExecuteTemplate(&buf, templateName, p)
	if err != nil {
		return errors.Wrap(err, "failed to execute template")
	}
	out, err := format.Source(buf.Bytes())
	if err != nil {
		return errors.Wrap(err, "failed running format.Source")
	}
	out, err = imports.Process("", out, nil)
	if err != nil {
		return errors.Wrap(err, "failed running imports.Process")
	}
	fl := filepath.Join(path, filename)
	return afero.WriteFile(fs, fl, out, 0644)
}

//ToArgName takes input like "foo-bar" and returns "FooBar"
func ToArgName(in string) string {
	out := in
	for _, separator := range []string{"_", "-"} {
		words := strings.Split(out, separator)
		for i, word := range words {
			words[i] = strings.Title(word)
		}
		out = strings.Join(words, "")
	}
	return out
}