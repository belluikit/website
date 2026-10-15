// Tree, TreeView and FileTree for hierarchical data.
package main

import (
	"os"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/storage"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Trees")
	w.Resize(geo.NewSize(720, 400))
	status := widget.NewLabel("")

	//docs:start strings
	animals := widget.NewTreeWithStrings(map[string][]string{
		"":        {"Mammals", "Birds"},
		"Mammals": {"Cat", "Dog"},
		"Birds":   {"Robin", "Owl"},
	})
	animals.OnSelected = func(id widget.TreeNodeID) { status.SetText(id) }
	animals.OpenBranch("Mammals")
	//docs:end strings

	//docs:start treeview
	project := widget.NewTreeView()
	project.AllowRename = true
	project.AllowReorder = true
	_ = project.AddNode("", widget.TreeViewNode{ID: "src", Text: "src", Icon: theme.FolderIcon()})
	_ = project.AddNode("src", widget.TreeViewNode{ID: "main.go", Text: "main.go", Icon: theme.DocumentIcon()})
	_ = project.AddNode("", widget.TreeViewNode{ID: "go.mod", Text: "go.mod", Icon: theme.FileIcon()})
	project.OnRenamed = func(id widget.TreeViewNodeID, name string) {
		status.SetText(id + " renamed to " + name)
	}
	project.OpenBranch("src")
	//docs:end treeview

	//docs:start filetree
	dir, _ := os.Getwd()
	files := widget.NewFileTree(storage.NewFileURI(dir))
	files.Watch = true
	files.OnSelected = func(id widget.TreeNodeID) { status.SetText(id) }
	//docs:end filetree

	w.SetContent(container.NewBorder(nil, status, nil, nil,
		container.NewGridWithColumns(3, animals, project, files)))
	bellui.ShowAndRun(w)
}
