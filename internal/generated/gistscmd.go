// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import (
	"github.com/octo-cli/octo-cli/internal"
)

type GistsCmd struct {
	CheckIsStarred GistsCheckIsStarredCmd `cmd:""`
	Create         GistsCreateCmd         `cmd:""`
	CreateComment  GistsCreateCommentCmd  `cmd:""`
	Delete         GistsDeleteCmd         `cmd:""`
	DeleteComment  GistsDeleteCommentCmd  `cmd:""`
	Fork           GistsForkCmd           `cmd:""`
	Get            GistsGetCmd            `cmd:""`
	GetComment     GistsGetCommentCmd     `cmd:""`
	GetRevision    GistsGetRevisionCmd    `cmd:""`
	List           GistsListCmd           `cmd:""`
	ListComments   GistsListCommentsCmd   `cmd:""`
	ListCommits    GistsListCommitsCmd    `cmd:""`
	ListForUser    GistsListForUserCmd    `cmd:""`
	ListForks      GistsListForksCmd      `cmd:""`
	ListPublic     GistsListPublicCmd     `cmd:""`
	ListStarred    GistsListStarredCmd    `cmd:""`
	Star           GistsStarCmd           `cmd:""`
	Unstar         GistsUnstarCmd         `cmd:""`
	Update         GistsUpdateCmd         `cmd:""`
	UpdateComment  GistsUpdateCommentCmd  `cmd:""`
}

type GistsCheckIsStarredCmd struct {
	GistId string `required:"" name:"gist_id"`
	internal.BaseCmd
}

func (c *GistsCheckIsStarredCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists/{gist_id}/star")
	c.UpdateURLPath("gist_id", c.GistId)
	return c.DoRequest("GET")
}

type GistsCreateCmd struct {
	Files       internal.JSONObject `required:"" name:"files"`
	Description string              `name:"description"`
	Public      bool                `name:"public"`
	internal.BaseCmd
}

func (c *GistsCreateCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists")
	c.UpdateBody("description", c.Description)
	c.UpdateBody("files", c.Files)
	c.UpdateBody("public", c.Public)
	return c.DoRequest("POST")
}

type GistsCreateCommentCmd struct {
	GistId string `required:"" name:"gist_id"`
	Body   string `required:"" name:"body"`
	internal.BaseCmd
}

func (c *GistsCreateCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists/{gist_id}/comments")
	c.UpdateURLPath("gist_id", c.GistId)
	c.UpdateBody("body", c.Body)
	return c.DoRequest("POST")
}

type GistsDeleteCmd struct {
	GistId string `required:"" name:"gist_id"`
	internal.BaseCmd
}

func (c *GistsDeleteCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists/{gist_id}")
	c.UpdateURLPath("gist_id", c.GistId)
	return c.DoRequest("DELETE")
}

type GistsDeleteCommentCmd struct {
	GistId    string `required:"" name:"gist_id"`
	CommentId int64  `required:"" name:"comment_id"`
	internal.BaseCmd
}

func (c *GistsDeleteCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists/{gist_id}/comments/{comment_id}")
	c.UpdateURLPath("gist_id", c.GistId)
	c.UpdateURLPath("comment_id", c.CommentId)
	return c.DoRequest("DELETE")
}

type GistsForkCmd struct {
	GistId string `required:"" name:"gist_id"`
	internal.BaseCmd
}

func (c *GistsForkCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists/{gist_id}/forks")
	c.UpdateURLPath("gist_id", c.GistId)
	return c.DoRequest("POST")
}

type GistsGetCmd struct {
	GistId string `required:"" name:"gist_id"`
	internal.BaseCmd
}

func (c *GistsGetCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists/{gist_id}")
	c.UpdateURLPath("gist_id", c.GistId)
	return c.DoRequest("GET")
}

type GistsGetCommentCmd struct {
	GistId    string `required:"" name:"gist_id"`
	CommentId int64  `required:"" name:"comment_id"`
	internal.BaseCmd
}

