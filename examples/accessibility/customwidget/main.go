// A custom widget that screen readers can name, read and operate, and that
// works with the keyboard.
package main

import (
	"fmt"
	"strings"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/event"
	"github.com/belluikit/bellui/widget"
)

// Rating is a row of five stars. Tap a star, or focus the widget and use the
// arrow keys.
type Rating struct {
	widget.BaseWidget
	Label string
	Value int // 0 to 5

	focused bool
	text    *widget.Label
}

func NewRating(label string) *Rating {
	r := &Rating{Label: label, text: widget.NewLabel("")}
	r.ExtendBaseWidget(r)
	return r
}

func (r *Rating) CreateRenderer() core.WidgetRenderer {
	r.update()
	return widget.NewSimpleRenderer(r.text)
}

func (r *Rating) update() {
	stars := strings.Repeat("★", r.Value) + strings.Repeat("☆", 5-r.Value)
	if r.focused {
		stars = "[" + stars + "]" // a visible focus indicator
	}
	r.text.SetText(stars)
}

func (r *Rating) set(v int) {
	r.Value = max(0, min(5, v))
	r.update()
}

//docs:start semantics
var (
	_ core.Accessible         = (*Rating)(nil)
	_ core.AccessibleValuer   = (*Rating)(nil)
	_ core.AccessibleActioner = (*Rating)(nil)
)

// AccessibleRole tells assistive technology what kind of control this is.
func (r *Rating) AccessibleRole() core.AccessibleRole { return core.RoleSlider }

// AccessibleLabel is the name a screen reader speaks.
func (r *Rating) AccessibleLabel() string { return r.Label }

// AccessibleValue is the value spoken after the name.
func (r *Rating) AccessibleValue() string { return fmt.Sprintf("%d of 5 stars", r.Value) }

// AccessibleRange gives the numeric range, so the value can be adjusted.
func (r *Rating) AccessibleRange() (min, max, current float64) {
	return 0, 5, float64(r.Value)
}

// AccessibleActions lists what assistive technology can do with the widget.
func (r *Rating) AccessibleActions() []core.AccessibleAction {
	return []core.AccessibleAction{
		{Name: core.ActionIncrement, Description: "More stars"},
		{Name: core.ActionDecrement, Description: "Fewer stars"},
	}
}

// PerformAccessibleAction runs an action by name.
func (r *Rating) PerformAccessibleAction(name string) {
	switch name {
	case core.ActionIncrement:
		r.set(r.Value + 1)
	case core.ActionDecrement:
		r.set(r.Value - 1)
	}
}

//docs:end semantics

//docs:start keyboard
var _ event.Focusable = (*Rating)(nil)

func (r *Rating) FocusGained()   { r.focused = true; r.update() }
func (r *Rating) FocusLost()     { r.focused = false; r.update() }
func (r *Rating) TypedRune(rune) {}

// TypedKey gives keyboard users the same control as the pointer.
func (r *Rating) TypedKey(ev *event.KeyEvent) {
	switch ev.Name {
	case event.KeyRight, event.KeyUp:
		r.set(r.Value + 1)
	case event.KeyLeft, event.KeyDown:
		r.set(r.Value - 1)
	case event.KeyHome:
		r.set(0)
	}
}

//docs:end keyboard

// Tapped sets the rating from where the row was tapped.
func (r *Rating) Tapped(ev *event.PointEvent) {
	if w := r.Size().Width; w > 0 {
		r.set(int(ev.Position.X/w*5) + 1)
	}
}

func main() {
	a := app.New()
	w := a.NewWindow("Rating")
	w.SetContent(container.NewVBox(
		widget.NewEntry(),
		NewRating("Food"),
		NewRating("Service"),
	))
	bellui.ShowAndRun(w)
}
