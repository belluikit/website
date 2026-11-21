package main

import (
	"testing"
	"time"

	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/test"
	"github.com/belluikit/bellui/widget"
)

//docs:start test
func TestIdle(t *testing.T) {
	test.NewTempApp(t) // starts LastUserActivity at the zero time

	tapped := false
	b := widget.NewButton("Open", func() { tapped = true })
	w := test.NewTempWindow(t, b)
	w.Resize(b.MinSize())

	if !isIdle(time.Now()) {
		t.Fatal("a fresh app should be idle")
	}

	// A simulated tap counts as user input, as a real one does.
	test.Tap(b)
	if !tapped || isIdle(time.Now()) {
		t.Fatal("the app should not be idle right after a tap")
	}

	// Ask about a moment past the idle time instead of waiting for it.
	if !isIdle(time.Now().Add(idleAfter)) {
		t.Error("the app should be idle two minutes after the tap")
	}
}

//docs:end test

//docs:start bench
func BenchmarkCard(b *testing.B) {
	test.NewTempApp(b)
	for b.Loop() {
		card := widget.NewCard("Title", "", container.NewVBox(
			widget.NewLabel("Body"), widget.NewButton("Open", nil)))
		test.WidgetRenderer(card).Layout(card.MinSize()) // builds the renderers
		test.DestroyRenderers(b, card)                   // or they stay cached
	}
}

//docs:end bench
