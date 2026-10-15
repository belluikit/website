// Candlesticks with a moving-average line on the same axes.
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

// prices returns n made-up daily candles and their 5-day moving average.
func prices(n int) ([]series.CandlestickData, []float64) {
	candles := make([]series.CandlestickData, n)
	avg := make([]float64, n)
	price := 100.0
	for i := range n {
		open := price
		close := open + 3*math.Sin(float64(i)*0.7) + 1.2*math.Cos(float64(i)*1.9) + 0.4
		high := math.Max(open, close) + 1 + math.Abs(math.Sin(float64(i)))*2
		low := math.Min(open, close) - 1 - math.Abs(math.Cos(float64(i)))*2
		candles[i] = series.CandlestickData{Open: open, High: high, Low: low, Close: close}
		price = close

		sum, k := 0.0, 0
		for j := max(0, i-4); j <= i; j++ {
			sum += candles[j].Close
			k++
		}
		avg[i] = sum / float64(k)
	}
	return candles, avg
}

func main() {
	a := app.New()
	w := a.NewWindow("Candlestick series")
	w.Resize(geo.NewSize(640, 420))

	candles, avg := prices(30)
	days := make([]string, len(candles))
	for i := range days {
		days[i] = fmt.Sprintf("%d", i+1)
	}

	//docs:start candlestick
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "ACME", SubText: "Daily prices, 5-day average"},
		XAxis: []axis.Config{{Type: axis.TypeCategory, Data: days}},
		YAxis: []axis.Config{{Type: axis.TypeValue}},
		Series: []series.Series{
			&series.CandlestickSeries{
				SeriesName: "ACME",
				Data:       candles, // one Open/High/Low/Close per category
				BarWidth:   0.6,
			},
			&series.LineSeries{SeriesName: "MA5", Data: avg},
		},
	})
	//docs:end candlestick

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
