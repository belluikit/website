---
title: Right-to-left layout
description: How Bellui mirrors layouts for Arabic, Hebrew and other right-to-left languages, and how to force or scope the direction.
weight: 7
---

In a right-to-left (RTL) language such as Arabic or Hebrew, the whole
interface reads from right to left. Bellui mirrors the layout for you:

- Horizontal containers place their first item on the right.
- Text aligned to the leading edge sits on the right.
- Paragraphs take a right-to-left base direction, so punctuation and mixed
  Latin text land in the right place.
- Keyboard focus moves through each row from right to left.
- Directional icons, such as the back and next arrows, are drawn mirrored.

By default the direction comes from the app's locale. A locale whose script
is written right to left, such as Arabic, Hebrew, Syriac or Thaana, gives a
right-to-left layout; everything else is left to right.
[`lang.SetLocale`](https://pkg.go.dev/github.com/belluikit/bellui/lang#SetLocale)
changes the locale, and with it the direction. See
[Localisation](/docs/localisation/).

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
focus order and the base direction of text. Scopes can be nested, and the
innermost one wins. `DirectionAuto` makes a scope inherit the direction around
it. To change a scope's direction later, call its `SetDirection` method.

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
as punctuation. Arabic letter joining and Hebrew letter forms come from the
text itself and do not depend on it.

The full program:

{{< example "theming/rtl/main.go" >}}
