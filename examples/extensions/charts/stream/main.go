// Live data: append values from a goroutine and keep the newest in view.
package main

import (
	"math/rand/v2"
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/axis"
	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/interact"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Live data")
	w.Resize(geo.NewSize(800, 450))

	//docs:start stream
	c := chart.New(chart.Option{
		DataZoom: &interact.ZoomConfig{Enabled: true, MinSpan: 0.05, MaxSpan: 1},
		XAxis:    []axis.Config{{Type: axis.TypeValue}},
		YAxis:    []axis.Config{{Type: axis.TypeValue}},
		Series:   []series.Series{&series.LineSeries{SeriesName: "Price", Data: []float64{100}}},
	})
	c.SetFollowLatest(true) // snap the zoom window to the newest data

	go func() {
		price := 100.0
		for range time.Tick(200 * time.Millisecond) {
			price += rand.Float64()*2 - 1
			c.AppendData(0, price) // safe from any goroutine
		}
	}()
	//docs:end stream

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
