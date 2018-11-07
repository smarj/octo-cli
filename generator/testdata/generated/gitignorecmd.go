// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import "github.com/octo-cli/octo-cli/internal"

type GitignoreCmd struct {
	ListTemplates GitignoreListTemplatesCmd `cmd:"" help:"Listing available templates - https://developer.github.com/v3/gitignore/#listing-available-templates"`
	GetTemplate   GitignoreGetTemplateCmd   `cmd:"" help:"Get a single template - https://developer.github.com/v3/gitignore/#get-a-single-template"`
}

type GitignoreListTemplatesCmd struct {
	internal.BaseCmd
}

func (c *GitignoreListTemplatesCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gitignore/templates")
	return c.DoRequest("GET")
}

type GitignoreGetTemplateCmd struct {
	internal.BaseCmd
	Name string `required:"" name:"name"`
}

func (c *GitignoreGetTemplateCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gitignore/templates/:name")
	c.UpdateURLPath("name", c.Name)
	return c.DoRequest("GET")
}
