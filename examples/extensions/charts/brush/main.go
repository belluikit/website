// Brush selection: drag a rectangle over a scatter chart to select points.
package main

import (
	"fmt"
	"math"
	"math/rand/v2"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/charts/axis"
	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/interact"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Brush")
	w.Resize(geo.NewSize(900, 560))

	// x, y pairs: height in cm and weight in kg.
	r := rand.New(rand.NewPCG(3, 5))
	people := make([]float64, 0, 400)
	for range 200 {
		h := 150 + 45*r.Float64()
		people = append(people, math.Round(h), math.Round(h*0.9-90+12*r.NormFloat64()))
	}

	//docs:start brush
	c := chart.New(chart.Option{
		Title:  &chart.Title{Text: "Height and weight", SubText: "Drag to select"},
		Brush:  &interact.BrushConfig{Enabled: true},
		Grid:   &chart.Grid{Left: 56, Right: 20, Top: 12, Bottom: 30},
		XAxis:  []axis.Config{{Type: axis.TypeValue}},
		YAxis:  []axis.Config{{Type: axis.TypeValue}},
		Series: []series.Series{&series.ScatterSeries{SeriesName: "People", Data: people, SymbolSize: 7}},
	})

	result := widget.NewLabel("Drag a rectangle over the points.")
	c.Events.On(interact.EventBrushSelect, func(evt interface{}) {
		e := evt.(interact.BrushSelectEvent)
		if len(e.Selected) == 0 {
			result.SetText("Nothing selected")
			return
		}
		sum := 0.0
		for _, p := range e.Selected {
			sum += p.Value.Y // p.SeriesIndex and p.DataIndex locate the point
		}
		result.SetText(fmt.Sprintf("%d people selected, mean weight %.1f kg",
			len(e.Selected), sum/float64(len(e.Selected))))
	})
	//docs:end brush

	w.SetContent(container.NewBorder(result, nil, nil, nil, c))
	bellui.ShowAndRun(w)
}
