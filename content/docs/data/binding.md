---
title: Data binding
description: Create bound values, connect them to widgets, convert between types, and update them from any goroutine.
weight: 1
---

A **binding** is a value that tells its listeners when it changes. Widgets
that are bound to it show its current value. Input widgets also write back
to it when the user changes them.

## Bind a string

{{< example file="data/binding/main.go" region="string" >}}

- [`binding.NewString`](https://pkg.go.dev/github.com/belluikit/bellui/data/binding#NewString)
  creates a string binding that holds its own value. `Set` stores a new value
  and `Get` reads it. Both return an error, which conversions use to report a
  value that cannot be converted.
- [`widget.NewEntryWithData`](https://pkg.go.dev/github.com/belluikit/bellui/widget#NewEntryWithData)
  shows the binding's value and sets it as the user types.
- [`binding.NewSprintf`](https://pkg.go.dev/github.com/belluikit/bellui/data/binding#NewSprintf)
  makes a new string from a format and other bindings. The label follows
  every keystroke.

There are bindings for `Bool`, `Bytes`, `Float`, `Int`, `Rune`, `String`,
`URI` and `Untyped` (any value). For your own types, use
[`binding.NewItem`](https://pkg.go.dev/github.com/belluikit/bellui/data/binding#NewItem)
with a function that says when two values are equal.

## Widgets that accept bindings

| Constructor | Binding |
| --- | --- |
| `widget.NewLabelWithData` | `binding.String` |
| `widget.NewEntryWithData` | `binding.String` |
| `widget.NewSelectWithData` | `binding.String` |
| `widget.NewCheckWithData` | `binding.Bool` |
| `widget.NewSliderWithData` | `binding.Float` |
| `widget.NewProgressBarWithData` | `binding.Float` |
| `widget.NewListWithData` | `binding.DataList` |
| `widget.NewGridWrapWithData` | `binding.DataList` |
| `widget.NewTreeWithData` | `binding.DataTree` |

The label, entry, select, check, slider and progress bar also have a `Bind`
method to connect an existing widget, and `Unbind` to disconnect it. Calling
`Bind` again replaces the previous binding.

## Convert between types

A widget needs a binding of its own type. A label needs a string, for
example, while your value may be a float. Conversion functions wrap one
binding in another:

{{< example file="data/binding/main.go" region="float" >}}

The slider, the progress bar and the label all share `volume`. Dragging the
slider moves the bar and changes the text.

- `FloatToString`, `IntToString` and `BoolToString` convert to text, and
  their `…WithFormat` forms take a `fmt` format. Bind one to an entry to let
  the user type a number. Text that does not parse is not written to the
  source, and the entry shows it as invalid.
- `StringToFloat`, `StringToInt`, `StringToBool` and `StringToURI` go the
  other way: they read a string binding as a number, a boolean or a URI.
- `IntToFloat` and `FloatToInt` convert between numbers.

Conversions work both ways: setting the converted binding parses the value
and sets the source.

Boolean bindings can be combined:

{{< example file="data/binding/main.go" region="bool" >}}

[`binding.Not`](https://pkg.go.dev/github.com/belluikit/bellui/data/binding#Not)
inverts a binding, so the two checks always disagree.
[`binding.And`](https://pkg.go.dev/github.com/belluikit/bellui/data/binding#And)
and [`binding.Or`](https://pkg.go.dev/github.com/belluikit/bellui/data/binding#Or)
combine several. Setting an `And` or `Or` binding sets every binding it was
made from to the same value.

## Update from a background goroutine

You can call `Set` from any goroutine. Bellui runs the listeners, including
the widget updates, on the UI goroutine:

{{< example file="data/binding/main.go" region="background" >}}

This is the main reason to prefer a binding over calling `SetText` from a
goroutine, which would need `bellui.Do`.

## Listen yourself

To run your own code when a binding changes, add a listener.
[`binding.NewDataListener`](https://pkg.go.dev/github.com/belluikit/bellui/data/binding#NewDataListener)
turns a function into one:

{{< example file="data/binding/main.go" region="listen" >}}

The listener runs once as soon as it is added, with the current value, and
again after each change. It always runs on the UI goroutine, so it can update
widgets. Keep the listener value if you want to remove it later with
`RemoveListener`. Adding the same listener twice has no effect.

A binding only notifies when the value actually changes. Setting the value it
already has does nothing.

The full program:

{{< example "data/binding/main.go" >}}
