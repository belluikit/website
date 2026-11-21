---
title: Custom themes
description: Change colours, sizes, fonts and icons by wrapping the default theme, generating colours from a seed colour, or loading a theme written in PXF.
weight: 2
---

A theme is any value that implements
[`bellui.Theme`](https://pkg.go.dev/github.com/belluikit/bellui#Theme). It has
four methods, one for each kind of value a widget asks for:

| Method | Returns | Looked up by |
| --- | --- | --- |
| `Color(name, variant)` | a `color.Color` | `bellui.ThemeColorName`, plus the light or dark variant |
| `Size(name)` | a `float32` | `bellui.ThemeSizeName` |
| `Font(style)` | a font resource | `bellui.TextStyle` |
| `Icon(name)` | an icon resource | `bellui.ThemeIconName` |

The names are constants in the
[`theme`](https://pkg.go.dev/github.com/belluikit/bellui/theme) package, such
as `theme.ColorNamePrimary`, `theme.SizeNamePadding` and
`theme.IconNameHome`.

## Wrap the default theme

You rarely want to define every value. Embed an existing theme and override
only what you change. Every method you do not write falls through to the
embedded theme:

{{< example file="theming/customtheme/main.go" region="theme" >}}

Install it once, before you show the first window:

{{< example file="theming/customtheme/main.go" region="apply" >}}

Some things to know:

- **Check the variant.** `Color` receives the variant the app is using.
  Return a different colour for `theme.VariantDark` when the light one would
  not have enough contrast.
- **Keep the modern look.** Widgets choose between the "classic" and "modern"
  drawing style by asking whether the theme implements
  [`theme.RenderStyleTheme`](https://pkg.go.dev/github.com/belluikit/bellui/theme#RenderStyleTheme).
  Embedding `bellui.Theme` does not pass that method through, so a wrapper
  without its own `RenderStyle` method gets the classic style. Add the method,
  as above.
- **Pair colours.** When you change a colour, also change the colour drawn on
  top of it. `ColorNamePrimary` goes with `ColorNameForegroundOnPrimary`, for
  example.
- **Sizes are in device-independent units**, not pixels. See
  [Scaling](../scaling/).

## Apply a theme to part of the window

[`container.NewThemeOverride`](https://pkg.go.dev/github.com/belluikit/bellui/container#NewThemeOverride)
draws its content with a different theme from the rest of the app:

{{< example file="theming/customtheme/main.go" region="override" >}}

Everything inside the container uses the theme you pass, including widgets in
nested containers. If you add objects to the content later, call `Refresh` on
the `ThemeOverride` so they pick up its theme. Lists, tables and trees inside
an override take its theme for the rows and cells they make as they scroll.
Use overrides sparingly: a window with several looks is harder to read.

An override reaches only the objects that exist when it is applied. A
[custom widget](/docs/custom-widgets/building/) that makes objects as it
grows, such as rows or dividers, calls
[`theme.MatchScope(parent, objs...)`](https://pkg.go.dev/github.com/belluikit/bellui/theme#MatchScope)
with itself as `parent` and the new objects. They then draw with the
override around the widget, and not the app theme. Objects that already
have it are skipped, so calling it on every layout is cheap. Outside an
override it does nothing.

## Generate colours from a seed colour

The default theme offers eight named primary colours. To start from any other
colour, generate the MD3 scheme yourself and serve it from your theme:

{{< example file="theming/seedcolor/main.go" region="seed" >}}

- [`theme.CorePalettesFromPrimary`](https://pkg.go.dev/github.com/belluikit/bellui/theme#CorePalettesFromPrimary)
  turns the seed colour into the six MD3 tonal palettes.
- [`theme.LightSchemeFromPalettes`](https://pkg.go.dev/github.com/belluikit/bellui/theme#LightSchemeFromPalettes)
  and
  [`theme.DarkSchemeFromPalettes`](https://pkg.go.dev/github.com/belluikit/bellui/theme#DarkSchemeFromPalettes)
  assign tones from those palettes to each colour role, giving an
  [`MD3Scheme`](https://pkg.go.dev/github.com/belluikit/bellui/theme#MD3Scheme).
- [`theme.ColorFromMD3Scheme`](https://pkg.go.dev/github.com/belluikit/bellui/theme#ColorFromMD3Scheme)
  looks a colour name up in a scheme. It understands both the MD3 role names
  and the general names such as `ColorNameBackground`. It returns `nil` for
  names the scheme does not cover, such as `ColorNameWarning`, so fall back
  to the embedded theme.

The seed colour sets the hue. The primary palette always gets enough colour
intensity to read as an accent, even from a greyish seed.

{{< example file="theming/seedcolor/main.go" region="apply" >}}

## Load a theme from PXF

[`theme.FromPXF`](https://pkg.go.dev/github.com/belluikit/bellui/theme#FromPXF)
builds a theme from text in PXF, the data format Bellui uses for its own
files. Anything the data does not set comes from the default theme:

{{< example file="theming/pxftheme/main.go" region="pxf" >}}

The fields are `colors`, `colors_light`, `colors_dark`, `sizes`, `fonts` and
`icons`. Inside each, the keys are the string values of the theme names:
`primary` for `ColorNamePrimary`, `text` for `SizeNameText`, and so on.

- `colors` applies to both variants. `colors_light` and `colors_dark` apply
  to one variant and win over `colors`.
- Colours are hex strings: `#rgb`, `#rgba`, `#rrggbb` or `#rrggbbaa`, with or
  without the `#`.
- `sizes` holds numbers.
- `fonts` and `icons` map names to resource URIs, such as `file:///…`. Icon
  keys are icon names; see [Icons](../icons/#replace-icons-in-a-custom-theme).
  Font keys are `regular`, `bold`, `boldItalic` and `monospace`. Other styles
  use one of these: italic, medium and symbol text use `regular`, and bold
  monospace text uses `bold`.

If the data cannot be parsed, `FromPXF` returns the error together with the
default theme, so the app still has a usable theme:

{{< example file="theming/pxftheme/main.go" region="load" >}}

[`theme.FromPXFWithFallback`](https://pkg.go.dev/github.com/belluikit/bellui/theme#FromPXFWithFallback)
takes the theme to fall back to instead. `FromPXFReader` and
`FromPXFReaderWithFallback` read from an `io.Reader`.

{{< callout type="warning" >}}
A PXF theme does not select the modern widget style, so widgets are drawn in
the classic style. To keep the modern style, wrap the PXF theme in a Go type
with a `RenderStyle` method, as in [Wrap the default theme](#wrap-the-default-theme).
{{< /callout >}}

On desktop, `theme.DefaultTheme()` also reads a user-wide theme from the
file `theme.pxf` in Bellui's configuration folder, the folder that holds
the global `settings.pxf`. When there is no `theme.pxf`, it reads an older
`theme.json` there instead.

### Themes written in JSON

`theme.FromJSON`, `FromJSONWithFallback`, `FromJSONReader` and
`FromJSONReaderWithFallback` still read the same data written as JSON, with
the fields named `Colors`, `Colors-light`, `Colors-dark`, `Sizes`, `Fonts`
and `Icons`. They are deprecated: use the `FromPXF` functions for new code.

The whole program:

{{< example "theming/pxftheme/main.go" >}}

## Check a theme

[`theme.ValidateTheme`](https://pkg.go.dev/github.com/belluikit/bellui/theme#ValidateTheme)
asks a theme for every known colour, in both variants, and every known size.
It returns one line for each colour that is `nil` or transparent and each
size that is 0. Call it from a test of your theme. A size you set to 0 on
purpose, such as a corner radius, is listed too.
