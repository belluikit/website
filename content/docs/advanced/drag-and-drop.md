---
title: Drag and drop
description: Drag widgets within your app, drag text and files out to other apps, and accept files dropped from the desktop.
weight: 7
---

Bellui supports three kinds of drag and drop. Each one is a set of small
interfaces that your widget implements:

| You want to | Source implements | Target implements |
| --- | --- | --- |
| move something within your app | `event.Draggable` plus `core.DragPayloadProvider` and/or `core.DragImageProvider` | `core.DragTarget` |
| drag text, files or an image out to another app | `event.Draggable`, calling `desktop.StartDrag` | the other app |
| accept files dragged in from the desktop | the file manager | `bellui.Droppable`, optionally `bellui.DropHoverable` |

The example at the end of the page shows all three in one window.

## Within your app

A drag source is a widget that implements
[`event.Draggable`](https://pkg.go.dev/github.com/belluikit/bellui/event#Draggable).
To take part in drag and drop, it also implements one or both of these:

- [`core.DragPayloadProvider`](https://pkg.go.dev/github.com/belluikit/bellui/core#DragPayloadProvider):
  `DragPayload()` returns the value the target receives, of any type.
- [`core.DragImageProvider`](https://pkg.go.dev/github.com/belluikit/bellui/core#DragImageProvider):
  `DragImage()` returns an object that Bellui draws under the pointer during
  the drag (a "ghost"). Return nil for none.

Both are read once, when the drag starts.

{{< example file="advanced/dragdrop/main.go" region="card" >}}

A drop target implements
[`core.DragTarget`](https://pkg.go.dev/github.com/belluikit/bellui/core#DragTarget).
Bellui finds the deepest target under the pointer and calls it as the drag
moves:

| Method | When |
| --- | --- |
| `DragEnter` | the pointer moves onto the target |
| `DragOver` | the pointer moves while over it |
| `DragLeave` | the pointer moves off it, or the drag ends elsewhere |
| `Drop` | the user releases over it |

Each call except `DragLeave` gets a
[`DragTargetEvent`](https://pkg.go.dev/github.com/belluikit/bellui/core#DragTargetEvent)
with the payload and the pointer position. `Position` is relative to the
target, and `AbsolutePosition` to the window. Use `DragEnter` and `DragLeave` to highlight the target:

{{< example file="advanced/dragdrop/main.go" region="bin" >}}

Only drags whose source provides a payload or an image reach drop targets. A
slider or a scroll bar is also `Draggable`, and dragging it never triggers
your targets.

## Out to other apps

[`desktop.StartDrag`](https://pkg.go.dev/github.com/belluikit/bellui/driver/desktop#StartDrag)
hands the current drag to the operating system, so the user can drop into a
file manager, a text editor or a browser. Call it from your widget's `Dragged`
method:

{{< example file="advanced/dragdrop/main.go" region="export" >}}

[`DragData`](https://pkg.go.dev/github.com/belluikit/bellui/driver/desktop#DragData)
holds what is dragged. Set any of:

- `Text`: plain text.
- `URIs`: files. Only `file://` URIs are carried; others are skipped.
- `Image`: an image. It is also the picture that follows the pointer.
- `DragImageObject`: a Bellui object to draw as that picture instead, when
  `Image` is nil. It is only a picture and is never dropped as data.
- `OnCompleted`: called on the UI thread when the drag ends, with whether a
  target accepted the drop. A move can delete the original here.

`StartDrag` returns false when the platform cannot start a drag. Your widget
then keeps its in-app behavior. Once it returns true, the operating system owns
the gesture and Bellui ends its own drag, calling `DragEnd`.

| Platform | Drag out |
| --- | --- |
| macOS | yes |
| Windows | yes |
| Linux, X11 | yes |
| Linux, Wayland | yes, but `OnCompleted` is never called and no drag picture is shown |
| BSDs (X11) | yes |
| Web | yes, through `DragOutProvider` (below) |
| iOS, Android | no; `StartDrag` returns false |

### On the web

In a browser, `Dragged` is not called once the browser's own drag has begun,
so `StartDrag` from `Dragged` cannot work there. Implement
[`desktop.DragOutProvider`](https://pkg.go.dev/github.com/belluikit/bellui/driver/desktop#DragOutProvider)
as well. Its `DragOutData()` returns the data and true, or false to keep an
ordinary in-app drag. The web build asks for it when the pointer goes down on
your widget. The example's `export` widget implements both, so it works on the
desktop and on the web.

## Files from the desktop

A widget that implements
[`bellui.Droppable`](https://pkg.go.dev/github.com/belluikit/bellui#Droppable)
receives files the user drops on it:

{{< example file="advanced/dragdrop/main.go" region="files" >}}

The [`DropEvent`](https://pkg.go.dev/github.com/belluikit/bellui#DropEvent)
has the drop position, relative to the widget and to the window, and the items
as `URIs`. On the desktop these are file URIs.

Bellui gives the drop to the deepest `Droppable` under the pointer. A drop that
no widget claims goes to the window's callback, if you set one:

{{< example file="advanced/dragdrop/main.go" region="window" >}}

To highlight a widget while files are dragged over it, implement
[`bellui.DropHoverable`](https://pkg.go.dev/github.com/belluikit/bellui#DropHoverable)
too. It adds `DragEnter`, `DragOver` and `DragLeave`, each with only the
pointer position: the operating system does not reveal the files until the
drop. These hover calls arrive on macOS, Windows and X11. Elsewhere the widget
still gets `Dropped`, just without the hover calls.

{{< callout type="info" >}}
One type cannot be both a `bellui.DropHoverable` and a `core.DragTarget`.
Both interfaces have methods named `DragEnter` and `DragOver`, with different
argument types. Use two widgets, for example one wrapped in the other, when an
area should accept both files and in-app drags.
{{< /callout >}}

## The whole example

{{< example "advanced/dragdrop/main.go" >}}
