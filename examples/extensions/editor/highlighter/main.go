// A highlighter of your own, written as a function.
package main

import (
	"strings"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/editor"
)

//docs:start highlighter
var iniHighlighter = editor.HighlighterFunc(func(line string, state int) ([]editor.StyleRun, int) {
	// Colour an INI file: [sections], key = value, and ; comments.
	// No construct spans lines, so the returned state is always 0.
	trimmed := strings.TrimSpace(line)
	switch {
	case strings.HasPrefix(trimmed, ";"):
		return []editor.StyleRun{{Text: line, ColorName: editor.ColorNameSyntaxComment, Italic: true}}, 0
	case strings.HasPrefix(trimmed, "["):
		return []editor.StyleRun{{Text: line, ColorName: editor.ColorNameSyntaxKeyword, Bold: true}}, 0
	}
	key, value, ok := strings.Cut(line, "=")
	if !ok {
		return []editor.StyleRun{{Text: line}}, 0
	}
	// The runs must cover the whole line, in order.
	return []editor.StyleRun{
		{Text: key, ColorName: editor.ColorNameSyntaxType},
		{Text: "=", ColorName: editor.ColorNameSyntaxOperator},
		{Text: value, ColorName: editor.ColorNameSyntaxString},
	}, 0
})

//docs:end highlighter

func main() {
	a := app.New()
	w := a.NewWindow("settings.ini")
	w.Resize(geo.NewSize(600, 350))

	ed := editor.NewEditor()
	ed.Highlighter = iniHighlighter
	ed.Gutter = editor.NewLineNumberGutter()
	ed.SetText("; App settings\n[window]\nwidth = 800\nheight = 600\n\n[theme]\nvariant = dark\n")

	w.SetContent(ed)
	bellui.ShowAndRun(w)
}
