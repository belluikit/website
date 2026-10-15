// Zoom with the mouse wheel, the built-in buttons, a slider and from code.
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
	w := a.NewWindow("Zoom")
	w.Resize(geo.NewSize(900, 560))

	days := make([]string, 365)
	visits := make([]float64, 365)
	for i := range days {
		if i%30 == 0 {
			days[i] = fmt.Sprintf("Day %d", i+1)
		}
		visits[i] = 1000 + 400*math.Sin(float64(i)/58) + 150*math.Sin(float64(i)/3.5) + 2*float64(i)
	}

	//docs:start zoom
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Daily visits"},
		// Enabled turns on the wheel, the drag-to-pan and the + − Reset buttons.
		// MinSpan and MaxSpan bound the visible fraction of the data.
		DataZoom: &interact.ZoomConfig{Enabled: true, MinSpan: 0.05, MaxSpan: 1, ZoomSensitivity: 0.1},
		Tooltip:  &component.TooltipConfig{Trigger: component.TooltipTriggerAxis},
		Grid:     &chart.Grid{Left: 56, Right: 20, Top: 12, Bottom: 30},
		XAxis:    []axis.Config{{Type: axis.TypeCategory, Data: days}},
		YAxis:    []axis.Config{{Type: axis.TypeValue}},
		Series:   []series.Series{&series.LineSeries{SeriesName: "Visits", Data: visits}},
	})
	//docs:end zoom

	label := widget.NewLabel("")
	show := func(win interact.DataWindow) {
		first := int(win.Start * float64(len(days)-1))
		last := int(win.End * float64(len(days)-1))
		label.SetText(fmt.Sprintf("Showing day %d to day %d", first+1, last+1))
	}

	//docs:start slider
	slider := component.NewDataZoomSlider(func(win interact.DataWindow) {
		c.SetZoomWindow(win) // does not emit EventZoomChange, so no loop
		show(win)
	})
	c.Events.On(interact.EventZoomChange, func(evt interface{}) {
		e := evt.(interact.ZoomChangeEvent)
		slider.SetWindow(interact.DataWindow{Start: e.Start, End: e.End})
		show(interact.DataWindow{Start: e.Start, End: e.End})
	})
	//docs:end slider

	//docs:start code
	lastQuarter := widget.NewButton("Last 90 days", func() {
		win := interact.DataWindow{Start: 1 - 90.0/365, End: 1}
		c.SetZoomWindow(win)
		slider.SetWindow(win)
		show(win)
	})
	all := widget.NewButton("All", func() {
		c.SetZoomWindow(interact.FullWindow())
		slider.SetWindow(interact.FullWindow())
		show(interact.FullWindow())
	})
	//docs:end code

	start := interact.DataWindow{Start: 0.55, End: 0.85}
	c.SetZoomWindow(start)
	slider.SetWindow(start)
	show(start)

	w.SetContent(container.NewBorder(
		container.NewHBox(lastQuarter, all, label),
		container.NewPadded(slider), nil, nil, c))
	bellui.ShowAndRun(w)
}
