---
title: Files and URIs
description: Name resources with URIs and read, write, list, copy and delete them with the storage package.
weight: 1
---

Bellui names files and other resources with a
[`bellui.URI`](https://pkg.go.dev/github.com/belluikit/bellui#URI) rather than
a path. A URI starts with a scheme, such as `file:`, and the scheme decides
which code handles it. That handler is called a **repository**. This is how a
file chosen on an Android phone (a `content:` URI) and a file on a desktop disk
(a `file:` URI) can go through the same functions.

The functions in the [`storage`](https://pkg.go.dev/github.com/belluikit/bellui/storage)
package look up the repository for a URI and call it. Use them rather than the
`os` package when your code has to work on every platform.

## Making URIs

{{< example file="storage/files/main.go" region="uris" >}}

- [`storage.NewFileURI`](https://pkg.go.dev/github.com/belluikit/bellui/storage#NewFileURI)
  turns a file path into a `file:` URI. A relative path is made absolute first.
- [`storage.ParseURI`](https://pkg.go.dev/github.com/belluikit/bellui/storage#ParseURI)
  parses a URI string.
- [`storage.Child`](https://pkg.go.dev/github.com/belluikit/bellui/storage#Child)
  and [`storage.Parent`](https://pkg.go.dev/github.com/belluikit/bellui/storage#Parent)
  move down and up the hierarchy. `Parent` of a root returns
  `repository.ErrURIRoot`.

A URI also tells you its `Name()`, `Extension()` (with the dot), `MimeType()`,
`Scheme()` and `Path()`. Compare two URIs with
[`storage.EqualURI`](https://pkg.go.dev/github.com/belluikit/bellui/storage#EqualURI).

## Writing and reading

[`storage.Writer`](https://pkg.go.dev/github.com/belluikit/bellui/storage#Writer)
opens a resource for writing. For a file it creates the file, or empties it if
it already exists. Always check the error from `Close`: it is part of saving.

{{< example file="storage/files/main.go" region="write" >}}

[`storage.Reader`](https://pkg.go.dev/github.com/belluikit/bellui/storage#Reader)
opens it for reading.

{{< example file="storage/files/main.go" region="read" >}}

[`storage.Appender`](https://pkg.go.dev/github.com/belluikit/bellui/storage#Appender)
opens a resource to add to its end. `CanRead` and `CanWrite` tell you in
advance whether reading or writing is possible.

If you only need the bytes as a
[`bellui.Resource`](https://pkg.go.dev/github.com/belluikit/bellui#Resource),
for example to show an image,
[`storage.LoadResourceFromURI`](https://pkg.go.dev/github.com/belluikit/bellui/storage#LoadResourceFromURI)
reads the whole resource at once.

## Folders and listing

A resource that can contain others, such as a folder, is called **listable**.
[`storage.CreateListable`](https://pkg.go.dev/github.com/belluikit/bellui/storage#CreateListable)
creates one. For files this is "make a folder"; it fails if the parent does
not exist and does not create parents for you.

{{< example file="storage/files/main.go" region="mkdir" >}}

[`storage.List`](https://pkg.go.dev/github.com/belluikit/bellui/storage#List)
returns the URIs inside a listable resource. `CanList` tells you whether a URI
is listable, and
[`storage.ListerForURI`](https://pkg.go.dev/github.com/belluikit/bellui/storage#ListerForURI)
converts a URI into a
[`bellui.ListableURI`](https://pkg.go.dev/github.com/belluikit/bellui#ListableURI),
the type the folder dialog returns.

{{< example file="storage/files/main.go" region="list" >}}

[`storage.IsHidden`](https://pkg.go.dev/github.com/belluikit/bellui/storage#IsHidden)
applies the platform's rule for hidden files: a leading dot on Unix-like
systems, the hidden attribute on Windows. It always reports `false` on mobile
and the web.

## Copying, moving and deleting

{{< example file="storage/files/main.go" region="copy" >}}

[`storage.Copy`](https://pkg.go.dev/github.com/belluikit/bellui/storage#Copy),
[`storage.Move`](https://pkg.go.dev/github.com/belluikit/bellui/storage#Move),
[`storage.Delete`](https://pkg.go.dev/github.com/belluikit/bellui/storage#Delete)
and [`storage.DeleteAll`](https://pkg.go.dev/github.com/belluikit/bellui/storage#DeleteAll)
do what their names say. `DeleteAll` removes a folder and everything in it.

## Which schemes work where

Not every repository supports every operation. When an operation is not
available for a URI's scheme, the function returns
[`repository.ErrOperationNotSupported`](https://pkg.go.dev/github.com/belluikit/bellui/storage/repository#pkg-variables).
The schemes Bellui registers are:

| Scheme | Where | Notes |
| --- | --- | --- |
| `file` | Desktop, Android, iOS | Local files. Full read, write and list support on the desktop. |
| `content` | Android | Files handed over by the Android document picker. |
| `http`, `https` | Everywhere | Reading from the web. |
| `idbfile` | Web | The browser's IndexedDB, used for app storage on the web. |

To add your own scheme, see [Custom repositories](../repositories/).

## The complete example

{{< example "storage/files/main.go" >}}
