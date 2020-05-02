// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import "github.com/octo-cli/octo-cli/internal"

type SearchCmd struct {
	Code                  SearchCodeCmd                  `cmd:""`
	Commits               SearchCommitsCmd               `cmd:""`
	EmailLegacy           SearchEmailLegacyCmd           `cmd:""`
	IssuesAndPullRequests SearchIssuesAndPullRequestsCmd `cmd:""`
	IssuesLegacy          SearchIssuesLegacyCmd          `cmd:""`
	Labels                SearchLabelsCmd                `cmd:""`
	Repos                 SearchReposCmd                 `cmd:""`
	ReposLegacy           SearchReposLegacyCmd           `cmd:""`
	Topics                SearchTopicsCmd                `cmd:""`
	Users                 SearchUsersCmd                 `cmd:""`
	UsersLegacy           SearchUsersLegacyCmd           `cmd:""`
}

type SearchCodeCmd struct {
	Order   string `name:"order"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	Q       string `required:"" name:"q"`
	Sort    string `name:"sort"`
	internal.BaseCmd
}

func (c *SearchCodeCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/search/code")
	c.UpdateURLQuery("order", c.Order)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("q", c.Q)
	c.UpdateURLQuery("sort", c.Sort)
	return c.DoRequest("GET")
}

type SearchCommitsCmd struct {
	Cloak   bool   "name:\"cloak-preview\" required:\"\" help:\"The Commit Search API is currently available for developers to preview. During the preview period, the APIs may change without advance notice. Please see the [blog post](https://developer.github.com/changes/2017-01-05-commit-search-api/) for full details.\n\nTo access the API you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.cloak-preview\n```\""
	Order   string `name:"order"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	Q       string `required:"" name:"q"`
	Sort    string `name:"sort"`
	internal.BaseCmd
}

func (c *SearchCommitsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/search/commits")
	c.UpdatePreview("cloak", c.Cloak)
	c.UpdateURLQuery("order", c.Order)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("q", c.Q)
	c.UpdateURLQuery("sort", c.Sort)
	return c.DoRequest("GET")
}

type SearchEmailLegacyCmd struct {
	Email string `required:"" name:"email"`
	internal.BaseCmd
}

func (c *SearchEmailLegacyCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/legacy/user/email/:email")
	c.UpdateURLPath("email", c.Email)
	return c.DoRequest("GET")
}

type SearchIssuesAndPullRequestsCmd struct {
	Order   string `name:"order"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	Q       string `required:"" name:"q"`
	Sort    string `name:"sort"`
	internal.BaseCmd
}

func (c *SearchIssuesAndPullRequestsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/search/issues")
	c.UpdateURLQuery("order", c.Order)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("q", c.Q)
	c.UpdateURLQuery("sort", c.Sort)
	return c.DoRequest("GET")
}

type SearchIssuesLegacyCmd struct {
	Keyword    string `required:"" name:"keyword"`
	Owner      string `required:"" name:"owner"`
	Repository string `required:"" name:"repository"`
	State      string `required:"" name:"state"`
	internal.BaseCmd
}

func (c *SearchIssuesLegacyCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/legacy/issues/search/:owner/:repository/:state/:keyword")
	c.UpdateURLPath("keyword", c.Keyword)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repository", c.Repository)
	c.UpdateURLPath("state", c.State)
	return c.DoRequest("GET")
}

type SearchLabelsCmd struct {
	Order        string `name:"order"`
	Q            string `required:"" name:"q"`
	RepositoryId int64  `required:"" name:"repository_id"`
	Sort         string `name:"sort"`
	internal.BaseCmd
}

func (c *SearchLabelsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/search/labels")
	c.UpdateURLQuery("order", c.Order)
	c.UpdateURLQuery("q", c.Q)
	c.UpdateURLQuery("repository_id", c.RepositoryId)
	c.UpdateURLQuery("sort", c.Sort)
	return c.DoRequest("GET")
}

type SearchReposCmd struct {
	Mercy   bool   "name:\"mercy-preview\" help:\"The `topics` property for repositories on GitHub is currently available for developers to preview. To view the `topics` property in calls that return repository results, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.mercy-preview+json\n```\""
	Order   string `name:"order"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	Q       string `required:"" name:"q"`
	Sort    string `name:"sort"`
	internal.BaseCmd
}

func (c *SearchReposCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/search/repositories")
	c.UpdatePreview("mercy", c.Mercy)
	c.UpdateURLQuery("order", c.Order)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("q", c.Q)
	c.UpdateURLQuery("sort", c.Sort)
	return c.DoRequest("GET")
}

type SearchReposLegacyCmd struct {
	Keyword   string `required:"" name:"keyword"`
	Language  string `name:"language"`
	Order     string `name:"order"`
	Sort      string `name:"sort"`
	StartPage string `name:"start_page"`
	internal.BaseCmd
}

func (c *SearchReposLegacyCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/legacy/repos/search/:keyword")
	c.UpdateURLPath("keyword", c.Keyword)
	c.UpdateURLQuery("language", c.Language)
	c.UpdateURLQuery("order", c.Order)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("start_page", c.StartPage)
	return c.DoRequest("GET")
}

type SearchTopicsCmd struct {
	Mercy bool   "name:\"mercy-preview\" help:\"The `topics` property for repositories on GitHub is currently available for developers to preview. To view the `topics` property in calls that return repository results, you must provide a custom [media type](https://developer.github.com/v3/media) in the `Accept` header:\n```shell\napplication/vnd.github.mercy-preview+json\n```\""
	Q     string `required:"" name:"q"`
	internal.BaseCmd
}

func (c *SearchTopicsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/search/topics")
	c.UpdatePreview("mercy", c.Mercy)
	c.UpdateURLQuery("q", c.Q)
	return c.DoRequest("GET")
}

type SearchUsersCmd struct {
	Order   string `name:"order"`
	Page    int64  `name:"page"`
	PerPage int64  `name:"per_page"`
	Q       string `required:"" name:"q"`
	Sort    string `name:"sort"`
	internal.BaseCmd
}

func (c *SearchUsersCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/search/users")
	c.UpdateURLQuery("order", c.Order)
	c.UpdateURLQuery("page", c.Page)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("q", c.Q)
	c.UpdateURLQuery("sort", c.Sort)
	return c.DoRequest("GET")
}

type SearchUsersLegacyCmd struct {
	Keyword   string `required:"" name:"keyword"`
	Order     string `name:"order"`
	Sort      string `name:"sort"`
	StartPage string `name:"start_page"`
	internal.BaseCmd
}

func (c *SearchUsersLegacyCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/legacy/user/search/:keyword")
	c.UpdateURLPath("keyword", c.Keyword)
	c.UpdateURLQuery("order", c.Order)
	c.UpdateURLQuery("sort", c.Sort)
	c.UpdateURLQuery("start_page", c.StartPage)
	return c.DoRequest("GET")
}
