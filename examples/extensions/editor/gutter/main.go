// A gutter with breakpoints, line numbers and fold controls, plus diagnostics.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/editor"
	"github.com/belluikit/bellui/x/editor/highlight/starlark"
)

const script = `# Rules for the build.
# Edit with care.

def evaluate(ctx):
    total = 0
    for item in ctx.items:
        total += item.price
    if total == None:
        return 0
    return total
`

func main() {
	a := app.New()
	w := a.NewWindow("Gutter")
	w.Resize(geo.NewSize(700, 450))

	status := widget.NewLabel("Click left of a line number to set a breakpoint.")

	//docs:start gutter
	ed := editor.NewEditor()
	ed.Highlighter = starlark.New()
	ed.SetText(script)

	breakpoints := editor.NewBreakpointGutter(func(lines map[int]bool) {
		status.SetText(fmt.Sprintf("%d breakpoints", len(lines)))
	})
	// NewCompositeGutter puts the columns in the usual order:
	// markers, then line numbers, then fold controls.
	ed.Gutter = editor.NewCompositeGutter(
		editor.NewLineNumberGutter(),
		ed.FoldGutter(),
		ed.DiagnosticGutter(),
		breakpoints,
	)
	//docs:end gutter

	//docs:start folds
	// Rows and columns count from 0.
	ed.SetFoldRegions([]editor.FoldRegion{
		{Range: editor.TextRange{StartRow: 0, EndRow: 1}, Kind: editor.FoldKindComment},
		{Range: editor.TextRange{StartRow: 3, EndRow: 9}, Kind: editor.FoldKindBlock},
	})
	//docs:end folds

	//docs:start diagnostics
	ed.SetDiagnostics([]editor.Diagnostic{
		{
			Range:    editor.TextRange{StartRow: 7, StartCol: 7, EndRow: 7, EndCol: 20},
			Severity: editor.SeverityWarning,
			Message:  "compare with None using 'is'",
			Source:   "lint",
		},
	})
	ed.SetDiagnosticClickToReveal(true) // clicking a marker moves to the problem
	//docs:end diagnostics

	w.SetContent(container.NewBorder(nil, status, nil, nil, ed))
	bellui.ShowAndRun(w)
}
