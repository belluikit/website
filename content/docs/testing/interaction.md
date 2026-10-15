---
title: Simulating input
description: Build your UI in a test window, tap and type into it, and check the result.
weight: 1
---

## Make the UI reachable from a test

A test needs to get at the widgets it taps and reads. The simplest way is to
build the UI in a function that returns them, and call that function from both
`main` and the test:

{{< example file="testing/counter/main.go" region="ui" >}}

## Tap, type, check

Put the test next to the code, in a `_test.go` file of the same package:

{{< example file="testing/counter/counter_test.go" region="tap" >}}

Run it with `go test ./...`. No window opens.

- [`test.Type`](https://pkg.go.dev/github.com/belluikit/bellui/test#Type)
  focuses the widget and then sends the text one character at a time, as
  typing would.
- [`test.Tap`](https://pkg.go.dev/github.com/belluikit/bellui/test#Tap)
  taps the widget. It runs the widget's tap handling, including moving focus,
  so the button's `OnTapped` function runs.
- After that, read the widgets' fields, such as `Label.Text`, and compare them
  with what you expect.

Resizing the window to its content's minimum size, as above, gives every
widget a real size and position. Do this when the test depends on layout, or
when it taps a position on the canvas rather than a widget.

## Other input

| Function | Simulates |
| --- | --- |
| `Tap`, `TapAt(obj, pos)` | A primary click or tap on a widget, optionally at a position inside it. |
| `DoubleTap` | A double click. |
| `TapSecondary`, `TapSecondaryAt` | A right click. |
| `LongPress`, `LongPressAt` | A long press on a touch screen. |
| `TapCanvas(c, pos)` | A tap at an absolute position on a canvas, delivered to whatever is there. |
| `Type(obj, text)` | Typing into a focusable widget. |
| `TypeOnCanvas(c, text)` | Typing to the canvas's own key handler. |
| `Drag(c, pos, dx, dy)` | A drag that starts at `pos` and moves by `dx`, `dy`. `DragSteps` does it in several steps. |
| `Scroll(c, pos, dx, dy)` | A scroll wheel or trackpad scroll at a position. |
| `MoveMouse(c, pos)` | Moving the pointer, so hover effects run. |
| `FocusNext(c)`, `FocusPrevious(c)` | Moving focus to the next or previous widget, as Tab and Shift-Tab do. |
| `Compose`, `ComposeCommit`, `ComposeCancel` | Text entry through an input method, as used for Chinese, Japanese and Korean. |

All of these are in the [`test`](https://pkg.go.dev/github.com/belluikit/bellui/test)
package. The canvas `c` is usually `w.Canvas()` of your test window.

## Assertion helpers

The package has a few assertions for common checks. They report a failure
through `t`, and most also return `false`, so you can stop the test early if
you need to.

{{< example file="testing/counter/counter_test.go" region="asserts" >}}

| Function | Checks |
| --- | --- |
| `AssertVisible`, `AssertHidden` | The object's `Visible()` state. |
| `AssertEnabled`, `AssertDisabled` | A widget that can be disabled. |
| `AssertWidgetSize(t, obj, w, h)` | The current size, to within 1.5 units. |
| `AssertWidgetPosition(t, obj, x, y)` | The current position, to within 1.5 units. |
| `AssertMinSizeAtLeast(t, obj, w, h)` | The minimum size is at least this big. |
| `AssertFitsWithin(t, obj, w, h)` | The current size is no bigger than this. |
| `AssertCanvasTappableAt(t, c, pos)` | Something tappable is at this position. |
| `AssertNotificationSent(t, n, f)` | Running `f` sends the notification `n`. |

## Themes in tests

The test app starts with a test theme whose colours are deliberately plain.
To check how your UI looks with another theme,
[`test.ApplyTheme(t, th)`](https://pkg.go.dev/github.com/belluikit/bellui/test#ApplyTheme)
switches the theme and switches it back when the test ends.

## The complete test

{{< example "testing/counter/counter_test.go" >}}
