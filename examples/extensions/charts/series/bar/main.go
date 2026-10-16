// Bar series: rounded bars with a negative value, and a line on the same axes.
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
	w := a.NewWindow("Bar series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start bar
	quarters := []string{"Q1 25", "Q2 25", "Q3 25", "Q4 25", "Q1 26", "Q2 26"}
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Operating profit", SubText: "Millions, with a four-quarter average"},
		XAxis: []axis.Config{{Type: axis.TypeCategory, Data: quarters}},
		YAxis: []axis.Config{{Type: axis.TypeValue}},
		Series: []series.Series{
			&series.BarSeries{
				SeriesName:   "Profit",
				Data:         []float64{4.2, 5.1, -1.8, 6.3, 7.0, 5.6},
				BarWidth:     0.5,
				CornerRadius: 4,
			},
			&series.LineSeries{
				SeriesName: "Average",
				Data:       []float64{3.1, 3.6, 2.9, 3.5, 4.2, 4.3},
				Symbol:     series.SymbolCircle,
			},
		},
	})
	//docs:end bar

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
