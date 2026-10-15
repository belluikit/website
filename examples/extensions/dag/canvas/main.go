// The interactive canvas: show a graph, lay it out, react to selection.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/dag/dag"
	"github.com/belluikit/bellui/x/dag/dagwidget"
)

func pipeline() *dag.Graph {
	g := dag.NewGraph("build", "Build")
	port := func(name string, dir dag.PortDirection) []dag.Port {
		return []dag.Port{{Name: name, Direction: dir}}
	}
	for _, n := range []*dag.Node{
		{ID: "src", Label: "Checkout", Outputs: port("out", dag.PortOutput)},
		{ID: "test", Label: "Test", Inputs: port("in", dag.PortInput), Outputs: port("out", dag.PortOutput)},
		{ID: "lint", Label: "Lint", Inputs: port("in", dag.PortInput), Outputs: port("out", dag.PortOutput)},
		{ID: "ship", Label: "Ship", Inputs: []dag.Port{{Name: "tests", Direction: dag.PortInput}, {Name: "lint", Direction: dag.PortInput}}},
	} {
		_ = g.AddNode(n)
	}
	edge := func(id, from, fromPort, to, toPort string) {
		_ = g.AddEdge(&dag.Edge{ID: id,
			Source: dag.PortID{NodeID: from, Name: fromPort},
			Target: dag.PortID{NodeID: to, Name: toPort}})
	}
	edge("e1", "src", "out", "test", "in")
	edge("e2", "src", "out", "lint", "in")
	edge("e3", "test", "out", "ship", "tests")
	edge("e4", "lint", "out", "ship", "lint")
	return g
}

func main() {
	a := app.New()
	w := a.NewWindow("DAG canvas")
	w.Resize(geo.NewSize(900, 500))

	//docs:start canvas
	g := pipeline() // nodes and edges, all at position (0, 0)
	c := dagwidget.NewDAGCanvas(g)
	if dag.NeedsLayout(g) {
		_ = c.AutoLayout() // left-to-right layers
	}

	status := widget.NewLabel("Click a node.")
	c.OnNodeSelected = func(id string) { status.SetText("Selected " + g.Nodes[id].Label) }
	c.OnGraphChanged = func() { status.SetText("The graph changed.") }
	//docs:end canvas

	//docs:start readonly
	editable := widget.NewCheck("Editable", func(on bool) {
		if on {
			c.Capabilities = dagwidget.AllCapabilities()
		} else {
			c.Capabilities = dagwidget.Capabilities{} // read-only
		}
		c.Refresh()
	})
	editable.SetChecked(true)
	//docs:end readonly

	w.SetContent(container.NewBorder(container.NewHBox(editable, status), nil, nil, nil, c))
	bellui.ShowAndRun(w)
}
