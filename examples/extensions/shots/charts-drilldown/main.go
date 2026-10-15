// Capture-only copy of charts/drilldown that opens one level down, so the
// screenshot shows the breadcrumb.
package main

import (
	"strings"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/charts/axis"
	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/interact"
	"github.com/belluikit/bellui/x/charts/series"
)

var quarterly = map[string][]float64{
	"Europe":   {310, 280, 350, 420},
	"Americas": {450, 470, 430, 510},
	"Asia":     {220, 260, 300, 340},
}

// region returns the child chart for one pie slice, or nil to not drill.
func region(hit interact.HitResult, path []string) *chart.Option {
	if len(path) > 0 {
		return nil // already one level down: the bars do not drill further
	}
	name := hit.Value.Label // a pie hit carries the slice name
	data, ok := quarterly[name]
	if !ok {
		return nil
	}
	return &chart.Option{
		Title:  &chart.Title{Text: name, SubText: "Revenue by quarter"},
		Grid:   &chart.Grid{Left: 56, Right: 20, Top: 12, Bottom: 30},
		XAxis:  []axis.Config{{Type: axis.TypeCategory, Data: []string{"Q1", "Q2", "Q3", "Q4"}}},
		YAxis:  []axis.Config{{Type: axis.TypeValue}},
		Series: []series.Series{&series.BarSeries{SeriesName: name, Data: data, CornerRadius: 4}},
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Drill-down")
	w.Resize(geo.NewSize(900, 560))

	slices := []series.PieDataItem{}
	for _, name := range []string{"Europe", "Americas", "Asia"} {
		total := 0.0
		for _, v := range quarterly[name] {
			total += v
		}
		slices = append(slices, series.PieDataItem{Name: name, Value: total})
	}
	c := chart.New(chart.Option{
		Title:     &chart.Title{Text: "Revenue by region", SubText: "Click a slice"},
		Series:    []series.Series{&series.PieSeries{SeriesName: "Revenue", Data: slices, ShowLabels: true}},
		Drilldown: &chart.DrilldownConfig{DataProvider: region},
	})

	where := widget.NewLabel("Top level")
	c.Events.OnFiltered(interact.EventFilter{
		EventTypes: []interact.EventType{interact.EventDrilldown, interact.EventDrillUp},
	}, func(evt interface{}) {
		e := evt.(interact.DrilldownEvent)
		where.SetText("Path: " + strings.Join(e.Path, " > "))
	})
	back := widget.NewButton("Back", c.DrillUp)
	w.SetContent(container.NewBorder(container.NewHBox(back, where), nil, nil, nil, c))
	hit := interact.HitResult{SeriesName: "Revenue", Value: interact.DataValue{Label: "Europe"}}
	c.DrillInto("Europe", *region(hit, nil))
	bellui.ShowAndRun(w)
}
