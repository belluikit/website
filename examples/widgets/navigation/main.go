// An app shell: top app bar, navigation bar on narrow windows, a rail beside
// the content, and a navigation drawer behind the menu icon.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

var pages = []string{"Home", "Search", "Profile"}

func main() {
	a := app.New()
	w := a.NewWindow("Navigation")
	w.Resize(geo.NewSize(480, 400))
	content := widget.NewLabel("Home")

	//docs:start bars
	var bar *widget.NavigationBar
	var rail *widget.NavigationRail
	show := func(i int) {
		content.SetText(pages[i])
		bar.Selected, rail.Selected = i, i
		bar.Refresh()
		rail.Refresh()
	}

	bar = widget.NewNavigationBar([]*widget.NavigationBarItem{
		{Label: "Home", Icon: theme.HomeIcon()},
		{Label: "Search", Icon: theme.SearchIcon()},
		{Label: "Profile", Icon: theme.AccountIcon()},
	}, show)

	rail = widget.NewNavigationRail([]*widget.NavigationRailItem{
		{Label: "Home", Icon: theme.HomeIcon()},
		{Label: "Search", Icon: theme.SearchIcon()},
		{Label: "Profile", Icon: theme.AccountIcon()},
	}, show)
	//docs:end bars

	//docs:start top
	drawer := widget.NewNavigationDrawer([]*widget.NavigationDrawerItem{
		{Label: "Home", Icon: theme.HomeIcon()},
		{Label: "Search", Icon: theme.SearchIcon()},
		{Label: "Profile", Icon: theme.AccountIcon()},
		{IsDivider: true},
		{Label: "About", Icon: theme.InfoIcon()},
	}, func(i int) {
		// i indexes Items, dividers included, so "About" is 4.
		if i < len(pages) {
			show(i)
			return
		}
		content.SetText("About")
	})

	top := widget.NewTopAppBar("My app", theme.MenuIcon(), func() {
		drawer.Toggle(w.Canvas())
	})
	top.Actions = []widget.TopAppBarAction{
		{Icon: theme.SettingsIcon(), OnTapped: func() { content.SetText("Settings") }},
	}
	//docs:end top

	w.SetContent(container.NewBorder(top, bar, rail, nil, content))
	bellui.ShowAndRun(w)
}
