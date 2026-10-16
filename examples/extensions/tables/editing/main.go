// Editable cells with a text, a dropdown and a checkbox editor, and undo.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/tables"
	"github.com/belluikit/bellui/x/tables/cell"
)

func main() {
	a := app.New()
	w := a.NewWindow("Editing")
	w.Resize(geo.NewSize(600, 400))

	//docs:start editing
	people := []map[string]any{
		{"name": "Alice", "role": "Engineer", "active": true},
		{"name": "Bob", "role": "Designer", "active": true},
		{"name": "Chen", "role": "Manager", "active": false},
	}

	editable := true
	grid := tables.NewDataGrid(
		[]*tables.ColumnDef{
			{Field: "name", HeaderName: "Name", Width: 150, Editable: &editable},
			{
				Field: "role", HeaderName: "Role", Width: 150, Editable: &editable,
				Editor: cell.NewSelectEditor([]string{"Engineer", "Designer", "Manager"}),
			},
			{
				Field: "active", HeaderName: "Active", Type: tables.ColumnBoolean, Width: 80,
				Editable: &editable, Editor: cell.NewCheckboxEditor(),
			},
		},
		func() int { return len(people) },
		func(row int) any { return people[row] },
	)

	status := widget.NewLabel("Double-click a cell or press F2 to edit.")
	grid.OnCellValueChanged = func(id tables.CellID, old, new any) {
		// The grid has already written new into the row's map.
		status.SetText(fmt.Sprintf("Changed %v to %v", old, new))
	}
	//docs:end editing

	undo := widget.NewButton("Undo", grid.Undo)
	redo := widget.NewButton("Redo", grid.Redo)

	w.SetContent(container.NewBorder(
		container.NewHBox(undo, redo, status), nil, nil, nil,
		grid,
	))
	bellui.ShowAndRun(w)
}
