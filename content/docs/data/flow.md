---
title: Reactive state with flow
description: Hold state in observable Go values, derive values from them, and connect both to widgets.
weight: 3
---

The [`flow`](https://pkg.go.dev/github.com/belluikit/bellui/flow) package
holds state as typed values that tell listeners when they change. It is plain
Go with generics and imports nothing from Bellui, so you can keep your state
in a package that has no UI code and test it on its own.

## Values and derived values

{{< example file="data/flow/main.go" region="state" >}}

- [`flow.New`](https://pkg.go.dev/github.com/belluikit/bellui/flow#New)
  creates a
  [`flow.Value`](https://pkg.go.dev/github.com/belluikit/bellui/flow#Value)
  with a starting value. The type is inferred from it and must be comparable.
  `Get` reads it and `Set` changes it.
- [`flow.Derive`](https://pkg.go.dev/github.com/belluikit/bellui/flow#Derive)
  creates a
  [`flow.Derived`](https://pkg.go.dev/github.com/belluikit/bellui/flow#Derived)
  value. It runs the function once straight away, then again each time one of
  the listed sources changes. A derived value can itself be a source for
  another one.

List every value the function reads. `Derive` does not discover them by
itself, and a source you leave out will not trigger a recompute.

`Set` does nothing when the new value equals the current one, and a derived
value only notifies its listeners when the result actually changes. For a
read-modify-write, use `Modify`, which does both steps under one lock:

{{< example file="data/flow/main.go" region="format" >}}

## Connect values to widgets

The `widget` package has helpers that connect a flow value to a widget:

{{< example file="data/flow/main.go" region="widgets" >}}

| Helper | Widget | Direction |
| --- | --- | --- |
| [`widget.BindEntry`](https://pkg.go.dev/github.com/belluikit/bellui/widget#BindEntry) | Entry and `Value[string]` | both ways |
| [`widget.BindCheck`](https://pkg.go.dev/github.com/belluikit/bellui/widget#BindCheck) | Check and `Value[bool]` | both ways |
| [`widget.BindSlider`](https://pkg.go.dev/github.com/belluikit/bellui/widget#BindSlider) | Slider and `Value[float64]` | both ways |
| [`widget.BindProgressBar`](https://pkg.go.dev/github.com/belluikit/bellui/widget#BindProgressBar) | ProgressBar from `Value[float64]` | value to widget |
| [`widget.BindLabel`](https://pkg.go.dev/github.com/belluikit/bellui/widget#BindLabel) | Label from a `StringSource` | value to widget |

`BindLabel` takes a
[`widget.StringSource`](https://pkg.go.dev/github.com/belluikit/bellui/widget#StringSource),
not a flow value. Wrap the value to get one:

- `widget.StringValue{Value: v}` for a `*flow.Value[string]`,
- `widget.StringDerived{Derived: d}` for a `*flow.Derived[string]`,
- `widget.FormatValue(v, format)` for any `*flow.Value`, formatted with
  `fmt.Sprintf`.

Each helper sets the widget from the value straight away and returns a
function that disconnects them. The two-way helpers install their own
`OnChanged` on the widget. They keep the `OnChanged` you set before calling
them and call it too, and restore it when you disconnect.

## Goroutines

`flow` values are safe to read and set from any goroutine. Listeners,
however, run on the goroutine that called `Set`, straight away. When a value
is connected to widgets, its listeners change those widgets, so set it on
the UI goroutine. From background work, wrap the change in
[`bellui.Do`](https://pkg.go.dev/github.com/belluikit/bellui#Do):

{{< example file="data/flow/main.go" region="background" >}}

Data binding does this for you; see [which one to use](../#which-one-to-use).

## Listen to changes

`OnChange` registers a function to call after each change and returns a
function that unregisters it:

{{< example file="data/flow/main.go" region="listen" >}}

Unlike a binding listener, it does not run when you register it.

`Value` and `Derived` also have `Listen(func(old, new T))`. In this version
both arguments receive the new value. If you need the previous value, keep a
copy yourself in an `OnChange` listener.

A derived value keeps listening to its sources until you call its `Close`
method. Call it when a derived value is no longer needed but its sources live
on.

The full program:

{{< example "data/flow/main.go" >}}
