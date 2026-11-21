// A strip that pans sideways inside a page that scrolls up and down. On a
// touch screen it takes only sideways drags, so an upward swipe that starts on
// it still scrolls the page, and it keeps moving after a fling.
package main

import (
	"fmt"
	"image/color"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/canvas"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/event"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"
)

//docs:start strip
// strip pans its content left and right when dragged.
type strip struct {
	widget.BaseWidget
	offset float32
	text   *canvas.Text
}

func newStrip() *strip {
	s := &strip{text: canvas.NewText("", color.NRGBA{R: 0x3b, G: 0x6e, B: 0xc8, A: 0xff})}
	s.ExtendBaseWidget(s)
	return s
}

// event.Draggable: move the content with the pointer.
func (s *strip) Dragged(e *event.DragEvent) {
	s.offset += e.Dragged.DX
	s.Refresh()
}

func (s *strip) DragEnd() {}

// event.DragClaimer: on touch, take a drag only when it starts sideways.
// An upward or downward drag goes on to the scroll around the strip.
func (s *strip) ClaimsDrag(d geo.Delta) bool {
	return abs(d.DX) > abs(d.DY)
}

// event.MomentumDraggable: on touch, keep moving after a fast release,
// because the strip pans content rather than being moved itself.
func (s *strip) DragMomentum() bool { return true }

//docs:end strip

func abs(v float32) float32 {
	if v < 0 {
		return -v
	}
	return v
}

func (s *strip) CreateRenderer() core.WidgetRenderer {
	bg := canvas.NewRectangle(color.NRGBA{R: 0x80, G: 0x80, B: 0x80, A: 0x30})
	bg.SetMinSize(geo.NewSize(200, 80))
	r := &stripRenderer{s: s, bg: bg}
	r.Refresh()
	return r
}

type stripRenderer struct {
	s  *strip
	bg *canvas.Rectangle
}

func (r *stripRenderer) Layout(size geo.Size) {
	r.bg.Resize(size)
	r.s.text.Move(geo.NewPos(r.s.offset, (size.Height-r.s.text.MinSize().Height)/2))
}

func (r *stripRenderer) MinSize() geo.Size { return r.bg.MinSize() }

func (r *stripRenderer) Refresh() {
	r.s.text.Text = fmt.Sprintf("Offset %.0f", r.s.offset)
	r.s.text.Refresh()
	r.Layout(r.s.Size())
}

func (r *stripRenderer) Objects() []core.CanvasObject { return []core.CanvasObject{r.bg, r.s.text} }

func (r *stripRenderer) Destroy() {}

func main() {
	a := app.NewWithID("com.example.drag")
	w := a.NewWindow("Drag directions")

	page := container.NewVBox()
	for i := 1; i <= 20; i++ {
		page.Add(widget.NewLabel(fmt.Sprintf("Row %d", i)))
		if i == 3 {
			page.Add(newStrip())
		}
	}

	w.SetContent(container.NewVScroll(page))
	w.Resize(geo.NewSize(320, 480))
	bellui.ShowAndRun(w)
}
