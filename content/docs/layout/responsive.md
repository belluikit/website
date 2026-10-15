---
title: Responsive layouts
description: Rows that reflow with the available width, using size classes, plus grids that follow device orientation and rows that wrap.
weight: 5
---

A responsive layout changes its arrangement as the space changes: one card
per row on a phone, four across a wide desktop window. Bellui's main tool for
this is [`layout.ResponsiveLayout`](https://pkg.go.dev/github.com/belluikit/bellui/layout#ResponsiveLayout).

## Size classes

The responsive layout sorts the width it is given into one of five **size
classes**. The default boundaries, in
[`layout.DefaultResponsiveBreakpoints`](https://pkg.go.dev/github.com/belluikit/bellui/layout#DefaultResponsiveBreakpoints),
follow the Material Design 3 window size classes:

| Class | Width |
| --- | --- |
| 1. compact | below 600 |
| 2. medium | 600 to below 840 |
| 3. expanded | 840 to below 1200 |
| 4. large | 1200 to below 1600 |
| 5. extra-large | 1600 and above |

The width is the **container's** width, not the window's. A responsive
layout inside one pane of a [split](../split/) responds to that pane.

## Give each child a width per class

Wrap each child with
[`layout.Responsive(object, ratios...)`](https://pkg.go.dev/github.com/belluikit/bellui/layout#Responsive).
Each ratio is the share of the full width the child takes in one size class,
in class order. Then put the children in
[`layout.NewResponsiveLayout`](https://pkg.go.dev/github.com/belluikit/bellui/layout#NewResponsiveLayout):

{{< example file="layout/responsive/main.go" region="responsive" >}}

Here each card is full width when compact, half width when medium, and a
quarter width when expanded. The large and extra-large classes have no ratio
given, so they reuse the last one, a quarter.

The layout fills rows from left to right. When the next child's ratio would
take the row past the full width, it starts a new row. Each row is as tall as
its tallest child.

Rules for the ratios:

- A ratio must be greater than 0 and at most 1. Other values are logged and
  treated as 1.
- A class with no ratio uses the previous class's ratio. With no ratios at
  all, the child is full width in every class.
- A child that is not wrapped with `Responsive` takes a full row in every
  class.
- Ratios beyond the fifth are ignored, with a logged error.

`NewResponsiveLayout` returns a container, not a layout. Put it wherever a
canvas object goes. On a long page, put it in a
[vertical scroll](../scrolling/) so the rows can grow downwards.

## Your own breakpoints

Set the `Breakpoints` field to use different boundaries, such as Bootstrap's.
Each value is the upper bound of one class; widths at or past the last value
fall in the class after it.

{{< example file="layout/responsive/main.go" region="breakpoints" >}}

A ratio configuration always has five entries. With fewer than four
breakpoints the extra entries are unused; with more, the last one repeats.

The gallery app's Responsive page shows the classes live: it labels the
current width and class as you resize the window.

## Adaptive grid

[`container.NewAdaptiveGrid(n, ...)`](https://pkg.go.dev/github.com/belluikit/bellui/container#NewAdaptiveGrid)
is a grid of `n` columns when the device is in landscape orientation and `n`
rows when it is in portrait. On a phone or tablet it changes as the device
rotates.

{{< example file="layout/responsive/main.go" region="adaptive" >}}

{{< callout type="info" >}}
On desktop the device always reports landscape, so an adaptive grid is
always `n` columns there, however you shape the window. To respond to window
width on desktop, use the responsive layout.
{{< /callout >}}

## Row wrap

[`layout.NewRowWrapLayout()`](https://pkg.go.dev/github.com/belluikit/bellui/layout#NewRowWrapLayout)
places children left to right at their own minimum width, and starts a new
row when the next one does not fit. It suits items of similar height and
different widths, such as tags or chips. Unlike
[GridWrap](../containers/#gridwrap), it does not force every child to one
size.

{{< example file="layout/responsive/main.go" region="rowwrap" >}}

`NewRowWrapLayoutWithCustomPadding(h, v)` sets the gaps between items and
between rows.

Every row is as tall as the tallest child. Until the layout has run once, its
minimum height assumes one child per row; after that it reports the height
of the rows it actually made.

## The whole program

{{< example "layout/responsive/main.go" >}}
