// Reference lines, bands and an annotation drawn over a line chart.
package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/axis"
	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Overlays")
	w.Resize(geo.NewSize(900, 560))

	hours := make([]string, 48)
	latency := make([]float64, 48)
	for i := range hours {
		if i%6 == 0 { // label every sixth hour only
			hours[i] = fmt.Sprintf("%02d:00", i%24)
		}
		latency[i] = 120 + 35*math.Sin(float64(i)/4) + 10*math.Cos(float64(i)*1.7)
	}
	latency[31] = 290 // a spike

	//docs:start overlays
	red := color.NRGBA{R: 229, G: 57, B: 53, A: 220}
	overlays := []chart.Overlay{
		// Horizontal: constant Y, across the whole plot.
		&chart.ReferenceLine{
			Axis: chart.AxisHorizontal, Value: 250, Label: "SLO 250 ms",
			Color: red, Dashed: true, DashPattern: []float32{8, 4},
		},
		&chart.ReferenceBand{Axis: chart.AxisHorizontal, From: 80, To: 160, Label: "Normal"},
		// Vertical: constant X. On a category axis X is the category index.
		&chart.ReferenceBand{
			Axis: chart.AxisVertical, From: 19, To: 22, Label: "Maintenance",
			Color: color.NRGBA{R: 255, G: 167, B: 38, A: 50},
		},
		&chart.ReferenceLine{Axis: chart.AxisVertical, Value: 12, Label: "Deploy", Width: 2},
		// An annotation sits at an X, Y data point.
		&chart.Annotation{X: 31, Y: 290, Text: "Cache miss storm", Arrow: true},
	}
	//docs:end overlays

	c := chart.New(chart.Option{
		Title:    &chart.Title{Text: "API latency", SubText: "p95, last 48 hours"},
		Grid:     &chart.Grid{Left: 56, Right: 20, Top: 12, Bottom: 30},
		XAxis:    []axis.Config{{Type: axis.TypeCategory, Data: hours}},
		YAxis:    []axis.Config{{Type: axis.TypeValue}},
		Series:   []series.Series{&series.LineSeries{SeriesName: "p95 ms", Data: latency}},
		Overlays: overlays,
	})

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
