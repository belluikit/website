---
title: Packaging and distribution
description: Use the bellui command to build, package, install and release your app on every platform.
weight: 12
---

`go build` gives you an executable. To ship an app you usually need more: an
application bundle on macOS, an icon and version information in a Windows
executable, a desktop entry on Linux, an APK for Android, an app bundle for
iOS, or a folder of files for the web. The `bellui` command builds these for
you.

## Install the command

The command is its own Go module, separate from the Bellui library:

```sh
go install github.com/belluikit/bellui/cmd/bellui@latest
```

Check that it is on your `PATH`:

```sh
bellui version
```

This prints the version of the command and, when run inside your app's module,
the version of the Bellui library your app uses.

## Commands

| Command | What it does |
| --- | --- |
| `bellui init [module]` | Starts a new project in the current folder: `go.mod`, `main.go`, `manifest.pxf` and `translations/en.json`. |
| `bellui build` | Builds the executable. In this release it does not compile your app's metadata in; see [App metadata](metadata/). |
| `bellui package` | Builds and wraps the app in the package format of the target platform. |
| `bellui install` | Packages the app and installs it on this computer, or on a connected phone. |
| `bellui release` | Packages and signs the app for an app store or other public distribution. |
| `bellui serve` | Builds the web version and serves it on `http://localhost:8080`. |
| `bellui bundle` | Turns files into Go resources compiled into your app. |
| `bellui translate` | Collects the strings your code translates into a translation file. |
| `bellui env` | Prints the Go, OS and Bellui versions, for bug reports. |

Run `bellui help <command>` for every flag of a command.

{{< callout type="warning" >}}
In this release the `main.go` that `bellui init` writes ends with
`w.ShowAndRun()`, which does not compile. Change that line to
`bellui.ShowAndRun(w)`.
{{< /callout >}}

{{< cards >}}
  {{< card link="metadata/" title="App metadata" subtitle="The manifest.pxf file: ID, name, version, icon and more." >}}
  {{< card link="package/" title="Packaging and installing" subtitle="What bellui package produces on each platform, and bellui install." >}}
  {{< card link="release/" title="Releasing" subtitle="Signing and store packages for macOS, Windows, Android and iOS." >}}
  {{< card link="icons/" title="App icons" subtitle="Where the icon comes from and how each platform uses it." >}}
  {{< card link="resources/" title="Bundling resources" subtitle="Compile images and other files into your app with bellui bundle." >}}
  {{< card link="settings-app/" title="The settings app" subtitle="Let users choose theme, colour and scale for all Bellui apps." >}}
{{< /cards >}}
