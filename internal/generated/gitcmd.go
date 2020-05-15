// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import (
	"github.com/octo-cli/octo-cli/internal"
)

type GitCmd struct {
	CreateBlob       GitCreateBlobCmd       `cmd:""`
	CreateCommit     GitCreateCommitCmd     `cmd:""`
	CreateRef        GitCreateRefCmd        `cmd:""`
	CreateTag        GitCreateTagCmd        `cmd:""`
	CreateTree       GitCreateTreeCmd       `cmd:""`
	DeleteRef        GitDeleteRefCmd        `cmd:""`
	GetBlob          GitGetBlobCmd          `cmd:""`
	GetCommit        GitGetCommitCmd        `cmd:""`
	GetRef           GitGetRefCmd           `cmd:""`
	GetTag           GitGetTagCmd           `cmd:""`
	GetTree          GitGetTreeCmd          `cmd:""`
	ListMatchingRefs GitListMatchingRefsCmd `cmd:""`
	UpdateRef        GitUpdateRefCmd        `cmd:""`
}

type GitCreateBlobCmd struct {
	Repo     string `required:"" name:"repo"`
	Content  string `required:"" name:"content"`
	Encoding string `name:"encoding"`
	internal.BaseCmd
}

func (c *GitCreateBlobCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/git/blobs")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("content", c.Content)
	c.UpdateBody("encoding", c.Encoding)
	return c.DoRequest("POST")
}

type GitCreateCommitCmd struct {
	Repo           string   `required:"" name:"repo"`
	Message        string   `required:"" name:"message"`
	Parents        []string `required:"" name:"parents"`
	Tree           string   `required:"" name:"tree"`
	AuthorDate     string   `name:"author.date"`
	AuthorEmail    string   `name:"author.email"`
	AuthorName     string   `name:"author.name"`
	CommitterDate  string   `name:"committer.date"`
	CommitterEmail string   `name:"committer.email"`
	CommitterName  string   `name:"committer.name"`
	Signature      string   `name:"signature"`
	internal.BaseCmd
}

func (c *GitCreateCommitCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/git/commits")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("author.date", c.AuthorDate)
	c.UpdateBody("author.email", c.AuthorEmail)
	c.UpdateBody("author.name", c.AuthorName)
	c.UpdateBody("committer.date", c.CommitterDate)
	c.UpdateBody("committer.email", c.CommitterEmail)
	c.UpdateBody("committer.name", c.CommitterName)
	c.UpdateBody("message", c.Message)
	c.UpdateBody("parents", c.Parents)
	c.UpdateBody("signature", c.Signature)
	c.UpdateBody("tree", c.Tree)
	return c.DoRequest("POST")
}

type GitCreateRefCmd struct {
	Repo string `required:"" name:"repo"`
	Ref  string `required:"" name:"ref"`
	Sha  string `required:"" name:"sha"`
	internal.BaseCmd
}

func (c *GitCreateRefCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/git/refs")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("ref", c.Ref)
	c.UpdateBody("sha", c.Sha)
	return c.DoRequest("POST")
}

type GitCreateTagCmd struct {
	Repo        string `required:"" name:"repo"`
	Message     string `required:"" name:"message"`
	Object      string `required:"" name:"object"`
	Tag         string `required:"" name:"tag"`
	Type        string `required:"" name:"type"`
	TaggerDate  string `name:"tagger.date"`
	TaggerEmail string `name:"tagger.email"`
	TaggerName  string `name:"tagger.name"`
	internal.BaseCmd
}

func (c *GitCreateTagCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/git/tags")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("message", c.Message)
	c.UpdateBody("object", c.Object)
	c.UpdateBody("tag", c.Tag)
	c.UpdateBody("tagger.date", c.TaggerDate)
	c.UpdateBody("tagger.email", c.TaggerEmail)
	c.UpdateBody("tagger.name", c.TaggerName)
	c.UpdateBody("type", c.Type)
	return c.DoRequest("POST")
}

type GitCreateTreeCmd struct {
	Repo     string                `required:"" name:"repo"`
	Tree     []internal.JSONObject `required:"" name:"tree"`
	BaseTree string                `name:"base_tree"`
	internal.BaseCmd
}

func (c *GitCreateTreeCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/git/trees")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("base_tree", c.BaseTree)
	c.UpdateBody("tree", c.Tree)
	return c.DoRequest("POST")
}

type GitDeleteRefCmd struct {
	Repo string `required:"" name:"repo"`
	Ref  string `required:"" name:"ref"`
	internal.BaseCmd
}

func (c *GitDeleteRefCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/git/refs/{ref}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("ref", c.Ref)
	return c.DoRequest("DELETE")
}

type GitGetBlobCmd struct {
	Repo    string `required:"" name:"repo"`
	FileSha string `required:"" name:"file_sha"`
	internal.BaseCmd
}

func (c *GitGetBlobCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/git/blobs/{file_sha}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("file_sha", c.FileSha)
	return c.DoRequest("GET")
}

type GitGetCommitCmd struct {
	Repo      string `required:"" name:"repo"`
	CommitSha string `required:"" name:"commit_sha"`
	internal.BaseCmd
}

func (c *GitGetCommitCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/git/commits/{commit_sha}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("commit_sha", c.CommitSha)
	return c.DoRequest("GET")
}

type GitGetRefCmd struct {
	Repo string `required:"" name:"repo"`
	Ref  string `required:"" name:"ref"`
	internal.BaseCmd
}

func (c *GitGetRefCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/git/ref/{ref}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("ref", c.Ref)
	return c.DoRequest("GET")
}

type GitGetTagCmd struct {
	Repo   string `required:"" name:"repo"`
	TagSha string `required:"" name:"tag_sha"`
	internal.BaseCmd
}

func (c *GitGetTagCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/git/tags/{tag_sha}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("tag_sha", c.TagSha)
	return c.DoRequest("GET")
}

type GitGetTreeCmd struct {
	Repo      string `required:"" name:"repo"`
	TreeSha   string `required:"" name:"tree_sha"`
	Recursive int64  `name:"recursive"`
	internal.BaseCmd
}

func (c *GitGetTreeCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/git/trees/{tree_sha}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("tree_sha", c.TreeSha)
	c.UpdateURLQuery("recursive", c.Recursive)
	return c.DoRequest("GET")
}

type GitListMatchingRefsCmd struct {
	Repo    string `required:"" name:"repo"`
	Ref     string `required:"" name:"ref"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	internal.BaseCmd
}

func (c *GitListMatchingRefsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/git/matching-refs/{ref}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("ref", c.Ref)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type GitUpdateRefCmd struct {
	Repo  string `required:"" name:"repo"`
	Ref   string `required:"" name:"ref"`
	Sha   string `required:"" name:"sha"`
	Force bool   `name:"force"`
	internal.BaseCmd
}

func (c *GitUpdateRefCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/{repo}/git/refs/{ref}")
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("ref", c.Ref)
	c.UpdateBody("force", c.Force)
	c.UpdateBody("sha", c.Sha)
	return c.DoRequest("PATCH")
}
