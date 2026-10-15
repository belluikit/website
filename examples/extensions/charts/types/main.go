// Four series types side by side: stacked bars, a doughnut, candlesticks and a gauge.
package main

import (
	"image/color"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/axis"
	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/component"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Series types")
	w.Resize(geo.NewSize(900, 650))

	days := []string{"Mon", "Tue", "Wed", "Thu", "Fri"}

	//docs:start stacked
	stacked := chart.New(chart.Option{
		Title:  &chart.Title{Text: "Stacked bars"},
		Legend: &component.LegendConfig{Position: component.LegendBottom},
		XAxis:  []axis.Config{{Type: axis.TypeCategory, Data: days}},
		YAxis:  []axis.Config{{Type: axis.TypeValue}},
		Series: []series.Series{
			&series.BarSeries{SeriesName: "Email", Data: []float64{120, 132, 101, 134, 90}, Stack: "total"},
			&series.BarSeries{SeriesName: "Ads", Data: []float64{220, 182, 191, 234, 290}, Stack: "total"},
		},
	})
	//docs:end stacked

	//docs:start pie
	doughnut := chart.New(chart.Option{
		Title: &chart.Title{Text: "Doughnut"},
		Series: []series.Series{
			&series.PieSeries{
				SeriesName:  "Browsers",
				Radius:      1.0,
				InnerRadius: 0.5,
				ShowLabels:  true,
				Data: []series.PieDataItem{
					{Name: "Chrome", Value: 65},
					{Name: "Safari", Value: 18},
					{Name: "Firefox", Value: 8},
					{Name: "Other", Value: 9},
				},
			},
		},
	})
	//docs:end pie

	//docs:start candles
	candles := chart.New(chart.Option{
		Title: &chart.Title{Text: "Candlestick"},
		XAxis: []axis.Config{{Type: axis.TypeCategory, Data: days}},
		YAxis: []axis.Config{{Type: axis.TypeValue}},
		Series: []series.Series{
			&series.CandlestickSeries{
				SeriesName: "Stock",
				Data: []series.CandlestickData{
					{Open: 100, High: 115, Low: 95, Close: 110},
					{Open: 110, High: 130, Low: 105, Close: 125},
					{Open: 125, High: 135, Low: 110, Close: 112},
					{Open: 112, High: 120, Low: 100, Close: 105},
					{Open: 105, High: 118, Low: 98, Close: 115},
				},
			},
		},
	})
	//docs:end candles

	//docs:start gauge
	gauge := chart.New(chart.Option{
		Title: &chart.Title{Text: "Gauge"},
		Series: []series.Series{
			&series.GaugeSeries{
				SeriesName: "Load",
				Value:      72.5, Min: 0, Max: 100,
				ShowValue: true,
				Sectors: []series.GaugeSector{
					{Min: 0, Max: 70, Color: color.NRGBA{R: 38, G: 186, B: 121, A: 120}},
					{Min: 70, Max: 100, Color: color.NRGBA{R: 239, G: 83, B: 80, A: 120}},
				},
			},
		},
	})
	//docs:end gauge

	w.SetContent(container.NewGridWithColumns(2, stacked, doughnut, candles, gauge))
	bellui.ShowAndRun(w)
}
