---
title: Resources and images
description: Embed files in your app as resources, use the theme's icons, and load images from files and URIs.
weight: 5
---

A **resource** is a named block of bytes, such as an image, an icon or a
font. Anything in Bellui that shows an image takes a
[`bellui.Resource`](https://pkg.go.dev/github.com/belluikit/bellui#Resource):
`canvas.NewImageFromResource`, `widget.NewIcon`, `widget.NewButtonWithIcon`,
`App.SetIcon` and more. The interface has two methods, `Name()` and
`Content()`. The name identifies the resource, and its extension tells
Bellui the format, so an SVG file's name should end in `.svg`.

## Embed files in your binary

Go's `embed` package puts a file into the compiled program.
[`bellui.NewStaticResource`](https://pkg.go.dev/github.com/belluikit/bellui#NewStaticResource)
turns the bytes into a resource:

{{< example file="concepts/resources/main.go" region="embed" >}}

The file `logo.svg` sits next to `main.go`. Embedded resources need no file
at run time, so this works the same on desktop, mobile and web.

Once you have a resource, use it wherever an image is wanted:

{{< example file="concepts/resources/main.go" region="use" >}}

## Theme icons

The [`theme`](https://pkg.go.dev/github.com/belluikit/bellui/theme) package
has a function for each standard icon: `theme.HomeIcon()`,
`theme.SettingsIcon()`, `theme.SearchIcon()`, `theme.DeleteIcon()`,
`theme.InfoIcon()` and many more. `theme.Icon(name)` looks one up by its
`theme.IconName...` constant. The standard icons take their colour from the
current theme, so they look right in light and dark mode.

To give your own single-colour SVG the same behaviour, wrap it.
`theme.NewThemedResource` uses the foreground colour;
`theme.NewPrimaryThemedResource`, `NewErrorThemedResource`,
`NewWarningThemedResource` and `NewSuccessThemedResource` use those colours.
`theme.NewColoredResource(res, name)` takes any theme colour name.

{{< example file="concepts/resources/main.go" region="theme" >}}

## Load images at run time

When the image is not known at build time, load it:

| Function | Source |
| --- | --- |
| [`canvas.NewImageFromFile(path)`](https://pkg.go.dev/github.com/belluikit/bellui/canvas#NewImageFromFile) | A file path |
| [`canvas.NewImageFromURI(uri)`](https://pkg.go.dev/github.com/belluikit/bellui/canvas#NewImageFromURI) | A [`bellui.URI`](https://pkg.go.dev/github.com/belluikit/bellui#URI): a file, or another scheme such as `https` |
| [`canvas.NewImageFromReader(r, name)`](https://pkg.go.dev/github.com/belluikit/bellui/canvas#NewImageFromReader) | Any `io.Reader`; the name identifies it and gives its format |
| [`canvas.NewImageFromImage(img)`](https://pkg.go.dev/github.com/belluikit/bellui/canvas#NewImageFromImage) | A Go `image.Image` you already have |

{{< example file="concepts/resources/main.go" region="file" >}}

[`storage.NewFileURI`](https://pkg.go.dev/github.com/belluikit/bellui/storage#NewFileURI)
makes a URI from a path. For a `file` URI, `NewImageFromURI` is the same as
`NewImageFromFile` with that path.

{{< callout type="warning" >}}
For any scheme other than `file`, `NewImageFromURI` downloads the data
before it returns. Do not call it with a network URI from a callback, or the
window freezes until the download ends. Load it on a goroutine and set the
image with [`bellui.Do`](../ui-thread/).
{{< /callout >}}

To make a resource rather than an image:

- [`bellui.LoadResourceFromPath(path)`](https://pkg.go.dev/github.com/belluikit/bellui#LoadResourceFromPath)
  reads a file into a static resource.
- [`storage.LoadResourceFromURI(uri)`](https://pkg.go.dev/github.com/belluikit/bellui/storage#LoadResourceFromURI)
  does the same for a URI.
- [`bellui.NewLazyResource(name, loader)`](https://pkg.go.dev/github.com/belluikit/bellui#NewLazyResource)
  and `NewLazyResourceFromPath(path)` load their bytes on first use and keep
  them. `Release()` frees the bytes; the next use loads them again.
- [`storage.NewAssetBundle(dir)`](https://pkg.go.dev/github.com/belluikit/bellui/storage#NewAssetBundle)
  serves a directory of files as lazily loaded resources, by relative name.

## The whole example

{{< example "concepts/resources/main.go" >}}

Run it with an image path to also show an image loaded from disk:

```sh
go run . ~/Pictures/photo.png
```
