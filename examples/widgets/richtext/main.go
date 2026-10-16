// RichText from segments, Markdown and HTML.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"
)

const notes = "# Release 1.2\n\n" +
	"Fixes **two** crashes and adds *dark mode*.\n\n" +
	"- Faster start-up\n" +
	"- [x] Translations updated\n\n" +
	"```go\nfmt.Println(\"hello\")\n```\n"

func main() {
	a := app.New()
	w := a.NewWindow("Rich text")
	w.Resize(geo.NewSize(480, 520))

	//docs:start segments
	status := widget.NewRichText(
		&widget.TextSegment{Text: "Build ", Style: widget.RichTextStyleInline},
		&widget.TextSegment{Text: "passed", Style: widget.RichTextStyleStrong},
		&widget.TextSegment{Text: " in 42 s. ", Style: widget.RichTextStyleInline},
		&widget.TextSegment{Text: "go test ./...", Style: widget.RichTextStyleCodeInline},
	)
	status.Wrapping = bellui.TextWrapWord
	//docs:end segments

	//docs:start markdown
	md := widget.NewRichTextFromMarkdown(notes)
	md.Wrapping = bellui.TextWrapWord
	//docs:end markdown

	//docs:start html
	html := widget.NewRichTextFromHTML(
		`<p>Read the <a href="https://bellui.com">guide</a> or <code>go doc</code>.</p>` +
			`<ul><li><b>Bold</b> item</li><li><i>Italic</i> item</li></ul>`)
	//docs:end html

	w.SetContent(container.NewVScroll(container.NewVBox(status, widget.NewSeparator(), md,
		widget.NewSeparator(), html)))
	bellui.ShowAndRun(w)
}
