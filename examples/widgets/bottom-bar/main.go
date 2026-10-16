// A bottom app bar with actions and a floating action button.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Bottom app bar")
	w.Resize(geo.NewSize(400, 300))
	status := widget.NewLabel("")

	//docs:start bottom
	fab := widget.NewFAB(theme.ContentAddIcon(), func() { status.SetText("New") })
	bottom := widget.NewBottomAppBar([]*widget.BottomAppBarAction{
		{Icon: theme.SearchIcon(), OnTapped: func() { status.SetText("Search") }},
		{Icon: theme.DeleteIcon(), OnTapped: func() { status.SetText("Delete") }},
	}, fab)
	//docs:end bottom

	w.SetContent(container.NewBorder(nil, bottom, nil, nil, status))
	bellui.ShowAndRun(w)
}
