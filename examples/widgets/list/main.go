// List: a scrolling, recycling list of rows.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("List")
	w.Resize(geo.NewSize(360, 420))

	var mails []string
	for i := 1; i <= 1000; i++ {
		mails = append(mails, fmt.Sprintf("Message %d", i))
	}
	status := widget.NewLabel("")

	//docs:start list
	list := widget.NewList(
		func() int { return len(mails) },
		func() core.CanvasObject {
			// A template row; it is reused for many items.
			return container.NewHBox(widget.NewIcon(theme.MailComposeIcon()), widget.NewLabel(""))
		},
		func(id widget.ListItemID, row core.CanvasObject) {
			row.(*bellui.Container).Objects[1].(*widget.Label).SetText(mails[id])
		},
	)
	list.OnSelected = func(id widget.ListItemID) {
		status.SetText("Opened " + mails[id])
	}
	//docs:end list

	w.SetContent(container.NewBorder(nil, status, nil, nil, list))
	bellui.ShowAndRun(w)
}
