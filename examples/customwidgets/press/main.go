// A widget of your own that shows the MD3 pressed state: a ripple that
// spreads from the press and fades on release, as on bellui's own widgets.
package main

import (
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

// Tile is a rounded, tappable surface with a label.
//
//docs:start widget
type Tile struct {
	widget.DisableableWidget
	widget.PressTarget // the press events and the ripple's state

	Text     string
	OnTapped func()
}

func NewTile(text string, tapped func()) *Tile {
	t := &Tile{Text: text, OnTapped: tapped}
	t.ExtendBaseWidget(t)
	return t
}

// Tapped runs the action. TapPress shows a press for a tap that had none
// before it, such as a key or an accessibility action.
func (t *Tile) Tapped(ev *event.PointEvent) {
	if t.Disabled() {
		return
	}
	widget.TapPress(&t.PressTarget, ev)
	if t.OnTapped != nil {
		t.OnTapped()
	}
}

// HoverOut ends a press when the pointer leaves: a mouse released
// elsewhere sends the tile no MouseUp.
func (t *Tile) HoverIn(*event.HoverEvent)    {}
func (t *Tile) HoverMoved(*event.HoverEvent) {}
func (t *Tile) HoverOut()                    { widget.CancelPress(&t.PressTarget) }

//docs:end widget

const tileRadius = 12

//docs:start renderer
func (t *Tile) CreateRenderer() core.WidgetRenderer {
	t.ExtendBaseWidget(t)
	v := bellui.CurrentApp().Settings().ThemeVariant()
	bg := canvas.NewRectangle(t.Theme().Color(theme.ColorNameSecondaryContainer, v))
	bg.CornerRadius = tileRadius
	label := canvas.NewText(t.Text, t.Theme().Color(theme.ColorNameOnSecondaryContainer, v))

	// AttachPress asks this function at each press where the state layer
	// goes and in which colour; false refuses the press.
	ripple := widget.AttachPress(&t.PressTarget, func() (widget.PressShape, bool) {
		return widget.PressShape{
			Size:         t.Size(),
			CornerRadius: tileRadius,
			Content:      label.Color, // the layer is this colour, faded
		}, !t.Disabled()
	})

	// The ripple goes above the background and below the content.
	return &tileRenderer{bg: bg, ripple: ripple, label: label,
		objects: []core.CanvasObject{bg, ripple, label}}
}

//docs:end renderer

type tileRenderer struct {
	bg      *canvas.Rectangle
	ripple  *canvas.Ripple
	label   *canvas.Text
	objects []core.CanvasObject
}

func (r *tileRenderer) Layout(size geo.Size) {
	r.bg.Resize(size)
	r.ripple.Resize(size)
	ls := r.label.MinSize()
	r.label.Resize(ls)
	r.label.Move(geo.NewPos((size.Width-ls.Width)/2, (size.Height-ls.Height)/2))
}

func (r *tileRenderer) MinSize() geo.Size {
	ls := r.label.MinSize()
	return geo.NewSize(ls.Width+48, 56)
}

func (r *tileRenderer) Refresh()                     { canvas.Refresh(r.bg) }
func (r *tileRenderer) Objects() []core.CanvasObject { return r.objects }
func (r *tileRenderer) Destroy()                     {}

func main() {
	a := app.New()
	w := a.NewWindow("Pressed state")
	status := widget.NewLabel("Press the tile")
	tile := NewTile("Press me", func() { status.SetText("Tapped") })
	w.SetContent(container.NewVBox(tile, status))
	w.Resize(geo.NewSize(320, 200))
	bellui.ShowAndRun(w)
}
