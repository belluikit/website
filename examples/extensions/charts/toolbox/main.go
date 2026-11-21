// The toolbox: line/bar switching, and saving the chart as a PNG.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/charts/axis"
	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/component"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Toolbox")
	w.Resize(geo.NewSize(900, 560))

	//docs:start toolbox
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Signups", SubText: "Switch between line and bar"},
		Toolbox: &component.ToolboxConfig{
			MagicType: []string{"line", "bar"}, // one button per type
		},
		Grid:  &chart.Grid{Left: 56, Right: 20, Top: 12, Bottom: 30},
		XAxis: []axis.Config{{Type: axis.TypeCategory, Data: []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}}},
		YAxis: []axis.Config{{Type: axis.TypeValue}},
		Series: []series.Series{
			&series.BarSeries{SeriesName: "Signups", Data: []float64{120, 200, 150, 80, 70, 110, 130}},
		},
	})
	//docs:end toolbox

	//docs:start save
	// The same as the toolbox's Save button: a save dialog, then a PNG of
	// the chart. Call it on the main goroutine, as a button does.
	save := widget.NewButton("Save PNG", c.SaveAsImage)
	//docs:end save

	w.SetContent(container.NewBorder(container.NewHBox(save), nil, nil, nil, c))
	bellui.ShowAndRun(w)
}
