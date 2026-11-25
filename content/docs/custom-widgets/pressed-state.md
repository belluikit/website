---
title: Pressed state
description: Give a widget of your own the MD3 ripple that Bellui's widgets show when pressed.
weight: 3
---

Bellui's tappable widgets, such as `Hyperlink`, `Chip`, tree rows and tabs,
show the MD3 pressed state: a ripple that spreads from the press and fades on
release. A widget you build can show the same ripple, with the same timing
and colours, by embedding
[`widget.PressTarget`](https://pkg.go.dev/github.com/belluikit/bellui/widget#PressTarget).

## The widget

{{< example file="customwidgets/press/main.go" region="widget" >}}

Embedding `PressTarget` gives the widget the touch and mouse press events.
You add two calls in the widget's own event methods:

- In `Tapped`, call
  [`widget.TapPress`](https://pkg.go.dev/github.com/belluikit/bellui/widget#TapPress).
  A tap with no press before it, from a key, an accessibility action or your
  own code, then shows a press too. After a pointer press it does nothing.
- In `HoverOut`, call
  [`widget.CancelPress`](https://pkg.go.dev/github.com/belluikit/bellui/widget#CancelPress).
  A mouse released outside the widget sends it no `MouseUp`, so the press
  would otherwise stay. Call it as well when the widget stops taking presses
  in the middle of one.

If your widget defines `TouchDown` or `MouseDown` itself, call the
`PressTarget`'s method from yours.

## The renderer

{{< example file="customwidgets/press/main.go" region="renderer" >}}

Call [`widget.AttachPress`](https://pkg.go.dev/github.com/belluikit/bellui/widget#AttachPress)
from `CreateRenderer`. It returns the ripple to draw: put it above the
widget's background and below its content, and resize it in `Layout` like
the background.

The function you pass is asked at each press. It returns a
[`PressShape`](https://pkg.go.dev/github.com/belluikit/bellui/widget#PressShape):
where the state layer shows, in the widget's coordinates, its corner radius,
and `Content`, the colour of the widget's text or icon. The layer is that
colour at the theme's pressed opacity. Return false to refuse the press, for
example while the widget is disabled.

## The whole program

{{< example "customwidgets/press/main.go" >}}
