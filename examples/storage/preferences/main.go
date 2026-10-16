// Typed app preferences that persist between runs.
package main

import (
	"strconv"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"
)

func main() {
	//docs:start setup
	// Preferences are stored per app, so the app needs a unique ID.
	a := app.NewWithID("com.example.preferences")
	prefs := a.Preferences()
	//docs:end setup
	w := a.NewWindow("Preferences")

	//docs:start read
	launches := bellui.PrefOr(prefs, "launches", 0) + 1
	bellui.SetPref(prefs, "launches", launches)

	name, ok := bellui.Pref[string](prefs, "name")
	if !ok {
		name = "stranger"
	}
	//docs:end read

	greeting := widget.NewLabel("Hello " + name + ", launch number " + strconv.Itoa(launches))

	//docs:start write
	entry := widget.NewEntry()
	entry.SetText(name)
	entry.OnChanged = func(s string) {
		bellui.SetPref(prefs, "name", s)
	}

	dark := widget.NewCheck("Dark mode", func(on bool) {
		bellui.SetPref(prefs, "dark", on)
	})
	dark.SetChecked(bellui.PrefOr(prefs, "dark", false))
	//docs:end write

	//docs:start lists
	recent := bellui.PrefListOr(prefs, "recent", []string{})
	recent = append([]string{"opened at launch " + strconv.Itoa(launches)}, recent...)
	if len(recent) > 3 {
		recent = recent[:3]
	}
	bellui.SetPrefList(prefs, "recent", recent)
	//docs:end lists

	//docs:start listen
	changes := widget.NewLabel("")
	prefs.AddChangeListener(func() {
		// Runs on the goroutine that made the change: the UI thread when your
		// code sets a value, a background one when the file changes on disk.
		bellui.RunOnMain(func() {
			changes.SetText("Name is now " + bellui.PrefOr(prefs, "name", ""))
		})
	})
	//docs:end listen

	reset := widget.NewButton("Forget my name", func() {
		prefs.RemoveValue("name")
	})

	w.SetContent(container.NewVBox(greeting, entry, dark, reset, changes,
		widget.NewLabel("Recent: "+joinLines(recent))))
	bellui.ShowAndRun(w)
}

func joinLines(items []string) string {
	s := ""
	for _, it := range items {
		s += "\n  " + it
	}
	return s
}
