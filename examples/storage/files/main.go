// Reading, writing and listing files through URIs.
package main

import (
	"io"
	"os"
	"path/filepath"
	"strings"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/storage"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Files and URIs")
	out := widget.NewLabel("")

	run := widget.NewButton("Run", func() {
		var log strings.Builder
		if err := demo(&log); err != nil {
			log.WriteString("error: " + err.Error() + "\n")
		}
		out.SetText(log.String())
	})

	w.SetContent(container.NewVBox(run, out))
	bellui.ShowAndRun(w)
}

func demo(log *strings.Builder) error {
	//docs:start uris
	dir := storage.NewFileURI(filepath.Join(os.TempDir(), "bellui-demo"))
	file, err := storage.Child(dir, "hello.txt")
	if err != nil {
		return err
	}
	log.WriteString(file.String() + "\n")    // file:///…/bellui-demo/hello.txt
	log.WriteString(file.Name() + "\n")      // hello.txt
	log.WriteString(file.Extension() + "\n") // .txt
	log.WriteString(file.MimeType() + "\n")  // text/plain
	//docs:end uris

	//docs:start mkdir
	if ok, err := storage.Exists(dir); err != nil {
		return err
	} else if !ok {
		if err := storage.CreateListable(dir); err != nil {
			return err
		}
	}
	//docs:end mkdir

	//docs:start write
	wr, err := storage.Writer(file)
	if err != nil {
		return err
	}
	if _, err := wr.Write([]byte("Hello from Bellui\n")); err != nil {
		wr.Close()
		return err
	}
	if err := wr.Close(); err != nil {
		return err
	}
	//docs:end write

	//docs:start read
	r, err := storage.Reader(file)
	if err != nil {
		return err
	}
	defer r.Close()
	data, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	log.WriteString("read: " + string(data))
	//docs:end read

	//docs:start list
	children, err := storage.List(dir)
	if err != nil {
		return err
	}
	for _, child := range children {
		log.WriteString("child: " + child.Name() + "\n")
	}
	//docs:end list

	//docs:start copy
	backup, _ := storage.Child(dir, "hello.bak")
	if err := storage.Copy(file, backup); err != nil {
		return err
	}
	if err := storage.Delete(backup); err != nil {
		return err
	}
	//docs:end copy
	return nil
}
