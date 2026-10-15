---
title: Lists, tables and trees
description: List, Table, GridWrap, Tree, TreeView and FileTree show large collections by reusing a few row widgets.
weight: 9
---

These widgets show collections that can be far larger than the window. They
do not create a widget per item. Instead you give them three functions:

1. **length**: how many items there are;
2. **create**: build one template widget, such as a label;
3. **update**: fill a template with the data for one item.

The widget creates only as many templates as fit on screen, and reuses them as
the user scrolls. So `update` must set everything that differs between items,
and must not change the collection widget itself. When your data changes,
call `Refresh()`, or `RefreshItem(id)` for a single item.

Each of these widgets scrolls by itself, so do not put it in a
`container.NewScroll`. Give it space with a layout that stretches it, such as
the centre of `container.NewBorder`.

## Which one to use

| Your data is | Use |
| --- | --- |
| A flat list of rows | `List` |
| Rows and columns | `Table` |
| Equally sized tiles, such as photos, that wrap to the width | `GridWrap` |
| A hierarchy you already hold in your own structures | `Tree` |
| A hierarchy the user edits: rename, drag to reorder, delete | `TreeView` |
| Files and folders on disk | `FileTree` |

## List

{{< example file="widgets/list/main.go" region="list" >}}

[`List`](https://pkg.go.dev/github.com/belluikit/bellui/widget#List) stacks
rows vertically. Every row takes the template's height, unless you set a
different one with `SetItemHeight(id, height)`.

- `OnSelected` and `OnUnselected` report taps. `Select`, `Unselect` and
  `UnselectAll` change the selection from code. The arrow keys move through a
  focused list.
- `ScrollTo(id)`, `ScrollToTop()` and `ScrollToBottom()` scroll from code.
- `HideSeparators` removes the lines between rows.

`NewListWithData` takes a `binding.DataList` instead of the length and update
functions. To drive a list from a sortable, filterable model, see
`BindListModel` in [Data binding](/docs/data/).

## Table

{{< example file="widgets/table/main.go" region="table" >}}

[`Table`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Table)
works like a list in two directions. `length` returns rows and columns, and
`update` receives a
[`TableCellID`](https://pkg.go.dev/github.com/belluikit/bellui/widget#TableCellID)
with `Row` and `Col`. Columns take the template's width; make the template
label's text as long as a typical value, or call `SetColumnWidth`.
`SetRowHeight` does the same for rows.

- `ShowHeaderRow` and `ShowHeaderColumn` add headers that stay in place while
  the content scrolls. By default they show "A, B, C" and "1, 2, 3". To show
  your own, set both `CreateHeader` and `UpdateHeader`. In `UpdateHeader`,
  a header-row cell has `Row` -1 and a header-column cell has `Col` -1.
- `NewTableWithHeaders` turns both default headers on.
- `StickyRowCount` and `StickyColumnCount` keep the first data rows or
  columns in place as well.
- `OnSelected` reports the tapped cell. `ScrollTo(cell)` scrolls without
  selecting.

A table has no built-in sorting. For sortable columns over a model, see
[`BindTableModel`](https://pkg.go.dev/github.com/belluikit/bellui/widget#BindTableModel)
and [`BindTableSort`](https://pkg.go.dev/github.com/belluikit/bellui/widget#BindTableSort).

## GridWrap

{{< example file="widgets/gridwrap/main.go" region="grid" >}}

[`GridWrap`](https://pkg.go.dev/github.com/belluikit/bellui/widget#GridWrap)
has the same functions and selection methods as `List`, but lays its items
out in rows that wrap to the widget's width. `ColumnCount()` says how many
columns currently fit. `NewGridWrapWithData` takes a `binding.DataList`.

## Tree

[`Tree`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Tree) shows a
hierarchy. Every node has a unique string ID (`TreeNodeID`), and the root's ID
is the empty string `""` unless you set `Root`. `NewTree` takes four
functions: the child IDs of a node, whether a node is a branch, and the usual
create and update, which are told whether the node is a branch.

For a small, fixed tree, `NewTreeWithStrings` takes a map from each branch's
ID to its children, with the root's children under `""`. The IDs are also
shown as the text:

{{< example file="widgets/tree/main.go" region="strings" >}}

`OpenBranch`, `CloseBranch`, `ToggleBranch`, `OpenAllBranches` and
`CloseAllBranches` expand and collapse from code. `OnBranchOpened` and
`OnBranchClosed` report the user doing it. `ScrollTo(id)` opens the branches
above a node and scrolls to it. `NewTreeWithData` takes a `binding.DataTree`.

## TreeView

[`TreeView`](https://pkg.go.dev/github.com/belluikit/bellui/widget#TreeView)
is an MD3-styled tree that holds its own nodes. Each
[`TreeViewNode`](https://pkg.go.dev/github.com/belluikit/bellui/widget#TreeViewNode)
has an `ID`, `Text`, an optional `Icon` and any `Data` you want to keep with
it. You add nodes with `AddNode(parentID, node)` (use `""` for the top level)
or `InsertNode`, and the tree draws an icon and label for each.

{{< example file="widgets/tree/main.go" region="treeview" >}}

Editing features are off until you turn them on:

| Field | Lets the user |
| --- | --- |
| `AllowReorder` | drag a node to a new place under the same parent |
| `AllowReparent` | drag a node into another branch |
| `AllowRename` | double-click a node to rename it |
| `AllowDelete` | delete a node with the Delete key or the context menu |

Each change is reported to `OnMoved`, `OnRenamed` or `OnDeleted`, so you can
update your own data. `ShowGuideLines` draws indent guides, and
`ContextActions` adds your own entries to a node's right-click menu.
`RenameNode`, `MoveNode`, `DeleteNode` and `RemoveNode` make the same changes
from code.

## FileTree

[`FileTree`](https://pkg.go.dev/github.com/belluikit/bellui/widget#FileTree)
is a `Tree` of the files and folders under a root URI. It lists folders in the
background and shows "Loading..." until each listing arrives, so a slow disk
or network share does not freeze the window. Node IDs are the URIs as
strings.

{{< example file="widgets/tree/main.go" region="filetree" >}}

Set these fields before the tree is shown:

- `Filter` limits which files appear; folders always appear.
- `ShowHidden` includes hidden files.
- `Sorter` orders each folder's children. Without it, the order is whatever
  the file system returns.
- `Watch` re-lists a folder when its contents change on disk, where the
  platform supports it.

`Refresh()` re-lists every folder the tree has loaded.
