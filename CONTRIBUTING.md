# Contributing to octo-cli

## Requirements

- Go 1.11+
- dep
- golint
- golangci-lint (installed by script/bootstrap)
- bash shell for `./scripts/*`

## Scripts

Octo-cli is using [Scripts to Rule Them All](https://githubengineering.com/scripts-to-rule-them-all/). Or at least a subset of the idea.

- `script/bootstrap` - installs dependencies. This is currently lacking and only installs golangci-lint.  You are on your own for the rest.
- `script/build` - builds the main binary and writes it to `bin/octo`
- `script/cibuild` - This is what ci runs. You can also run it locally so you aren't surprised by a ci failure.  It currently takes about 13 seconds to run on my laptop.
- `script/generate` - Generates the code in `./internal/generated/`.  We'll get to what that is further down the page.
- `script/golint` - runs golint on everything that isn't vendor
- `script/latestversion` - returns the last git tagged version
- `script/lint` - runs all the default linters from golangci-lint
- `script/newversion` - tells you what version to use for a new release flag
- `script/test` - runs tests
- `script/update-routes` - updates `./routes.json` with the latest from https://octokit.github.io/routes/index.json
- `script/update-testdata` - copies `./routes.json` to `./generator/testdata/routes.json` and runs generator to regenerate `./generator/testdata/generated`

## Issues before PRs

Please create issues to discuss enhancements before spending any significant
time writing code. This project has a very specific purpose, and pull
requests that don't move that forward will be rejected. We don't want to
say no to a PR that you have spent hours on because it doesn't fit the project's
vision. Please don't make us do that.

## Questions and Help

If you have questions or get stuck, you can create an issue here asking
for clarification. You can also ask on the [#octo-cli](https://invite.slack.golangbridge.org/)
channel on [gophers slack](https://invite.slack.golangbridge.org/). I
created that channel right before typing this, so if you act fast, you
may just be first post.

## Project structure

Let's start with what this project is. Octo-cli is a command-line interface
for GitHub's REST API that is generated from [Octokit routes](https://octokit.github.io/routes/).
That last bit about being generated is key to understanding the project
structure.

### `./routes.json`
`routes.json` is what we download from [Octokit routes](https://octokit.github.io/routes/).
It is json that describes all the REST endpoints covered in [GitHub's documentation](https://developer.github.com/v3/).
Updating routes.json will eventually be automated. Until then, only trusted
octo-cli developers can update routes.json. PRs with changes to routes.json
will be politely rejected.

`routes-last-modified.txt` is the Last-Modified header from when routes.json
was last downloaded.

### `./generator/`
`generator/` is the code that parses routes.json and creates commands
for each defined endpoint. Most of the action here is in the Generate
function and `const tmplt`. Generate is currently a bit of an oversized
mess. Please bear with us until it is broken up into more easily grokked
pieces.

Generator is not well tested. Tests consist generating new commands and
checking whether they match the expected output. Unit tests are needed here.

### `./generator/testdata`
As you probably guessed this is test data for generator's few tests. This
should only be modified by running `./script/update-testdata`.

### `./internal/`

Just a note about [internal directories](https://golang.org/doc/go1.4#internalpackages)
in a go project. Packages that have a directory named `internal` cannot
be imported by packages in other projects. I've run across other go devs
who don't know this, so I wanted to include the reasoning for this name
here.

### `./internal/generated`

This is where you will find all of our generated code. Generated code
should never be edited manually, and if it doesn't match what is output
by `script/generate` then ci will fail.

These files are used by [kong](https://github.com/alecthomas/kong) to
create octo-cli's sub-commands.

Each struct represents one subcommand.
The top level subcommands are in the `CLI` struct in `cli.go`. These have
their own subcommands that each represent an endpoint. For instance the
`CLI` contains `Issues` with the type `IssuesCmd`. `IssuesCmd` then contains
`Create` with the type `IssuesCreateCmd`.  `IssuesCreateCmd` represents
the command-line `octo issues create`. `IssuesCreateCmd` contains fields
representing the flags for that command.

Each executable command struct has a `Run()` method. This is what is
executed when the command is run. Here flag values are used to modify the
API request then the request is performed by `DoRequest` on the last line.

### `./internal/generated/basecmd.go`

This contains the `baseCmd` struct and it's receivers. `baseCmd` is an
anonymous member of every executable command struct.  This can be thought
of as the low-level api client that the generated code uses.

### `./tests/`

`./tests/` contains tests. Not many of them yet.