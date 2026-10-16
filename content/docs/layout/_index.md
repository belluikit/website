---
title: Layout
description: Arrange widgets with containers and layouts - boxes, borders, grids, scrolling, splits, tabs, responsive rows and your own layouts.
weight: 2
---

A window shows one object. To show more, you put them in a **container**. A
container holds a list of children and a **layout**, and the layout decides
where each child goes and how big it is. Containers are
[canvas objects](/docs/concepts/canvas-objects/) themselves, so you nest them
to build a screen: a border around a scrolling list, a box of buttons in its
top edge, and so on.

The [`container`](https://pkg.go.dev/github.com/belluikit/bellui/container)
package has a constructor for each common arrangement, such as
`container.NewVBox` and `container.NewBorder`. Each returns a
[`*bellui.Container`](https://pkg.go.dev/github.com/belluikit/bellui#Container)
with a layout from the [`layout`](https://pkg.go.dev/github.com/belluikit/bellui/layout)
package. `container.New(layout, objects...)` combines any layout with any
children.

This app shows the standard containers side by side:

{{< example file="layout/containers/main.go" region="overview" >}}

The functions it calls, and the rest of the program, are on the
[Containers and layouts](containers/) page.

## How sizing works

Every object reports a **minimum size** (`MinSize`). A layout reads the
minimum sizes of its children, and its own `MinSize` is what it needs to fit
them. On desktop, the user cannot make a window smaller than its content's
minimum size.

When the container is resized, the layout places the children. Some layouts
keep children at their minimum size in one direction (a VBox keeps each
child's height); others stretch them to fill (a Stack, a grid cell). Knowing
which one does what is most of the work of laying out a screen.

{{< cards >}}
  {{< card link="containers/" title="Containers and layouts" subtitle="VBox, HBox, Border, Grid, GridWrap, Center, Stack, Padded and Form." >}}
  {{< card link="scrolling/" title="Scrolling" subtitle="Show content larger than the space it has." >}}
  {{< card link="split/" title="Split containers" subtitle="Panes the user resizes by dragging a divider." >}}
  {{< card link="tabs/" title="Tabs" subtitle="App tabs for navigation, document tabs the user opens and closes." >}}
  {{< card link="responsive/" title="Responsive layouts" subtitle="Rows that reflow by window width." >}}
  {{< card link="custom-layout/" title="Writing a custom layout" subtitle="Implement core.Layout to place children your own way." >}}
{{< /cards >}}
