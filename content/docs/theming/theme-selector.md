---
title: Theme selector
description: Add a System, Light or Dark choice and a primary colour picker to your settings page, and restore the user's choice at startup.
weight: 6
---

[`widget.ThemeSelector`](https://pkg.go.dev/github.com/belluikit/bellui/widget#ThemeSelector)
is a ready-made appearance setting. It shows a System, Light and Dark switch
above a row of swatches for the eight standard primary colours.

{{< screenshot name="theme" alt="A theme selector with System selected and the blue swatch ringed" >}}

## Add the selector

The widget only reports what the user picks. It does not change the theme or
save anything itself, so you decide what happens:

{{< example file="theming/selector/main.go" region="selector" >}}

- `OnChanged` receives a
  [`widget.ThemeSelection`](https://pkg.go.dev/github.com/belluikit/bellui/widget#ThemeSelection):
  the variant choice (`widget.ThemeVariantSystem`, `ThemeVariantLight` or
  `ThemeVariantDark`) and the name of the primary colour.
- [`ThemeSelection.Apply`](https://pkg.go.dev/github.com/belluikit/bellui/widget#ThemeSelection.Apply)
  re-themes the running app. An empty `Primary` leaves the primary colour as
  it is.
- [`widget.SaveThemeSelection`](https://pkg.go.dev/github.com/belluikit/bellui/widget#SaveThemeSelection)
  stores the choice in the preferences you pass, normally the app's own.

A new selector starts on System, with the app's current primary colour (blue
if none is set). To show a saved choice, assign it to the `Selection` field
before the selector is displayed. Assigning the field does not call
`OnChanged`. `SetSelection` does, when the value changes.

## Restore the choice at startup

Load the saved selection and apply it before you show the first window, so
the app does not open in one theme and then switch:

{{< example file="theming/selector/main.go" region="startup" >}}

[`widget.LoadThemeSelection`](https://pkg.go.dev/github.com/belluikit/bellui/widget#LoadThemeSelection)
returns `false` when nothing was saved. The app then follows the system
appearance, as it does by default.

Preferences need an app ID, so create the app with `app.NewWithID` (or set
the ID in the app's metadata). Without one, `Preferences()` logs an error and
nothing is stored between runs.

## What it does not do

- **It does not change other apps.** The choice applies to your app and is
  saved in your app's preferences. Bellui's system-wide settings, shared by
  all Bellui apps, are not touched.
- **System is resolved when it is applied.** `Apply` with
  `ThemeVariantSystem` reads the system appearance at that moment.
- **No custom colours.** The swatches are the eight named primary colours.
  For any other colour, build a [custom theme](../custom/#generate-colours-from-a-seed-colour).

The full program:

{{< example "theming/selector/main.go" >}}
