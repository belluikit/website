---
title: Split containers
description: Divide space between panes the user resizes by dragging - Split for two panes, Paned for any number.
weight: 3
---

A split container divides its space between panes, with a divider the user
drags to change the share. Bellui has two:

- [`container.Split`](https://pkg.go.dev/github.com/belluikit/bellui/container#Split)
  holds exactly two children.
- [`container.Paned`](https://pkg.go.dev/github.com/belluikit/bellui/container#Paned)
  holds any number, with a divider between each pair.

{{< screenshot name="containment" alt="The gallery's Containment page: a tree on the left and scrolling cards on the right, divided by a draggable split bar" >}}

## Split

{{< example file="layout/split/main.go" region="split" >}}

`container.NewHSplit(leading, trailing)` puts the panes side by side with a
vertical divider. `container.NewVSplit(top, bottom)` stacks them with a
horizontal divider.

`SetOffset` sets the divider position as a fraction of the space, from `0.0`
to `1.0`. A new split starts at `0.5`, sharing the space equally. The offset
is limited so that neither pane is smaller than its minimum size. If the
user drags the divider, the `Offset` field follows.

If you hide one pane, the other one gets all the space and the divider is
not shown.

## Paned

{{< example file="layout/split/main.go" region="paned" >}}

`container.NewHPaned(panels...)` arranges panes left to right;
`container.NewVPaned(panels...)` top to bottom. Each pane has a weight, 1 by
default, and gets a share of the space in proportion to its weight.
`SetWeight(i, w)` changes one weight. `AddPanel` appends a pane and
`RemovePanel(i)` removes one.

As with Split, no pane is made smaller than its minimum size.

## Panes that shrink

A pane cannot be dragged smaller than its content's minimum size. A long
label or a wide form can make a pane hold more space than you want. Two
ways around it:

- Put the content in a [scroll container](../scrolling/), whose minimum size
  is small.
- Wrap it in `container.New(layout.NewShrinkLayout(), content)`. The shrink
  layout reports a minimum size of zero and always gives the content the
  pane's full size, even when that is below the content's minimum.

## The whole program

The example nests a Split inside a Paned: a sidebar and editor on top, and
two panels below.

{{< example "layout/split/main.go" >}}
