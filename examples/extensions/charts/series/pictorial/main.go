// A pictorial bar chart: bars drawn as stacks of symbols.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/axis"
	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Pictorial series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start pictorial
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Trees planted", SubText: "Hundreds, per district"},
		XAxis: []axis.Config{{Type: axis.TypeCategory, Data: []string{"North", "East", "South", "West", "Centre"}}},
		YAxis: []axis.Config{{Type: axis.TypeValue}},
		Series: []series.Series{
			&series.PictorialSeries{
				SeriesName: "Trees",
				Data:       []float64{18, 31, 24, 12, 27},
				Symbol:     series.SymbolCircle, // a column of circles
				SymbolSize: 14,
			},
		},
	})
	//docs:end pictorial

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
