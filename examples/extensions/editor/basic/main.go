// A Go editor with line numbers and a save handler.
package main

import (
	"log"
	"os"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/editor"
	"github.com/belluikit/bellui/x/editor/highlight/golang"
)

const source = `package main

import "fmt"

// main prints a greeting.
func main() {
	name := "Bellui"
	fmt.Printf("Hello, %s!\n", name)
}
`

func main() {
	a := app.New()
	w := a.NewWindow("main.go")
	w.Resize(geo.NewSize(700, 450))

	//docs:start editor
	ed := editor.NewEditor()
	ed.Highlighter = golang.New()
	ed.Gutter = editor.NewLineNumberGutter()
	ed.SetText(source)

	ed.OnSave = func() { // Cmd+S on macOS
		if err := os.WriteFile("main.go.txt", []byte(ed.Text()), 0o644); err != nil {
			log.Println(err)
		}
	}
	//docs:end editor

	w.SetContent(ed)
	bellui.ShowAndRun(w)
}
