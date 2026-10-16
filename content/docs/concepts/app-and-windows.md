---
title: The app and its windows
description: Create the app, open and size windows, run the event loop, and control what happens when windows close.
weight: 1
---

Every Bellui program has one app. The app owns the windows, the settings and
the event loop. You create it once at the start of `main`.

## Create the app

{{< example file="concepts/windows/main.go" region="app" >}}

- [`app.NewWithID`](https://pkg.go.dev/github.com/belluikit/bellui/app#NewWithID)
  creates the app with a unique ID, written like a reverse domain name. The
  [preferences](https://pkg.go.dev/github.com/belluikit/bellui#App) API needs
  this ID to know where to store data.
- [`app.New`](https://pkg.go.dev/github.com/belluikit/bellui/app#New) creates
  the app without an ID. It takes the ID from the app's `manifest.pxf`
  metadata, if there is one. Prefer `NewWithID`.
- `a.NewWindow(title)` creates a window. It is not shown yet.
- `w.Resize` asks for a content size. The platform may not honour it
  exactly.

The app also gives you access to its
[`Settings`](https://pkg.go.dev/github.com/belluikit/bellui#Settings),
[`Preferences`](https://pkg.go.dev/github.com/belluikit/bellui#Preferences),
[`Storage`](https://pkg.go.dev/github.com/belluikit/bellui#Storage),
[`Clipboard`](https://pkg.go.dev/github.com/belluikit/bellui#Clipboard) and
[`Lifecycle`](https://pkg.go.dev/github.com/belluikit/bellui#Lifecycle) hooks.
`bellui.CurrentApp()` returns the app from anywhere in your program.

## Show a window and run

The complete program puts content in the window with `w.SetContent` and then
calls [`bellui.ShowAndRun(w)`](https://pkg.go.dev/github.com/belluikit/bellui#ShowAndRun):

{{< example "concepts/windows/main.go" >}}

`bellui.ShowAndRun(w)` is the same as calling `w.Show()` and then
`a.Run()`. `Run` starts the event loop and blocks until the app quits, so it
goes at the end of `main`. Show any other window with its own `w.Show()`.

A [`Window`](https://pkg.go.dev/github.com/belluikit/bellui#Window) has one
content object. To show several things, put them in a
[container](/docs/layout/) and set the container as the content. Other
useful window methods:

| Method | What it does |
| --- | --- |
| `SetTitle`, `Title` | Change or read the title bar text |
| `SetFixedSize(true)` | Hint that the user should not resize the window |
| `CenterOnScreen()` | Centre the window on its current monitor |
| `SetFullScreen(true)` | Go full screen |
| `SetPadded(false)` | Remove the inner padding around the content |
| `SetIcon` | Set the window icon (defaults to the app icon) |
| `SetMainMenu` | Add a menu bar |
| `Hide()` | Hide the window without closing it |
| `Canvas()` | The window's canvas, for keyboard handlers and shortcuts |

## Multiple windows

Call `a.NewWindow` again for each extra window. How the extra windows appear
depends on the platform:

- **Desktop:** every window is a separate operating system window. Any number
  can be open at once.
- **Web:** the first window fills the browser canvas. Extra windows appear as
  draggable windows inside it.
- **Mobile:** windows form a stack and only the top one is visible. Opening a
  window covers the previous one, and closing it shows the previous one again.

In the example, the button opens a new window each time it is tapped:

{{< example file="concepts/windows/main.go" region="more" >}}

`SetOnClosed` runs a function when that window closes.

## When the app quits

On desktop, the app quits when its last window closes. You can change that in
two ways.

**Make one window the master.** Closing a master window quits the app, even
if other windows are still open. `SetMaster` is on the optional
[`bellui.MasterWindow`](https://pkg.go.dev/github.com/belluikit/bellui#MasterWindow)
interface, so check for it:

{{< example file="concepts/windows/main.go" region="master" >}}

**Intercept the close.** With
[`SetCloseIntercept`](https://pkg.go.dev/github.com/belluikit/bellui#Window),
the close button runs your function instead of closing the window. The window
only closes when you call `w.Close()`. Use it to ask about unsaved work:

{{< example file="concepts/windows/main.go" region="intercept" >}}

`a.Quit()` closes every window and stops the app. On mobile it does nothing:
the operating system manages the app's lifetime there.

{{< callout type="info" >}}
`Hide()` only hides a window. It does not close it, and it does not make the
app quit.
{{< /callout >}}
