// A node-link graph with fixed node positions and weighted edges.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Graph series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start graph
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Service map", SubText: "Edge width is requests per second"},
		Series: []series.Series{
			&series.GraphSeries{
				SeriesName: "Services",
				NodeSize:   10,
				// X and Y are fractions of the plot area, from 0 to 1.
				Nodes: []series.GraphNode{
					{Name: "gateway", X: 0.05, Y: 0.5},
					{Name: "auth", X: 0.35, Y: 0.1},
					{Name: "catalog", X: 0.35, Y: 0.5},
					{Name: "orders", X: 0.35, Y: 0.9},
					{Name: "search", X: 0.7, Y: 0.3},
					{Name: "postgres", X: 0.95, Y: 0.6},
					{Name: "queue", X: 0.7, Y: 0.95},
				},
				Edges: []series.GraphEdge{
					{Source: 0, Target: 1, Value: 2},
					{Source: 0, Target: 2, Value: 5},
					{Source: 0, Target: 3, Value: 3},
					{Source: 2, Target: 4, Value: 4},
					{Source: 2, Target: 5, Value: 3},
					{Source: 3, Target: 5, Value: 2},
					{Source: 3, Target: 6, Value: 1},
				},
			},
		},
	})
	//docs:end graph

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
