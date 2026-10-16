// A line chart: one category axis, one value axis, one series.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/axis"
	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/component"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Weekly sales")
	w.Resize(geo.NewSize(700, 450))

	c := chart.New(chart.Option{
		Title:   &chart.Title{Text: "Weekly sales", SubText: "Revenue per day"},
		Tooltip: &component.TooltipConfig{Trigger: component.TooltipTriggerAxis},
		XAxis: []axis.Config{{
			Type: axis.TypeCategory,
			Data: []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
		}},
		YAxis: []axis.Config{{Type: axis.TypeValue}},
		Series: []series.Series{
			&series.LineSeries{
				SeriesName: "Revenue",
				Data:       []float64{150, 230, 224, 218, 135, 147, 260},
				Symbol:     series.SymbolCircle,
				AreaStyle:  &series.AreaStyle{Opacity: 0.15},
			},
		},
		Animation: true,
	})

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