func (c *GistsGetCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists/{gist_id}/comments/{comment_id}")
	c.UpdateURLPath("gist_id", c.GistId)
	c.UpdateURLPath("comment_id", c.CommentId)
	return c.DoRequest("GET")
}

type GistsGetRevisionCmd struct {
	GistId string `required:"" name:"gist_id"`
	Sha    string `required:"" name:"sha"`
	internal.BaseCmd
}

func (c *GistsGetRevisionCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists/{gist_id}/{sha}")
	c.UpdateURLPath("gist_id", c.GistId)
	c.UpdateURLPath("sha", c.Sha)
	return c.DoRequest("GET")
}

type GistsListCmd struct {
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	Since   string `name:"since"`
	internal.BaseCmd
}

func (c *GistsListCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists")
	c.UpdateURLQuery("since", c.Since)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type GistsListCommentsCmd struct {
	GistId  string `required:"" name:"gist_id"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *GistsListCommentsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists/{gist_id}/comments")
	c.UpdateURLPath("gist_id", c.GistId)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type GistsListCommitsCmd struct {
	GistId  string `required:"" name:"gist_id"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *GistsListCommitsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists/{gist_id}/commits")
	c.UpdateURLPath("gist_id", c.GistId)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type GistsListForUserCmd struct {
	Username string `required:"" name:"username"`
	Page     int64  `name:"page"`
	PerPage  int64  `name:"per_page"`
	Since    string `name:"since"`
	internal.BaseCmd
}

func (c *GistsListForUserCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/users/{username}/gists")
	c.UpdateURLPath("username", c.Username)
	c.UpdateURLQuery("since", c.Since)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type GistsListForksCmd struct {
	GistId  string `required:"" name:"gist_id"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *GistsListForksCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists/{gist_id}/forks")
	c.UpdateURLPath("gist_id", c.GistId)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type GistsListPublicCmd struct {
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	Since   string `name:"since"`
	internal.BaseCmd
}

func (c *GistsListPublicCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists/public")
	c.UpdateURLQuery("since", c.Since)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type GistsListStarredCmd struct {
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	Since   string `name:"since"`
	internal.BaseCmd
}

func (c *GistsListStarredCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists/starred")
	c.UpdateURLQuery("since", c.Since)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type GistsStarCmd struct {
	GistId string `required:"" name:"gist_id"`
	internal.BaseCmd
}

func (c *GistsStarCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists/{gist_id}/star")
	c.UpdateURLPath("gist_id", c.GistId)
	return c.DoRequest("PUT")
}

type GistsUnstarCmd struct {
	GistId string `required:"" name:"gist_id"`
	internal.BaseCmd
}

func (c *GistsUnstarCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists/{gist_id}/star")
	c.UpdateURLPath("gist_id", c.GistId)
	return c.DoRequest("DELETE")
}

type GistsUpdateCmd struct {
	GistId      string              `required:"" name:"gist_id"`
	Description string              `name:"description"`
	Files       internal.JSONObject `name:"files"`
	internal.BaseCmd
}

func (c *GistsUpdateCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists/{gist_id}")
	c.UpdateURLPath("gist_id", c.GistId)
	c.UpdateBody("description", c.Description)
	c.UpdateBody("files", c.Files)
	return c.DoRequest("PATCH")
}

type GistsUpdateCommentCmd struct {
	GistId    string `required:"" name:"gist_id"`
	CommentId int64  `required:"" name:"comment_id"`
	Body      string `required:"" name:"body"`
	internal.BaseCmd
}

func (c *GistsUpdateCommentCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/gists/{gist_id}/comments/{comment_id}")
	c.UpdateURLPath("gist_id", c.GistId)
	c.UpdateURLPath("comment_id", c.CommentId)
	c.UpdateBody("body", c.Body)
	return c.DoRequest("PATCH")
}
