// A small pipeline editor: a palette of components, the canvas, and a properties panel.
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

func main() {
	a := app.New()
	w := a.NewWindow("Pipeline editor")
	w.Resize(geo.NewSize(1000, 600))

	//docs:start registry
	text := dag.StreamType{Name: "text/line"}
	reg := dag.NewRegistry()
	reg.MustRegister(&dag.Component{
		ID: "file_source", Name: "File", Category: "Sources",
		Outputs:    []dag.PortDef{{Name: "lines", Direction: dag.PortOutput, Stream: text}},
		Properties: []dag.PropertyDef{{Key: "path", Label: "Path", Kind: dag.PropertyString, Default: "input.txt"}},
	})
	reg.MustRegister(&dag.Component{
		ID: "grep", Name: "Grep", Category: "Transforms",
		Inputs:     []dag.PortDef{{Name: "in", Direction: dag.PortInput, Stream: text}},
		Outputs:    []dag.PortDef{{Name: "out", Direction: dag.PortOutput, Stream: text}},
		Properties: []dag.PropertyDef{{Key: "pattern", Label: "Pattern", Kind: dag.PropertyString, Required: true}},
	})
	reg.MustRegister(&dag.Component{
		ID: "stdout", Name: "Print", Category: "Sinks",
		Inputs: []dag.PortDef{{Name: "in", Direction: dag.PortInput, Stream: text}},
	})

	g := dag.NewGraph("p1", "Untitled")
	g.Registry = reg // edges between incompatible stream types are refused
	//docs:end registry

	//docs:start panels
	c := dagwidget.NewDAGCanvas(g)
	palette := dagwidget.NewPalette(reg, c) // click or drag a component onto the canvas
	props := dagwidget.NewPropertiesPanel(reg, c.Commands)
	props.CanvasRef = c

	c.OnNodeSelected = func(id string) { props.SetNode(g.Nodes[id]) }
	c.OnGroupSelected = func(id string) { props.SetGroup(g.Groups[id]) }
	c.OnAnnotationSelected = func(id string) { props.SetAnnotation(g.Annotations[id]) }
	c.OnSelectionCleared = func() { props.SetNode(nil) }
	//docs:end panels

	toolbar := container.NewHBox(
		widget.NewButton("Undo", func() { c.Commands.Undo() }),
		widget.NewButton("Redo", func() { c.Commands.Redo() }),
		widget.NewButton("Group", c.AddGroupAtCenter),
		widget.NewButton("Note", c.AddAnnotationAtCenter),
		widget.NewButton("Arrange", func() { _ = c.AutoLayout() }),
	)

	split := container.NewHSplit(c, props)
	split.Offset = 0.75
	w.SetContent(container.NewBorder(toolbar, nil, palette, nil, split))
	bellui.ShowAndRun(w)
}
