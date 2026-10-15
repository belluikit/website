---
cascade:
  type: docs
title: Tools
description: Separate tools from the Bellui team. A file transfer service, a code search index, and a code graph for AI agents.
weight: 5
---

These tools come from the same team as Bellui. None of them is part of the
toolkit, and you do not need any of them to build a Bellui app. Each has its
own repository, its own release cycle and its own license.

{{< cards >}}
  {{< card link="uploader" title="uploader" subtitle="A file transfer service over gRPC, with a Go client. Uploads and downloads that are cut off resume where they stopped." icon="cloud-upload" >}}
  {{< card link="clair" title="Clair" subtitle="A searchable index of a Go framework's API and of your own project, with optional questions and answers through a language model." icon="document-search" >}}
  {{< card link="alidade" title="alidade" subtitle="A graph of your whole repository for AI coding agents: who calls what, across Go, TypeScript, Python, .proto, SQL and Markdown." icon="share" >}}
{{< /cards >}}

| Tool | What you get | Current release | License |
| --- | --- | --- | --- |
| [Uploader](uploader/) | A server binary and a Go SDK | v0.1.0 | BSD-3-Clause |
| [Clair](clair/) | A Go library and the `kbtool` command | No tagged release yet | Apache-2.0 |
| [Alidade](alidade/) | One binary, an MCP server | v0.13.0 | MIT |

The [gallery app](/gallery/) uses two of them. Its File Transfer page is a
Bellui front end for an Uploader server, and its Clair page searches an index
of Bellui's own widgets.
