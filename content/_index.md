---
title: Bellui
layout: hextra-home
description: A cross-platform Go UI toolkit with Material Design 3 widgets and WebGPU rendering, for desktop, mobile and the web.
---

<div class="bellui-hero">
<img src="images/logo-256.webp" alt="The Bellui gopher holding a toolkit, in front of an app window" width="160" height="160" />
<div>
{{< hextra/hero-headline >}}
  Native apps in Go,&nbsp;<br class="hx:sm:block hx:hidden" />on every screen
{{< /hextra/hero-headline >}}
</div>
</div>

<div class="hx:mb-12">
{{< hextra/hero-subtitle >}}
  Bellui is a UI toolkit for Go with Material Design 3 widgets and WebGPU rendering.&nbsp;<br class="hx:sm:block hx:hidden" />One codebase runs on macOS, Windows, Linux, Android, iOS and the browser.
{{< /hextra/hero-subtitle >}}
</div>

<div class="hx:mb-6">
{{< hextra/hero-button text="Get started" link="develop" >}}
{{< hextra/hero-button text="Live demo" link="demo/" style="background: transparent; color: inherit; border: 1px solid currentColor;" >}}
{{< hextra/hero-button text="View on GitHub" link="https://github.com/belluikit/bellui" style="background: transparent; color: inherit; border: 1px solid currentColor;" >}}
</div>

<div class="hx:mt-6"></div>

{{< screenshot name="actions" alt="The Bellui gallery app showing Material Design 3 buttons" >}}

<p style="text-align:center;opacity:.75;margin-top:-0.5rem">Every widget, theme and extension in one app: <a href="gallery/">take the gallery tour</a>.</p>

{{< hextra/feature-grid >}}
  {{< hextra/feature-card
    title="Material Design 3"
    subtitle="A full MD3 component set and theme system: navigation bars and rails, sheets, chips, date and time pickers, snackbars, rich text and more, with light and dark schemes."
  >}}
  {{< hextra/feature-card
    title="WebGPU rendering"
    subtitle="One renderer drives Metal, Vulkan, D3D12 and browser WebGPU. The OpenGL backend stays available with the `gl` build tag."
  >}}
  {{< hextra/feature-card
    title="Desktop, mobile and web"
    subtitle="Build for macOS, Windows and Linux, package for Android and iOS, or compile to WebAssembly, all from the same Go code."
  >}}
  {{< hextra/feature-card
    title="Plain Go"
    subtitle="No markup language, no JavaScript bridge, no generated code. Widgets are Go values, and the native libraries ship inside Go modules."
  >}}
  {{< hextra/feature-card
    title="Built for concurrent Go"
    subtitle="Do slow work in goroutines as usual. `bellui.Do`, `DoAndWait` and `RunOnMain` hand the UI update back to the main thread."
  >}}
  {{< hextra/feature-card
    title="Batteries included"
    subtitle="Data binding and reactive state, dialogs, storage, translations, accessibility, sign-in with OIDC, an embedded web view, developer tools and testing helpers."
  >}}
{{< /hextra/feature-grid >}}

<h2 class="bellui-home-heading">Extensions</h2>

<p style="margin-bottom:1.5rem;opacity:.75">Larger components ship as optional Go modules, so an app only pays for what it imports. All of them are drawn by Bellui itself: no HTML, no JavaScript.</p>

{{< hextra/feature-grid >}}
  {{< hextra/feature-card
    title="Charts"
    subtitle="Line, bar, scatter, pie, financial, hierarchical, 3D and geographic charts, with zoom, tooltips and animation."
    link="extensions/charts/"
    image="images/gallery/charts-light.png"
    imageClass="hx:top-[40%] hx:left-[24px] hx:w-[180%] hx:sm:w-[110%] hx:dark:opacity-80"
    class="hx:aspect-auto hx:md:aspect-[1.1/1] hx:max-md:min-h-[340px]"
  >}}
  {{< hextra/feature-card
    title="Data grid"
    subtitle="A spreadsheet-class table: virtual scrolling, sorting, filtering, editing, pinned columns, grouping and sparklines."
    link="extensions/tables/"
    image="images/gallery/data-grid-light.png"
    imageClass="hx:top-[40%] hx:left-[24px] hx:w-[180%] hx:sm:w-[110%] hx:dark:opacity-80"
    class="hx:aspect-auto hx:md:aspect-[1.1/1] hx:max-md:min-h-[340px]"
  >}}
  {{< hextra/feature-card
    title="Code editor"
    subtitle="Syntax highlighting, a gutter with line numbers, breakpoints and diagnostics, folding, find, completion and hover."
    link="extensions/editor/"
    image="images/gallery/code-editor-light.png"
    imageClass="hx:top-[40%] hx:left-[24px] hx:w-[180%] hx:sm:w-[110%] hx:dark:opacity-80"
    class="hx:aspect-auto hx:md:aspect-[1.1/1] hx:max-md:min-h-[340px]"
  >}}
  {{< hextra/feature-card
    title="Graphs"
    subtitle="An interactive node-and-edge canvas for directed acyclic graphs, over a graph model you can use without any UI."
    link="extensions/dag/"
  >}}
  {{< hextra/feature-card
    title="Media"
    subtitle="Video and audio playback with subtitles and device capture, decoded by FFmpeg."
    link="extensions/mmedia/"
  >}}
  {{< hextra/feature-card
    title="Libraries and tools"
    subtitle="The Go libraries Bellui is built on, each usable on its own, and tools we publish alongside it: Alidade, Clair and Uploader."
    link="tools/"
  >}}
{{< /hextra/feature-grid >}}

<h2 class="bellui-home-heading">Hello, Bellui</h2>

{{< example "hello/main.go" >}}

```sh
go mod init example.com/hello
go mod tidy
go run .
```

<p><a class="bellui-link" href="develop/">Install the prerequisites and build your first app &rarr;</a></p>

<div class="bellui-spotlight">
<div class="bellui-spotlight-label">Also from the Bellui team</div>
<h2 class="bellui-home-heading" style="margin-top:0.25rem">Alidade: a map of your code for AI agents</h2>

<p>Alidade turns a whole repository into one graph and serves it to AI coding agents over MCP. An agent can ask who calls a function, what implements an interface, or which code uses a database column, and get the answer the compiler would give: Go through the Go type checker, TypeScript through its own compiler, plus <code>.proto</code>, SQL, Python and Markdown in the same graph. No model reads your code.</p>

<p>In its published benchmark, 47 call-graph questions over four public repositories, agents using Alidade were as accurate as agents using grep, at about half the cost per correct answer.</p>

```sh
brew install belluikit/tap/alidade
alidade setup
```

<p><a href="tools/alidade/">Read about Alidade &rarr;</a> &nbsp;·&nbsp; <a href="https://github.com/belluikit/alidade">github.com/belluikit/alidade</a></p>
</div>
