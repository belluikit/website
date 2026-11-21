---
title: Containers and layouts
description: The standard containers - VBox, HBox, Border, Grid, GridWrap, Center, Stack, Padded and Form - and how each one sizes its children.
weight: 1
---

Each section below shows one container, taken from a single program that
draws them all. In the examples, `box(name)` is a small helper that returns a
tinted rectangle with a label on it.

| Constructor | Children are | Use it for |
| --- | --- | --- |
| `NewVBox`, `NewHBox` | In a column or a row, at their minimum height (or width) | Lists of controls, toolbars |
| `NewBorder` | Around the edges, with the rest in the middle | The outer frame of most screens |
| `NewGridWithColumns`, `NewGridWithRows` | In equal cells | Button pads, dashboards |
| `NewGridWrap` | At one fixed size, wrapping onto new rows | Thumbnails, tiles |
| `NewCenter` | At their minimum size, centred | A single control in a large space |
| `NewStack` | On top of each other, all full size | Backgrounds, overlays |
| `NewPadded` | Full size, inset by the theme padding | Breathing room around content |
| `New(layout.NewFormLayout(), ...)` | In label and value columns | Forms |

## VBox and HBox

`container.NewVBox` places children top to bottom. Each child is as wide as
the container and as tall as its own minimum height.

{{< example file="layout/containers/main.go" region="vbox" >}}

`container.NewHBox` does the same left to right: each child is as tall as the
container and as wide as its minimum width.

