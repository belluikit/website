// Sorting, column filters and a quick filter, set from code.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/tables"
	"github.com/belluikit/bellui/x/tables/filter"
)

func main() {
	a := app.New()
	w := a.NewWindow("Sort and filter")
	w.Resize(geo.NewSize(700, 500))

	sectors := []string{"Technology", "Energy", "Health", "Retail"}
	var rows []map[string]any
	for i := range 200 {
		rows = append(rows, map[string]any{
			"symbol": fmt.Sprintf("SYM%03d", i),
			"sector": sectors[i%len(sectors)],
			"price":  float64(10 + (i*37)%490),
		})
	}

	grid := tables.NewDataGrid(
		[]*tables.ColumnDef{
			{Field: "symbol", HeaderName: "Symbol", Width: 100},
			{Field: "sector", HeaderName: "Sector", Width: 140},
			{Field: "price", HeaderName: "Price", Type: tables.ColumnNumber, Width: 100},
		},
		func() int { return len(rows) },
		func(row int) any { return rows[row] },
	)

	//docs:start sortfilter
	api := grid.API()

	// Sort by sector, then by price, highest first.
	api.SetSort(
		tables.SortItem{ColID: "sector", Direction: tables.SortAscending},
		tables.SortItem{ColID: "price", Direction: tables.SortDescending},
	)

	// Keep rows whose price is at least 100.
	atLeast100 := &filter.NumberFilter{Mode: filter.NumberGreaterThanOrEqual, Value: 100}
	api.SetColumnFilter("price", atLeast100.Matches)

	// A quick filter matches its text against every column.
	search := widget.NewEntry()
	search.SetPlaceHolder("Search all columns")
	search.OnChanged = api.SetQuickFilter
	//docs:end sortfilter

	clear := widget.NewButton("Clear filters", func() {
		api.ClearFilters()
		search.SetText("")
	})

	w.SetContent(container.NewBorder(
		container.NewBorder(nil, nil, nil, clear, search), nil, nil, nil,
		grid,
	))
	bellui.ShowAndRun(w)
}
