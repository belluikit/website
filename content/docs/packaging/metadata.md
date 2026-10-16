---
title: App metadata
description: The manifest.pxf file holds your app's ID, name, version, icon and platform details.
weight: 1
---

Your app's metadata lives in a file named `manifest.pxf` next to its `main`
package. It is written in PXF, a small text format of `key = value` pairs,
`{ … }` blocks and `[ … ]` lists. The `bellui` command reads it when it builds
or packages the app, so you do not repeat the same flags every time.

Here is a manifest with the fields most apps use:

{{< example file="packaging/metadata/manifest.pxf" lang="text" >}}

## The fields

### `details`

| Field | Meaning |
| --- | --- |
| `id` | The app's unique ID in reverse-domain form, like `com.example.notes`. Preferences and storage are keyed on it, and Android, iOS and Windows store releases require it. |
| `name` | The name people see: the bundle name on macOS, the desktop entry name on Linux. |
| `version` | The version, as `x`, `x.y` or `x.y.z`. |
| `build` | A build number that increases with every package. |
| `icon` | The icon file, relative to the manifest. See [App icons](../icons/). |
| `file_associations` | Document types the app opens, each with `extensions`, `mime_type`, `description`, `role` and `icon`. |
| `url_schemes` | Custom URL schemes the app handles, such as `myapp` for `myapp://…` links. |
| `ios_usage_descriptions` | The text iOS shows when the app asks for a permission, keyed by the `Info.plist` name, for example `NSCameraUsageDescription`. The macOS bundle uses the same entries. |
| `associated_domains` | Domains whose links should open the app (iOS Universal Links). |

### Other top-level fields

| Field | Meaning |
| --- | --- |
| `website`, `description` | Descriptive text about the app. |
| `development`, `release` | Your own key–value settings. The app sees one set or the other, depending on how it was built. |
| `linux_and_bsd` | Desktop entry details: `generic_name`, `categories`, `comment`, `keywords` and `exec_params`. |
| `adaptive_icon` | Android adaptive icon layers: `foreground`, `background` and `monochrome`. |
| `languages` | The languages the app is translated into. |
| `source` | `repo` and `dir` of the app's source code, written into the Linux desktop entry. |

## Reading the metadata in your app

[`App.Metadata`](https://pkg.go.dev/github.com/belluikit/bellui#AppMetadata)
returns the values. Your own settings are in `Custom`.

{{< example file="packaging/metadata/main.go" region="meta" >}}

When you start the app with `go run` from the folder that holds the manifest,
Bellui reads `manifest.pxf` at start-up and fills in the metadata, using the
`development` settings. `app.New` then takes the app ID from it.

{{< callout type="warning" >}}
In this release, `bellui build` and `bellui package` use the manifest to build
the package (its name, ID, version and icon) but do not compile the metadata
into the executable: a version check in the command predates Bellui 1.0 and
skips that step. A packaged app started outside its source folder therefore
has empty metadata, and `app.New` has no ID. Until this is fixed, create your
app with `app.NewWithID` and do not rely on `Metadata()` or the `release`
settings in packaged builds.
{{< /callout >}}

The `--metadata key=value` flag of `bellui build`, `package` and `release` is
meant to add a custom value for one build without writing it in the manifest;
it is affected by the same limitation.

## The build number

Every successful `bellui package` or `bellui release` increases `build` in the
manifest by one. The command edits only that number and keeps the rest of the
file, including your comments and layout. Commit the change along with your
release.

## Flags still win

The command-line flags `--name`, `--app-id`, `--app-version`, `--app-build`
and `--icon` override the manifest for one run.

## The complete example

{{< example "packaging/metadata/main.go" >}}