A [`layout.NewSpacer()`](https://pkg.go.dev/github.com/belluikit/bellui/layout#NewSpacer)
in a box takes up the free space. Two spacers share it equally. Use one to
push items to the far end, or two to centre something:

{{< example file="layout/containers/main.go" region="hbox" >}}

Boxes never stretch a child along their main direction. If you want a child to
fill the remaining space, use a Border instead. Between children, boxes put
the theme's padding; `layout.NewCustomPaddedVBoxLayout(p)` and
`NewCustomPaddedHBoxLayout(p)` use your own value.

## Border

`container.NewBorder(top, bottom, left, right, objects...)` puts the first
four objects along the edges and stretches the rest over the middle. Pass
`nil` for an edge you do not need.

{{< example file="layout/containers/main.go" region="border" >}}

The top and bottom objects get their minimum height and the full width. The
left and right objects get their minimum width and the height between top and
bottom. Everything after the fourth argument fills the middle; if you pass
several, they are stacked on top of each other.

A Border is the usual frame for a screen: a toolbar on top, a status line at
the bottom, and a scrolling list in the middle that gets all the remaining
space.

## Grid

`container.NewGridWithColumns(n, ...)` lays children out in `n` columns,
adding rows as needed. Every cell is the same size, and each child fills its
cell.

{{< example file="layout/containers/main.go" region="grid" >}}

`container.NewGridWithRows(n, ...)` fixes the number of rows instead.
`container.NewAdaptiveGrid(n, ...)` switches between the two with the device
orientation; see [Responsive layouts](../responsive/#adaptive-grid).

## GridWrap

`container.NewGridWrap(size, ...)` gives every child the same fixed size and
places them in rows, wrapping to a new row when the next one does not fit.

{{< example file="layout/containers/main.go" region="gridwrap" >}}

When the children have different natural widths, such as a set of tags, use
`layout.NewRowWrapLayout()` instead; see
[Responsive layouts](../responsive/#row-wrap).

## Center

`container.NewCenter` keeps each child at its minimum size, in the middle of
the space.

{{< example file="layout/containers/main.go" region="center" >}}

## Stack

`container.NewStack` makes every child the full size of the container and
draws them in order, so later children are on top. Use it to put a
background behind something, or a badge over it.

{{< example file="layout/containers/main.go" region="stack" >}}

## Padded

`container.NewPadded` makes each child the full size of the container minus
the theme's padding on all four sides.

{{< example file="layout/containers/main.go" region="padded" >}}

For different padding on each side, use
`container.New(layout.NewCustomPaddedLayout(top, bottom, left, right), ...)`.

## Form

The form layout has two columns. Children are taken in pairs: the first of
each pair is the label, the second its value. The label column is as wide as
the widest label.

{{< example file="layout/containers/main.go" region="form" >}}

For a form with a submit button and validation, use the
[`widget.Form`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Form)
widget instead.

## Other layouts

The [`layout`](https://pkg.go.dev/github.com/belluikit/bellui/layout) package
has a few more, used with `container.New`:

- `NewFlexHBoxLayout(flex)` and `NewFlexVBoxLayout(flex)`: a row or column
  where each child has a grow factor. Children with factor 0 keep their
  minimum size; the rest share the free space in proportion. The
  `MinWidth` and `MaxWidth` fields (or `MinHeight` and `MaxHeight`) clamp
  individual children. See [Flex boxes](#flex-boxes) below.
- `NewAspectStackLayout(ratio)`: stacks its children, as a Stack does, in
  the largest box of the given width to height ratio that fits, centred. A
  ratio of 1 is a square. Use it for a video or a photo frame.
- `NewShrinkLayout()`: reports a minimum size of zero and fills the space.
  Wrap content in it when a parent such as a split must be able to shrink
  below the content's natural minimum size.
- `NewResponsiveLayout` and `NewRowWrapLayout`: see
  [Responsive layouts](../responsive/).

### Flex boxes

By default a flex row stretches every child to the row's height. Set
`CrossAlign` to keep each child at its own minimum height instead, and place
it at the top (`layout.CrossStart`), in the middle (`CrossCenter`) or at the
bottom (`CrossEnd`). The default is `CrossStretch`. A flex column does the
same across its width.

Between children, a flex box puts the theme's padding, and a `Padding` of 0
also means the theme's padding. Set `ExactPadding` to use `Padding` exactly
as given, so 0 puts the children flush against each other.

{{< example file="layout/extras/main.go" region="flex" >}}

A fixed aspect ratio, with `NewAspectStackLayout`:

{{< example file="layout/extras/main.go" region="aspect" >}}

## Rounded clipping and translucency

A container draws its children in full, even where they reach past its
edges. Set `ClipToBounds` on a
[`*bellui.Container`](https://pkg.go.dev/github.com/belluikit/bellui#Container)
to cut them at its bounds. Add `ClipCornerRadius` to cut them to a rounded
rectangle, as a card cuts its content. Without `ClipToBounds`,
`ClipCornerRadius` does nothing.

{{< example file="layout/extras/main.go" region="clip" >}}

[`container.NewOpacity(content, opacity)`](https://pkg.go.dev/github.com/belluikit/bellui/container#Opacity)
draws everything in its content translucent, from 0 (not drawn) to 1
(unchanged). Nested `Opacity` containers multiply. Each object is faded on its
own, so where two objects overlap, the lower one shows through. Change the
value with `SetOpacity`. It redraws without refreshing the content, so you
can animate it.

{{< example file="layout/extras/main.go" region="opacity" >}}

A widget of your own can be faded the same way by implementing
[`core.OpacityGroup`](https://pkg.go.dev/github.com/belluikit/bellui/core#OpacityGroup).

## Changing a container's children

A container's children are in its `Objects` field. Change them with methods,
so the layout runs again:

- `Add(obj)` appends a child. It is safe to call from any goroutine.
- `Remove(obj)` removes one child, and `RemoveAll()` removes them all. Call
  these on the [UI thread](/docs/concepts/ui-thread/).

If you assign to `Objects` directly, call `Refresh()` afterwards.
`Hide()` on a child removes it from the layout: the layouts above skip
hidden children.

### Swapping pages by hand

`Window.SetContent` frees what the toolkit keeps for the content it
replaces: widget renderers, records about the objects, and the textures they
were drawn with. If you switch pages by changing a container's children
instead, call
[`bellui.Release(page)`](https://pkg.go.dev/github.com/belluikit/bellui#Release)
with the page you dropped. The next frame then frees what was kept for it.
Without the call, dropped content is freed only at a frame at least a
minute after it was last drawn, and an idle app draws no frames.

{{< example file="layout/extras/main.go" region="release" >}}

Call `Release` on the UI thread, and only for content you will not show
again. If you show it again later, its widgets get new renderers. A focused
widget on a page you take out of the window loses the focus, with or
without `Release`, so keys typed afterwards do not reach it.

The program with the flex box, aspect, clipping, opacity and page-swap
examples:

{{< example "layout/extras/main.go" >}}

## The whole program

{{< example "layout/containers/main.go" >}}
