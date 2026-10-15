// Floating action buttons and the FAB menu (speed dial).
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/layout"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("FAB")
	w.Resize(geo.NewSize(420, 360))
	status := widget.NewLabel("")

	//docs:start fab
	compose := widget.NewExtendedFAB("Compose", theme.MailComposeIcon(), func() {
		status.SetText("Compose")
	})
	compose.Style = widget.FABStyleTertiary
	//docs:end fab

	//docs:start menu
	menu := widget.NewFABMenu(theme.ContentAddIcon(), []*widget.FABMenuItem{
		{Label: "Photo", Icon: theme.MediaPhotoIcon(), OnTapped: func() { status.SetText("Photo") }},
		{Label: "Document", Icon: theme.DocumentIcon(), OnTapped: func() { status.SetText("Document") }},
	})

	var add *widget.FAB
	add = widget.NewFAB(theme.ContentAddIcon(), func() {
		// The menu opens above and to the left of the point you pass:
		// here, the FAB's top-right corner in window coordinates.
		pos := bellui.CurrentApp().Driver().AbsolutePositionForObject(add)
		menu.Toggle(w.Canvas(), pos.AddXY(add.Size().Width, 0))
	})
	//docs:end menu

	w.SetContent(container.NewBorder(status, nil, nil, nil,
		container.NewVBox(layout.NewSpacer(),
			container.NewHBox(compose, layout.NewSpacer(), add))))
	bellui.ShowAndRun(w)
}
