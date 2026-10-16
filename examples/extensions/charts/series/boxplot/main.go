// A box-and-whisker chart: one five-number summary per category.
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
	w := a.NewWindow("Boxplot series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start boxplot
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Request latency", SubText: "Milliseconds, per region"},
		XAxis: []axis.Config{{Type: axis.TypeCategory, Data: []string{"eu-west", "eu-north", "us-east", "us-west", "ap-south"}}},
		YAxis: []axis.Config{{Type: axis.TypeValue}},
		Series: []series.Series{
			&series.BoxplotSeries{
				SeriesName: "Latency",
				BoxWidth:   0.4,
				Data: []series.BoxplotData{
					{Min: 18, Q1: 26, Median: 31, Q3: 38, Max: 55},
					{Min: 22, Q1: 30, Median: 34, Q3: 41, Max: 62},
					{Min: 35, Q1: 48, Median: 57, Q3: 66, Max: 90},
					{Min: 40, Q1: 52, Median: 60, Q3: 74, Max: 104},
					{Min: 61, Q1: 79, Median: 92, Q3: 110, Max: 151},
				},
			},
		},
	})
	//docs:end boxplot

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
