// A bound list shown in a widget.List, with items added and removed at runtime.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/data/binding"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Bound list")

	//docs:start list
	tasks := binding.NewStringList()
	_ = tasks.Set([]string{"Write the report", "Book the room"})

	list := widget.NewListWithData(tasks,
		func() core.CanvasObject {
			return widget.NewLabel("template")
		},
		func(item binding.DataItem, row core.CanvasObject) {
			row.(*widget.Label).Bind(item.(binding.String))
		},
	)
	//docs:end list

	//docs:start edit
	n := 0
	add := widget.NewButton("Add", func() {
		n++
		_ = tasks.Append(fmt.Sprintf("New task %d", n))
	})
	removeFirst := widget.NewButton("Remove first", func() {
		if tasks.Length() == 0 {
			return
		}
		first, _ := tasks.GetValue(0)
		_ = tasks.Remove(first)
	})
	//docs:end edit

	w.SetContent(container.NewBorder(container.NewHBox(add, removeFirst), nil, nil, nil, list))
	bellui.ShowAndRun(w)
}
