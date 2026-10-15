// A trading chart: candlesticks with two moving averages above a volume
// pane, a shared crosshair and zoom, and a toolbar of drawing tools.
package main

import (
	"fmt"
	"image/color"
	"math"
	"math/rand/v2"
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/charts/axis"
	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/component"
	"github.com/belluikit/bellui/x/charts/interact"
	"github.com/belluikit/bellui/x/charts/series"
)

// market returns n made-up trading days: candles, volumes and date labels.
func market(n int) ([]series.CandlestickData, []float64, []string) {
	r := rand.New(rand.NewPCG(7, 11))
	candles := make([]series.CandlestickData, n)
	volume := make([]float64, n)
	dates := make([]string, n)
	day := time.Date(2026, 3, 2, 0, 0, 0, 0, time.UTC)
	price := 180.0
	for i := range n {
		for day.Weekday() == time.Saturday || day.Weekday() == time.Sunday {
			day = day.AddDate(0, 0, 1)
		}
		dates[i] = day.Format("Jan 2")
		day = day.AddDate(0, 0, 1)

		open := price
		close := open * (1 + 0.004*math.Sin(float64(i)/9) + 0.018*(r.Float64()-0.5))
		high := math.Max(open, close) * (1 + 0.01*r.Float64())
		low := math.Min(open, close) * (1 - 0.01*r.Float64())
		candles[i] = series.CandlestickData{Open: open, High: high, Low: low, Close: close}
		volume[i] = 2 + 6*r.Float64()*math.Abs(close-open)/open*40 // millions
		price = close
	}
	return candles, volume, dates
}

//docs:start ma
// movingAverage returns the simple moving average of the closes over n days.
// The first n-1 values are NaN, which a line series leaves undrawn.
func movingAverage(candles []series.CandlestickData, n int) []float64 {
	out := make([]float64, len(candles))
	sum := 0.0
	for i, c := range candles {
		sum += c.Close
		if i >= n {
			sum -= candles[i-n].Close
		}
		if i < n-1 {
			out[i] = math.NaN()
		} else {
			out[i] = sum / float64(n)
		}
	}
	return out
}

//docs:end ma

