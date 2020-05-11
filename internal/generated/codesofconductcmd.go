// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import (
	"github.com/octo-cli/octo-cli/internal"
)

type CodesOfConductCmd struct {
	GetAllCodesOfConduct CodesOfConductGetAllCodesOfConductCmd `cmd:""`
	GetConductCode       CodesOfConductGetConductCodeCmd       `cmd:""`
	GetForRepo           CodesOfConductGetForRepoCmd           `cmd:""`
}

type CodesOfConductGetAllCodesOfConductCmd struct {
	ScarletWitch bool `required:"" name:"scarlet-witch-preview"`
	internal.BaseCmd
}

func (c *CodesOfConductGetAllCodesOfConductCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/codes_of_conduct")
	c.UpdatePreview("scarlet-witch", c.ScarletWitch)
	return c.DoRequest("GET")
}

type CodesOfConductGetConductCodeCmd struct {
	Key          string `required:"" name:"key"`
	ScarletWitch bool   `required:"" name:"scarlet-witch-preview"`
	internal.BaseCmd
}

func (c *CodesOfConductGetConductCodeCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/codes_of_conduct/{key}")
	c.UpdateURLPath("key", c.Key)
	c.UpdatePreview("scarlet-witch", c.ScarletWitch)
	return c.DoRequest("GET")
}

type CodesOfConductGetForRepoCmd struct {
	Owner        string `name:"owner"`
	Repo         string `required:"" name:"repo"`
	ScarletWitch bool   `required:"" name:"scarlet-witch-preview"`
	internal.BaseCmd
}

func (c *CodesOfConductGetForRepoCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{owner}/{repo}/community/code_of_conduct")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdatePreview("scarlet-witch", c.ScarletWitch)
	return c.DoRequest("GET")
}
