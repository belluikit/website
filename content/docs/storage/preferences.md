---
title: Preferences
description: Save and load typed settings with the generic Pref, PrefOr and SetPref functions.
weight: 2
---

Preferences are small values your app wants to remember between runs: a
user's name, the last tab they opened, whether they turned on dark mode. You
set a value and Bellui saves it for you.

## Getting the preferences

{{< example file="storage/preferences/main.go" region="setup" >}}

[`App.Preferences`](https://pkg.go.dev/github.com/belluikit/bellui#App)
returns the app's
[`bellui.Preferences`](https://pkg.go.dev/github.com/belluikit/bellui#Preferences).
Saving to disk needs an app ID, so create the app with `app.NewWithID`.
Without an ID, values still work while the app runs but are not saved.

## Reading and writing values

You read and write with generic functions in the `bellui` package. The type
parameter is the type you want back.

{{< example file="storage/preferences/main.go" region="read" >}}

| Function | What it does |
| --- | --- |
| [`Pref[T](p, key)`](https://pkg.go.dev/github.com/belluikit/bellui#Pref) | Returns the value and `true`, or the zero value and `false` if the key is missing or holds another type. |
| [`PrefOr(p, key, fallback)`](https://pkg.go.dev/github.com/belluikit/bellui#PrefOr) | Returns the value, or `fallback` if it is missing. The type is taken from `fallback`. |
| [`SetPref(p, key, value)`](https://pkg.go.dev/github.com/belluikit/bellui#SetPref) | Stores the value. |
| `p.RemoveValue(key)` | Deletes the key. |

The types you can store are `bool`, `int`, `float64` and `string`, and slices
of those. The compiler rejects anything else. Numbers are converted for you
where needed: an `int` saved earlier can be read back as an `int` even though
the stored form may not keep the difference between whole numbers and
decimals.

Writing from a widget's callback is the usual pattern:

{{< example file="storage/preferences/main.go" region="write" >}}

## Lists

[`PrefList`](https://pkg.go.dev/github.com/belluikit/bellui#PrefList),
[`PrefListOr`](https://pkg.go.dev/github.com/belluikit/bellui#PrefListOr) and
[`SetPrefList`](https://pkg.go.dev/github.com/belluikit/bellui#SetPrefList)
do the same for slices. The type parameter is the element type.

{{< example file="storage/preferences/main.go" region="lists" >}}

## Reacting to changes

`AddChangeListener` registers a function that runs after any preference
changes. It does not tell you which key changed, so read the values you care
about.

{{< example file="storage/preferences/main.go" region="listen" >}}

On the desktop, Bellui watches the preferences file and reloads it when
another process changes it. Listeners then run on that watcher's goroutine,
not the UI thread. `bellui.RunOnMain` runs your UI update inline when you are
already on the UI thread, and moves it there otherwise.

## Where preferences are stored

On desktop and mobile platforms the values are written to a file named
`preferences.pxf` in the app's configuration folder (the same folder as
[app storage](../app-storage/)). The file is written when a value changes, with
rapid bursts of changes batched together, and once more when the app exits. On
the web the values are kept in the browser's `localStorage`.

## The complete example

{{< example "storage/preferences/main.go" >}}
