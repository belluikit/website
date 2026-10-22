---
title: Scaling
description: How theme sizes become pixels, and how users and developers change the scale.
weight: 5
---

Sizes in Bellui, in themes and in your layouts, are in device-independent
units, not pixels. A text size of 14 is the same physical size on a normal
screen and on a high-density one. Bellui multiplies every size by a **scale**
when it draws.

You do not need to do anything for this to work. Write sizes for a normal
screen, and Bellui scales them for the display the window is on.

## Where the scale comes from

On desktop, the scale of a window is the product of two numbers, rounded to
one decimal place:

- **The system scale.** On macOS this is 1, because the system scales the
  rendered image itself. On Windows it is the monitor's content scale from
  the display settings. On Wayland it is 1, because the compositor scales
  the content. Elsewhere Bellui works it out from the monitor's size and
  resolution, with 120 dpi as 1.0 and never less than 1.0.
- **The user scale.** A preference for making everything larger or smaller.
  It comes from the `BELLUI_SCALE` environment variable when set, otherwise
  from Bellui's user settings, otherwise 1.

In a browser the system scale is the page's `devicePixelRatio`.

`Settings().Scale()` returns the user scale from the settings. A canvas
reports its own effective scale through its `Scale()` method.

## Try a different scale

Set `BELLUI_SCALE` to check that your layout works for users who need larger
text:

```sh
BELLUI_SCALE=1.5 go run .
```

`BELLUI_SCALE=auto` ignores the user setting and
uses only the system scale.

Users can set their scale, theme and primary colour for all Bellui apps with
the settings app in the bellui module:

```sh
go run github.com/belluikit/bellui/cmd/bellui_settings@v1.2.0
```

Running desktop apps watch these settings and apply changes straight away.
