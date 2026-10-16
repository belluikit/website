// An Arabic, right-to-left app: translations, a mirrored layout, mixed
// Arabic and Latin text, and an optional font for Arabic letters.
package main

import (
	"embed"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/lang"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

//go:embed translations
var translations embed.FS

func main() {
	//docs:start font
	// No bundled text font has Arabic or Hebrew letters. Desktop systems
	// usually supply them; elsewhere, register a font of your own before
	// building the UI.
	if res, err := bellui.LoadResourceFromPath("NotoSansArabic-Regular.ttf"); err == nil {
		bellui.RegisterFontFallback(res)
	}
	//docs:end font

	if err := lang.AddTranslationsFS(translations, "translations"); err != nil {
		bellui.LogError("Could not load translations", err)
	}

	//docs:start locale
	// A real app follows the system locale and needs no call here. This
	// example forces Arabic so that it shows right to left on any machine.
	lang.SetLocale("ar")
	//docs:end locale

	a := app.New()
	w := a.NewWindow(lang.L("Order"))

	//docs:start ui
	bar := widget.NewTopAppBar(lang.L("Order"), theme.NavigateBackIcon(), func() {})

	name := widget.NewEntry()
	name.SetText("ليلى Haddad")
	size := widget.NewSelect([]string{lang.L("Small"), lang.L("Large")}, nil)
	size.SetSelectedIndex(0)
	form := widget.NewForm(
		widget.NewFormItem(lang.L("Name"), name),
		widget.NewFormItem(lang.L("Size"), size),
		widget.NewFormItem("", widget.NewCheck(lang.L("Gift wrap"), nil)),
	)

	// Arabic with a Latin name and numbers in it.
	note := widget.NewLabel(lang.L("Order 1024 ships with Bellui 1.0 tomorrow."))

	buttons := container.NewHBox(
		widget.NewButtonWithIcon(lang.L("Previous"), theme.NavigateBackIcon(), func() {}),
		widget.NewButtonWithIcon(lang.L("Next"), theme.NavigateNextIcon(), func() {}),
	)
	//docs:end ui

	progress := widget.NewProgressBar()
	progress.SetValue(0.3)

	w.SetContent(container.NewBorder(bar, nil, nil, nil,
		container.NewVBox(form, note, progress, buttons)))
	bellui.ShowAndRun(w)
}
