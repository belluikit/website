// A funnel chart: stages of a process, each narrower than the last.
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
	w := a.NewWindow("Funnel series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start funnel
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Checkout funnel", SubText: "Visitors per stage"},
		Series: []series.Series{
			&series.FunnelSeries{
				SeriesName: "Checkout",
				// Stages are drawn top to bottom in this order.
				Data: []series.FunnelData{
					{Name: "Product page", Value: 10000},
					{Name: "Added to cart", Value: 4200},
					{Name: "Checkout", Value: 2600},
					{Name: "Payment", Value: 1900},
					{Name: "Order placed", Value: 1700},
				},
			},
		},
	})
	//docs:end funnel

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
