// A heatmap: a grid of values mapped to colours.
package main

import (
	"fmt"
	"image/color"
	"math"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Heatmap series")
	w.Resize(geo.NewSize(640, 420))

	// Commits per hour of the day (columns) for each weekday (rows).
	const rows, cols = 7, 24
	data := make([]float64, rows*cols)
	busy := []float64{0.8, 1, 1, 0.9, 0.6, 0.15, 0.1} // Monday to Sunday
	for r := range rows {
		weekday := r < 5
		for c := range cols {
			v := 2 * (1 + float64(r%3)/2) * math.Exp(-math.Pow(float64(c-21), 2)/8) // evening work
			if weekday {
				v += 10 * math.Exp(-math.Pow(float64(c-11), 2)/10)
				v += 8 * math.Exp(-math.Pow(float64(c-15), 2)/8)
			}
			data[r*cols+c] = v * (0.5 + busy[r])
		}
	}

	hours := make([]string, cols)
	for i := range hours {
		hours[i] = fmt.Sprintf("%02d:00", i)
	}

	//docs:start heatmap
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Commits by hour", SubText: "Monday (top) to Sunday, 00:00 to 23:00"},
		Series: []series.Series{
			&series.HeatmapSeries{
				SeriesName: "Commits",
				Rows:       rows,
				Cols:       cols,
				Data:       data, // row-major: Data[row*Cols+col]
				XLabels:    hours,
				YLabels:    []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
				ColorStops: []series.HeatmapColorStop{
					{Position: 0, Color: color.NRGBA{R: 235, G: 244, B: 250, A: 255}},
					{Position: 0.5, Color: color.NRGBA{R: 66, G: 146, B: 198, A: 255}},
					{Position: 1, Color: color.NRGBA{R: 8, G: 48, B: 107, A: 255}},
				},
			},
		},
	})
	//docs:end heatmap

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
