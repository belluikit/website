// Row grouping by a column, and tree data from parent IDs.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/tables"
	"github.com/belluikit/bellui/x/tables/pipeline"
)

func main() {
	a := app.New()
	w := a.NewWindow("Grouping and trees")
	w.Resize(geo.NewSize(900, 450))

	//docs:start group
	sales := []map[string]any{
		{"region": "North", "product": "Widget", "revenue": 1200.0},
		{"region": "South", "product": "Widget", "revenue": 800.0},
		{"region": "North", "product": "Gadget", "revenue": 450.0},
		{"region": "East", "product": "Gadget", "revenue": 990.0},
		{"region": "South", "product": "Gizmo", "revenue": 310.0},
	}
	grouped := tables.NewDataGrid(
		[]*tables.ColumnDef{
			{Field: "region", HeaderName: "Region", Width: 120},
			{Field: "product", HeaderName: "Product", Width: 120},
			{Field: "revenue", HeaderName: "Revenue", Type: tables.ColumnNumber, Width: 100},
		},
		func() int { return len(sales) },
		func(row int) any { return sales[row] },
	)
	grouped.API().SetRowGroupColumn("region")
	//docs:end group

	//docs:start pivot
	pivoted := tables.NewDataGrid(
		[]*tables.ColumnDef{
			{Field: "region", HeaderName: "Region", Width: 100},
			{Field: "product", HeaderName: "Product", Width: 100},
			{Field: "revenue", HeaderName: "Revenue", Type: tables.ColumnNumber, Width: 100},
		},
		func() int { return len(sales) },
		func(row int) any { return sales[row] },
	)
	// One row per region, one column per product, cells summing revenue.
	pivoted.API().SetPivot("region", "product", "revenue", pipeline.AggSum)
	//docs:end pivot

	//docs:start tree
	staff := []map[string]any{
		{"id": 1, "parent": -1, "name": "Ada", "title": "CEO"},
		{"id": 2, "parent": 1, "name": "Grace", "title": "CTO"},
		{"id": 3, "parent": 1, "name": "Linus", "title": "CFO"},
		{"id": 4, "parent": 2, "name": "Ken", "title": "Engineer"},
		{"id": 5, "parent": 2, "name": "Barbara", "title": "Engineer"},
	}
	tree := tables.NewDataGrid(
		[]*tables.ColumnDef{
			{Field: "name", HeaderName: "Name", Width: 180},
			{Field: "title", HeaderName: "Title", Width: 140},
		},
		func() int { return len(staff) },
		func(row int) any { return staff[row] },
	)
	tree.API().SetTreeData(
		func(row any) any { return row.(map[string]any)["id"] },
		func(row any) int { return row.(map[string]any)["parent"].(int) },
	)
	tree.API().ExpandAllNodes()
	//docs:end tree

	w.SetContent(container.NewGridWithColumns(3, grouped, pivoted, tree))
	bellui.ShowAndRun(w)
}
