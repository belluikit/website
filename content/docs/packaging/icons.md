---
title: App icons
description: Where the bellui command finds your app icon, and what it makes from it for each platform.
weight: 4
---

## Choosing the icon

The `bellui` command looks for the icon in this order:

1. the `--icon` flag;
2. `icon` in the `details` block of your [manifest](../metadata/), relative to
   the manifest;
3. a file named `Icon.png` in the app's source folder.

If none of these exists, packaging still works: the command uses the Bellui
logo and prints a message saying so.

The largest size any platform needs is 1024 × 1024 pixels (for iOS and
macOS), so start from a square image at least that big. A PNG is best; a JPEG
or another image format the command can read is converted to PNG first.

## What each platform gets

| Platform | What the command does with the icon |
| --- | --- |
| macOS | Converts it to `icon.icns` in the bundle. It first scales the image down and clips it to the rounded square that macOS icons use. Pass `--use-raw-icon` to skip that step if your icon is already shaped. |
| Windows | Converts it to an `.ico` and embeds it in the executable. |
| Linux and BSD | Copies it into `share/pixmaps`, named after the app ID (or the app name if there is no ID). |
| Android | Puts it in the APK. For an adaptive icon, set `adaptive_icon` in the manifest with `foreground`, `background` and, optionally, `monochrome` images. |
| iOS | Resizes it to the sizes iOS needs, from 76 to 1024 pixels, and compiles them into the app. |
| Web | Copies it as `icon.png`, used as the page icon and on the loading screen. |

## The icon at run time

Inside the app, the icon set with
[`App.SetIcon`](https://pkg.go.dev/github.com/belluikit/bellui#App) is used for
windows and, on the desktop, as the default system tray icon. A resource made
with [`bellui bundle`](../resources/) works well for this.
