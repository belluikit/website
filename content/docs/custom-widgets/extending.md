---
title: Extending a widget
description: Embed an existing widget to add event handling or change how it reacts to input, keeping everything else it does.
weight: 2
---

To change one thing about a standard widget, embed it in your own struct
and add or replace methods. Your type keeps the widget's look, its fields and
its methods, and adds whatever you define.

## Add a behaviour: a tappable icon

[`widget.Icon`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Icon)
shows an image but does not respond to taps. Embedding it and adding a
`Tapped` method makes it [`Tappable`](https://pkg.go.dev/github.com/belluikit/bellui/event#Tappable):

{{< example file="customwidgets/extend/main.go" region="icon" >}}

The important line is `icon.ExtendBaseWidget(icon)`. Call it with your outer
type, before anything else uses the widget, and in place of the embedded
widget's constructor (`widget.NewIcon` here). Bellui finds the event methods
by looking at the object it was given in `ExtendBaseWidget`. If it only knew
about the inner `widget.Icon`, it would never see `Tapped`.

That is also why the constructor calls `SetResource` after
`ExtendBaseWidget`, rather than calling `widget.NewIcon`: `NewIcon` would
register the inner icon, and the outer type's methods would be ignored.

## Change a behaviour: an entry that takes only digits

Define a method with the same name as one of the widget's, and yours is
called instead. To keep the original behaviour for some input, call the
embedded widget's method explicitly:

{{< example file="customwidgets/extend/main.go" region="entry" >}}

[`widget.Entry`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Entry)
receives typed characters through `TypedRune`. The replacement passes digits
on to `e.Entry.TypedRune` and drops everything else.

{{< callout type="warning" >}}
Overriding `TypedRune` filters typing only. Text can still arrive by other
routes, such as pasting (a shortcut, handled by `TypedShortcut`) or your own
code calling `SetText`. If the content must be valid, also check it where
you read it, or set the entry's `Validator`.
{{< /callout >}}

## Which methods you can replace

Go has no virtual methods. Your method is used when **Bellui** calls it
through an interface on your object: the event methods (`Tapped`,
`TypedRune`, `TypedKey`, `Dragged` and the others in
[Input events](/docs/concepts/input-events/)), and the `core.Widget` methods
such as `CreateRenderer` and `MinSize`.

Your method is **not** used when the embedded widget calls a method on
itself. Inside `widget.Entry`, a call to `e.TypedRune` means the Entry's own
method, not yours. So replacing a method changes how the widget reacts to
the toolkit, not how it works internally.

When the change you want is in how a widget looks or lays itself out, an
extension is usually the wrong tool. Wrap the widget in a
[container](/docs/layout/) or [build a new widget](../building/) that uses it.

## The whole program

{{< example "customwidgets/extend/main.go" >}}
