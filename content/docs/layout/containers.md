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
  individual children.
- `NewShrinkLayout()`: reports a minimum size of zero and fills the space.
  Wrap content in it when a parent such as a split must be able to shrink
  below the content's natural minimum size.
- `NewResponsiveLayout` and `NewRowWrapLayout`: see
  [Responsive layouts](../responsive/).

## Changing a container's children

A container's children are in its `Objects` field. Change them with methods,
so the layout runs again:

- `Add(obj)` appends a child. It is safe to call from any goroutine.
- `Remove(obj)` removes one child, and `RemoveAll()` removes them all. Call
  these on the [UI thread](/docs/concepts/ui-thread/).

If you assign to `Objects` directly, call `Refresh()` afterwards.
`Hide()` on a child removes it from the layout: the layouts above skip
hidden children.

## The whole program

{{< example "layout/containers/main.go" >}}
