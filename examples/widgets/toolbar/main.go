// A toolbar with actions, a separator, a spacer and a two-state action.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Toolbar")
	status := widget.NewLabel("")

	//docs:start toolbar
	undo := widget.NewToolbarAction(theme.ContentUndoIcon(), func() {
		status.SetText("Undo")
	})
	play := widget.NewTwoStateToolbarAction(theme.MediaPlayIcon(), theme.MediaPauseIcon(),
		func(on bool) {
			status.SetText(fmt.Sprint("Playing: ", on))
		})

	bar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() { status.SetText("New") }),
		widget.NewToolbarAction(theme.DocumentSaveIcon(), func() { status.SetText("Save") }),
		widget.NewToolbarSeparator(),
		undo,
		play,
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() { status.SetText("Help") }),
	)
	undo.Disable()
	//docs:end toolbar

	w.SetContent(container.NewBorder(bar, nil, nil, nil, status))
	bellui.ShowAndRun(w)
}
