---
title: Testing
description: Test your UI with go test, an in-memory app and window, simulated input and golden files.
weight: 11
---

The [`test`](https://pkg.go.dev/github.com/belluikit/bellui/test) package lets
you test a Bellui UI with ordinary `go test`. It replaces the real app with an
in-memory one: windows are never shown, nothing needs a GPU or a display, and
tests run on any machine, including CI servers.

A UI test usually does three things:

1. Creates a test app and puts the UI in a test window.
2. Simulates what a user does: taps, typing, dragging.
3. Checks the result, either by reading widget state or by comparing the
   rendering with a saved golden file.

{{< cards >}}
  {{< card link="interaction/" title="Simulating input" subtitle="Test apps and windows, taps, typing, drags, and the assertion helpers." >}}
  {{< card link="rendering/" title="Golden files" subtitle="Compare a widget's rendering with saved markup or images." >}}
{{< /cards >}}

## What the test app provides

[`test.NewApp`](https://pkg.go.dev/github.com/belluikit/bellui/test#NewApp)
creates the app and makes it the current one, so code that calls
`bellui.CurrentApp()` works. It comes with:

- a test driver whose windows live only in memory;
- in-memory preferences, so tests never touch the user's real settings;
- app storage rooted in the system's temporary folder;
- its own clipboard;
- a fixed test theme, so results do not depend on the user's theme settings.

On the test driver, `bellui.Do` and `bellui.DoAndWait` run the function
straight away and return when it has finished. Code that hands work to the UI
thread can be tested without waiting for an event loop.

Prefer [`test.NewTempApp(t)`](https://pkg.go.dev/github.com/belluikit/bellui/test#NewTempApp)
and [`test.NewTempWindow(t, content)`](https://pkg.go.dev/github.com/belluikit/bellui/test#NewTempWindow)
in tests. They do the same as `NewApp` and `NewWindow`, and clean up when the
test ends.
