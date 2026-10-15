// Keep data in a dataset, filter it, and plot two of its columns.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/axis"
	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/dataset"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Dataset")
	w.Resize(geo.NewSize(700, 450))

	//docs:start dataset
	ds := dataset.New(
		dataset.Dimension{Name: "height", Type: dataset.DimFloat},
		dataset.Dimension{Name: "weight", Type: dataset.DimFloat},
		dataset.Dimension{Name: "name", Type: dataset.DimString},
	)
	_ = ds.AppendRow(161.2, 51.6, "Ana")
	_ = ds.AppendRow(167.5, 59.0, "Ben")
	_ = ds.AppendRow(159.5, 49.2, "Cleo")
	_ = ds.AppendRow(174.0, 69.8, "Dev")
	_ = ds.AppendRow(181.5, 74.3, "Eli")

	// Keep rows taller than 160, sorted by height.
	tall, err := (&dataset.Pipeline{Steps: []dataset.Transform{
		&dataset.FilterTransform{Dimension: 0, Predicate: func(h float64) bool { return h > 160 }},
		&dataset.SortTransform{Dimension: 0},
	}}).Apply(ds)
	if err != nil {
		panic(err)
	}

	xs, ys := dataset.ExtractXY(tall, dataset.DefaultEncode()) // column 0 → X, column 1 → Y
	points := make([]float64, 0, 2*len(xs))
	for i := range xs {
		points = append(points, xs[i], ys[i])
	}
	//docs:end dataset

	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Height and weight"},
		XAxis: []axis.Config{{Type: axis.TypeValue}},
		YAxis: []axis.Config{{Type: axis.TypeValue}},
		Series: []series.Series{&series.ScatterSeries{
			SeriesName: "People", Data: points, Symbol: series.SymbolCircle, SymbolSize: 10,
		}},
	})

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
