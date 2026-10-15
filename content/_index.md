---
title: Bellui
layout: hextra-home
description: A cross-platform Go UI toolkit with Material Design 3 widgets and WebGPU rendering, for desktop, mobile and the web.
---

<div class="hx:mt-6 hx:mb-6">
{{< hextra/hero-headline >}}
  Native apps in Go,&nbsp;<br class="hx:sm:block hx:hidden" />on every screen
{{< /hextra/hero-headline >}}
</div>

<div class="hx:mb-12">
{{< hextra/hero-subtitle >}}
  Bellui is a UI toolkit for Go with Material Design 3 widgets and WebGPU rendering.&nbsp;<br class="hx:sm:block hx:hidden" />One codebase runs on macOS, Windows, Linux, Android, iOS and the browser.
{{< /hextra/hero-subtitle >}}
</div>

<div class="hx:mb-6">
{{< hextra/hero-button text="Get started" link="develop" >}}
{{< hextra/hero-button text="View on GitHub" link="https://github.com/belluikit/bellui" style="background: transparent; color: inherit; border: 1px solid currentColor;" >}}
</div>

<div class="hx:mt-6"></div>

<img src="images/widgets-light.png" alt="The Bellui gallery app showing Material Design 3 widgets" class="hx:dark:hidden" style="max-width:100%; margin: 2rem auto;" />
<img src="images/widgets-dark.png" alt="The Bellui gallery app showing Material Design 3 widgets in dark mode" class="hx:hidden hx:dark:block" style="max-width:100%; margin: 2rem auto;" />

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
    subtitle="Data binding and reactive state, dialogs, storage, translations, accessibility, testing helpers, and extensions for charts, tables, a code editor, graphs and media."
  >}}
{{< /hextra/feature-grid >}}

## Hello, Bellui

{{< example "hello/main.go" >}}

```sh
go mod init example.com/hello
go mod tidy
go run .
```

[Install the prerequisites and build your first app &rarr;](develop/)
