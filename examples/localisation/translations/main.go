// Translating an app's own strings: loading translation files, plain and
// templated strings, plural forms, keys and grammatical variants.
package main

import (
	"embed"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/lang"
	"github.com/belluikit/bellui/widget"
)

//docs:start embed
//go:embed translations
var translations embed.FS

//docs:end embed

//docs:start types
// User fills {Name} from its exported field.
type User struct {
	Name string
}

// Stats has a method, which a placeholder cannot call.
type Stats struct {
	files []string
}

func (s Stats) Count() int { return len(s.files) }

//docs:end types

func main() {
	//docs:start load
	if err := lang.AddTranslationsFS(translations, "translations"); err != nil {
		bellui.LogError("Could not load translations", err)
	}
	//docs:end load

	a := app.New()
	w := a.NewWindow("Translations")

	//docs:start use
	name := "Ada"
	welcome := lang.L("Welcome, {Name}!", map[string]any{"Name": name})

	n := 3
	selected := lang.N("files_selected", n, map[string]any{"Count": n})

	save := lang.X("button.save", "Save")

	gender := "female"
	update := lang.LSelect("profile_updated", gender, map[string]any{"Name": name})
	//docs:end use

	//docs:start several
	user := User{Name: "Ada"}
	stats := Stats{files: []string{"a.txt", "b.txt"}}
	sent := lang.L("{Name} sent {Count} files", user, map[string]any{"Count": stats.Count()})
	//docs:end several

	//docs:start fill
	progress := lang.Fill("{Done} of {Total}", map[string]any{"Done": 2, "Total": 5})
	//docs:end fill

	//docs:start locale
	info := widget.NewLabel("System locale: " + lang.SystemLocale().String() +
		", in use: " + lang.CurrentLocale().String())
	//docs:end locale

	w.SetContent(container.NewVBox(
		widget.NewLabel(welcome),
		widget.NewLabel(selected),
		widget.NewButton(save, func() {}),
		widget.NewLabel(update),
		widget.NewLabel(sent),
		widget.NewLabel(progress),
		info,
	))
	bellui.ShowAndRun(w)
}
