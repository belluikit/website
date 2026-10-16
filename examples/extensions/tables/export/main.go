// Export what the grid shows to CSV, Excel and PDF.
package main

import (
	"bytes"
	"os"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/dialog"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/tables"
	"github.com/belluikit/bellui/x/tables/export"
)

func main() {
	a := app.New()
	w := a.NewWindow("Export")
	w.Resize(geo.NewSize(600, 400))

	rows := []map[string]any{
		{"item": "Paper", "qty": 10},
		{"item": "Pens", "qty": 25},
		{"item": "Stapler", "qty": 2},
	}
	grid := tables.NewDataGrid(
		[]*tables.ColumnDef{
			{Field: "item", HeaderName: "Item", Width: 200},
			{Field: "qty", HeaderName: "Quantity", Type: tables.ColumnNumber, Width: 100},
		},
		func() int { return len(rows) },
		func(row int) any { return rows[row] },
	)

	//docs:start export
	save := func(name string, write func(*bytes.Buffer) error) {
		var buf bytes.Buffer
		if err := write(&buf); err != nil {
			dialog.ShowError(err, w)
			return
		}
		if err := os.WriteFile(name, buf.Bytes(), 0o644); err != nil {
			dialog.ShowError(err, w)
		}
	}
	api := grid.API()
	csv := widget.NewButton("CSV", func() {
		save("stock.csv", func(b *bytes.Buffer) error { return api.ExportCSV(b) })
	})
	xlsx := widget.NewButton("Excel", func() {
		save("stock.xlsx", func(b *bytes.Buffer) error { return api.ExportExcel(b) })
	})
	pdf := widget.NewButton("PDF", func() {
		save("stock.pdf", func(b *bytes.Buffer) error {
			return api.ExportPDF(b, export.PDFOptions{Title: "Stock", Landscape: true})
		})
	})
	//docs:end export

	w.SetContent(container.NewBorder(container.NewHBox(csv, xlsx, pdf), nil, nil, nil, grid))
	bellui.ShowAndRun(w)
}
