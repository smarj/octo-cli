// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import (
	"github.com/octo-cli/octo-cli/internal"
)

type EmojisCmd struct {
	Get EmojisGetCmd `cmd:""`
}

type EmojisGetCmd struct {
	internal.BaseCmd
}

func (c *EmojisGetCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/emojis")
	return c.DoRequest("GET")
}
