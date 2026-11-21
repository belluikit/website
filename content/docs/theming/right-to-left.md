---
title: Right-to-left layout
description: Force the layout direction, keep part of a window in one direction, and support right-to-left in your own layouts and text styles.
weight: 7
---

This page covers the API that controls layout direction: forcing it,
scoping it to part of a window, and supporting it in your own layouts and
text styles.

If you are translating an app into Arabic, Hebrew, Persian or Urdu, start
with [Right-to-left languages](/docs/localisation/right-to-left/) in the
Localisation section. It covers what Bellui mirrors and what it does not,
fonts, and the limits of text direction.

By default the direction comes from the app's locale. A locale whose script
is written right to left, such as Arabic or Hebrew, gives a right-to-left
layout; everything else is left to right. Horizontal containers then place
their first item on the right, leading-aligned text sits on the right, and
directional icons are drawn mirrored.

## Force a direction

[`bellui.SetLayoutDirection`](https://pkg.go.dev/github.com/belluikit/bellui#SetLayoutDirection)
overrides the direction for the whole app, whatever the locale:

{{< example file="theming/rtl/main.go" region="global" >}}

- `bellui.DirectionLTR` and `bellui.DirectionRTL` force a direction.
- `bellui.DirectionAuto`, the default, follows the locale again.

When the direction changes while the app runs, every window is laid out and
painted again, as for a theme change. Call it on the UI goroutine, for
example from a widget callback.
[`bellui.CurrentLayoutDirection`](https://pkg.go.dev/github.com/belluikit/bellui#CurrentLayoutDirection)
returns the setting, which may be `DirectionAuto`.
[`lang.IsRTLDirection`](https://pkg.go.dev/github.com/belluikit/bellui/lang#IsRTLDirection)
returns the direction actually in effect, as a boolean.

Forcing RTL in an English app is a quick way to test that your layout
mirrors well before you have any translations.

## Keep part of the window in one direction

Some content must not mirror: source code, a phone number field, a media
timeline. Wrap it in
[`container.NewDirectionScope`](https://pkg.go.dev/github.com/belluikit/bellui/container#NewDirectionScope):

{{< example file="theming/rtl/main.go" region="scope" >}}

The scope sets the direction for everything inside it: layout, alignment,
focus order, input such as swipes and arrow keys, and the base direction of
text. Pop-ups that a widget inside the scope opens, such as a `Select`'s menu
or a date picker, take the scope's direction too. The scope keeps its
direction when the content is laid out again, for example after a theme
change. Scopes can be nested, and the innermost one wins. `DirectionAuto`
makes a scope inherit the direction around it. To change a scope's direction
later, call its `SetDirection` method.

A pop-up you show yourself with `widget.PopUp` keeps its window's direction.
Wrap its content in a `DirectionScope` when it needs another.

A single window can also have its own direction. Its canvas has an optional
`SetDirection(bellui.Direction)` method, which you reach with a type
assertion on `w.Canvas()`.

## Ask which direction applies

[`container.DirectionOf(obj)`](https://pkg.go.dev/github.com/belluikit/bellui/container#DirectionOf)
returns the direction that applies where an object sits: its innermost
`DirectionScope`, else its window's direction, else the app's. The answer is
always `bellui.DirectionLTR` or `bellui.DirectionRTL`.

`lang.IsRTLDirection` gives the same answer only while a layout or a paint
is running. A widget's renderer can use it in `Layout` and `Refresh`. Code
that runs outside those, such as an event handler, calls `DirectionOf`:

{{< example file="theming/rtl/main.go" region="directionof" >}}

## Custom layouts

The standard layouts mirror themselves. If you write your own
`core.Layout`, position the children left to right as usual, then, when
`lang.IsRTLDirection()` is true, call
[`layout.MirrorChildren`](https://pkg.go.dev/github.com/belluikit/bellui/layout#MirrorChildren)
with the children and the container's width. It reflects each visible child
horizontally and leaves vertical positions and sizes alone.

## Text direction

The direction of a piece of text is set by `Direction` in its
[`bellui.TextStyle`](https://pkg.go.dev/github.com/belluikit/bellui#TextStyle).
Left at its default, it follows the layout direction. Set it to
`bellui.TextDirectionLTR` or `TextDirectionRTL` for text whose direction
differs from the interface, such as a Hebrew quotation in an English
document.

The direction only sets the base order of runs and neutral characters such
as punctuation. Arabic letter joining comes from the text itself and does
not depend on it. `Language` in the same style overrides the language used
for shaping; left empty, it follows the locale. For the limits of mixed
right-to-left and left-to-right text, see
[Mixed Arabic and Latin text](/docs/localisation/right-to-left/#mixed-arabic-and-latin-text).

The full program:

{{< example "theming/rtl/main.go" >}}
