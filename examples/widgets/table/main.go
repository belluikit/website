// Table: a scrolling grid of cells with a header row.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"
)

var rows = [][]string{
	{"Ada", "Engineering", "London"},
	{"Grace", "Research", "New York"},
	{"Linus", "Kernel", "Portland"},
	{"Margaret", "Apollo", "Boston"},
}

var headers = []string{"Name", "Team", "City"}

func main() {
	a := app.New()
	w := a.NewWindow("Table")
	w.Resize(geo.NewSize(480, 260))

	//docs:start table
	table := widget.NewTable(
		func() (int, int) { return len(rows), len(headers) },
		func() core.CanvasObject { return widget.NewLabel("template text") },
		func(id widget.TableCellID, cell core.CanvasObject) {
			cell.(*widget.Label).SetText(rows[id.Row][id.Col])
		},
	)
	table.ShowHeaderRow = true
	table.CreateHeader = func() core.CanvasObject {
		return widget.NewLabelWithStyle("", bellui.TextAlignLeading, bellui.TextStyle{Bold: true})
	}
	table.UpdateHeader = func(id widget.TableCellID, cell core.CanvasObject) {
		if id.Row == -1 { // the header row
			cell.(*widget.Label).SetText(headers[id.Col])
		}
	}
	table.SetColumnWidth(0, 120)
	table.OnSelected = func(id widget.TableCellID) {
		w.SetTitle("Selected " + rows[id.Row][0])
	}
	//docs:end table

	w.SetContent(table)
	bellui.ShowAndRun(w)
}
