// The graph model on its own: nodes with ports, validated edges, ordering and saving.
package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/dag/dag"
)

func main() {
	a := app.New()
	w := a.NewWindow("Graph model")

	var out strings.Builder

	//docs:start model
	g := dag.NewGraph("etl", "Nightly import")

	in := func(name string) dag.Port { return dag.Port{Name: name, Direction: dag.PortInput} }
	outp := func(name string) dag.Port { return dag.Port{Name: name, Direction: dag.PortOutput} }

	_ = g.AddNode(&dag.Node{ID: "read", Label: "Read CSV", Inputs: []dag.Port{in("retry")}, Outputs: []dag.Port{outp("rows")}})
	_ = g.AddNode(&dag.Node{ID: "clean", Label: "Clean", Inputs: []dag.Port{in("in")}, Outputs: []dag.Port{outp("out")}})
	_ = g.AddNode(&dag.Node{ID: "load", Label: "Load", Inputs: []dag.Port{in("in")}, Outputs: []dag.Port{outp("done")}})

	// An edge runs from an output port to an input port. A port is found by
	// node and name, so the ports of one node need different names.
	_ = g.AddEdge(&dag.Edge{ID: "e1",
		Source: dag.PortID{NodeID: "read", Name: "rows"},
		Target: dag.PortID{NodeID: "clean", Name: "in"}})
	_ = g.AddEdge(&dag.Edge{ID: "e2",
		Source: dag.PortID{NodeID: "clean", Name: "out"},
		Target: dag.PortID{NodeID: "load", Name: "in"}})

	// The graph refuses an edge that would close a loop.
	err := g.AddEdge(&dag.Edge{ID: "e3",
		Source: dag.PortID{NodeID: "load", Name: "done"},
		Target: dag.PortID{NodeID: "read", Name: "retry"}})
	fmt.Fprintln(&out, "loop refused:", errors.Is(err, dag.ErrCycleDetected))
	//docs:end model

	//docs:start order
	order, err := g.TopologicalOrder()
	if err == nil {
		for _, n := range order {
			fmt.Fprintln(&out, "run", n.Label)
		}
	}
	//docs:end order

	//docs:start save
	data, err := dag.MarshalGraph(g) // PXF text
	if err != nil {
		panic(err)
	}
	g2, err := dag.UnmarshalGraph(data)
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(&out, "reloaded %d nodes and %d edges\n", len(g2.Nodes), len(g2.Edges))
	//docs:end save

	w.SetContent(widget.NewLabel(out.String()))
	bellui.ShowAndRun(w)
}
