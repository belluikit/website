// A pad that reports every pointer event it receives, plus a window-wide
// keyboard shortcut and key handler.
package main

import (
	"fmt"
	"image/color"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/canvas"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/driver/desktop"
	"github.com/belluikit/bellui/event"
	"github.com/belluikit/bellui/widget"
)

//docs:start pad
// pad is a widget that implements several event interfaces. The toolkit
// checks which interfaces an object has and delivers matching events.
type pad struct {
	widget.BaseWidget
	bg  *canvas.Rectangle
	log func(string)
}

func newPad(log func(string)) *pad {
	p := &pad{bg: canvas.NewRectangle(color.NRGBA{R: 0x3b, G: 0x6e, B: 0xc8, A: 0xff}), log: log}
	p.ExtendBaseWidget(p)
	return p
}

func (p *pad) CreateRenderer() core.WidgetRenderer {
	return widget.NewSimpleRenderer(p.bg)
}

// event.Tappable: a click or a tap.
func (p *pad) Tapped(e *event.PointEvent) { p.log(fmt.Sprintf("Tapped at %v", e.Position)) }

// event.SecondaryTappable: a right-click on desktop.
func (p *pad) TappedSecondary(*event.PointEvent) { p.log("TappedSecondary") }

// event.LongPressable: press and hold.
func (p *pad) LongPressed(*event.PointEvent) { p.log("LongPressed") }

// event.Hoverable: the pointer enters, moves over and leaves (desktop and web).
func (p *pad) HoverIn(*event.HoverEvent)    { p.log("HoverIn") }
func (p *pad) HoverMoved(*event.HoverEvent) {}
func (p *pad) HoverOut()                    { p.log("HoverOut") }

// event.Draggable: the pointer moves while pressed.
func (p *pad) Dragged(e *event.DragEvent) { p.log(fmt.Sprintf("Dragged by %v", e.Dragged)) }
func (p *pad) DragEnd()                   { p.log("DragEnd") }

// event.Scrollable: a mouse wheel or touchpad scroll.
func (p *pad) Scrolled(e *event.ScrollEvent) { p.log(fmt.Sprintf("Scrolled %v", e.Scrolled)) }

//docs:end pad

func main() {
	a := app.NewWithID("com.example.events")
	w := a.NewWindow("Input events")

	last := widget.NewLabel("Interact with the blue pad")
	show := func(s string) { last.SetText(s) }

	//docs:start keys
	// A shortcut for the whole window: Ctrl+S, or Cmd+S on macOS.
	save := &desktop.CustomShortcut{KeyName: event.KeyS, Modifier: event.KeyModifierShortcutDefault}
	w.Canvas().AddShortcut(save, func(core.Shortcut) {
		show("Save shortcut")
	})

	// Keys that no focused widget consumed.
	w.Canvas().SetOnTypedKey(func(e *event.KeyEvent) {
		show(fmt.Sprintf("Key %s", e.Name))
	})
	//docs:end keys

	w.SetContent(container.NewBorder(nil, last, nil, nil, newPad(show)))
	bellui.ShowAndRun(w)
}
