// Code generated by octo-cli/generator; DO NOT EDIT.

package generated

import "github.com/octo-cli/octo-cli/internal"

type ChecksCmd struct {
	ListForRef       ChecksListForRefCmd       `cmd:"" help:"List check runs for a specific ref - https://developer.github.com/v3/checks/runs/#list-check-runs-for-a-specific-ref"`
	ListForSuite     ChecksListForSuiteCmd     `cmd:"" help:"List check runs in a check suite - https://developer.github.com/v3/checks/runs/#list-check-runs-in-a-check-suite"`
	Get              ChecksGetCmd              `cmd:"" help:"Get a single check run - https://developer.github.com/v3/checks/runs/#get-a-single-check-run"`
	ListAnnotations  ChecksListAnnotationsCmd  `cmd:"" help:"List annotations for a check run - https://developer.github.com/v3/checks/runs/#list-annotations-for-a-check-run"`
	GetSuite         ChecksGetSuiteCmd         `cmd:"" help:"Get a single check suite - https://developer.github.com/v3/checks/suites/#get-a-single-check-suite"`
	ListSuitesForRef ChecksListSuitesForRefCmd `cmd:"" help:"List check suites for a specific ref - https://developer.github.com/v3/checks/suites/#list-check-suites-for-a-specific-ref"`
	CreateSuite      ChecksCreateSuiteCmd      `cmd:"" help:"Create a check suite - https://developer.github.com/v3/checks/suites/#create-a-check-suite"`
	RerequestSuite   ChecksRerequestSuiteCmd   `cmd:"" help:"Rerequest check suite - https://developer.github.com/v3/checks/suites/#rerequest-check-suite"`
}

