// Load rows a page at a time as the user scrolls towards the end.
package main

import (
	"fmt"
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/tables"
)

const pageSize = 100

// fetchPage stands in for a database or network call.
func fetchPage(from int) []map[string]any {
	time.Sleep(200 * time.Millisecond)
	page := make([]map[string]any, pageSize)
	for i := range page {
		page[i] = map[string]any{"n": from + i + 1, "name": fmt.Sprintf("Item %d", from+i+1)}
	}
	return page
}

func main() {
	a := app.New()
	w := a.NewWindow("Paging")
	w.Resize(geo.NewSize(500, 400))

	//docs:start paging
	rows := fetchPage(0)
	loading := false

	grid := tables.NewDataGrid(
		[]*tables.ColumnDef{
			{Field: "n", HeaderName: "#", Type: tables.ColumnNumber, Width: 80},
			{Field: "name", HeaderName: "Name", Flex: 1},
		},
		func() int { return len(rows) },
		func(row int) any { return rows[row] },
	)
	grid.OnVisibleRowsChanged = func(first, last int) {
		if loading || last < grid.RowCount()-20 {
			return // not near the end yet
		}
		loading = true
		from := len(rows)
		go func() {
			page := fetchPage(from)
			bellui.Do(func() {
				rows = append(rows, page...)
				loading = false
				grid.RefreshData()
			})
		}()
	}
	//docs:end paging

	w.SetContent(grid)
	bellui.ShowAndRun(w)
}
