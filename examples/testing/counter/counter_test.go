package main

import (
	"testing"

	"github.com/belluikit/bellui/test"
)

//docs:start tap
func TestCounter_TypeAndTap(t *testing.T) {
	test.NewTempApp(t) // an in-memory app, torn down when the test ends

	c := newCounter()
	w := test.NewTempWindow(t, c.content())
	w.Resize(w.Content().MinSize())

	test.Type(c.name, "Ada")
	test.Tap(c.add)
	test.Tap(c.add)

	if got, want := c.output.Text, "Ada clicked 2 times"; got != want {
		t.Errorf("label = %q, want %q", got, want)
	}
}

//docs:end tap

//docs:start asserts
func TestCounter_Layout(t *testing.T) {
	test.NewTempApp(t)

	c := newCounter()
	w := test.NewTempWindow(t, c.content())
	w.Resize(w.Content().MinSize())

	test.AssertVisible(t, c.add)
	test.AssertEnabled(t, c.add)
	test.AssertMinSizeAtLeast(t, c.add, 40, 20)

	c.add.Disable()
	test.AssertDisabled(t, c.add)
}

//docs:end asserts
