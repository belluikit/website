// An area chart: a line series with an AreaStyle fill underneath.
package main

import (
	"fmt"
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
	w := a.NewWindow("Area series")
	w.Resize(geo.NewSize(640, 420))

	hours := make([]string, 24)
	load := make([]float64, 24)
	for i := range hours {
		hours[i] = fmt.Sprintf("%02d", i)
		// A daily curve: quiet at night, busy in the afternoon.
		load[i] = 35 + 25*math.Sin(float64(i-7)*math.Pi/12) + 4*math.Sin(float64(i))
	}

	//docs:start area
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "CPU load", SubText: "Percent, hourly average"},
		XAxis: []axis.Config{{Type: axis.TypeCategory, Data: hours}},
		YAxis: []axis.Config{{Type: axis.TypeValue}},
		Series: []series.Series{
			&series.LineSeries{
				SeriesName: "Load",
				Data:       load,
				AreaStyle:  &series.AreaStyle{Opacity: 0.25},
			},
		},
	})
	//docs:end area

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
