---
title: Custom repositories
description: Register a repository so the storage functions, and widgets built on them, work with your own URI scheme.
weight: 4
---

A repository is the code behind one URI scheme. When you call
`storage.Reader(u)`, the storage package looks up the repository registered
for `u.Scheme()` and asks it for a reader. By registering your own repository
you can make, say, `mem://` or `s3://` URIs work with every function in the
[`storage`](https://pkg.go.dev/github.com/belluikit/bellui/storage) package.

## What a repository implements

Every repository implements
[`repository.Repository`](https://pkg.go.dev/github.com/belluikit/bellui/storage/repository#Repository):
`Exists`, `Reader`, `CanRead`, and `Destroy`, which is called when the
repository is replaced. Extra abilities come from extra interfaces. The storage
functions check for them and return `repository.ErrOperationNotSupported` when
one is missing. `storage.DeleteAll` is the exception: without
`DeleteAllRepository` it falls back to `repository.GenericDeleteAll`.

| Interface | Methods | Enables |
| --- | --- | --- |
| `WritableRepository` | `Writer`, `CanWrite`, `Delete` | `storage.Writer`, `CanWrite`, `Delete` |
| `AppendableRepository` | `Appender` | `storage.Appender` |
| `ListableRepository` | `CanList`, `List`, `CreateListable` | `storage.List`, `CanList`, `CreateListable` |
| `HierarchicalRepository` | `Parent`, `Child` | `storage.Parent`, `storage.Child` |
| `CopyableRepository` | `Copy` | `storage.Copy` |
| `MovableRepository` | `Move` | `storage.Move` |
| `DeleteAllRepository` | `DeleteAll` | A faster `storage.DeleteAll` |
| `CustomURIRepository` | `ParseURI` | Your own parsing in `storage.ParseURI` |
| `WatchableRepository` | `WatchURI` | Change notifications, for example to refresh a file tree |

The package also has helpers such as `repository.GenericCopy`,
`GenericMove`, `GenericParent`, `GenericChild` and `GenericDeleteAll`, built
from the basic operations, for repositories that have no faster way.

## An in-memory repository

This repository keeps files in a map. It implements the required methods and
`WritableRepository`:

{{< example file="storage/repository/main.go" region="repo" >}}

Register it once, before you use the scheme. Without a `CustomURIRepository`,
`storage.ParseURI` parses the string as a standard URI.

{{< example file="storage/repository/main.go" region="register" >}}

From then on the ordinary storage functions work with `mem://` URIs. Operations
the repository does not offer fail cleanly:

{{< example file="storage/repository/main.go" region="use" >}}

## Things to know

- Registering a scheme a second time replaces the earlier repository and calls
  its `Destroy`.
- Scheme names are matched in lower case.
- A storage function calls your repository on the caller's goroutine. If your
  app uses storage from background goroutines, guard shared state, as the
  mutex does here.
- A `WatchableRepository` callback can run on any goroutine. Use `bellui.Do` to
  touch the UI from it.

## The complete example

{{< example "storage/repository/main.go" >}}
