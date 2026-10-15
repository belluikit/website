// Specialised entries: numbers, a combo box, search and completion.
package main

import (
	"fmt"
	"strings"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"
)

var fruits = []string{"Apple", "Apricot", "Banana", "Blueberry", "Cherry", "Grape", "Lemon"}

func main() {
	a := app.New()
	w := a.NewWindow("Special entries")
	status := widget.NewLabel("")

	//docs:start numeric
	price := widget.NewNumericEntry()
	price.SetPlaceHolder("Price")
	price.OnValueChanged = func(v float64) {
		status.SetText(fmt.Sprintf("Price %.2f", v))
	}

	qty := widget.NewIntegerEntry()
	qty.SetBounds(1, 99)
	qty.SetValue(1)
	//docs:end numeric

	//docs:start choice
	colour := widget.NewSelectEntry([]string{"Red", "Green", "Blue"})
	colour.SetPlaceHolder("Pick or type a colour")
	colour.OnOptionSelected = func(s string) {
		status.SetText("Picked " + s)
	}

	country := widget.NewSearchSelect(
		[]string{"Austria", "Belgium", "Brazil", "Canada", "Germany", "Japan"},
		func(s string) { status.SetText("Country " + s) })
	country.Preferred = []string{"Germany", "Austria"}
	country.PlaceHolder = "Country"
	//docs:end choice

	//docs:start search
	search := widget.NewSearchField(func(s string) {
		status.SetText("Filter: " + s)
	})
	search.OnSearched = func(s string) {
		status.SetText("Search for " + s)
	}
	//docs:end search

	//docs:start completion
	fruit := widget.NewEntry()
	fruit.SetPlaceHolder("Fruit")
	fruit.OnChanged = func(s string) {
		var items []widget.CompletionItem
		for _, f := range fruits {
			if s != "" && strings.HasPrefix(strings.ToLower(f), strings.ToLower(s)) {
				items = append(items, widget.CompletionItem{Text: f})
			}
		}
		fruit.ShowCompletion("", items) // no items hides the popup
	}
	fruit.OnCompletionSelected = func(item widget.CompletionItem) {
		fruit.SetText(item.Text)
		fruit.HideCompletion()
	}
	//docs:end completion

	w.SetContent(container.NewVBox(price, qty, colour, country, search, fruit, status))
	bellui.ShowAndRun(w)
}
