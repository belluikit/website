// Rendering Chinese, Japanese and Korean text with the bundled CJK fallback.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"

	//docs:start import
	_ "github.com/belluikit/bellui/theme/cjk" // registers a CJK fallback font
	//docs:end import
)

func main() {
	a := app.New()
	w := a.NewWindow("CJK")
	w.SetContent(container.NewVBox(
		widget.NewLabel("你好，世界"),
		widget.NewLabel("こんにちは"),
		widget.NewLabel("안녕하세요"),
	))
	bellui.ShowAndRun(w)
}