type ChecksListForRefCmd struct {
	internal.BaseCmd
	Antiope   bool   "name:\"antiope-preview\" required:\"\" help:\"The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.antiope-preview+json\n\n```\""
	Owner     string `required:"" name:"owner"`
	Repo      string `required:"" name:"repo"`
	Ref       string `required:"" name:"ref"`
	CheckName string "name:\"check_name\" help:\"Returns check runs with the specified `name`.\""
	Status    string "name:\"status\" help:\"Returns check runs with the specified `status`. Can be one of `queued`, `in_progress`, or `completed`.\""
	Filter    string "name:\"filter\" help:\"Filters check runs by their `completed_at` timestamp. Can be one of `latest` (returning the most recent check runs) or `all`.\""
	PerPage   int64  `name:"per_page" help:"Results per page (max 100)"`
	Page      int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *ChecksListForRefCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/commits/:ref/check-runs")
	c.UpdatePreview("antiope", c.Antiope)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("ref", c.Ref)
	c.UpdateURLQuery("check_name", c.CheckName)
	c.UpdateURLQuery("status", c.Status)
	c.UpdateURLQuery("filter", c.Filter)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ChecksListForSuiteCmd struct {
	internal.BaseCmd
	Antiope      bool   "name:\"antiope-preview\" required:\"\" help:\"The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.antiope-preview+json\n\n```\""
	Owner        string `required:"" name:"owner"`
	Repo         string `required:"" name:"repo"`
	CheckSuiteId int64  `required:"" name:"check_suite_id"`
	CheckName    string "name:\"check_name\" help:\"Returns check runs with the specified `name`.\""
	Status       string "name:\"status\" help:\"Returns check runs with the specified `status`. Can be one of `queued`, `in_progress`, or `completed`.\""
	Filter       string "name:\"filter\" help:\"Filters check runs by their `completed_at` timestamp. Can be one of `latest` (returning the most recent check runs) or `all`.\""
	PerPage      int64  `name:"per_page" help:"Results per page (max 100)"`
	Page         int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *ChecksListForSuiteCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/check-suites/:check_suite_id/check-runs")
	c.UpdatePreview("antiope", c.Antiope)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("check_suite_id", c.CheckSuiteId)
	c.UpdateURLQuery("check_name", c.CheckName)
	c.UpdateURLQuery("status", c.Status)
	c.UpdateURLQuery("filter", c.Filter)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ChecksGetCmd struct {
	internal.BaseCmd
	Antiope    bool   "name:\"antiope-preview\" required:\"\" help:\"The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.antiope-preview+json\n\n```\""
	Owner      string `required:"" name:"owner"`
	Repo       string `required:"" name:"repo"`
	CheckRunId int64  `required:"" name:"check_run_id"`
}

func (c *ChecksGetCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/check-runs/:check_run_id")
	c.UpdatePreview("antiope", c.Antiope)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("check_run_id", c.CheckRunId)
	return c.DoRequest("GET")
}

type ChecksListAnnotationsCmd struct {
	internal.BaseCmd
	Antiope    bool   "name:\"antiope-preview\" required:\"\" help:\"The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.antiope-preview+json\n\n```\""
	Owner      string `required:"" name:"owner"`
	Repo       string `required:"" name:"repo"`
	CheckRunId int64  `required:"" name:"check_run_id"`
	PerPage    int64  `name:"per_page" help:"Results per page (max 100)"`
	Page       int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *ChecksListAnnotationsCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/check-runs/:check_run_id/annotations")
	c.UpdatePreview("antiope", c.Antiope)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("check_run_id", c.CheckRunId)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ChecksGetSuiteCmd struct {
	internal.BaseCmd
	Antiope      bool   "name:\"antiope-preview\" required:\"\" help:\"The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.antiope-preview+json\n\n```\""
	Owner        string `required:"" name:"owner"`
	Repo         string `required:"" name:"repo"`
	CheckSuiteId int64  `required:"" name:"check_suite_id"`
}

func (c *ChecksGetSuiteCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/check-suites/:check_suite_id")
	c.UpdatePreview("antiope", c.Antiope)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("check_suite_id", c.CheckSuiteId)
	return c.DoRequest("GET")
}

type ChecksListSuitesForRefCmd struct {
	internal.BaseCmd
	Antiope   bool   "name:\"antiope-preview\" required:\"\" help:\"The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.antiope-preview+json\n\n```\""
	Owner     string `required:"" name:"owner"`
	Repo      string `required:"" name:"repo"`
	Ref       string `required:"" name:"ref"`
	AppId     int64  "name:\"app_id\" help:\"Filters check suites by GitHub App `id`.\""
	CheckName string `name:"check_name" help:"Filters checks suites by the name of the [check run](https://developer.github.com/v3/checks/runs/)."`
	PerPage   int64  `name:"per_page" help:"Results per page (max 100)"`
	Page      int64  `name:"page" help:"Page number of the results to fetch."`
}

func (c *ChecksListSuitesForRefCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/commits/:ref/check-suites")
	c.UpdatePreview("antiope", c.Antiope)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("ref", c.Ref)
	c.UpdateURLQuery("app_id", c.AppId)
	c.UpdateURLQuery("check_name", c.CheckName)
	c.UpdateURLQuery("per_page", c.PerPage)
	c.UpdateURLQuery("page", c.Page)
	return c.DoRequest("GET")
}

type ChecksCreateSuiteCmd struct {
	internal.BaseCmd
	Antiope bool   "name:\"antiope-preview\" required:\"\" help:\"The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.antiope-preview+json\n\n```\""
	Owner   string `required:"" name:"owner"`
	Repo    string `required:"" name:"repo"`
	HeadSha string `required:"" name:"head_sha" help:"The sha of the head commit."`
}

func (c *ChecksCreateSuiteCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/check-suites")
	c.UpdatePreview("antiope", c.Antiope)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateBody("head_sha", c.HeadSha)
	return c.DoRequest("POST")
}

type ChecksRerequestSuiteCmd struct {
	internal.BaseCmd
	Antiope      bool   "name:\"antiope-preview\" required:\"\" help:\"The Checks API is currently available for developers to preview. During the preview period, the API may change without advance notice. Please see the [blog post](/changes/2018-05-07-new-checks-api-public-beta/) for full details. To access the API during the preview period, you must provide a custom [media type](/v3/media) in the `Accept` header:\n\n```\n  application/vnd.github.antiope-preview+json\n\n```\""
	Owner        string `required:"" name:"owner"`
	Repo         string `required:"" name:"repo"`
	CheckSuiteId int64  `required:"" name:"check_suite_id"`
}

func (c *ChecksRerequestSuiteCmd) Run(isValueSetMap map[string]bool) error {
	c.SetIsValueSetMap(isValueSetMap)
	c.SetURLPath("/repos/:owner/:repo/check-suites/:check_suite_id/rerequest")
	c.UpdatePreview("antiope", c.Antiope)
	c.UpdateURLPath("owner", c.Owner)
	c.UpdateURLPath("repo", c.Repo)
	c.UpdateURLPath("check_suite_id", c.CheckSuiteId)
	return c.DoRequest("POST")
}
