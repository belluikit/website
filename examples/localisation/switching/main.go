// Switching the app's language while it runs, with text that re-translates
// itself.
package main

import (
	"embed"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/data/binding"
	"github.com/belluikit/bellui/lang"
	"github.com/belluikit/bellui/widget"
)

//go:embed translations
var translations embed.FS

func main() {
	_ = lang.AddTranslationsFS(translations, "translations")

	a := app.New()
	w := a.NewWindow("")

	//docs:start bound
	// A label whose text is re-translated on every language change.
	welcome := widget.NewLabelWithData(
		binding.NewLocalizedString("Welcome, {{.Name}}!", map[string]any{"Name": "Ada"}),
	)

	// A plural that follows both the language and the count.
	count := binding.NewInt()
	_ = count.Set(1)
	selected := widget.NewLabelWithData(
		// With no template data, {{.Count}} is the live count.
		binding.NewLocalizedPluralWithKey("files_selected", "{{.Count}} files selected", count),
	)
	more := widget.NewButton("+1", func() {
		n, _ := count.Get()
		_ = count.Set(n + 1)
	})
	//docs:end bound

	//docs:start bindtext
	// Surfaces without data binding: a button label and the window title.
	save := widget.NewButton("", func() {})
	lang.BindText("button.save", save.SetText)
	lang.BindText("Welcome, {{.Name}}!", w.SetTitle, map[string]any{"Name": "Ada"})
	//docs:end bindtext

	//docs:start switch
	language := widget.NewSelect([]string{"System", "English", "Deutsch", "العربية"}, func(s string) {
		switch s {
		case "English":
			lang.SetLocale("en-US")
		case "Deutsch":
			lang.SetLocale("de-DE")
		case "العربية":
			lang.SetLocale("ar-EG") // also switches the layout to right-to-left
		default:
			lang.SetLocale("") // back to the system locale
		}
	})
	//docs:end switch

	w.SetContent(container.NewVBox(language, welcome, container.NewHBox(selected, more), save))
	bellui.ShowAndRun(w)
}
