---
title: About dialog
description: Show your app's icon, name and version with some text and links, as a dialog or a window.
weight: 5
---

[`dialog.ShowAbout`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#ShowAbout)
builds an About box from your app's metadata. At the top it shows the app
icon, the name and the version. Your content goes under them, and your links
are placed in a row along the bottom.

{{< example file="dialogs/about/main.go" region="about" >}}

- The content is Markdown, so you can use bold text, lists and links in it.
  Images from the network are not loaded.
- The links are [`widget.Hyperlink`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Hyperlink)
  values.
- The title is "About" followed by the app name.

[`dialog.ShowAboutWindow`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#ShowAboutWindow)
shows the same content in a new window instead of a dialog over an existing
one. `NewAbout` and `NewAboutWindow` return the dialog or window without
showing it.

## Where the name, version and icon come from

They are read from [`App.Metadata`](https://pkg.go.dev/github.com/belluikit/bellui#AppMetadata),
which comes from your app's `manifest.pxf` file. When metadata is missing, the
dialog leaves out what it does not have, and the title is just "About". See
[App metadata](/docs/packaging/metadata/) for the manifest format and for when
the metadata is available in this release.

## The complete example

{{< example "dialogs/about/main.go" >}}
