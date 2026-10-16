---
title: App storage
description: Keep documents in a private folder that belongs to your app, and know where that folder is on each platform.
weight: 3
---

Every app with an ID gets a private storage area.
[`App.Storage`](https://pkg.go.dev/github.com/belluikit/bellui#App) returns a
[`bellui.Storage`](https://pkg.go.dev/github.com/belluikit/bellui#Storage) for
it. Use it for files the user does not pick themselves: drafts, caches, saved
state.

{{< example file="storage/appstorage/main.go" region="root" >}}

## Documents

The `Storage` methods work on named documents in a `Documents` folder inside
the storage area. The folder is created the first time you call `Create`.

| Method | What it does |
| --- | --- |
| `Create(name)` | Opens a new document for writing. Fails with [`storage.ErrAlreadyExists`](https://pkg.go.dev/github.com/belluikit/bellui/storage#pkg-variables) if it exists. |
| `Save(name)` | Opens an existing document for writing. Fails with `storage.ErrNotExists` if it does not exist. |
| `Open(name)` | Opens a document for reading. |
| `Remove(name)` | Deletes a document. |
| `List()` | Returns the names of all documents. |
| `RootURI()` | Returns the URI of the whole storage area. |

Writing is finished only when you close the writer. A save that should work
whether or not the document exists tries `Create` first:

{{< example file="storage/appstorage/main.go" region="save" >}}

Reading is the same as for any other URI:

{{< example file="storage/appstorage/main.go" region="load" >}}

`RootURI` is an ordinary URI, so you can also use the functions in
[Files and URIs](../files/) on it, for example to make your own subfolders
next to `Documents`.

## Where the storage area is

| Platform | Location |
| --- | --- |
| macOS | `~/Library/Preferences/bellui/<app ID>` |
| Windows | `%USERPROFILE%\AppData\Roaming\bellui\<app ID>` |
| Linux and BSD | `bellui/<app ID>` inside the user configuration folder, usually `~/.config` |
| Android | `bellui` inside the app's private files folder |
| iOS | `bellui` inside the app's Documents folder |
| Web | `idbfile:///bellui/`, kept in the browser's IndexedDB |

On Android and iOS the operating system already gives every app its own
sandbox, so the app ID is not part of the path. The same folder also holds the
app's [preferences](../preferences/) file.

{{< callout type="warning" >}}
Without an app ID there is no storage area. `RootURI` logs an error and
returns the system temporary folder, and the document methods return an error
asking you to use `app.NewWithID`.
{{< /callout >}}

## The complete example

{{< example "storage/appstorage/main.go" >}}
