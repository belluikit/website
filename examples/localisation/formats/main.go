// Locale-aware formatting of dates, times, numbers, money, lists and
// distances, and locale-aware sorting.
package main

import (
	"sort"
	"strings"
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/lang"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Formats")

	//docs:start formats
	now := time.Now()
	lines := []string{
		// month/day/year in en-US, day.month.year in de-DE
		"Date: " + lang.FormatDate(now),
		// "2:05 PM" in en-US, "14:05" in de-DE
		"Time: " + lang.FormatTime(now.Hour(), now.Minute()),
		// "MM/DD/YYYY" in en-US, "DD.MM.YYYY" in de-DE
		"Date hint: " + lang.DatePatternHint(),
		// "1,234,567.89" in en-US, "1.234.567,89" in de-DE
		"Number: " + lang.FormatDecimal(1234567.891, 2),
		// "$19.99" in en-US, "19,99 €" in de-DE
		"Price: " + lang.FormatCurrency(19.99, lang.DefaultCurrency()),
		// "red, green, and blue" in English, "red, green und blue" in German
		"Colours: " + lang.FormatList([]string{"red", "green", "blue"}),
		// "1.24 mi" in en-US, "2 km" in de-DE
		"Distance: " + lang.FormatDistance(2000),
	}
	//docs:end formats

	//docs:start clock
	if lang.Prefers24HourClock() {
		lines = append(lines, "This locale uses a 24-hour clock")
	}
	//docs:end clock

	//docs:start parse
	typed := "1.234,5" // what a German user might type
	if v, err := lang.ParseLocaleFloat(typed); err == nil {
		lines = append(lines, "Parsed: "+lang.FormatDecimal(v, -1))
	}
	//docs:end parse

	//docs:start sort
	names := []string{"Zoë", "Ärger", "apple", "Österreich", "zebra"}
	sort.Slice(names, func(i, j int) bool {
		return lang.Compare(names[i], names[j]) < 0
	})
	lines = append(lines, "Sorted: "+strings.Join(names, ", "))
	//docs:end sort

	w.SetContent(container.NewVBox(widget.NewLabel(strings.Join(lines, "\n"))))
	bellui.ShowAndRun(w)
}
