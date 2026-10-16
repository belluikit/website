---
title: Clair
description: A searchable index of a Go framework's API and of your own project, with an optional question-and-answer layer through a language model.
weight: 2
---

Clair is a knowledge base for code. It reads Go source, records every
declaration with its signature and doc comment, and puts them in a full-text
search index. You can then search that index offline, or let a language model
answer questions with the matching entries added to its prompt.

It is meant for people building developer tools: an editor's "search the API"
box, or a coding assistant that should know a framework's real API. It is a
Go library plus a small command, `kbtool`, for building and searching
indexes.

- Repository: [github.com/belluikit/clair](https://github.com/belluikit/clair)
- License: Apache-2.0
- Release: there is no tagged release yet. `@latest` gives you the newest
  commit on the main branch.

## How it is organised

Clair keeps two indexes and searches them together:

- **The framework index** is built once from a framework's source and opened
  read-only. For a Bellui app, that is Bellui's API.
- **The project index** covers your own code. It is read-write, and a file
  watcher keeps it up to date as you edit.

Searches go to both and the results are merged. The index is a
[Bleve](https://blevesearch.com/) index on disk. Only Go source can be parsed
today.

The model layer is optional, and search works without it. When you add one,
Clair has adapters for Anthropic, for Ollama, and for servers with an
OpenAI-compatible API.

A privacy policy decides which indexed files may be included in what is sent
to a remote model. The default policy leaves out `.env` files and files whose
names suggest credentials, secrets or keys (`credentials*`, `secrets*`,
`*_secret*`, `private_key*`, `*.pem`, `*.key`). It filters search results
only: your question itself is sent as you typed it.

## Install

Requires Go 1.26 or later.

The command:

```sh
go install github.com/belluikit/clair/cmd/kbtool@latest
```

The library, in your module:

```sh
go get github.com/belluikit/clair@latest
```

## Quick start: search Bellui's widgets

{{% steps %}}

### Find the Bellui source

Inside your app's module, Go already has the Bellui source in its module
cache. This prints where:

```sh
go list -m -f '{{.Dir}}' github.com/belluikit/bellui
```

### Build an index

Point `kbtool build` at a directory. `--module` is the framework's module
path, and `--output` is where the index is written:

```sh
kbtool build --source "$(go list -m -f '{{.Dir}}' github.com/belluikit/bellui)/widget" \
  --module github.com/belluikit/bellui --output bellui-widget.bleve
```

It prints how many declarations it found.

### Search it

```sh
kbtool search bellui-widget.bleve password entry
```

Each of the top ten hits shows its name, kind, package, signature and doc
comment. For this query the first hit is `widget.NewPasswordEntry`.

{{% /steps %}}

`kbtool` also has `inspect` and `validate` subcommands. In the current code
both report zero entries for an index that `search` finds results in, so do
not rely on them yet.

## Using the library

The README's [Library usage](https://github.com/belluikit/clair#library-usage)
section shows the outline: open the two indexes with `kb.OpenFramework` and
`kb.OpenProject`, combine them with `assistant.New`, and then call
`Builtin.Search` for offline search or `AI.Ask` for a streamed answer.
`assistant.WithProvider` adds a model and `assistant.WithPrivacy` sets the
policy.

One detail differs from that README: the provider constructors take a
`Config` struct, not a bare key. For Anthropic, that is
`anthropic.New(anthropic.Config{APIKey: key})`, with optional `Model` and
`MaxContextTokens` fields.

To keep a project index current, the `indexer` package indexes a set of
directories once and then watches them for changes.

## With Bellui

The [gallery app](/gallery/) uses Clair on two pages. Its Clair page indexes
Bellui's `widget` and `container` packages, and searches them when you press Enter.
Its Code Assistant page sends a question and a sample program to Anthropic's
API through Clair, with matching Bellui entries added as context. See
[the gallery tour](/gallery/#clair) for what each needs to run.
