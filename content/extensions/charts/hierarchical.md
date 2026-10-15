---
title: Hierarchies and flows
description: Tree, treemap, sunburst, sankey and graph series in x/charts.
weight: 3
---

These series show structure rather than values on an axis: a hierarchy, the
size of its parts, or flows between things. They ignore the chart's axes and
lay themselves out in the plot area.

Nodes take their colours from a fixed built-in palette unless you set a node's
own `Color`. The tree, treemap, sankey and graph series draw their labels in a
fixed light colour, so the labels are easy to read in the dark theme and faint
in the light theme. The sunburst draws its labels in dark grey.

## Tree

A [`TreeSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#TreeSeries)
draws a hierarchy as nodes joined by lines, such as an organisation chart or a
file tree.

{{< screenshot name="charts-tree" alt="A tree drawn left to right: a root node, three team nodes and seven leaf nodes" >}}

{{< example file="extensions/charts/series/tree/main.go" region="tree" >}}

`Data` is the root
[`TreeNode`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#TreeNode),
and each node lists its `Children`. `Orient` is `"TB"` for top to bottom (the
default) or `"LR"` for left to right. `NodeSize` is the radius of each dot
(default 6). Nodes are coloured by depth.

Leaves are spaced evenly and each parent is centred on its children. The
deepest level is placed short of the far edge, which leaves empty space on
that side, as in the screenshot. `TreeNode.Value` is not used for drawing.

## Treemap

A [`TreemapSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#TreemapSeries)
divides a rectangle into nested rectangles sized by value. Use it to see what
takes up the space, such as disk usage.

{{< screenshot name="charts-treemap" alt="A treemap of disk usage: Media, Projects, System and Apps, with Media and Projects split into their children" >}}

{{< example file="extensions/charts/series/treemap/main.go" region="treemap" >}}

`Data` lists the top-level
[`TreemapNode`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#TreemapNode)s.
A leaf's size is its `Value`. A node with `Children` is the sum of its
children, and its own `Value` is ignored. Siblings are sorted largest first and
placed side by side along the longer side of their parent's rectangle. A
rectangle's name is drawn when there is room for it.

## Sunburst

A [`SunburstSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#SunburstSeries)
shows the same kind of hierarchy as rings: the top level in the middle, each
deeper level one ring further out.

{{< screenshot name="charts-sunburst" alt="A sunburst of disk usage: an inner ring of four folders and an outer ring with the contents of Media and Projects" >}}

{{< example file="extensions/charts/series/sunburst/main.go" region="sunburst" >}}

[`SunburstNode`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#SunburstNode)
works like the treemap's node: a leaf has a `Value`, and a parent is the sum of
its children. Each segment spans an angle in proportion to its value. Each ring
colours its segments from the palette in order, lighter in the outer rings. A segment gets a label only
when it spans more than 20 degrees.

## Sankey

A [`SankeySeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#SankeySeries)
is meant to show quantities flowing from sources to targets, with the width of
each link proportional to its value.

{{< screenshot name="charts-sankey" alt="A sankey diagram with three sources on the left and three targets on the right; the links are straight horizontal bands that do not meet the nodes they join" >}}

{{< example file="extensions/charts/series/sankey/main.go" region="sankey" >}}

`Nodes` lists the
[`SankeyNode`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#SankeyNode)s
by index, and each
[`SankeyLink`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#SankeyLink)
joins a `Source` index to a `Target` index with a `Value`.

{{< callout type="warning" >}}
The sankey diagram does not render correctly in v1.0.0. It has only two
columns: nodes that are only sources on the left, every other node on the
right, so a flow with intermediate steps cannot be shown. Each link is drawn as
a straight horizontal band at the height halfway between its two nodes, so a
link between nodes at different heights does not meet either of them, as the
screenshot shows. Until this is fixed, consider a [graph](#graph) for flows.
{{< /callout >}}

## Graph

A [`GraphSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#GraphSeries)
draws nodes joined by edges, for networks such as services and their
dependencies. For graphs the user edits, see the [Graphs](../../dag/)
extension instead.

{{< screenshot name="charts-graph" alt="A service map: seven named nodes at fixed positions joined by edges of different widths" >}}

{{< example file="extensions/charts/series/graph/main.go" region="graph" >}}

The main fields:

- `Nodes` lists the
  [`GraphNode`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#GraphNode)s.
  `X` and `Y` place a node as fractions of the plot area, from 0 to 1. If every
  node is left at 0, 0, the nodes are arranged on a circle instead.
- `GraphNode.Value` makes a node larger; `NodeSize` is the base radius
  (default 8).
- `Edges` joins nodes by index. `GraphEdge.Value` makes an edge thicker.

Edges are plain lines without arrowheads, and there is no automatic
force-directed layout.
