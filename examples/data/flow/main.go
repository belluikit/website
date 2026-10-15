// Reactive state with the flow package: values, derived values and the
// widget helpers that connect them.
package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/flow"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Flow")

	//docs:start state
	first := flow.New("Ada")
	last := flow.New("Lovelace")
	shout := flow.New(false)

	full := flow.Derive(func() string {
		s := first.Get() + " " + last.Get()
		if shout.Get() {
			s = strings.ToUpper(s)
		}
		return s
	}, first, last, shout)
	//docs:end state

	//docs:start widgets
	firstEntry := widget.NewEntry()
	widget.BindEntry(firstEntry, first) // two-way
	lastEntry := widget.NewEntry()
	widget.BindEntry(lastEntry, last)
	shoutCheck := widget.NewCheck("Shout", nil)
	widget.BindCheck(shoutCheck, shout)

	result := widget.NewLabel("")
	widget.BindLabel(result, widget.StringDerived{Derived: full}) // one-way
	//docs:end widgets

	//docs:start format
	clicks := flow.New(0)
	counter := widget.NewLabel("")
	widget.BindLabel(counter, widget.FormatValue(clicks, "Clicked %d times"))
	button := widget.NewButton("Click", func() {
		clicks.Modify(func(n int) int { return n + 1 })
	})
	//docs:end format

	//docs:start background
	go func() {
		time.Sleep(3 * time.Second)
		bellui.Do(func() { // listeners update widgets, so Set on the UI goroutine
			last.Set("Byron")
		})
	}()
	//docs:end background

	//docs:start listen
	stop := full.OnChange(func() {
		fmt.Println("name is now", full.Get())
	})
	defer stop()
	//docs:end listen

	w.SetContent(container.NewVBox(firstEntry, lastEntry, shoutCheck, result, counter, button))
	bellui.ShowAndRun(w)
}
