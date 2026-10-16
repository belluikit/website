// A star-rating widget built from scratch: five dots that fill up to the
// value, set by tapping.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/canvas"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/event"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

//docs:start widget
const maxRating = 5

// Rating shows a value from 0 to 5 as a row of dots.
// The widget holds the state; its renderer draws it.
type Rating struct {
	widget.BaseWidget

	Value     int
	OnChanged func(int)
}

// NewRating creates the widget. ExtendBaseWidget must be called before
// the widget is used, so BaseWidget knows which object it belongs to.
func NewRating(value int) *Rating {
	r := &Rating{Value: value}
	r.ExtendBaseWidget(r)
	return r
}

// SetValue changes the value and redraws the widget.
func (r *Rating) SetValue(v int) {
	if v < 0 || v > maxRating || v == r.Value {
		return
	}
	r.Value = v
	r.Refresh()
	if r.OnChanged != nil {
		r.OnChanged(v)
	}
}

// Tapped makes the widget respond to clicks and taps (event.Tappable).
func (r *Rating) Tapped(e *event.PointEvent) {
	dot := int(e.Position.X/(r.Size().Width/maxRating)) + 1
	if dot == r.Value {
		dot = 0 // tapping the current value clears it
	}
	r.SetValue(dot)
}

// CreateRenderer is called when the widget is first needed, and again if
// its cached renderer was destroyed.
func (r *Rating) CreateRenderer() core.WidgetRenderer {
	rr := &ratingRenderer{rating: r}
	for i := range rr.dots {
		rr.dots[i] = canvas.NewCircle(nil)
		rr.dots[i].StrokeWidth = 2
	}
	rr.Refresh()
	return rr
}

//docs:end widget

//docs:start renderer
// ratingRenderer owns the canvas objects that draw a Rating.
type ratingRenderer struct {
	rating *Rating
	dots   [maxRating]*canvas.Circle
}

// MinSize is the smallest size the widget can be drawn at.
func (rr *ratingRenderer) MinSize() geo.Size {
	d := theme.IconInlineSize()
	p := theme.Padding()
	return geo.NewSize(maxRating*d+(maxRating-1)*p, d)
}

// Layout places the objects inside the widget's current size.
func (rr *ratingRenderer) Layout(size geo.Size) {
	slot := size.Width / maxRating
	d := min(slot-theme.Padding(), size.Height)
	for i, dot := range rr.dots {
		dot.Resize(geo.NewSquareSize(d))
		dot.Move(geo.NewPos(float32(i)*slot+(slot-d)/2, (size.Height-d)/2))
	}
}

// Refresh updates the objects from the widget's state and theme.
func (rr *ratingRenderer) Refresh() {
	th := rr.rating.Theme()
	v := bellui.CurrentApp().Settings().ThemeVariant()
	on := th.Color(theme.ColorNamePrimary, v)
	off := th.Color(theme.ColorNameDisabled, v)
	for i, dot := range rr.dots {
		dot.StrokeColor = off
		dot.FillColor = nil
		if i < rr.rating.Value {
			dot.StrokeColor, dot.FillColor = on, on
		}
		dot.Refresh()
	}
}

// Objects lists everything the renderer draws.
func (rr *ratingRenderer) Objects() []core.CanvasObject {
	objs := make([]core.CanvasObject, len(rr.dots))
	for i, dot := range rr.dots {
		objs[i] = dot
	}
	return objs
}

// Destroy releases resources when the renderer is no longer needed.
func (rr *ratingRenderer) Destroy() {}

//docs:end renderer

func main() {
	a := app.NewWithID("com.example.rating")
	w := a.NewWindow("Custom widget")

	//docs:start use
	label := widget.NewLabel("Rating: 3")
	rating := NewRating(3)
	rating.OnChanged = func(v int) {
		label.SetText(fmt.Sprintf("Rating: %d", v))
	}
	reset := widget.NewButton("Reset", func() { rating.SetValue(0) })
	//docs:end use

	w.SetContent(container.NewCenter(container.NewVBox(rating, label, reset)))
	bellui.ShowAndRun(w)
}
