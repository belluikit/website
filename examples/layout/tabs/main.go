// App tabs for top-level navigation, with document tabs the user can open
// and close inside one of them.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.NewWithID("com.example.tabs")
	w := a.NewWindow("Tabs")

	//docs:start doctabs
	docs := container.NewDocTabs(
		container.NewTabItem("untitled-1", widget.NewMultiLineEntry()),
	)
	n := 1
	// Setting CreateTab adds a "+" button to the tab bar.
	docs.CreateTab = func() *container.TabItem {
		n++
		return container.NewTabItem(fmt.Sprintf("untitled-%d", n), widget.NewMultiLineEntry())
	}
	docs.OnClosed = func(item *container.TabItem) {
		fmt.Println("closed", item.Text)
	}
	//docs:end doctabs

	//docs:start apptabs
	tabs := container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Welcome")),
		container.NewTabItem("Documents", docs),
		container.NewTabItemWithIcon("Settings", theme.SettingsIcon(), widget.NewLabel("Settings")),
	)
	tabs.SetTabLocation(container.TabLocationLeading)
	tabs.OnSelected = func(item *container.TabItem) {
		w.SetTitle("Tabs - " + item.Text)
	}
	//docs:end apptabs

	w.SetContent(tabs)
	w.Resize(geo.NewSize(640, 420))
	bellui.ShowAndRun(w)
}
