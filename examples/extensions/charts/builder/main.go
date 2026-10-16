// The fluent builder: the same Option, written as a chain of calls.
package main

import (
	"image/color"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/builder"
	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Builder")
	w.Resize(geo.NewSize(700, 450))

	//docs:start builder
	c := builder.LineChart().
		Title("Sales and cost").
		CategoryX("Q1", "Q2", "Q3", "Q4").
		ValueY().
		Legend().
		Zoom().
		Palette(
			color.NRGBA{R: 255, G: 99, B: 132, A: 255},
			color.NRGBA{R: 54, G: 162, B: 235, A: 255},
		).
		AddLine("Sales", 40, 65, 50, 80).Symbol(series.SymbolCircle).Smooth().Done().
		AddLine("Cost", 30, 45, 35, 50).AreaFill(0.2).Done().
		ReferenceLine(chart.AxisHorizontal, 60).Label("Average").Dashed().Done().
		Animate().
		Widget()
	//docs:end builder

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
