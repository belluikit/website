// Drag and drop: inside the app, out to other apps, and files in from the OS.
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
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"
)

var (
	idle  = color.NRGBA{R: 0x9E, G: 0x9E, B: 0x9E, A: 0x40}
	hover = color.NRGBA{R: 0x1E, G: 0x88, B: 0xE5, A: 0x60}
)

//docs:start card
// card is a drag source for in-app drag and drop.
type card struct {
	widget.BaseWidget
	name string
}

func newCard(name string) *card {
	c := &card{name: name}
	c.ExtendBaseWidget(c)
	return c
}

func (c *card) CreateRenderer() core.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewStack(canvas.NewRectangle(hover), widget.NewLabel(c.name)))
}

// event.Draggable: the driver calls these as the pointer moves and on release.
func (c *card) Dragged(*event.DragEvent) {}
func (c *card) DragEnd()                 {}

// core.DragPayloadProvider: the value handed to the drop target.
func (c *card) DragPayload() any { return c.name }

// core.DragImageProvider: the "ghost" painted under the pointer.
func (c *card) DragImage() core.CanvasObject { return widget.NewLabel(c.name) }

//docs:end card

//docs:start bin
// bin is an in-app drop target that highlights while a card is over it.
type bin struct {
	widget.BaseWidget
	bg     *canvas.Rectangle
	label  *widget.Label
	onDrop func(name string)
}

func newBin(onDrop func(string)) *bin {
	b := &bin{bg: canvas.NewRectangle(idle), label: widget.NewLabel("Drop a card here"), onDrop: onDrop}
	b.ExtendBaseWidget(b)
	return b
}

func (b *bin) CreateRenderer() core.WidgetRenderer {
	return widget.NewSimpleRenderer(container.NewStack(b.bg, container.NewCenter(b.label)))
}

func (b *bin) setColor(c color.Color) { b.bg.FillColor = c; b.bg.Refresh() }

// core.DragTarget
func (b *bin) DragEnter(*core.DragTargetEvent) { b.setColor(hover) }
func (b *bin) DragOver(*core.DragTargetEvent)  {}
func (b *bin) DragLeave()                      { b.setColor(idle) }
func (b *bin) Drop(e *core.DragTargetEvent) {
	b.setColor(idle)
	if name, ok := e.Payload.(string); ok {
		b.onDrop(name)
	}
}

//docs:end bin

//docs:start export
// export drags a piece of text out of the app, into an editor or browser.
type export struct {
	widget.BaseWidget
	text    string
	started bool
}

func newExport(text string) *export {
	e := &export{text: text}
	e.ExtendBaseWidget(e)
	return e
}

func (e *export) CreateRenderer() core.WidgetRenderer {
	return widget.NewSimpleRenderer(widget.NewLabel("Drag me to another app"))
}

func (e *export) data() desktop.DragData {
	return desktop.DragData{
		Text: e.text,
		OnCompleted: func(accepted bool) {
			fmt.Println("drop accepted:", accepted)
		},
	}
}

// Desktop: start the native drag from the first Dragged call of a gesture.
func (e *export) Dragged(*event.DragEvent) {
	if !e.started {
		e.started = desktop.StartDrag(e.data())
	}
}

func (e *export) DragEnd() { e.started = false }

// Web: the browser asks for the payload when the gesture starts.
func (e *export) DragOutData() (desktop.DragData, bool) { return e.data(), true }

//docs:end export

//docs:start files
// filesIn accepts files dropped from the desktop's file manager.
type filesIn struct {
	widget.BaseWidget
	label *widget.Label
}

func newFilesIn() *filesIn {
	f := &filesIn{label: widget.NewLabel("Drop files from your desktop here")}
	f.ExtendBaseWidget(f)
	return f
}

func (f *filesIn) CreateRenderer() core.WidgetRenderer {
	return widget.NewSimpleRenderer(f.label)
}

// bellui.Droppable
func (f *filesIn) Dropped(e *bellui.DropEvent) {
	f.label.SetText(fmt.Sprintf("Received %d item(s), first: %s", len(e.URIs), e.URIs[0].Name()))
}

//docs:end files

func main() {
	a := app.New()
	w := a.NewWindow("Drag and drop")

	status := widget.NewLabel("")
	target := newBin(func(name string) { status.SetText("Dropped " + name) })

	//docs:start window
	// Drops that no Droppable claims fall back to the window-wide callback.
	if dw, ok := w.(bellui.DroppableWindow); ok {
		dw.SetOnDropped(func(pos geo.Position, items []bellui.URI) {
			status.SetText(fmt.Sprintf("%d item(s) dropped on the window", len(items)))
		})
	}
	//docs:end window

	w.SetContent(container.NewPadded(container.NewVBox(
		container.NewHBox(newCard("Alpha"), newCard("Beta")),
		container.NewGridWrap(geo.NewSize(300, 80), target),
		newExport("Hello from Bellui"),
		newFilesIn(),
		status,
	)))
	w.Resize(geo.NewSize(420, 360))
	bellui.ShowAndRun(w)
}
