---
title: Lists, structs and preferences
description: Bind lists to list widgets, and bind your own variables, struct fields and saved preferences.
weight: 2
---

## Bound lists

A list binding holds a slice and notifies its listeners when items are
added, removed or changed. Pass it to
[`widget.NewListWithData`](https://pkg.go.dev/github.com/belluikit/bellui/widget#NewListWithData):

{{< example file="data/lists/main.go" region="list" >}}

The list asks for a row template once per visible row, then calls the
update function with the binding for one item and the row to fill. Rows are
reused as the list scrolls. Calling `Bind` on the row's label replaces
whatever it was bound to before, so a reused row always shows the right
item.

Change the list through the binding, and the widget follows:

{{< example file="data/lists/main.go" region="edit" >}}

A list binding has `Append`, `Prepend`, `Remove`, `Set` (replace all items),
`SetValue` (replace one item), `Get`, `GetValue` and `Length`. `Remove`
takes a value, not an index, and removes one item equal to it.

There are list bindings for every value type (`NewStringList`,
`NewFloatList`, `NewIntList` and so on), and
[`binding.NewList`](https://pkg.go.dev/github.com/belluikit/bellui/data/binding#NewList)
for any type. Trees work the same way with `NewStringTree` and
`widget.NewTreeWithData`, and `widget.NewGridWrapWithData` shows a list as a
wrapping grid.

## Bind to your own variable

The `Bind…` functions bind to a variable you already have instead of
holding their own value:

{{< example file="data/external/main.go" region="variable" >}}

Setting the binding writes to the variable. The reverse is not automatic:
when your code changes the variable directly, call `Reload` so the binding
notices and tells its listeners. `BindString`, `BindFloat`, `BindBool` and
the other `Bind…` functions, as well as `BindStringList` and the other list
forms, all work this way.

## Bind struct fields

[`binding.BindStruct`](https://pkg.go.dev/github.com/belluikit/bellui/data/binding#BindStruct)
takes a pointer to a struct and gives one binding per exported field:

{{< example file="data/external/main.go" region="struct" >}}

{{< example file="data/external/main.go" region="fields" >}}

`GetItem` returns the field's binding as a `binding.DataItem`. Bool, int,
float and string fields come back as `binding.Bool`, `binding.Int`,
`binding.Float` and `binding.String`, so assert to the type you need. Editing
the entries changes `person`. After you change `person` in code, call
`Reload` on the struct binding.

For a `map[string]any`, use `binding.NewUntypedMap` or
`binding.BindUntypedMap`.

## Bind a preference

Preference bindings read their value from the app's preferences and save
every change back:

{{< example file="data/external/main.go" region="prefs" >}}

The value is restored the next time the app starts. Preferences need an app
ID, so create the app with `app.NewWithID`. Bindings for the same key share
one underlying value, and binding one key with two different types logs an
error.

There are `BindPreferenceBool`, `…Float`, `…Int` and `…String`, and list
forms such as `BindPreferenceStringList`.

The full programs:

{{< example "data/lists/main.go" >}}

{{< example "data/external/main.go" >}}
