// Desktop integration: main menu, system tray, notifications, more windows.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/dialog"
	"github.com/belluikit/bellui/driver/desktop"
	"github.com/belluikit/bellui/event"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.NewWithID("com.example.desktop")
	w := a.NewWindow("Desktop")
	w.Resize(geo.NewSize(420, 300))

	//docs:start menu
	newItem := bellui.NewMenuItem("New window", func() { openSecondWindow(a) })
	newItem.Shortcut = &desktop.CustomShortcut{
		KeyName:  event.KeyN,
		Modifier: event.KeyModifierShortcutDefault, // Cmd on macOS, Ctrl elsewhere
	}
	file := bellui.NewMenu("File", newItem)
	help := bellui.NewMenu("Help",
		bellui.NewMenuItem("About", func() {
			dialog.ShowInformation("About", "A desktop example.", w)
		}))
	w.SetMainMenu(bellui.NewMainMenu(file, help))
	//docs:end menu

	//docs:start tray
	// The tray is reached through the driver: in v1.0.0 the app value
	// itself does not implement all of desktop.App.
	if desk, ok := a.Driver().(desktop.App); ok {
		desk.SetSystemTrayMenu(bellui.NewMenu("Tray",
			bellui.NewMenuItem("Show window", w.Show),
		))
		desk.SetSystemTrayIcon(theme.HomeIcon())
		desk.SetSystemTrayTooltip("Desktop example")

		// Tapping the tray icon shows w, and closing w now only hides it.
		desk.SetSystemTrayWindow(w)
	}
	//docs:end tray

	//docs:start notify
	notify := widget.NewButton("Send notification", func() {
		bellui.SendNotification(bellui.NewNotification("Export finished", "3 files written"))
	})
	rich := widget.NewButton("Send notification with actions", func() {
		bellui.SendNotificationWithOptions(&bellui.NotificationOptions{
			Title:    "Download ready",
			Content:  "report.pdf",
			OnTapped: w.Show,
			Actions: []bellui.NotificationAction{
				{Label: "Open", Callback: func() { w.Show() }},
			},
		})
	})
	//docs:end notify

	w.SetContent(container.NewVBox(
		widget.NewLabel("Use the File menu, the tray icon or the buttons."),
		notify, rich,
	))
	bellui.ShowAndRun(w)
}

//docs:start window
func openSecondWindow(a bellui.App) {
	w2 := a.NewWindow("Second window")
	w2.SetContent(widget.NewLabel("Each desktop window is its own OS window."))
	w2.Resize(geo.NewSize(300, 120))
	w2.Show()
}

//docs:end window
