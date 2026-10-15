---
title: Data and state
description: Keep widgets in step with your data, using data binding or the flow package for reactive state.
weight: 5
---

Most screens show values that change: a name the user types, a download's
progress, a list of messages. You can update widgets by hand in every
callback, but it is easy to miss one. Bellui has two ways to connect data to
widgets so that a change in one place shows up everywhere.

{{< cards >}}
  {{< card link="binding/" title="Data binding" subtitle="Bound values that widgets read and write, with conversions between types." >}}
  {{< card link="external-data/" title="Lists, structs and preferences" subtitle="Bind lists, your own variables, struct fields and saved preferences." >}}
  {{< card link="flow/" title="Reactive state with flow" subtitle="Plain Go values and derived values, connected to widgets with helpers." >}}
{{< /cards >}}

## Which one to use

**Data binding**, in the
[`data/binding`](https://pkg.go.dev/github.com/belluikit/bellui/data/binding)
package, is built into the widgets. Constructors such as
`widget.NewEntryWithData` and `widget.NewListWithData` take a binding
directly. Listeners always run on the UI goroutine, so you can set a binding
from any goroutine. It also covers lists, trees, struct fields and
preferences.

**The [`flow`](https://pkg.go.dev/github.com/belluikit/bellui/flow)
package** holds application state as typed Go values, and computes derived
values from them. It has no Bellui dependencies, so your state code does not
import the UI. Helpers in the `widget` package connect a value to an entry,
label, check, slider or progress bar. Listeners run on the goroutine that
changed the value.

| | Data binding | flow |
| --- | --- | --- |
| Types | Bool, Bytes, Float, Int, Rune, String, URI, any type with `NewItem` | any comparable type |
| Derived values | conversions (`IntToString`, `NewSprintf`, `Not`, `And`, `Or`, …) | `flow.Derive` with any function |
| Lists, trees, maps | yes | no |
| Widget support | `New…WithData` constructors and `Bind` methods | `widget.BindEntry`, `BindLabel`, `BindCheck`, `BindSlider`, `BindProgressBar` |
| Setting from a background goroutine | safe | wrap in `bellui.Do` when widgets listen |
| Errors | `Get` and `Set` return an error | none |

Use **data binding** when a widget shows the value directly, for lists, and
when values change from background work. Use **flow** for state shared
across a screen or the whole app, especially when several values are derived
from it.

The two are not exclusive. A screen can keep its state in flow values and
still use a bound list for its `widget.List`.

For lists and tables that need sorting, filtering or selection that survives
re-sorting, the [`data/model`](https://pkg.go.dev/github.com/belluikit/bellui/data/model)
package and `widget.BindListModel` build on the binding listeners.
