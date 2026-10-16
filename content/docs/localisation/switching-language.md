---
title: Switching language
description: Let users pick a language inside the app, and keep labels, buttons and window titles translated when it changes.
weight: 2
---

By default an app uses the operating system's language. To offer a language
setting inside the app, call
[`lang.SetLocale`](https://pkg.go.dev/github.com/belluikit/bellui/lang#SetLocale):

{{< example file="localisation/switching/main.go" region="switch" >}}

`SetLocale` takes a BCP 47 locale such as `"fr"`, `"de-DE"` or
`"sr-Cyrl-RS"`. It changes translations, formats, sorting and text direction
together. An empty locale goes back to the system locale.

When the locale changes, Bellui lays out and repaints every window, as for a
theme change. Call it on the UI goroutine, for example from a widget
callback. The choice is not saved: store it in the app's preferences and call
`SetLocale` again at startup.

## Text that follows the language

Repainting is not the same as re-translating. A label made with
`widget.NewLabel(lang.L("Cancel"))` received a finished string when it was
built, and keeps it after a language change. For text that should change,
use one of these instead.

**Widgets that take a binding.** A localized binding is a
`binding.String` whose value is the translation. It translates again on
every language change:

{{< example file="localisation/switching/main.go" region="bound" >}}

- [`binding.NewLocalizedString`](https://pkg.go.dev/github.com/belluikit/bellui/data/binding#NewLocalizedString)
  works like `lang.L`, and
  [`NewLocalizedStringWithKey`](https://pkg.go.dev/github.com/belluikit/bellui/data/binding#NewLocalizedStringWithKey)
  like `lang.X`.
- [`binding.NewLocalizedPlural`](https://pkg.go.dev/github.com/belluikit/bellui/data/binding#NewLocalizedPlural)
  and
  [`NewLocalizedPluralWithKey`](https://pkg.go.dev/github.com/belluikit/bellui/data/binding#NewLocalizedPluralWithKey)
  take the count as an `Int` binding, so the text also changes with the
  count. When you pass no template data, `{{.Count}}` in the message is the
  current count.

A localized binding only listens for language changes while something is
bound to it, so a binding you stop using does not leak.

**Everything else.** A button's text, a window title or a menu item has no
binding. [`lang.BindText`](https://pkg.go.dev/github.com/belluikit/bellui/lang#BindText)
calls a setter with the translation now and after every language change:

{{< example file="localisation/switching/main.go" region="bindtext" >}}

`BindText` returns a function that stops the updates. Call it when the
widget is thrown away. For things that live as long as the app, such as the
main window's title, you can ignore it.

For anything else that depends on the locale,
[`lang.AddLocaleChangeListener`](https://pkg.go.dev/github.com/belluikit/bellui/lang#AddLocaleChangeListener)
registers a function to call after each change.

The full program:

{{< example "localisation/switching/main.go" >}}
