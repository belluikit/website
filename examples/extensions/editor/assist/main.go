// Completion, hover cards, a current-line highlight and span decorations.
package main

import (
	"image/color"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/editor"
	"github.com/belluikit/bellui/x/editor/highlight/starlark"
)

func main() {
	a := app.New()
	w := a.NewWindow("Assist")
	w.Resize(geo.NewSize(700, 450))

	ed := editor.NewEditor()
	ed.Highlighter = starlark.New()
	ed.Gutter = editor.NewLineNumberGutter()
	ed.SetText("def greet(name):\n    return \"Hello, \" + name\n\ngreet(\"world\")\n")

	//docs:start assist
	ed.SetCompletionProvider(&editor.WordListCompletionProvider{
		Words: []string{"def", "return", "greet", "print", "len"},
		Kind:  editor.CompletionKeyword,
	})
	ed.SetHoverProvider(&editor.WordHoverProvider{
		Docs: map[string]string{
			"greet": "**greet(name)**\n\nReturns a greeting for `name`.",
		},
		Text: ed.Text,
	})
	//docs:end assist

	//docs:start decorate
	current := editor.NewCurrentLineDecorator(1) // 1-based
	ed.Decorators = []editor.RowDecorator{current}
	ed.OnCursorChanged = func(row, col int) { // 0-based
		current.Row = row + 1
		ed.RefreshDecorators()
	}

	ed.SpanDecorators = []editor.SpanDecorator{&editor.StaticSpanDecorator{
		SpansList: []editor.DecoratedSpan{{
			Range: editor.TextRange{StartRow: 3, StartCol: 0, EndRow: 3, EndCol: 5},
			Kind:  editor.SpanBox,
			Color: color.NRGBA{R: 33, G: 150, B: 243, A: 255},
		}},
	}}
	//docs:end decorate

	w.SetContent(ed)
	bellui.ShowAndRun(w)
}
