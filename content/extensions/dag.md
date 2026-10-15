---
title: Graphs
description: A directed acyclic graph model with ports and validation, and an interactive canvas for viewing and editing it.
weight: 4
---

`x/dag` is for apps that show or edit a directed acyclic graph: a data
pipeline, a build, a workflow. It has two packages:

- [`dag`](https://pkg.go.dev/github.com/belluikit/bellui/x/dag/dag) is the
  graph model. It has no user interface and no dependency on the rest of
  Bellui's widgets, so you can use it in a server or a command-line tool too.
- [`dagwidget`](https://pkg.go.dev/github.com/belluikit/bellui/x/dag/dagwidget)
  draws a graph as cards joined by curved edges, and lets the user pan, zoom,
  select, move, connect and delete, with undo.

## Install

```sh
go get github.com/belluikit/bellui/x/dag@latest
```

## The graph model

A [`dag.Graph`](https://pkg.go.dev/github.com/belluikit/bellui/x/dag/dag#Graph)
holds nodes and edges. A node has named input and output **ports**, and an
edge always runs from one node's output port to another node's input port.
An edge names a port by its node ID and port name, so give the ports of one
node different names: if an input and an output share a name, the edge finds
the input.

{{< example file="extensions/dag/model/main.go" region="model" >}}

`AddEdge` checks every edge before adding it and returns an error instead if:

- an edge with the same ID is already in the graph (`ErrEdgeExists`);
- either node does not exist (`ErrNodeNotFound`), or either port does not
  exist on its node (`ErrPortNotFound`);
- the source is not an output or the target is not an input
  (`ErrPortDirection`);
- the target input already has an edge; an input accepts one connection
  (`ErrTargetOccupied`);
- the edge would create a cycle (`ErrCycleDetected`);
- the graph has a component registry and the two ports carry incompatible
  stream types (`ErrStreamIncompat`, see [Components](#components-palette-and-properties)).

Test for them with `errors.Is`. `WouldCycle` asks the cycle question without
adding anything. Removing a node also removes its edges and anything attached
to it.

`TopologicalOrder` returns the nodes in an order where every node comes after
the nodes it depends on, ready for execution:

{{< example file="extensions/dag/model/main.go" region="order" >}}

The graph's methods are safe to call from several goroutines.

Besides nodes and edges, a graph can hold:

- **Groups**, boxes around several nodes, which can be collapsed. A node is in
  at most one group. **Group edges** say that everything in one group comes
  before everything in another, and `TopologicalOrder` respects them.
- **Link edges**, plain lines between two nodes with an optional label. They
  carry no data, may form cycles, and are ignored by `TopologicalOrder`.
- **Annotations**, text notes placed freely or attached to a node, edge or
  group.

### Saving and loading

{{< example file="extensions/dag/model/main.go" region="save" >}}

`MarshalGraph` writes the graph, including positions, groups, link edges and
annotations, as PXF text, the same format Bellui uses for its own settings
files. `GraphToProto` and `GraphFromProto` convert to and from the Protocol
Buffers message in `dag/dagpb` if you prefer another encoding. Loading goes
through the same validation as `AddEdge`: an entry that is malformed, would
create a cycle or refers to something missing is dropped, and the rest of the
graph loads. Only data that cannot be decoded at all makes `UnmarshalGraph`
fail.

### Layout

`dag.LayeredLayout` computes a position for every node. It puts the nodes in
layers by the longest path that reaches them from a node with no inputs, with
data flowing left to right, and orders each layer to reduce crossing edges. It returns the positions
without changing the graph. `NeedsLayout` reports whether a graph arrived
without positions, for example from code that never set them.

## The canvas

[`dagwidget.NewDAGCanvas`](https://pkg.go.dev/github.com/belluikit/bellui/x/dag/dagwidget#NewDAGCanvas)
shows a graph and lets the user work with it:

{{< example file="extensions/dag/canvas/main.go" region="canvas" >}}

`AutoLayout` runs the layered layout and moves the nodes, as one step the user
can undo.

What the user can do with the mouse and keyboard:

- **Pan** by dragging empty space; **zoom** with the scroll wheel, centred on
  the pointer.
- **Select** a node, edge, group or annotation by clicking it. Double-click or
  Enter **activates** a node or group (`OnNodeActivated`,
  `OnGroupActivated`), for example to open its settings.
- **Move** a node by dragging it, and **resize** it with its handle. Dropping a
  node into a group makes it a member.
- **Connect** by dragging from an output port to an input port. The pointer
  only snaps to ports where the edge would be valid. Dragging an edge's end
  away from its port detaches it.
- **Delete** the selection with Delete or Backspace. **Undo** and **redo**
  with the usual shortcuts.
- **Right-click** an element to fire its `On...SecondaryTapped` callback, which
  receives the pointer event so you can show a context menu there.

When you add or remove things from code, use the canvas's own methods
(`c.AddNode`, `c.RemoveNode`, `c.AddGroup`, and so on). They change the graph
and update the drawing together. After replacing the whole graph, call
`LoadGraph`.

### Read-only canvases

[`Capabilities`](https://pkg.go.dev/github.com/belluikit/bellui/x/dag/dagwidget#Capabilities)
decides which edits the user may make. A new canvas allows everything. The zero
value allows nothing, which suits a monitoring view: the user can still pan,
zoom and select.

{{< example file="extensions/dag/canvas/main.go" region="readonly" >}}

Each field covers one kind of edit, such as `ConnectEdges`, `MoveResize` or
`RemoveNodes`. Capabilities limit only the user. Your own code can still change
the graph.

### Node appearance

Set `c.NodeStyler` to a function that returns a
[`NodeStyle`](https://pkg.go.dev/github.com/belluikit/bellui/x/dag/dagwidget#NodeStyle)
for each node: an icon, a subtitle and an accent colour. Base it on the node's
`TypeTag` or `Metadata`. Call `RefreshNodeStyles` after changing it.

## Components, palette and properties

For an editor where users build graphs from a fixed set of building blocks,
describe each block as a
[`dag.Component`](https://pkg.go.dev/github.com/belluikit/bellui/x/dag/dag#Component):
its ports, the stream type each port carries, and its configurable properties.
Put them in a registry, and give the registry to the graph:

{{< example file="extensions/dag/editor/main.go" region="registry" >}}

With `g.Registry` set, an edge is refused unless the source port's stream type
matches the target's. The stream type named `*` matches anything.

`dagwidget` has two widgets that work with the registry. The **palette** lists
the components by category; the user clicks one, or drags it onto the canvas,
to add a node. The **properties panel** shows a form for the selected node's
properties, a group's label and colour, or an annotation's text. Edits made
there go through the canvas's undo stack.

{{< example file="extensions/dag/editor/main.go" region="panels" >}}

A node made from a component stores its property values as JSON in
`Node.Metadata`, starting from each property's `Default`.

## Licence

`x/dag` is BSD 3-Clause licensed, like Bellui. Its Protocol Buffers messages
are declared in the package `bellui.dag.v1`.
