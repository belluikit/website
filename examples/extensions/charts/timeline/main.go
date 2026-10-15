// A timeline: step or play through one chart per year.
package main

import (
	"fmt"
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/charts/axis"
	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/component"
	"github.com/belluikit/bellui/x/charts/series"
)

var (
	years    = []string{"2021", "2022", "2023", "2024", "2025"}
	browsers = []string{"Chrome", "Safari", "Edge", "Firefox", "Other"}
	share    = [][]float64{
		{64, 19, 4, 4, 9},
		{65, 19, 4, 3, 9},
		{63, 20, 5, 3, 9},
		{65, 18, 5, 3, 9},
		{67, 17, 5, 2, 9},
	}
)

// optionFor builds the chart for one step of the timeline.
func optionFor(i int) chart.Option {
	zero, hundred := 0.0, 100.0
	return chart.Option{
		Title:     &chart.Title{Text: "Browser share, " + years[i], SubText: "Percent of page views"},
		Animation: true, // SetOption morphs from the previous year
		Grid:      &chart.Grid{Left: 56, Right: 20, Top: 12, Bottom: 30},
		XAxis:     []axis.Config{{Type: axis.TypeCategory, Data: browsers}},
		YAxis:     []axis.Config{{Type: axis.TypeValue, Min: &zero, Max: &hundred}},
		Series:    []series.Series{&series.BarSeries{SeriesName: "Share", Data: share[i], CornerRadius: 4}},
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Timeline")
	w.Resize(geo.NewSize(900, 560))

	c := chart.New(optionFor(0))
	year := widget.NewLabel(years[0])

	//docs:start timeline
	tl := component.NewTimelineState(len(years),
		component.TimelineConfig{Interval: 1500 * time.Millisecond, Loop: true},
		func(i int) {
			// Play calls this from its own goroutine: hop to the UI thread.
			bellui.Do(func() {
				c.SetOption(optionFor(i))
				year.SetText(fmt.Sprintf("%s (%d of %d)", years[i], i+1, len(years)))
			})
		})

	controls := container.NewHBox(
		widget.NewButton("Previous", func() { tl.Stop(); tl.Prev() }),
		widget.NewButton("Play", tl.Play),
		widget.NewButton("Stop", tl.Stop),
		widget.NewButton("Next", func() { tl.Stop(); tl.Next() }),
		year,
	)
	//docs:end timeline

	w.SetContent(container.NewBorder(controls, nil, nil, nil, c))
	bellui.ShowAndRun(w)
}
