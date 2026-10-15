// A small app whose UI is built by a function, so a test can build it too.
package main

import (
	"strconv"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/widget"
)

//docs:start ui
// counter holds the widgets a test needs to reach.
type counter struct {
	name   *widget.Entry
	add    *widget.Button
	output *widget.Label
	count  int
}

func newCounter() *counter {
	c := &counter{
		name:   widget.NewEntry(),
		output: widget.NewLabel("Nobody has clicked yet"),
	}
	c.add = widget.NewButton("Click", func() {
		c.count++
		c.output.SetText(c.name.Text + " clicked " + strconv.Itoa(c.count) + " times")
	})
	return c
}

func (c *counter) content() core.CanvasObject {
	return container.NewVBox(c.name, c.add, c.output)
}

func main() {
	a := app.New()
	w := a.NewWindow("Counter")
	w.SetContent(newCounter().content())
	bellui.ShowAndRun(w)
}

//docs:end ui
