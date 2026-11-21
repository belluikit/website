// Swipe areas: one that reports swipes, and a card that follows the finger.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Swipe")
	w.Resize(geo.NewSize(360, 480))
	status := widget.NewLabel("Swipe something")

	//docs:start swipe
	pages := widget.NewLabel("Swipe left or right")
	area := widget.NewSwipeArea(pages, func(d widget.SwipeDirection) {
		status.SetText(fmt.Sprint("Swiped ", d))
	})
	//docs:end swipe

	//docs:start card
	names := []string{"Ada", "Grace", "Linus", "Margaret"}
	next := 0
	card := widget.NewCard(names[0], "Swipe right to keep, left to skip", nil)

	deck := widget.NewSwipeArea(card, func(d widget.SwipeDirection) {
		status.SetText(fmt.Sprintf("%s: %s", names[next], d))
		// The card is back in place: show the next item on it.
		next = (next + 1) % len(names)
		card.SetTitle(names[next])
	})
	deck.Follow = true
	deck.Directions = map[widget.SwipeDirection]widget.SwipeCardMode{
		widget.SwipeUp:   widget.SwipeIgnore,
		widget.SwipeDown: widget.SwipeIgnore,
	}
	deck.OnSwipeProgress = func(d widget.SwipeDirection, p float32) {
		if p > 0 {
			status.SetText(fmt.Sprintf("%s %.0f%%", d, p*100))
		}
	}
	//docs:end card

	w.SetContent(container.NewBorder(area, status, nil, nil, deck))
	bellui.ShowAndRun(w)
}
