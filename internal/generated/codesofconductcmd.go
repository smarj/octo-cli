// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import "github.com/octo-cli/octo-cli/internal"

type CodesOfConductCmd struct {
	ListConductCodes CodesOfConductListConductCodesCmd `cmd:"" help:"List all codes of conduct - https://developer.github.com/v3/codes_of_conduct/#list-all-codes-of-conduct"`
	GetConductCode   CodesOfConductGetConductCodeCmd   `cmd:"" help:"Get an individual code of conduct - https://developer.github.com/v3/codes_of_conduct/#get-an-individual-code-of-conduct"`
	GetForRepo       CodesOfConductGetForRepoCmd       `cmd:"" help:"Get the contents of a repository's code of conduct - https://developer.github.com/v3/codes_of_conduct/#get-the-contents-of-a-repositorys-code-of-conduct"`
}

type CodesOfConductListConductCodesCmd struct {
	internal.BaseCmd
}

func (c *CodesOfConductListConductCodesCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/codes_of_conduct")
	return c.DoRequest("GET")
}

type CodesOfConductGetConductCodeCmd struct {
	internal.BaseCmd
	Key string `required:"" name:"key"`
}

func (c *CodesOfConductGetConductCodeCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/codes_of_conduct/:key")
	c.UpdateURLPath("key", c.Key)
	return c.DoRequest("GET")
}

type CodesOfConductGetForRepoCmd struct {
	internal.BaseCmd
	Owner string `required:"" name:"owner"`
	Repo  string `required:"" name:"repo"`
}

func (c *CodesOfConductGetForRepoCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/community/code_of_conduct")
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	return c.DoRequest("GET")
}
