// Keyboard shortcuts and raw key events on the desktop.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/driver/desktop"
	"github.com/belluikit/bellui/event"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.NewWithID("com.example.shortcuts")
	w := a.NewWindow("Shortcuts")
	status := widget.NewLabel("Press Ctrl+S (Cmd+S on macOS)")
	keys := widget.NewLabel("")

	//docs:start canvas
	save := &desktop.CustomShortcut{KeyName: event.KeyS, Modifier: event.KeyModifierShortcutDefault}
	w.Canvas().AddShortcut(save, func(core.Shortcut) {
		status.SetText("Saved")
	})
	//docs:end canvas

	//docs:start keys
	if dc, ok := w.Canvas().(desktop.Canvas); ok {
		dc.SetOnKeyDown(func(ev *event.KeyEvent) {
			keys.SetText("Down: " + string(ev.Name))
		})
		dc.SetOnKeyUp(func(ev *event.KeyEvent) {
			keys.SetText("Up: " + string(ev.Name))
		})
	}
	//docs:end keys

	//docs:start global
	if g, ok := a.Driver().(desktop.GlobalShortcutApp); ok {
		show := &desktop.CustomShortcut{
			KeyName:  event.KeySpace,
			Modifier: event.KeyModifierControl | event.KeyModifierAlt,
		}
		if err := g.RegisterGlobalShortcut(show, w.RequestFocus); err != nil {
			status.SetText("Global shortcut unavailable: " + err.Error())
		}
	}
	//docs:end global

	w.SetContent(container.NewVBox(status, keys))
	bellui.ShowAndRun(w)
}
