---
title: Alidade
description: A graph of your whole repository for AI coding agents, served over MCP. It answers who calls a function, what implements an interface, and which code uses a schema.
weight: 3
---

Alidade gives an AI coding agent a map of your repository. The agent can ask
who calls a function, what implements an interface, or which code uses a
database column, and get the answer from the compilers rather than from a
text search.

Alidade is not a Bellui library, and it works on any repository. It is
useful to you if you build a Bellui app, or anything else in Go, with an AI
agent such as Claude Code. A text search for a method name returns every type
that has a method with that name. Alidade knows which one a call actually
reaches.

- Repository: [github.com/belluikit/alidade](https://github.com/belluikit/alidade)
- License: MIT. A release binary also contains other projects' code under
  their own licenses, listed in
  [THIRD_PARTY_NOTICES.txt](https://github.com/belluikit/alidade/blob/main/THIRD_PARTY_NOTICES.txt).
- Current release: [v0.13.0](https://github.com/belluikit/alidade/releases/tag/v0.13.0).
  The project calls itself early.

## What it reads

| Layer | Read by |
| --- | --- |
| Go, with types resolved | the Go type checker |
| TypeScript, JavaScript, Svelte | TypeScript's own compiler |
| Python | jedi, in a Python interpreter built into the binary |
| `.proto` | a protobuf compiler |
| SQL (PostgreSQL, sqlc) | `sqlc.yaml` and the migrations |
| `.pxf` | a PXF parser |
| Markdown | a Markdown parser |

Generated code joins the layers. A `.proto` message links to its generated Go
and TypeScript types, and a SQL query to its generated Go function. A
question about one can then reach the code that uses the other.

No language model reads your code to build the graph, so building it costs
no tokens. Two builds of the same files give the same graph.

## Install

With Homebrew, on macOS or Linux:

```sh
brew install belluikit/tap/alidade
```

Or download a binary from the
[releases page](https://github.com/belluikit/alidade/releases). There are
builds for macOS and Linux, each on arm64 and amd64. The Linux builds need
glibc 2.29 or newer. To build from source, see
[Building Alidade](https://github.com/belluikit/alidade/blob/main/docs/building.md).

Check the install:

```sh
alidade version
```

The binary needs nothing else installed, except that it reads Go through the
`go` command on your `PATH`. A repository that `go build` can build is one
Alidade can read. A `go` newer than the one Alidade was built with is refused,
with an error naming both versions, so keep Alidade up to date with your Go.

## Quick start with Claude Code

{{% steps %}}

### Register it once

```sh
alidade setup
```

This registers Alidade with Claude Code for your user account, so it serves
every repository you open. Running it again after an upgrade is safe.

### Ask a question

Open a repository and ask how its code relates. The graph is built on the
first question, which takes a few seconds for a repository of a few hundred
packages.

### Optionally, make the agent use it

```sh
alidade setup -enforce
```

This adds a hook that refuses grep inside the repository and points the agent
to Alidade's tools instead. `alidade setup -no-enforce` removes the hook
again.

{{% /steps %}}

For any other MCP client, run `alidade serve -root /path/to/repo` as a stdio
server. `alidade setup -print` shows the configuration to use.

## What the agent gets

| Tool | What it answers |
| --- | --- |
| `subgraph_for_question` | The neighbourhood of whatever a question names, across all layers |
| `find_nodes` | An ID from a name, such as `Total` or `Workspace.ReferencesFor` |
| `neighbors` | What is within a few steps of a node: callers, callees, implementations |
| `shortest_path` | How two nodes connect |
| `get_node` | One node, with its position and how many edges it has |
| `search_text` | Text the graph does not hold, such as a string or a log message, with the declaration around each hit |
| `graph_stats` | What the graph holds, and which files it could not read |
| `graph_issues` | What the graph left out, and why |
| `rebuild` | Builds the graph again now. Rarely needed |

Answers stay current. Before each tool call, Alidade checks whether any file
changed and rebuilds what did, so there is no index to refresh after an edit
or a checkout.

Alidade says what it does not know. Each edge in the graph records how it was
found, an answer cut short by its size limit says what it left out, and
`graph_issues` lists what was dropped.

It runs on your machine and talks to the agent over standard input and
output. It opens no network port.

## Things that help with a Go repository

These flags go after `alidade serve`:

- **Several modules.** Every `go.mod` under the root becomes part of the
  graph, and calls between the modules are joined. A repository with several
  Go modules needs no extra setup.
- **Platform-specific files.** By default Alidade loads the Go build
  configuration of the machine it runs on. `-go-configs auto` also loads the
  configurations that bring in files the build tags exclude there. You can
  also pass a list such as `linux/amd64,windows/arm64`.
- **Tests.** Test files are left out unless you pass `-tests`.

Directories named `vendor` and `third_party` are skipped. To skip a different
list, put it in a `.alidade.pxf` file at the root of the repository.

## More

The repository's [Using Alidade](https://github.com/belluikit/alidade/blob/main/docs/usage.md)
guide covers every flag, what each kind of edge means, and the known limits.
One limit to know: a file whose path contains a space or `#` is skipped and
listed in `graph_stats`.
