// Zoom, a crosshair tooltip, brush selection and click events.
package main

import (
	"fmt"
	"math"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/charts/axis"
	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/component"
	"github.com/belluikit/bellui/x/charts/interact"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Interaction")
	w.Resize(geo.NewSize(800, 500))

	data := make([]float64, 200)
	for i := range data {
		data[i] = 50 + 30*math.Sin(float64(i)/12)
	}

	//docs:start zoom
	c := chart.New(chart.Option{
		Tooltip: &component.TooltipConfig{
			Trigger:   component.TooltipTriggerAxis,
			Crosshair: component.CrosshairBoth,
		},
		DataZoom: &interact.ZoomConfig{Enabled: true, MinSpan: 0.05, MaxSpan: 1, ZoomSensitivity: 0.1},
		XAxis:    []axis.Config{{Type: axis.TypeValue}},
		YAxis:    []axis.Config{{Type: axis.TypeValue}},
		Series:   []series.Series{&series.LineSeries{SeriesName: "Signal", Data: data}},
	})
	//docs:end zoom

	status := widget.NewLabel("Scroll to zoom, drag to pan, click a point.")

	//docs:start events
	c.Events.On(interact.EventClick, func(evt interface{}) {
		e := evt.(interact.DataPointEvent)
		status.SetText(fmt.Sprintf("Clicked %s, point %d: %.1f", e.SeriesName, e.DataIndex, e.Value.Y))
	})
	c.Events.On(interact.EventZoomChange, func(evt interface{}) {
		e := evt.(interact.ZoomChangeEvent)
		status.SetText(fmt.Sprintf("Showing %.0f%% to %.0f%% of the data", e.Start*100, e.End*100))
	})
	//docs:end events

	reset := widget.NewButton("Show all", func() {
		c.SetZoomWindow(interact.FullWindow())
	})

	//docs:start brush
	scatter := chart.New(chart.Option{
		Brush: &interact.BrushConfig{Enabled: true},
		XAxis: []axis.Config{{Type: axis.TypeValue}},
		YAxis: []axis.Config{{Type: axis.TypeValue}},
		Series: []series.Series{&series.ScatterSeries{
			SeriesName: "Samples",
			// x, y pairs
			Data:       []float64{10, 8.0, 8, 6.9, 13, 7.6, 9, 8.8, 11, 8.3, 14, 9.9, 6, 7.2, 4, 4.3, 12, 10.8},
			Symbol:     series.SymbolCircle,
			SymbolSize: 10,
		}},
	})
	scatter.Events.On(interact.EventBrushSelect, func(evt interface{}) {
		e := evt.(interact.BrushSelectEvent)
		status.SetText(fmt.Sprintf("Brushed %d points", len(e.Selected)))
	})
	//docs:end brush

	w.SetContent(container.NewBorder(
		container.NewHBox(reset, status), nil, nil, nil,
		container.NewGridWithColumns(2, c, scatter),
	))
	bellui.ShowAndRun(w)
}
