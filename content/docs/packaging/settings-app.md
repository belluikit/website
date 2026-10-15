---
title: The settings app
description: bellui_settings lets users choose the theme, main colour, animations and scale for all Bellui apps on their computer.
weight: 6
---

Bellui apps share a few user preferences: light or dark theme, the main
(primary) colour, whether widgets animate, and the interface scale. Users
change them with a small app called `bellui_settings`, which ships with Bellui.

## Installing it

```sh
go install github.com/belluikit/bellui/cmd/bellui_settings@latest
```

Then run `bellui_settings`. To get a proper application bundle with an icon,
package it like any other app: clone the Bellui repository and run
`bellui install` in `cmd/bellui_settings`.

## What it changes

The window has an **Appearance** tab with:

- **Scale**: the size of the whole interface;
- **Animations**: whether widgets animate;
- **Main Color**: the primary colour, picked from a row of swatches;
- **Theme**: dark or light, plus "system default" on macOS and Windows.

A preview shows the result. **Apply** saves the choices.

## Where the settings are kept

The choices are saved to a file named `settings.pxf` in Bellui's configuration
folder, the parent of each app's [storage folder](/docs/storage/app-storage/):
for example `~/Library/Preferences/bellui/settings.pxf` on macOS. On the
desktop, running Bellui apps watch this file and pick up changes straight
away, without a restart.

An app can still choose its own look. If it sets a theme in code with
`Settings().SetTheme`, Bellui keeps that theme, and the user's dark or light
choice selects which variant of it is shown. See [Theming](/docs/theming/).

To try a theme for one run without changing the saved settings, set the
`BELLUI_THEME` environment variable to `light` or `dark`:

```sh
BELLUI_THEME=dark go run .
```
