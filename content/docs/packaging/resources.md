---
title: Bundling resources
description: Compile images, fonts and other files into your app with bellui bundle.
weight: 5
---

An app that loads its images from files next to the executable breaks as soon
as it is moved, packaged for a phone or served on the web. Compiling the files
into the executable avoids that. `bellui bundle` writes a Go file that embeds
each file and exposes it as a
[`bellui.StaticResource`](https://pkg.go.dev/github.com/belluikit/bellui#StaticResource),
which you can pass anywhere Bellui takes a resource: icons, images, fonts,
themes.

## Generating the file

```sh
bellui bundle -o bundled.go logo.svg
bellui bundle -o bundled.go -a welcome.txt
```

The first command creates `bundled.go`; `-a` (append) adds to it without
writing the header again. Each resource gets a variable named from the
`resource` prefix and the file name. The result for these two files:

{{< example "packaging/bundle/bundled.go" >}}

The generated code uses Go's `//go:embed`, so the files must stay in the
package folder; `bellui bundle` does not copy their contents into the Go file.

| Flag | Meaning |
| --- | --- |
| `-o file` | Write to this file instead of standard output. |
| `-a` | Append to an existing file, without a header. |
| `--name` | The variable name, for a single file. |
| `--prefix` | The prefix for generated names. The default is `resource`. |
| `--package` | The package name in the header. The default is `main`. |

Given a folder instead of files, `bellui bundle` bundles every file in it
except `.go` files.

## Regenerating with go generate

Put the commands in `//go:generate` lines so that `go generate` rebuilds the
file when the assets change:

{{< example "packaging/bundle/main.go" >}}

## Using a resource

{{< example file="packaging/bundle/main.go" region="use" >}}

`Content()` returns the bytes and `Name()` the original file name. For your
own embedded files you can also call
[`bellui.NewStaticResource`](https://pkg.go.dev/github.com/belluikit/bellui#NewStaticResource)
directly; `bellui bundle` only saves you writing that code.
