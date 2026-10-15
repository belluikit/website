// A data grid over a slice of structs, with 100,000 rows.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/tables"
)

//docs:start rows
type Order struct {
	ID       int
	Customer string
	Total    float64
	Shipped  bool
}

//docs:end rows

func main() {
	a := app.New()
	w := a.NewWindow("Orders")
	w.Resize(geo.NewSize(700, 500))

	orders := make([]Order, 100_000)
	for i := range orders {
		orders[i] = Order{
			ID:       i + 1,
			Customer: fmt.Sprintf("Customer %d", i%977),
			Total:    float64(i%500) * 1.25,
			Shipped:  i%3 == 0,
		}
	}

	//docs:start grid
	grid := tables.NewDataGrid(
		[]*tables.ColumnDef{
			{Field: "ID", Type: tables.ColumnNumber, Width: 80},
			{Field: "Customer", Flex: 1},
			{
				Field: "Total", HeaderName: "Total", Type: tables.ColumnNumber, Width: 110,
				ValueFormatter: func(v any) string { return fmt.Sprintf("$%.2f", v) },
			},
			{Field: "Shipped", Type: tables.ColumnBoolean, Width: 90},
		},
		func() int { return len(orders) },
		func(row int) any { return orders[row] },
	)
	grid.Options.RowAlternation = true
	grid.Options.ShowStatusBar = true
	//docs:end grid

	w.SetContent(grid)
	bellui.ShowAndRun(w)
}
