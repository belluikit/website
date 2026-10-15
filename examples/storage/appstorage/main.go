// Documents in the app's private storage area.
package main

import (
	"errors"
	"io"
	"strings"
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/storage"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.NewWithID("com.example.appstorage")
	w := a.NewWindow("App storage")

	//docs:start root
	store := a.Storage()
	root := widget.NewLabel("Storage root: " + store.RootURI().String())
	//docs:end root

	files := widget.NewLabel("")
	refresh := func() {
		//docs:start list
		names := store.List() // the documents saved by this app
		//docs:end list
		files.SetText("Documents: " + strings.Join(names, ", "))
	}

	save := widget.NewButton("Save note", func() {
		if err := saveNote(store, "note.txt", "Saved at "+time.Now().Format(time.Kitchen)); err != nil {
			files.SetText("error: " + err.Error())
			return
		}
		refresh()
	})
	load := widget.NewButton("Load note", func() {
		text, err := loadNote(store, "note.txt")
		if err != nil {
			files.SetText("error: " + err.Error())
			return
		}
		files.SetText(text)
	})
	remove := widget.NewButton("Delete note", func() {
		_ = store.Remove("note.txt")
		refresh()
	})

	refresh()
	w.SetContent(container.NewVBox(root, save, load, remove, files))
	bellui.ShowAndRun(w)
}

//docs:start save
// saveNote creates the document the first time and overwrites it after that.
func saveNote(store bellui.Storage, name, text string) error {
	wr, err := store.Create(name)
	if errors.Is(err, storage.ErrAlreadyExists) {
		wr, err = store.Save(name)
	}
	if err != nil {
		return err
	}
	if _, err := wr.Write([]byte(text)); err != nil {
		wr.Close()
		return err
	}
	return wr.Close()
}

//docs:end save

//docs:start load
func loadNote(store bellui.Storage, name string) (string, error) {
	r, err := store.Open(name)
	if err != nil {
		return "", err
	}
	defer r.Close()
	data, err := io.ReadAll(r)
	return string(data), err
}

//docs:end load
