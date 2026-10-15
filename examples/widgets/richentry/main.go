// RichEntry: an editor whose text carries per-run styles.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("RichEntry")
	w.Resize(geo.NewSize(480, 300))
	count := widget.NewLabel("")

	//docs:start rich
	editor := widget.NewRichEntry()
	editor.Wrapping = bellui.TextWrapWord
	editor.SetSegments([]widget.RichTextSegment{
		&widget.TextSegment{Text: "Select some text and press ", Style: widget.RichTextStyleInline},
		&widget.TextSegment{Text: "Bold", Style: widget.RichTextStyleStrong},
		&widget.TextSegment{Text: ".", Style: widget.RichTextStyleInline},
	})
	editor.OnChanged = func() {
		count.SetText(editor.String())
	}

	bold := widget.NewButton("Bold", func() {
		editor.ApplyStyle(widget.RichTextStyleStrong)
	})
	italic := widget.NewButton("Italic", func() {
		editor.ApplyStyle(widget.RichTextStyleEmphasis)
	})
	//docs:end rich

	w.SetContent(container.NewBorder(container.NewHBox(bold, italic), count, nil, nil, editor))
	bellui.ShowAndRun(w)
}
