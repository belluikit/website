// File open, file save and folder dialogs.
package main

import (
	"io"
	"strconv"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/dialog"
	"github.com/belluikit/bellui/storage"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("File dialogs")
	text := widget.NewMultiLineEntry()
	status := widget.NewLabel("")

	//docs:start open
	open := widget.NewButton("Open…", func() {
		d := dialog.NewFileOpen(func(r bellui.URIReadCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if r == nil {
				return // the user cancelled
			}
			defer r.Close()

			data, err := io.ReadAll(r)
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			text.SetText(string(data))
			status.SetText("Opened " + r.URI().Name())
		}, w)
		d.SetFilter(storage.NewExtensionFileFilter([]string{".txt", ".md"}))
		d.Show()
	})
	//docs:end open

	//docs:start save
	save := widget.NewButton("Save as…", func() {
		d := dialog.NewFileSave(func(wr bellui.URIWriteCloser, err error) {
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			if wr == nil {
				return // the user cancelled
			}
			defer wr.Close()

			if _, err := wr.Write([]byte(text.Text)); err != nil {
				dialog.ShowError(err, w)
				return
			}
			status.SetText("Saved " + wr.URI().Name())
		}, w)
		d.SetFileName("notes.txt")
		d.Show()
	})
	//docs:end save

	//docs:start folder
	folder := widget.NewButton("Choose folder…", func() {
		dialog.ShowFolderOpen(func(dir bellui.ListableURI, err error) {
			if err != nil || dir == nil {
				return
			}
			items, err := dir.List()
			if err != nil {
				dialog.ShowError(err, w)
				return
			}
			status.SetText(dir.Name() + " holds " + strconv.Itoa(len(items)) + " items")
		}, w)
	})
	//docs:end folder

	//docs:start multiple
	several := widget.NewButton("Open several…", func() {
		dialog.ShowFileOpenMultiple(func(rs []bellui.URIReadCloser, err error) {
			if err != nil || rs == nil {
				return
			}
			for _, r := range rs {
				r.Close() // close every reader you are given
			}
			status.SetText("Picked " + strconv.Itoa(len(rs)) + " files")
		}, w)
	})
	//docs:end multiple

	buttons := container.NewHBox(open, save, folder, several)
	w.SetContent(container.NewBorder(buttons, status, nil, nil, text))
	bellui.ShowAndRun(w)
}
