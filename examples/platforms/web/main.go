// An app that also runs in the browser: build it with GOOS=js GOARCH=wasm.
package main

import (
	"net/url"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.NewWithID("com.example.web")
	w := a.NewWindow("Web")

	//docs:start web
	where := "a desktop or mobile app"
	if bellui.CurrentDevice().IsBrowser() {
		where = "a web page"
	}
	label := widget.NewLabel("Running as " + where)

	// In the browser a second window opens inside the page's canvas.
	more := widget.NewButton("Open another window", func() {
		w2 := a.NewWindow("Second")
		w2.SetContent(widget.NewLabel("Drag me by the title bar"))
		w2.Show()
	})

	docs, _ := url.Parse("https://bellui.com/docs/")
	link := widget.NewHyperlink("Read the docs", docs)
	//docs:end web

	w.SetContent(container.NewVBox(label, more, link))
	bellui.ShowAndRun(w)
}
