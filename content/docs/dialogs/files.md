---
title: File and folder dialogs
description: Let people open and save files, pick folders, and choose several files at once.
weight: 3
---

The file dialogs hand your callback a reader, a writer or a folder, not a file
path. That is what lets the same code work with a desktop file system, with
the document pickers on Android and iOS, and in a browser. Read
[Files and URIs](/docs/storage/files/) for how to use what you get back.

## Opening a file

[`dialog.NewFileOpen`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#NewFileOpen)
gives your callback a
[`bellui.URIReadCloser`](https://pkg.go.dev/github.com/belluikit/bellui#URIReadCloser):
an `io.ReadCloser` that also reports the chosen file's URI.

{{< example file="dialogs/files/main.go" region="open" >}}

Three rules for the callback:

- If the user cancels, the reader is `nil` and so is the error. Check for both.
- If the reader is not `nil`, you must close it.
- The callback runs on the UI thread, so it can update widgets directly.

`SetFilter` limits which files can be chosen.
[`storage.NewExtensionFileFilter`](https://pkg.go.dev/github.com/belluikit/bellui/storage#NewExtensionFileFilter)
matches extensions such as `.txt` (with the dot, ignoring case), and
[`storage.NewMimeTypeFileFilter`](https://pkg.go.dev/github.com/belluikit/bellui/storage#NewMimeTypeFileFilter)
matches MIME types, including wildcards like `image/*`.

## Saving a file

[`dialog.NewFileSave`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#NewFileSave)
works the same way but hands you a
[`bellui.URIWriteCloser`](https://pkg.go.dev/github.com/belluikit/bellui#URIWriteCloser).
If the user picks a file that already exists, they are asked to confirm before
it is overwritten. `SetFileName` suggests a name.

{{< example file="dialogs/files/main.go" region="save" >}}

## Choosing a folder

[`dialog.ShowFolderOpen`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#ShowFolderOpen)
returns a
[`bellui.ListableURI`](https://pkg.go.dev/github.com/belluikit/bellui#ListableURI),
whose `List` method returns the folder's contents.

{{< example file="dialogs/files/main.go" region="folder" >}}

## Choosing several files

[`dialog.ShowFileOpenMultiple`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#ShowFileOpenMultiple)
lets the user pick more than one file. The callback gets a slice of readers, or
`nil` if the user cancelled. Close every reader in the slice.

{{< example file="dialogs/files/main.go" region="multiple" >}}

In Bellui's own dialog, Ctrl-click (Cmd-click on macOS) adds or removes a file
and Shift-click selects a range. On a touch screen, a long press on a file
starts selection mode, where each tap adds or removes a file.

## Native pickers or Bellui's dialog

By default the file dialogs use the platform's own picker where there is one:

| Platform | Picker |
| --- | --- |
| macOS | The system open and save panels |
| Windows | The system file dialogs |
| Linux and BSD | The desktop's file chooser, through the XDG desktop portal |
| Android, iOS | The system document pickers |
| Web | The browser's file picker for opening; saving downloads the file |

When no native picker is available, for example on a Linux desktop with no
portal running, Bellui draws its own dialog instead, so the feature always
works. To use Bellui's dialog everywhere, call
[`dialog.UseNativeFileDialogs(false)`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#UseNativeFileDialogs)
once at start-up.

{{< callout type="warning" >}}
On Android and iOS only the system picker can reach files outside your app's
own storage. With native pickers turned off, Bellui's dialog on a phone can
browse only the app's private files. If you turn them off for a consistent desktop
look, do it only on desktop platforms.
{{< /callout >}}

Some details differ between the two:

- `SetShowHidden(true)`, which shows hidden files, only affects Bellui's own
  dialog; native pickers have no such option.
- A Flatpak build always uses the portal picker, whatever you set.
- In the browser, opening always uses the browser's picker and saving is
  always a download, because a web page has no file system to browse. Folder
  picking and multi-file picking fall back to Bellui's dialog there.
- In the browser the save callback receives a writer at once. The download is
  offered when you close it, and a callback set with `SetOnClosed` runs at that
  moment.

## More options

The [`FileDialog`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#FileDialog)
type also has:

- `SetLocation(bellui.ListableURI)` to choose the starting folder.
- `SetConfirmText`, `SetDismissText` and `SetTitleText` to relabel it.
- `SetView(dialog.ListView)` or `SetView(dialog.GridView)` for Bellui's dialog
  layout.
- `SetWindowed` and `SetWindowHosted`, as described in
  [Messages and confirmation](../messages/#options-every-dialog-shares).

## The complete example

{{< example "dialogs/files/main.go" >}}
