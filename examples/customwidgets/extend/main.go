// Two widgets made by extending existing ones: an icon that can be tapped,
// and an entry that only accepts digits.
package main

import (
	"unicode"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/event"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

//docs:start icon
// tappableIcon is a widget.Icon that also responds to taps.
type tappableIcon struct {
	widget.Icon
	onTapped func()
}

func newTappableIcon(res bellui.Resource, tapped func()) *tappableIcon {
	icon := &tappableIcon{onTapped: tapped}
	// Pass the outer type, so the toolkit sees tappableIcon (and its
	// Tapped method) rather than the embedded widget.Icon.
	icon.ExtendBaseWidget(icon)
	icon.SetResource(res)
	return icon
}

// Tapped adds event.Tappable to what widget.Icon already does.
func (t *tappableIcon) Tapped(*event.PointEvent) {
	t.onTapped()
}

//docs:end icon

//docs:start entry
// digitsEntry is a widget.Entry that ignores typed characters other than digits.
type digitsEntry struct {
	widget.Entry
}

func newDigitsEntry() *digitsEntry {
	e := &digitsEntry{}
	e.ExtendBaseWidget(e)
	return e
}

// TypedRune replaces the Entry's method; call the original for the runes
// you accept.
func (e *digitsEntry) TypedRune(r rune) {
	if unicode.IsDigit(r) {
		e.Entry.TypedRune(r)
	}
}

//docs:end entry

func main() {
	a := app.NewWithID("com.example.extend")
	w := a.NewWindow("Extending widgets")

	status := widget.NewLabel("Tap the icon")
	icon := newTappableIcon(theme.InfoIcon(), func() {
		status.SetText("Icon tapped")
	})
	number := newDigitsEntry()
	number.SetPlaceHolder("Digits only")

	w.SetContent(container.NewVBox(container.NewHBox(icon, status), number))
	bellui.ShowAndRun(w)
}
