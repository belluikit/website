// Small charts inside cells.
package main

import (
	"image/color"
	"math"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/tables"
	"github.com/belluikit/bellui/x/tables/sparkline"
)

func main() {
	a := app.New()
	w := a.NewWindow("Sparklines")
	w.Resize(geo.NewSize(600, 400))

	var rows []map[string]any
	for i := range 20 {
		trend := make([]float64, 14)
		change := make([]float64, 7)
		for j := range trend {
			trend[j] = 100 + 10*math.Sin(float64(i+j)/3)
		}
		for j := range change {
			change[j] = 5 * math.Cos(float64(i*j))
		}
		rows = append(rows, map[string]any{"name": string(rune('A' + i)), "trend": trend, "change": change})
	}

	//docs:start sparklines
	blue := color.NRGBA{R: 33, G: 150, B: 243, A: 255}
	grid := tables.NewDataGrid(
		[]*tables.ColumnDef{
			{Field: "name", HeaderName: "Name", Width: 80},
			{
				// The cell value must be a []float64.
				Field: "trend", HeaderName: "Trend", Width: 160,
				CellRenderer: sparkline.NewCellRenderer(sparkline.Area,
					sparkline.WithStroke(blue),
					sparkline.WithFill(color.NRGBA{R: 33, G: 150, B: 243, A: 50}),
				),
			},
			{
				Field: "change", HeaderName: "Daily change", Width: 160,
				CellRenderer: sparkline.NewCellRenderer(sparkline.Column,
					sparkline.WithStroke(blue),
					sparkline.WithAxis(),
					func(o *sparkline.Options) { o.NegativeColor = color.NRGBA{R: 244, G: 67, B: 54, A: 255} },
				),
			},
		},
		func() int { return len(rows) },
		func(row int) any { return rows[row] },
	)
	//docs:end sparklines

	w.SetContent(grid)
	bellui.ShowAndRun(w)
}