func main() {
	a := app.New()
	w := a.NewWindow("Trading chart")
	w.Resize(geo.NewSize(1000, 680))

	candles, volume, dates := market(120)

	// A category axis labels every category, so keep one label in ten.
	sparse := make([]string, len(dates))
	for i := range dates {
		if i%10 == 0 {
			sparse[i] = dates[i]
		}
	}

	//docs:start panes
	// crosshair builds the hover behaviour both panes share: crosshair lines
	// with value tags on the axes, and a text box only when the user clicks.
	crosshair := func(yTag string, text func(i int) string) *component.TooltipConfig {
		return &component.TooltipConfig{
			Trigger:       component.TooltipTriggerAxis,
			Crosshair:     component.CrosshairBoth,
			CrosshairOnly: true, // no text box while hovering; click to pin one
			ShowAxisTags:  true, // value tags where the crosshair meets the axes
			AxisTagFormatter: func(x, y float64) (string, string) {
				i := int(math.Round(x))
				if i < 0 || i >= len(dates) {
					return "", ""
				}
				return dates[i], fmt.Sprintf(yTag, y)
			},
			Formatter: func(hits []interact.HitResult) string {
				return text(hits[0].DataIndex)
			},
		}
	}
	priceTip := crosshair("%.2f", func(i int) string {
		c := candles[i]
		return fmt.Sprintf("%s\nO %.2f  H %.2f\nL %.2f  C %.2f", dates[i], c.Open, c.High, c.Low, c.Close)
	})
	volTip := crosshair("%.1fM", func(i int) string {
		return fmt.Sprintf("%s\nVolume %.1fM", dates[i], volume[i])
	})
	grid := func(top, bottom float32) *chart.Grid {
		// Same left and right margins in every pane, so the plots line up.
		return &chart.Grid{Left: 64, Right: 24, Top: top, Bottom: bottom}
	}

	// Axis tooltips only find line-like series, not candles or bars, so each
	// pane carries an invisible line through the values to hover on.
	hidden := series.LineStyle{Color: color.Transparent}
	closes := make([]float64, len(candles))
	for i, c := range candles {
		closes[i] = c.Close
	}

	price := chart.New(chart.Option{
		Title:    &chart.Title{Text: "ACME", SubText: "Daily, 10- and 30-day moving averages"},
		Tooltip:  priceTip,
		DataZoom: &interact.ZoomConfig{Enabled: true, MinSpan: 0.1, MaxSpan: 1, ZoomSensitivity: 0.08},
		Grid:     grid(12, 30),
		Color: []color.Color{
			color.NRGBA{R: 120, G: 120, B: 120, A: 255}, // legend swatch for the candles
			color.Transparent,
			color.NRGBA{R: 255, G: 167, B: 38, A: 255}, // MA10
			color.NRGBA{R: 66, G: 165, B: 245, A: 255}, // MA30
		},
		XAxis: []axis.Config{{Type: axis.TypeCategory, Data: make([]string, len(dates))}},
		YAxis: []axis.Config{{Type: axis.TypeValue, SplitCount: 5}},
		Series: []series.Series{
			&series.CandlestickSeries{SeriesName: "ACME", Data: candles, BarWidth: 0.7},
			&series.LineSeries{Data: closes, LineStyle: hidden}, // no name: not in the legend
			&series.LineSeries{SeriesName: "MA10", Data: movingAverage(candles, 10)},
			&series.LineSeries{SeriesName: "MA30", Data: movingAverage(candles, 30)},
		},
	})

	vol := chart.New(chart.Option{
		Tooltip: volTip,
		// Enabled: false gives the pane a zoom window that follows the group,
		// without its own zoom buttons.
		DataZoom: &interact.ZoomConfig{Enabled: false, MinSpan: 0.1, MaxSpan: 1},
		Grid:     grid(6, 26),
		XAxis:    []axis.Config{{Type: axis.TypeCategory, Data: sparse}},
		YAxis:    []axis.Config{{Type: axis.TypeValue, SplitCount: 2}},
		Series: []series.Series{
			&series.BarSeries{Data: volume, BarWidth: 0.7},
			&series.LineSeries{Data: volume, LineStyle: hidden},
		},
	})

	group := chart.NewChartGroup(
		chart.Pane{Chart: price, Height: 3},
		chart.Pane{Chart: vol, Height: 1},
	)
	//docs:end panes

	status := widget.NewLabel("Crosshair")

	//docs:start tools
	// setTooltip swaps the price pane's tooltip, keeping the rest of its option.
	setTooltip := func(cfg *component.TooltipConfig) {
		opt := price.Option()
		opt.Tooltip = cfg
		price.SetOption(opt)
	}
	noTooltip := &component.TooltipConfig{Trigger: component.TooltipTriggerNone}

	// Crosshair: plain viewing. Select: click a drawing to select it, drag it
	// to move it, press Delete to remove it. A click on a chart that shows a
	// tooltip pins the tooltip instead of selecting, so Select turns it off.
	crosshairMode := func() {
		price.SetDrawingTool(interact.DrawingToolNone)
		setTooltip(priceTip)
		status.SetText("Crosshair")
	}
	selectMode := func() {
		price.SetDrawingTool(interact.DrawingToolNone)
		setTooltip(noTooltip)
		status.SetText("Select: click a drawing, drag to move, Delete to remove")
	}
	bar := container.NewHBox(
		widget.NewButton("Crosshair", crosshairMode),
		widget.NewButton("Select", selectMode),
	)
	tools := []struct {
		label string
		tool  interact.DrawingToolMode
	}{
		{"H-line", interact.DrawingToolHLine},
		{"V-line", interact.DrawingToolVLine},
		{"Trend", interact.DrawingToolTrendLine},
		{"Ray", interact.DrawingToolRay},
		{"Rect", interact.DrawingToolRectangle},
		{"Channel", interact.DrawingToolChannel},
		{"Fib", interact.DrawingToolFibRetracement},
		{"Text", interact.DrawingToolTextLabel},
		{"Measure", interact.DrawingToolMeasure},
	}
	for _, t := range tools {
		bar.Add(widget.NewButton(t.label, func() {
			setTooltip(priceTip)
			price.SetDrawingTool(t.tool)
			status.SetText("Tool: " + t.label)
		}))
	}
	bar.Add(widget.NewButton("Clear", price.ClearDrawings))
	//docs:end tools

	//docs:start events
	price.Events.On(interact.EventDrawingCreated, func(evt interface{}) {
		e := evt.(interact.DrawingEvent)
		// Tools stay active after a drawing. The new drawing is selected,
		// so switch to Select mode to let the user move or delete it.
		selectMode()
		status.SetText(fmt.Sprintf("Created %s at %v", e.DrawingID, e.Drawing.DataPoints()))
	})
	price.Events.On(interact.EventDrawingSelected, func(evt interface{}) {
		status.SetText(fmt.Sprintf("Selected %s", evt.(interact.DrawingEvent).DrawingID))
	})
	price.Events.On(interact.EventDrawingDeleted, func(evt interface{}) {
		status.SetText(fmt.Sprintf("Deleted %s", evt.(interact.DrawingEvent).DrawingID))
	})
	//docs:end events

	w.SetContent(container.NewBorder(container.NewVBox(bar, status), nil, nil, nil, group))
	bellui.ShowAndRun(w)
}
