---
title: Building a widget
description: Write a widget from scratch with widget.BaseWidget and a renderer - state in the widget, drawing in the renderer.
weight: 1
---

A widget has two parts:

- The **widget** holds state and behaviour: its fields, its methods, and the
  [event methods](/docs/concepts/input-events/) it implements. It is what
  your app code creates and talks to.
- The **renderer** draws it. It owns a set of canvas objects and keeps them
  in step with the widget's state.

This page builds a rating widget: five dots that fill up to the current
value, which the user sets by tapping.

## The widget

{{< example file="customwidgets/rating/main.go" region="widget" >}}

Step by step:

1. **Embed [`widget.BaseWidget`](https://pkg.go.dev/github.com/belluikit/bellui/widget#BaseWidget).**
   It provides the [`CanvasObject`](https://pkg.go.dev/github.com/belluikit/bellui/core#CanvasObject)
   methods: size, position, visibility, `MinSize` and `Refresh`. The widget
   only has to add `CreateRenderer`.
2. **Call `ExtendBaseWidget(r)` in the constructor**, passing the widget
   itself. This tells `BaseWidget` which object it is part of, so that it
   can find the renderer and call your methods. Without it the widget does
   not render, and Bellui logs a warning that names the call site.
3. **Call `Refresh()` after changing state.** `SetValue` changes `Value`
   and then calls `Refresh`, which asks the renderer to update. Give your
   widget setter methods like this, so callers do not have to remember.
4. **Implement event interfaces** for the input you want. `Tapped` makes the
   widget [`Tappable`](https://pkg.go.dev/github.com/belluikit/bellui/event#Tappable).
   The event's `Position` is relative to the widget, so the code works out
   which dot was tapped from the x coordinate.
5. **`CreateRenderer`** creates the canvas objects and returns the renderer.
   Bellui calls it when the widget is first needed and caches the result. A
   cached renderer can be destroyed later, for example after its window
   closes or once the widget has been out of use for a while, and a new one
   created when it is needed again.

## The renderer

A renderer implements
[`core.WidgetRenderer`](https://pkg.go.dev/github.com/belluikit/bellui/core#WidgetRenderer):

{{< example file="customwidgets/rating/main.go" region="renderer" >}}

| Method | Called | Job |
| --- | --- | --- |
| `MinSize()` | When a layout measures the widget | Return the smallest size the widget can be drawn at. `BaseWidget.MinSize` returns this. |
| `Layout(size)` | When the widget is resized | Resize and move the objects to fit `size`. |
| `Refresh()` | When the widget's `Refresh` is called, and when the theme changes | Copy state from the widget onto the objects, then refresh them. |
| `Objects()` | When drawing and when finding what is under the pointer | Return the objects to draw, back to front. |
| `Destroy()` | When the renderer is no longer needed | Stop timers or release resources. Often empty. |

Some rules that keep a renderer correct:

- **Read colours and sizes from the theme in `Refresh`**, not once at
  creation. `rr.rating.Theme()` returns the theme that applies to this
  widget, which can differ inside a theme override. Because `Refresh` runs
  on theme changes, the widget then follows a switch between light and dark.
- **Do not create objects in `Refresh` or `Layout`.** Create them in
  `CreateRenderer` and change their fields afterwards. If the set of objects
  changes, update the slice that `Objects()` returns.
- **Keep state in the widget, not the renderer.** The renderer reads
  `rr.rating.Value`; it never stores a copy that could go stale. A renderer
  can be thrown away and recreated, and anything stored only there is lost.

For a widget drawn by a single object, you do not need your own renderer
type. [`widget.NewSimpleRenderer(obj)`](https://pkg.go.dev/github.com/belluikit/bellui/widget#NewSimpleRenderer)
returns one that draws `obj` at the widget's full size. The pad on the
[Input events](/docs/concepts/input-events/) page uses it.

## Use it

The widget behaves like any other. It goes in containers and reports its
minimum size to their layouts:

{{< example file="customwidgets/rating/main.go" region="use" >}}

## The whole program

{{< example "customwidgets/rating/main.go" >}}

## More you can add

- **State sync:** if a widget has internal state derived from its public
  fields, implement
  [`widget.StateSyncer`](https://pkg.go.dev/github.com/belluikit/bellui/widget#StateSyncer).
  `BaseWidget.Refresh` calls its `SyncState()` before the renderer's
  `Refresh`.
- **Disabling:** embed
  [`widget.DisableableWidget`](https://pkg.go.dev/github.com/belluikit/bellui/widget#DisableableWidget)
  instead of `BaseWidget` to get `Enable`, `Disable` and `Disabled`. Check
  `Disabled()` in your event methods. While the widget is disabled, Bellui
  draws it and everything inside it at 38% opacity, on every platform, so it
  looks disabled with no work from you. If your renderer draws a disabled look
  of its own, such as the theme's disabled colours, implement
  [`core.DisabledLook`](https://pkg.go.dev/github.com/belluikit/bellui/core#DisabledLook)
  with `DrawsDisabledLook()` returning true, so it is not dimmed a second
  time. A widget inside a disabled widget that has no look of its own is
  dimmed with it, whatever it reports.
- **Keyboard:** implement [`event.Focusable`](https://pkg.go.dev/github.com/belluikit/bellui/event#Focusable)
  to receive keys when focused.
- **Tooltips:** `BaseWidget.SetToolTipText` gives the widget a hover tooltip.
- **Touch drags:** a `Draggable` widget can take touch drags in some
  directions only, and keep moving after a fling. See
  [Drags on touch screens](/docs/concepts/input-events/#drags-on-touch-screens).
- **Direction:** in an event handler, ask
  [`container.DirectionOf`](/docs/theming/right-to-left/#ask-which-direction-applies)
  whether the widget sits in a right-to-left layout.
- **Theme overrides:** if the widget creates objects after it is first
  shown, such as rows as it grows, pass them to
  [`theme.MatchScope`](/docs/theming/custom/#apply-a-theme-to-part-of-the-window)
  so they follow a theme override around the widget.
- **Translucency:** implement
  [`core.OpacityGroup`](https://pkg.go.dev/github.com/belluikit/bellui/core#OpacityGroup)
  to draw the whole widget at an opacity of its own.
- **Work out of sight:** a widget that updates itself on a timer can skip
  the work while [`bellui.IsOnScreen`](/docs/concepts/canvas-objects/#where-an-object-is-drawn)
  reports it is not drawn.
- **Debugging:** set `widget.DebugLayout = true` to log every `MinSize` and
  `Resize`, or `widget.DebugRefresh = true` to log every `Refresh`.
