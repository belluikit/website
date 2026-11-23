// Inner windows: windows drawn inside the app's own window, which the user
// drags, resizes, raises, closes and maximises.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/dialog"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Inner windows")
	w.Resize(geo.NewSize(640, 420))

	//docs:start windows
	notes := container.NewInnerWindow("Notes", widget.NewMultiLineEntry())
	notes.Icon = theme.DocumentIcon()
	notes.Move(geo.NewPos(20, 20))
	notes.Resize(geo.NewSize(300, 220))

	info := container.NewInnerWindow("About", widget.NewLabel("Drag a title bar to move a window,\nand its bottom-right corner to resize it."))
	info.Move(geo.NewPos(260, 140))

	desk := container.NewMultipleWindows(notes, info)
	//docs:end windows

	//docs:start maximise
	var maximised bool
	var restorePos geo.Position
	var restoreSize geo.Size
	notes.OnMaximized = func() {
		if maximised {
			notes.Move(restorePos)
			notes.Resize(restoreSize)
		} else {
			restorePos, restoreSize = notes.Position(), notes.Size()
			notes.Move(geo.NewPos(0, 0))
			notes.Resize(desk.Size())
		}
		maximised = !maximised
		notes.SetMaximized(maximised) // swaps the maximise icon for restore
	}
	//docs:end maximise

	//docs:start close
	notes.CloseIntercept = func() {
		dialog.ShowConfirm("Close notes?", "Unsaved text will be lost.", func(ok bool) {
			if ok {
				notes.Close()
			}
		}, w)
	}
	//docs:end close

	w.SetContent(desk)
	bellui.ShowAndRun(w)
}
