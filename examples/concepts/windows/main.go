// An app with a main window that opens extra windows. Closing the main
// window quits the app; closing an extra window does not.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/dialog"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"
)

func main() {
	//docs:start app
	a := app.NewWithID("com.example.windows")
	w := a.NewWindow("Main window")
	w.Resize(geo.NewSize(420, 260))
	//docs:end app

	//docs:start master
	// Closing this window quits the app, even if other windows are open.
	if m, ok := w.(bellui.MasterWindow); ok {
		m.SetMaster()
	}
	//docs:end master

	//docs:start intercept
	// Ask before closing. The window only closes when Close is called.
	w.SetCloseIntercept(func() {
		dialog.ShowConfirm("Quit?", "Close the main window and quit?", func(ok bool) {
			if ok {
				w.Close()
			}
		}, w)
	})
	//docs:end intercept

	//docs:start more
	count := 0
	open := widget.NewButton("Open another window", func() {
		count++
		extra := a.NewWindow(fmt.Sprintf("Window %d", count))
		extra.SetContent(widget.NewLabel("Close me: the app keeps running."))
		extra.Resize(geo.NewSize(300, 120))
		extra.SetOnClosed(func() {
			fmt.Println("closed", extra.Title())
		})
		extra.Show()
	})
	//docs:end more

	w.SetContent(container.NewVBox(
		widget.NewLabel("This is the main window."),
		open,
	))
	bellui.ShowAndRun(w)
}
