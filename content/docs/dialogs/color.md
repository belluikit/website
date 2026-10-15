---
title: Colour picker
description: Let people choose a colour from swatches, or edit one precisely in advanced mode.
weight: 4
---

[`dialog.ShowColorPicker`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#ShowColorPicker)
shows rows of preset colours and greys. Tapping a swatch calls your callback
with that colour and closes the dialog. The picker remembers the colours
chosen recently and shows them in an extra row the next time.

{{< example file="dialogs/color/main.go" region="simple" >}}

The callback receives a `color.Color` from the standard `image/color`
package. Cancelling closes the dialog without calling it.

## Advanced mode

Set the `Advanced` field of a
[`ColorPickerDialog`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#ColorPickerDialog)
to `true` before you show it. The dialog then adds an expandable
**Advanced** section with a colour wheel, red, green, blue, hue, saturation,
lightness and alpha channels, a preview, and a hex entry. A **Confirm** button
delivers the colour you have built.

{{< example file="dialogs/color/main.go" region="advanced" >}}

`SetColor` sets the starting colour. It works only in advanced mode; in simple
mode it logs an error and does nothing, so set `Advanced` first.

## The complete example

{{< example "dialogs/color/main.go" >}}
