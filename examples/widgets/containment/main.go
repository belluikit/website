// Containment: cards, an accordion and a carousel.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Containment")

	//docs:start card
	order := widget.NewCard("Order #1042", "Shipped yesterday",
		widget.NewButton("Track parcel", func() {}))

	plan := widget.NewCard("Pro plan", "Billed yearly", widget.NewLabel("All features included"))
	plan.Style = widget.CardStyleOutlined
	//docs:end card

	//docs:start accordion
	faq := widget.NewAccordion(
		widget.NewAccordionItem("How do I reset my password?",
			widget.NewLabel("Use the link on the sign-in page.")),
		widget.NewAccordionItem("Can I export my data?",
			widget.NewLabel("Yes, from Settings > Export.")),
	)
	faq.MultiOpen = true
	faq.Open(0)
	//docs:end accordion

	//docs:start carousel
	var slides []core.CanvasObject
	for i := 1; i <= 4; i++ {
		slides = append(slides, widget.NewCard(fmt.Sprintf("Slide %d", i), "", nil))
	}
	carousel := widget.NewCarousel(slides...)
	carousel.OnChanged = func(i int) {
		w.SetTitle(fmt.Sprintf("Containment: slide %d", i+1))
	}
	//docs:end carousel

	w.SetContent(container.NewVBox(container.NewGridWithColumns(2, order, plan), faq, carousel))
	bellui.ShowAndRun(w)
}
