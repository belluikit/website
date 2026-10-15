// Master-detail: each row expands to show a widget of your own.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/tables"
)

func main() {
	a := app.New()
	w := a.NewWindow("Master-detail")
	w.Resize(geo.NewSize(600, 450))

	customers := []map[string]any{
		{"name": "Acme", "orders": 12},
		{"name": "Globex", "orders": 0},
		{"name": "Initech", "orders": 4},
	}

	//docs:start detail
	grid := tables.NewDataGrid(
		[]*tables.ColumnDef{
			{Field: "name", HeaderName: "Customer", Width: 200},
			{Field: "orders", HeaderName: "Orders", Type: tables.ColumnNumber, Width: 100},
		},
		func() int { return len(customers) },
		func(row int) any { return customers[row] },
	)
	grid.MasterDetail = tables.MasterDetailConfig{
		Enabled:      true,
		DetailHeight: 80,
		// Only customers with orders can be expanded.
		IsRowMaster: func(sourceRow int) bool { return customers[sourceRow]["orders"].(int) > 0 },
		DetailRenderer: func(sourceRow int) core.CanvasObject {
			c := customers[sourceRow]
			return widget.NewLabel(fmt.Sprintf("%s has %d open orders.", c["name"], c["orders"]))
		},
	}
	grid.ToggleDetail(0) // expand the first row from code
	//docs:end detail

	w.SetContent(grid)
	bellui.ShowAndRun(w)
}
