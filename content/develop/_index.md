---
cascade:
  type: docs
title: Get started
description: Install what Bellui needs, write your first app, and choose a renderer.
weight: 1
sidebar:
  open: true
---

Bellui apps are ordinary Go programs. You need Go, a C compiler for the native
parts, and on Linux a few development headers. After that, `go run .` opens a
window.

{{< cards >}}
  {{< card link="prerequisites" title="Prerequisites" subtitle="Go, a C compiler, and the system packages for your OS." icon="download" >}}
  {{< card link="first-app" title="Your first app" subtitle="A window with a label and a button, in twenty lines." icon="code" >}}
  {{< card link="rendering" title="Renderers" subtitle="WebGPU by default, OpenGL with a build tag, and where each one runs." icon="chip" >}}
  {{< card link="/docs" title="Documentation" subtitle="Concepts, layout, widgets, theming, data binding and more." icon="book-open" >}}
{{< /cards >}}
