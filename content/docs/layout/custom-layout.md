---
title: Writing a custom layout
description: Implement the core.Layout interface to position and size a container's children your own way.
weight: 6
---

When no standard layout places things the way you want, write your own. A
layout is any type that implements
[`core.Layout`](https://pkg.go.dev/github.com/belluikit/bellui/core#Layout),
which has two methods:

| Method | Job |
| --- | --- |
| `MinSize(objects []core.CanvasObject) geo.Size` | Report the smallest size that fits all the children. |
| `Layout(objects []core.CanvasObject, size geo.Size)` | Call `Resize` and `Move` on each child to place it inside `size`. |

The container calls both. `objects` is the container's child list, and
positions are relative to the container's top-left corner.

## An example: keep an aspect ratio

This layout makes its children as large as possible at a fixed width to
height ratio, and centres them. It is the shape of a video player or a photo
frame.

{{< example file="layout/customlayout/main.go" region="layout" >}}

Use it with `container.New`, like any other layout:

{{< example file="layout/customlayout/main.go" region="use" >}}

Given a 600×500 container, the rectangle becomes 600×337.5, centred
vertically. In
a wide, short space it is limited by the height instead.

## When the methods run

- `Layout` runs when the container is resized, when its children change
  through `Add`, `Remove` or `RemoveAll`, and when the container is
  refreshed.
- `MinSize` runs whenever something asks the container for its minimum size,
  which happens often: every time a parent lays itself out. Keep it cheap.

## Guidelines

- **Skip hidden children.** The standard layouts ignore objects whose
  `Visible()` is false, both when placing and when measuring. Users expect
  `Hide()` to take a child out of the layout.
- **Use the children's minimum sizes.** Do not make a child smaller than its
  `MinSize()` unless that is the point of your layout; the content will not
  fit.
- **Make `MinSize` agree with `Layout`.** If `MinSize` reports less than
  `Layout` needs, children are squeezed or overlap when the container is at its
  minimum size.
- **Space with the theme.** For gaps between children, use
  `theme.Padding()`, as the standard layouts do, so your layout matches the
  rest of the app.
- **Store settings in the layout value.** Configuration such as the ratio
  above lives in the layout's fields. One layout value can serve several
  containers if it has no per-container state.

A layout that needs to react to input or draw something of its own is no
longer just a layout. Write a [custom widget](/docs/custom-widgets/)
instead; its renderer has a `Layout` method of its own.

## The whole program

{{< example "layout/customlayout/main.go" >}}
