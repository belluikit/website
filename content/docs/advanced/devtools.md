---
title: Developer tools
description: See widget bounds, print the widget tree, trace refreshes, and inspect a running window.
weight: 4
---

Two packages help you find out why a layout looks the way it does:

- [`dev`](https://pkg.go.dev/github.com/belluikit/bellui/dev) adds keyboard
  shortcuts to a window: an outline overlay, a tree dump and refresh tracing.
  It also has a file watcher that rebuilds your app when the source changes.
- [`dev/inspector`](https://pkg.go.dev/github.com/belluikit/bellui/dev/inspector)
  opens a second window showing the live widget tree, with details for the
  selected object and a picker.

Both are for development on the desktop.

## Keyboard shortcuts

{{< example file="advanced/devtools/main.go" region="enable" >}}

After [`dev.Enable(w)`](https://pkg.go.dev/github.com/belluikit/bellui/dev#Enable):

| Key | What it does |
| --- | --- |
| F12 | Shows or hides an overlay that outlines every object in the window and labels it with its Go type. Each nesting level gets its own color. |
| F11 | Prints the widget tree to the log: each object's type, size, position and minimum size. |
| F10 | Turns tracing on or off. While it is on, every widget refresh, every minimum-size calculation and every resize is logged. |

The F11 output is the same text
[`widget.InspectLayout`](https://pkg.go.dev/github.com/belluikit/bellui/widget#InspectLayout)
returns, so you can also print it from code for any object. F10 flips the
`widget.DebugRefresh` and `widget.DebugLayout` variables, which you can also
set yourself.

Things to know:

- **The keys only work when no widget has keyboard focus.** While a text field
  or another focusable widget has focus, it receives the keys instead. Click an
  empty part of the window first.
- **The overlay is a snapshot.** It shows the layout at the moment you pressed
  F12. Press F12 twice to redraw it after the layout changes.
- `Enable` uses the canvas key handler (`SetOnTypedKey`). It keeps a handler
  you set before calling it and passes other keys on to it. A handler you set
  afterwards replaces the shortcuts.

## The inspector

{{< example file="advanced/devtools/main.go" region="inspector" >}}

[`inspector.Show(w)`](https://pkg.go.dev/github.com/belluikit/bellui/dev/inspector#Show)
opens a second window titled "Bellui Inspector". The example opens
it only when the environment variable is set, so you can leave the code in:

```sh
BELLUI_INSPECT=1 go run .
```

The inspector window has:

- **A tree** of everything in the target window, including the objects inside
  each widget. It does not update by itself; press **Refresh** after the UI
  changes.
- **Details** of the selected object: position, size and minimum size, its
  bounds in the window, whether it is visible, its clip area, a scroll
  container's offset, and the role, label and value it reports to screen
  readers.
- **A highlight** drawn around the selected object in your app's window.
- **Pick** mode: press **Pick**, move the pointer over your app to preview the
  object under it, and click to select it in the tree. Escape cancels.

The inspector is read-only: it does not change your widgets. The highlight and
the picker are overlays, and closing the inspector removes them.

The inspector needs a second window, so it works on desktop platforms only.
Call `Show` on the UI thread, for example before `bellui.ShowAndRun`.

## Rebuild on save

[`dev.Watch(dir, onChange)`](https://pkg.go.dev/github.com/belluikit/bellui/dev#Watch)
checks the `.go` files under `dir` twice a second and calls `onChange` when one
is added, changed or removed. It skips `.git`, `vendor`, `node_modules` and
`testdata` folders. It runs in its own goroutine and returns at once.

[`dev.WatchAndRun(watchDir, pkgPath)`](https://pkg.go.dev/github.com/belluikit/bellui/dev#WatchAndRun)
goes further. It builds the package at `pkgPath` with `go build`, starts it,
and on every change stops it, rebuilds and starts it again. Build errors are
printed and the old process stays stopped until the next change. It never
returns, so call it from a small helper program of its own rather than from
your app.
