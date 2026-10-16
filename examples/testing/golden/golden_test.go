package main

import (
	"testing"

	"github.com/belluikit/bellui/test"
)

//docs:start markup
func TestStatusCard_Markup(t *testing.T) {
	test.NewTempApp(t)

	card := statusCard("Backup", "Last run: today")
	test.AssertObjectRendersToMarkup(t, "status_card.xml", card)
}

//docs:end markup

//docs:start image
func TestStatusCard_Image(t *testing.T) {
	test.NewTempApp(t)

	card := statusCard("Backup", "Last run: today")
	test.AssertObjectRendersToImage(t, "status_card.png", card)
}

//docs:end image
