---
title: Icons
description: Use the theme's icons in widgets, colour your own SVG icons from the theme, and replace icons in a custom theme.
weight: 3
---

The theme supplies the icons widgets draw: the arrow on a back button, the
check mark in a checkbox, the clear button in a search field. You can use the
same icons in your own widgets, so they change along with the theme.

## Use a theme icon

Each icon has a name, a
[`bellui.ThemeIconName`](https://pkg.go.dev/github.com/belluikit/bellui#ThemeIconName)
constant such as `theme.IconNameHome`. There are two ways to get it:

{{< example file="theming/icons/main.go" region="lookup" >}}

- [`theme.Icon(name)`](https://pkg.go.dev/github.com/belluikit/bellui/theme#Icon)
  looks the name up in the current theme.
- The shortcut functions, such as `theme.HomeIcon()`, `theme.SettingsIcon()`
  or `theme.DeleteIcon()`, do the same for one icon each.

The [`theme` package reference](https://pkg.go.dev/github.com/belluikit/bellui/theme#pkg-constants)
lists every `IconName…` constant. They cover common actions (add, remove,
copy, paste, undo, search), files and folders, media controls, mail,
navigation, visibility, volume, zoom and window controls.

Theme icons are SVG images. Most are drawn in the theme's foreground colour,
so they are dark in the light variant and light in the dark variant.

## Directional icons

Some icons point in a reading direction: navigate back and next, the submenu
arrow, and mail reply, reply all and forward. When the app's layout is
right-to-left, the theme draws these icons mirrored. Other icons, such as
search or settings, are never mirrored. See
[Right-to-left layout](../right-to-left/).

## Colour your own icons

To make your own SVG icon follow the theme, wrap it in a themed resource. The
wrapper replaces the SVG's fill colours with a theme colour each time it is
drawn:

{{< example file="theming/icons/main.go" region="custom" >}}

| Constructor | Colour |
| --- | --- |
| [`theme.NewThemedResource`](https://pkg.go.dev/github.com/belluikit/bellui/theme#NewThemedResource) | foreground |
| [`theme.NewPrimaryThemedResource`](https://pkg.go.dev/github.com/belluikit/bellui/theme#NewPrimaryThemedResource) | primary |
| [`theme.NewErrorThemedResource`](https://pkg.go.dev/github.com/belluikit/bellui/theme#NewErrorThemedResource) | error |
| [`theme.NewSuccessThemedResource`](https://pkg.go.dev/github.com/belluikit/bellui/theme#NewSuccessThemedResource) | success |
| [`theme.NewWarningThemedResource`](https://pkg.go.dev/github.com/belluikit/bellui/theme#NewWarningThemedResource) | warning |
| [`theme.NewDisabledResource`](https://pkg.go.dev/github.com/belluikit/bellui/theme#NewDisabledResource) | disabled |
| [`theme.NewColoredResource`](https://pkg.go.dev/github.com/belluikit/bellui/theme#NewColoredResource) | any colour name you pass |

This works for SVG icons only. For anything that is not SVG, the wrapper logs
an error and draws the image with its own colours.

## Replace icons in a custom theme

A custom theme's `Icon` method can return a different resource for any icon
name. Widgets that use that icon then draw yours. For an example, see the
`Icon` method in [Custom themes](../custom/#wrap-the-default-theme). In a JSON
theme, the `Icons` object maps icon names to resource URIs.

Return a themed resource from `Icon` if your replacement should follow the
theme colours like the built-in icons do.

The full program:

{{< example "theming/icons/main.go" >}}
