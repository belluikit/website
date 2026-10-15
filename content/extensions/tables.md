---
title: Tables
description: A data grid with sorting, filtering, editing, pinned and grouped columns, row grouping, tree data, master-detail rows, sparklines and export.
weight: 2
---

`x/tables` provides
[`DataGrid`](https://pkg.go.dev/github.com/belluikit/bellui/x/tables#DataGrid),
a spreadsheet-like table for large data sets. It only creates widgets for the
cells you can see, so it scrolls through hundreds of thousands of rows. Users
can sort, filter, resize, reorder and pin columns, select rows or cells, edit
values, and copy and paste. Sorting, filtering, selection, editing and column
changes are also available from code.

{{< screenshot name="data-grid" alt="The gallery's data grid page" >}}

## Install

```sh
go get github.com/belluikit/bellui/x/tables@latest
```

## A first grid

A grid needs three things: the column definitions, a function returning the
number of rows, and a function returning one row. The grid never copies your
data; it calls the functions when it needs a row.

A row can be a struct, a pointer to a struct, or a `map[string]any`:

{{< example file="extensions/tables/basic/main.go" region="rows" >}}

Each column's `Field` names a struct field or a map key:

{{< example file="extensions/tables/basic/main.go" region="grid" >}}

The full program builds 100,000 orders. The grid creates cells only for the
rows in view and reuses them as you scroll.

- `HeaderName` is the header text. It defaults to `Field`.
- `Type` chooses how the column draws, sorts and filters its values:
  `ColumnText` (the default), `ColumnNumber`, `ColumnDate` or
  `ColumnBoolean`.
- `Width` is in pixels. `Flex` instead shares out the remaining width, like a
  flex-grow factor.
- `ValueFormatter` turns a value into display text. Without one, the column's
  renderer shows the raw value: `fmt.Sprint` for text and number columns, and
  the `2006-01-02` layout for `time.Time` values in date columns.
- `Field` accepts dot notation for nested values, such as `"address.city"`.
  For anything else, set `ValueGetter` to a function that extracts the value
  from the row.

`grid.Options` holds the grid-wide settings: `RowAlternation` stripes the rows,
`ShowStatusBar` adds a footer with the row count, the selection count and the
current page, and `Density` chooses compact, comfortable or spacious row
spacing. When your data changes, call
[`RefreshData`](https://pkg.go.dev/github.com/belluikit/bellui/x/tables#DataGrid.RefreshData)
so the grid reads it again.

## Sorting and filtering

Click a header to sort by that column; click again to reverse or clear the sort.
Shift-click adds the column to a multi-column sort. Right-click a header for a
menu with sort, filter, pin and hide commands. The menu's filter matches cells
containing the text you type. `Options.FloatingFilters` also adds a filter field
under every header.

From code, use the grid's
[`GridAPI`](https://pkg.go.dev/github.com/belluikit/bellui/x/tables#GridAPI),
returned by `grid.API()`:

{{< example file="extensions/tables/sortfilter/main.go" region="sortfilter" >}}

A column filter is any `func(value any) bool`. The
[`filter`](https://pkg.go.dev/github.com/belluikit/bellui/x/tables/filter)
package has ready-made ones: `TextFilter`, `NumberFilter`, `DateFilter` and
`SetFilter`, each with a `Matches` method you can pass directly. Set
`Comparator` on a column to sort it your own way. Set `Sortable` or
`Filterable` to a false value to turn either off for a column.

{{< callout type="info" >}}
Columns are sortable, filterable and resizable unless you say otherwise. If you
set `grid.DefaultColDef`, though, its fields become the defaults for every
column, and a zero `DefaultColumnDef` has all of them false. Set every field
you want on, not only the one you came to change.
{{< /callout >}}

## Selection

`grid.SelectionMode` chooses what a click selects: `SelectionSingleCell` (the
default), `SelectionMultiCell`, `SelectionSingleRow`, `SelectionMultiRow` or
`SelectionNone`. With multi-row selection, shift-click selects a range and
ctrl-click or cmd-click toggles one row. `Options.CheckboxSelection`
adds a checkbox column.

`OnSelectionChanged` receives the selected rows as indexes into your data, not
screen positions, so they stay correct after the user sorts or filters.
`OnCellClicked`, `OnCellDoubleClicked` and `OnRowActivated` (Enter or Space on
a row) report clicks and activation.

The arrow keys and Tab move the cursor. The copy shortcut copies whole rows as
tab-separated text, every column of each selected row, or of the selected
cell's row. The paste shortcut puts the clipboard text into
the current cell, if that cell is editable.

## Editing

Make a column editable with `Editable`. It is a `*bool`, so a column without
it inherits from `DefaultColDef`. The user starts editing with a double-click
or F2, commits with Enter, and cancels with Escape.

{{< example file="extensions/tables/editing/main.go" region="editing" >}}

The default editor is a text field, whatever the column type. Choose another
with `Editor`: `cell.NewSelectEditor` for a fixed list of choices,
`cell.NewCheckboxEditor` for booleans, `cell.NewDateEditor` for dates, or your
own [`cell.EditorFactory`](https://pkg.go.dev/github.com/belluikit/bellui/x/tables/cell#EditorFactory).
`CanEdit` refuses editing for individual rows of an editable column.

The undo and redo shortcuts undo and redo edits and pastes; `grid.Undo` and
`grid.Redo` do the same from code.

{{< callout type="warning" >}}
The grid writes an edited value back only when the row is a
`map[string]any`; it stores it under the column's `Field`. For struct rows,
copy the value into your data in `OnCellValueChanged`. The text editor
produces a `string`, even in a number column, so convert it there if you need
a number.
{{< /callout >}}

## Columns

{{< example file="extensions/tables/columns/main.go" region="columns" >}}

- **Pinned columns** stay in place while the rest scroll sideways. `Pinned:
  tables.PinLeft` or `tables.PinRight` sets this up; the header menu lets the
  user change it.
- **Column groups** put a shared header over several columns. A `ColumnDef`
  with `Children` is a group, and its `ColID` names it.
- **Hidden columns** are defined but not shown until you make them visible.
- Users **resize** a column by dragging the edge of its header, and **move** it
  by dragging the header. `MinWidth` and `MaxWidth` bound resizing;
  `Resizable` and `Movable` turn each off.

The same changes can be made from code:

{{< example file="extensions/tables/columns/main.go" region="columnapi" >}}

`GetGridState` returns the column widths, order, pinning and visibility, and
the sort, as a JSON-tagged struct. Save it and pass it to `SetGridState` the
next time the grid opens. `SetGridState` restores the widths, pinning,
visibility and sort, but not the column order, and neither call saves the
page.

## Grouping, pivots and tree data

**Row grouping** collects rows with the same value in one column under a group
row that the user expands and collapses:

{{< example file="extensions/tables/grouping/main.go" region="group" >}}

`ExpandAllGroups` and `CollapseAllGroups` open and close every group; pass
`""` to `SetRowGroupColumn` to remove the grouping.

A **pivot** turns the values of one column into columns of their own and
aggregates a third column into the cells, with `pipeline.AggSum`, `AggAvg`,
`AggMin`, `AggMax`, `AggCount`, `AggFirst` or `AggLast`:

{{< example file="extensions/tables/grouping/main.go" region="pivot" >}}

**Tree data** shows rows as a hierarchy, indented under their parents, with a
disclosure arrow on each parent row. You supply two functions: one returns a
row's ID, the other its parent's ID, or a negative number for a top-level row.

{{< example file="extensions/tables/grouping/main.go" region="tree" >}}

The parent function returns an `int`, and it is looked up among the values the
ID function returns. Return the ID as an `int` too, or parents will not be
found and every row will appear at the top level.

**Pagination** splits the rows into pages: `EnablePagination(50)`, then
`SetPage` and `GetPageCount` to move between them.

## Master-detail rows

A master row expands to show a band beneath it with any widget you like:

{{< example file="extensions/tables/detail/main.go" region="detail" >}}

Master-detail needs `Enabled`, a `DetailRenderer` and a `DetailHeight` above
zero. The user expands a row with the arrow at the left of its first column.
Expansion follows the row, not its position, so it survives sorting and
filtering.

## Sparklines

The [`sparkline`](https://pkg.go.dev/github.com/belluikit/bellui/x/tables/sparkline)
package draws a small chart in each cell of a column whose values are
`[]float64`. There are four kinds: `Line`, `Area`, `Bar` and `Column`.

{{< example file="extensions/tables/sparklines/main.go" region="sparklines" >}}

The same charts are available outside a grid as widgets, with
`sparkline.NewLine` and the other constructors. Hovering such a widget
highlights the nearest point and shows its value.

## Export

The API exports the rows the grid currently shows, after filtering and
sorting, with formatted values:

{{< example file="extensions/tables/export/main.go" region="export" >}}

| Method | Writes |
| --- | --- |
| `ExportCSV` | Comma-separated values with a header row |
| `ExportTSV` | Tab-separated values |
| `ExportExcel` | An `.xlsx` workbook |
| `ExportPDF` | A paginated report: headers repeated on every page, columns scaled to fit, page numbers, an optional title. `PDFOptions` sets the paper size, orientation, font size and a row limit. |
| `ExportPXF` | A typed PXF document with an embedded schema. Number columns need `ExportType` set. |

Each writes to an `io.Writer`, so you choose where the file goes.

## Licence

`x/tables` is BSD 3-Clause licensed, like Bellui. Its feature set is
modelled on that of [AG Grid](https://www.ag-grid.com/), a JavaScript data
grid.
