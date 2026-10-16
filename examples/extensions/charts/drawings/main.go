// Drawings created from code, saved to JSON and restored.
package main

import (
	"encoding/json"
	"fmt"
	"image/color"
	"math"
	"math/rand/v2"

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

// candles returns n made-up daily candles.
func candles(n int) []series.CandlestickData {
	r := rand.New(rand.NewPCG(7, 11))
	out := make([]series.CandlestickData, n)
	price := 180.0
	for i := range out {
		open := price
		close := open * (1 + 0.004*math.Sin(float64(i)/9) + 0.018*(r.Float64()-0.5))
		high := math.Max(open, close) * (1 + 0.01*r.Float64())
		low := math.Min(open, close) * (1 - 0.01*r.Float64())
		out[i] = series.CandlestickData{Open: open, High: high, Low: low, Close: close}
		price = close
	}
	return out
}

// extreme returns the index and value of the highest high (or lowest low)
// among data[from:to].
func extreme(data []series.CandlestickData, from, to int, highest bool) interact.DataPoint {
	best := interact.DataPoint{X: float64(from), Y: data[from].Low}
	if highest {
		best.Y = data[from].High
	}
	for i := from; i < to; i++ {
		if highest && data[i].High > best.Y {
			best = interact.DataPoint{X: float64(i), Y: data[i].High}
		}
		if !highest && data[i].Low < best.Y {
			best = interact.DataPoint{X: float64(i), Y: data[i].Low}
		}
	}
	return best
}

//docs:start save
// saved is one drawing in a form that encoding/json can write and read.
type saved struct {
	Tool   interact.DrawingToolMode
	Points []interact.DataPoint
	Label  string      `json:",omitempty"`
	Text   string      `json:",omitempty"` // text labels only
	Color  *color.NRGBA `json:",omitempty"`
}

func save(drawings []interact.Drawing) ([]byte, error) {
	out := make([]saved, 0, len(drawings))
	for _, d := range drawings {
		s := saved{Tool: d.Type(), Points: d.DataPoints(), Label: d.DrawingLabel()}
		if t, ok := d.(*interact.TextLabelDrawing); ok {
			s.Text = t.Text
		}
		if c := d.DrawingColor(); c != nil {
			n := color.NRGBAModel.Convert(c).(color.NRGBA)
			s.Color = &n
		}
		out = append(out, s)
	}
	return json.Marshal(out)
}

//docs:end save

//docs:start restore
func restore(data []byte) ([]interact.Drawing, error) {
	var in []saved
	if err := json.Unmarshal(data, &in); err != nil {
		return nil, err
	}
	var out []interact.Drawing
	for i, s := range in {
		// Use IDs of your own: the chart numbers the drawings the user makes
		// "drawing-1", "drawing-2"... and does not skip IDs already in use.
		base := interact.DrawingBase{ID_: interact.DrawingID(fmt.Sprintf("saved-%d", i+1)), Label_: s.Label}
		if s.Color != nil {
			base.Color_ = *s.Color
		}
		var d interact.Drawing
		switch s.Tool {
		case interact.DrawingToolHLine:
			d = &interact.HLineDrawing{DrawingBase: base}
		case interact.DrawingToolVLine:
			d = &interact.VLineDrawing{DrawingBase: base}
		case interact.DrawingToolTrendLine:
			d = &interact.TrendLineDrawing{DrawingBase: base}
		case interact.DrawingToolRay:
			d = &interact.RayDrawing{DrawingBase: base}
		case interact.DrawingToolRectangle:
			d = &interact.RectangleDrawing{DrawingBase: base}
		case interact.DrawingToolChannel:
			d = &interact.ChannelDrawing{DrawingBase: base}
		case interact.DrawingToolFibRetracement:
			d = &interact.FibRetracementDrawing{DrawingBase: base}
		case interact.DrawingToolTextLabel:
			d = &interact.TextLabelDrawing{DrawingBase: base, Text: s.Text}
		default:
			continue
		}
		d.SetDataPoints(s.Points)
		out = append(out, d)
	}
	return out, nil
}

//docs:end restore

func main() {
	a := app.New()
	w := a.NewWindow("Chart drawings")
	w.Resize(geo.NewSize(900, 600))

	data := candles(120)
	c := chart.New(chart.Option{
		Title:   &chart.Title{Text: "ACME", SubText: "Drawings added from code"},
		Tooltip: &component.TooltipConfig{Trigger: component.TooltipTriggerNone},
		Grid:    &chart.Grid{Left: 64, Right: 24, Top: 12, Bottom: 30},
		XAxis:   []axis.Config{{Type: axis.TypeCategory, Data: make([]string, len(data))}},
		YAxis:   []axis.Config{{Type: axis.TypeValue, SplitCount: 5}},
		Series:  []series.Series{&series.CandlestickSeries{SeriesName: "ACME", Data: data, BarWidth: 0.7}},
	})

	//docs:start add
	// X is the category index (0 is the first candle), Y is the price.
	high1, high2 := extreme(data, 0, 40, true), extreme(data, 70, 100, true)
	low1, low2 := extreme(data, 40, 70, false), extreme(data, 100, 120, false)
	orange := color.NRGBA{R: 255, G: 152, B: 0, A: 230}
	blue := color.NRGBA{R: 66, G: 165, B: 245, A: 230}

	c.AddDrawing(&interact.HLineDrawing{
		DrawingBase: interact.DrawingBase{ID_: "resistance", Label_: "Resistance"},
		Y:           high1.Y,
	})
	c.AddDrawing(&interact.TrendLineDrawing{
		DrawingBase: interact.DrawingBase{ID_: "highs", Color_: orange},
		Start:       high1,
		End:         high2,
	})
	c.AddDrawing(&interact.RayDrawing{
		DrawingBase: interact.DrawingBase{ID_: "lows", Color_: blue},
		Start:       low1,
		Direction:   low2, // the ray runs from Start through this point
	})
	c.AddDrawing(&interact.RectangleDrawing{
		DrawingBase: interact.DrawingBase{ID_: "base", Color_: blue},
		Corner1:     interact.DataPoint{X: 45, Y: extreme(data, 45, 65, true).Y},
		Corner2:     interact.DataPoint{X: 65, Y: extreme(data, 45, 65, false).Y},
	})
	c.AddDrawing(&interact.FibRetracementDrawing{
		DrawingBase: interact.DrawingBase{ID_: "fib"},
		HighY:       high2.Y,
		LowY:        low2.Y,
		X:           high2.X,
	})
	c.AddDrawing(&interact.TextLabelDrawing{
		DrawingBase: interact.DrawingBase{ID_: "note", Color_: orange},
		Pos:         interact.DataPoint{X: high2.X - 6, Y: high2.Y + 1.5},
		Text:        "Lower high",
	})
	//docs:end add

	status := widget.NewLabel(fmt.Sprintf("%d drawings", len(c.Drawings())))
	var stored []byte

	//docs:start buttons
	saveBtn := widget.NewButton("Save", func() {
		b, err := save(c.Drawings())
		if err != nil {
			status.SetText(err.Error())
			return
		}
		stored = b // a real app writes this to a file or to preferences
		status.SetText(fmt.Sprintf("Saved %d drawings, %d bytes of JSON", len(c.Drawings()), len(b)))
	})
	clearBtn := widget.NewButton("Clear", func() {
		c.ClearDrawings()
		status.SetText("Cleared")
	})
	restoreBtn := widget.NewButton("Restore", func() {
		ds, err := restore(stored)
		if err != nil {
			status.SetText(err.Error())
			return
		}
		c.ClearDrawings()
		for _, d := range ds {
			c.AddDrawing(d)
		}
		status.SetText(fmt.Sprintf("Restored %d drawings", len(ds)))
	})
	//docs:end buttons

	w.SetContent(container.NewBorder(
		container.NewHBox(saveBtn, clearBtn, restoreBtn, status), nil, nil, nil, c))
	bellui.ShowAndRun(w)
}
