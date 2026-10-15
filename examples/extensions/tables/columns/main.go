// Pinned columns, a collapsible column group, and a hidden column.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/tables"
)

func main() {
	a := app.New()
	w := a.NewWindow("Columns")
	w.Resize(geo.NewSize(700, 450))

	var rows []map[string]any
	for i := range 50 {
		rows = append(rows, map[string]any{
			"id": i + 1, "first": fmt.Sprintf("First %d", i), "last": fmt.Sprintf("Last %d", i),
			"city": "Lisbon", "country": "Portugal", "phone": "+351 000 000", "balance": float64(i * 10),
		})
	}

	//docs:start columns
	grid := tables.NewDataGrid(
		[]*tables.ColumnDef{
			{Field: "id", HeaderName: "ID", Type: tables.ColumnNumber, Width: 60, Pinned: tables.PinLeft},
			{
				HeaderName: "Name", ColID: "name", OpenByDefault: true,
				Children: []*tables.ColumnDef{
					{Field: "first", HeaderName: "First", Width: 120},
					{Field: "last", HeaderName: "Last", Width: 120},
				},
			},
			{Field: "city", HeaderName: "City", Width: 120},
			{Field: "country", HeaderName: "Country", Width: 120},
			{Field: "phone", HeaderName: "Phone", Width: 140, Hidden: true},
			{Field: "balance", HeaderName: "Balance", Type: tables.ColumnNumber, Width: 100, Pinned: tables.PinRight},
		},
		func() int { return len(rows) },
		func(row int) any { return rows[row] },
	)
	//docs:end columns

	//docs:start columnapi
	toggleGroup := widget.NewButton("Collapse or expand Name", func() {
		grid.ToggleColumnGroup("name")
	})
	showPhone := widget.NewButton("Show phone", func() {
		grid.API().SetColumnVisible("phone", true)
	})
	fit := widget.NewButton("Size to fit", grid.API().SizeColumnsToFit)
	//docs:end columnapi

	w.SetContent(container.NewBorder(
		container.NewHBox(toggleGroup, showPhone, fit), nil, nil, nil,
		grid,
	))
	bellui.ShowAndRun(w)
}
