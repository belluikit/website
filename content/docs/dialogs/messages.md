---
title: Messages and confirmation
description: Information, error and confirmation dialogs, plus the options every dialog shares.
weight: 1
---

## Information and errors

[`dialog.ShowInformation`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#ShowInformation)
shows a title, a message and an **OK** button.
[`dialog.ShowError`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#ShowError)
takes an `error` instead. Its title is "Error", and the message is the error's
text with the first letter capitalised. Do not pass a nil error.

{{< example file="dialogs/messages/main.go" region="info" >}}

A message dialog sizes itself to its text. It is never wider than 600 units,
nor wider than 90% of the parent window, and longer text wraps.

## Asking for confirmation

[`dialog.NewConfirm`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#NewConfirm)
shows **No** and **Yes** buttons. The callback receives `true` when the user
chose **Yes**, and `false` when they chose **No** or dismissed the dialog in
any other way, for example with the Escape key.

{{< example file="dialogs/messages/main.go" region="confirm" >}}

The returned
[`ConfirmDialog`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#ConfirmDialog)
lets you relabel both buttons with `SetConfirmText` and `SetDismissText`. Use
`SetConfirmImportance` to change how prominent the confirm button looks; it is
`widget.HighImportance` by default, and `widget.DangerImportance` suits a
destructive action. You can also close the dialog from code with `Confirm()`
(as if the user chose yes) or `Dismiss()` (as if they chose no).

## Options every dialog shares

The [`Dialog`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#Dialog)
interface is implemented by all the dialogs in the package. Besides `Show`,
`Hide` and `Resize`, it has:

| Method | What it does |
| --- | --- |
| `SetOnClosed(func())` | Runs your function whenever the dialog closes, whatever the answer. Calling it twice keeps both functions. |
| `SetDismissText(string)` | Relabels the dismiss button. |
| `Dismiss()` | Closes the dialog without an affirmative answer. |
| `SetWindowed(bool)` | Windowed (the default) gives the dialog a frame the user can drag by its title bar and resize from its corners. Pass `false` for a fixed, centred dialog. Call it before `Show`. |
| `SetWindowHosted(bool)` | Opens the dialog in its own OS window instead of inside the parent window. Call it before `Show`. |

### Dialogs in their own window

By default a dialog is drawn inside the parent window, over its content. With
`SetWindowHosted(true)` it opens in a separate window that the operating system
can move and resize, and that appears in the window switcher. While it is open
the parent window is covered by a scrim (a translucent layer) that blocks
input; tapping the scrim brings the dialog window to the front. Your app's
other windows keep working.

{{< example file="dialogs/messages/main.go" region="hosted" >}}

Closing the dialog window, with its close button or the Escape key, counts as
dismissing the dialog. On mobile and in the browser an app has only one window,
so the setting is ignored there and the dialog is drawn inside the window as
usual.

## The complete example

{{< example "dialogs/messages/main.go" >}}
