---
title: The default theme
description: How the Material Design 3 default theme picks light or dark, generates its colours from a primary colour, and how to switch to the dense theme.
weight: 1
---

An app uses
[`theme.DefaultTheme()`](https://pkg.go.dev/github.com/belluikit/bellui/theme#DefaultTheme)
unless you set another theme. It implements the Material Design 3 colour
system and type scale. It also selects the "modern" widget style: gradient
button backgrounds, filled card backgrounds and deeper shadows.

## Light and dark

The theme has two variants,
[`theme.VariantLight`](https://pkg.go.dev/github.com/belluikit/bellui/theme#VariantLight)
and
[`theme.VariantDark`](https://pkg.go.dev/github.com/belluikit/bellui/theme#VariantDark).
At startup the app picks the one that matches the operating system's
appearance setting.

While the app runs, it also follows changes to that setting on macOS,
Windows, Linux (through the desktop's colour-scheme setting) and the web
(through the browser's `prefers-color-scheme`). On Android and iOS the app
picks up the system appearance when the mobile driver reports a
configuration change.

To choose a variant yourself, call `SetThemeVariant` on the app's settings:

{{< example file="theming/variants/main.go" region="variant" >}}

`Settings().ThemeVariant()` returns the variant in use. The call takes effect
for the running app only. It is not saved.

{{< callout type="info" >}}
To let users choose between System, Light and Dark and remember their choice,
use the [theme selector](../theme-selector/) instead of building it yourself.
{{< /callout >}}

For a quick check during development, the `BELLUI_THEME` environment variable
forces a variant. It accepts `light` or `dark`:

```sh
BELLUI_THEME=dark go run .
```

## Primary colours

The default theme builds its whole colour scheme from one **primary
colour**. From it, the theme generates six tonal palettes (primary, secondary,
tertiary, error, neutral and neutral variant) and assigns tones from them to
about 30 named colour roles, separately for light and dark. That is why a
change of primary colour also changes surfaces, outlines and containers
slightly, not just buttons.

The primary colour is chosen by name from eight standard ones:
[`theme.PrimaryColorNames()`](https://pkg.go.dev/github.com/belluikit/bellui/theme#PrimaryColorNames)
returns red, orange, yellow, green, blue, purple, brown and gray. Blue is
used when nothing is set.

The app's settings implement
[`bellui.PrimaryColorSettings`](https://pkg.go.dev/github.com/belluikit/bellui#PrimaryColorSettings),
which can change it at runtime:

{{< example file="theming/variants/main.go" region="primary" >}}

Like the variant, this changes the running app only.

To use any other colour as the starting point, generate the scheme yourself
from a seed colour. See [Custom themes](../custom/#generate-colours-from-a-seed-colour).

## Colour names

Widgets look colours up by name, such as
[`theme.ColorNamePrimary`](https://pkg.go.dev/github.com/belluikit/bellui/theme#ColorNamePrimary)
or
[`theme.ColorNameBackground`](https://pkg.go.dev/github.com/belluikit/bellui/theme#ColorNameBackground).
There are two sets:

- The MD3 roles, for example `ColorNamePrimaryContainer`,
  `ColorNameSurfaceContainerHigh`, `ColorNameOnSurfaceVariant`,
  `ColorNameOutline` and `ColorNameScrim`.
- The general names that all Bellui themes have, for example
  `ColorNameForeground`, `ColorNameButton`, `ColorNameInputBackground` and
  `ColorNameSeparator`. The default theme maps each of them to an MD3 role,
  so `ColorNameBackground` is the MD3 surface colour.

`ColorNameWarning` is not part of MD3. The default theme draws it from a
separate palette that does not change with the primary colour.

To read a colour in your own code, call
[`theme.Color(name)`](https://pkg.go.dev/github.com/belluikit/bellui/theme#Color).
It uses the current theme and variant.

## The dense theme

[`theme.DenseTheme()`](https://pkg.go.dev/github.com/belluikit/bellui/theme#DenseTheme)
is a compact theme for desktop apps that show a lot of data at once, such as
dashboards or trading screens. It uses the same MD3 colours as the default
theme, with smaller sizes:

| Size | Default theme | Dense theme |
| --- | --- | --- |
| Body text | 14 | 12 |
| Padding | 6 | 2 |
| Inner padding | 10 | 4 |
| Inline icons | 20 | 16 |
| Button and input corner radius | 8 | 4 |

It also turns shadows off. Switch to it with `SetTheme`, and back with
`theme.DefaultTheme()`:

{{< example file="theming/variants/main.go" region="dense" >}}

The full program:

{{< example "theming/variants/main.go" >}}
