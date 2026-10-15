// Consent-gated analytics: nothing is recorded until the user opts in.
package main

import (
	"log"
	"os"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/analytics"
	"github.com/belluikit/bellui/analytics/posthog"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"
)

//docs:start sink
// logSink is a Sink that prints each batch instead of sending it anywhere.
type logSink struct{}

func (logSink) Send(batch []analytics.Event) error {
	for _, e := range batch {
		log.Printf("analytics: %s %q %v", e.Kind, e.Name, e.Props)
	}
	return nil
}

//docs:end sink

//docs:start choose
// sink picks PostHog when a project key is configured, else the log sink.
func sink() analytics.Sink {
	if key := os.Getenv("POSTHOG_KEY"); key != "" {
		return posthog.New(key)
	}
	return logSink{}
}

//docs:end choose

func main() {
	a := app.New()
	w := a.NewWindow("Analytics")

	//docs:start setup
	// Create the tracker after app.New, so its queue is saved in the
	// app's storage. It starts disabled.
	t := analytics.New(sink())
	analytics.SetDefault(t)
	analytics.Attach(t, a.Lifecycle())
	analytics.EnableAutoCapture(t)
	//docs:end setup

	//docs:start consent
	consent := widget.NewCheck("Share anonymous usage data", func(on bool) {
		analytics.SetEnabled(on)
	})
	//docs:end consent

	//docs:start track
	export := widget.NewButton("Export", func() {
		analytics.Track("export", map[string]any{"format": "csv"})
	})
	//docs:end track

	//docs:start nocapture
	// Taps inside this region are never auto-captured.
	private := analytics.NoCapture(container.NewVBox(
		widget.NewLabel("Account"),
		widget.NewButton("Change password", func() {}),
	))
	//docs:end nocapture

	w.SetContent(container.NewPadded(container.NewVBox(consent, export, private)))
	bellui.ShowAndRun(w)

	_ = analytics.Flush()
	_ = analytics.Close(t)
}
