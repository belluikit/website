// Plain text and small display widgets: Label, Hyperlink, Icon and Separator.
package main

import (
	"net/url"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Text")

	//docs:start label
	title := widget.NewLabelWithStyle("Release notes", bellui.TextAlignLeading,
		bellui.TextStyle{Bold: true})
	title.SizeName = theme.SizeNameHeadingText

	body := widget.NewLabel("Long text wraps at word boundaries when the label has a fixed width.")
	body.Wrapping = bellui.TextWrapWord

	path := widget.NewLabel("/home/user/projects/a/very/long/path/to/a/file.txt")
	path.Truncation = bellui.TextTruncateEllipsis
	path.ShowTooltipWhenTruncated = true

	warning := widget.NewLabel("Disk almost full")
	warning.Importance = widget.WarningImportance

	id := widget.NewLabel("Order ID: 7F3A-22C1")
	id.Selectable = true
	//docs:end label

	//docs:start link
	u, _ := url.Parse("https://bellui.com")
	site := widget.NewHyperlink("bellui.com", u)

	help := widget.NewHyperlink("Show help", nil)
	help.OnTapped = func() { body.SetText("Help was tapped.") }
	//docs:end link

	//docs:start icon
	row := container.NewHBox(widget.NewIcon(theme.InfoIcon()), widget.NewLabel("Saved to cloud"))
	//docs:end icon

	w.SetContent(container.NewVBox(title, body, path, warning, id,
		widget.NewSeparator(), site, help, row))
	bellui.ShowAndRun(w)
}
