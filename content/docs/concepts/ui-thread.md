---
title: The UI thread
description: Your callbacks run on the main goroutine; use bellui.Do, DoAndWait and RunOnMain to update the UI from other goroutines.
weight: 3
---

Bellui runs the user interface on one goroutine, the **main goroutine** (the
UI thread). Input events, widget callbacks such as a button's tap function,
and drawing all happen there, one at a time. This has two consequences:

- Code in a callback can change the UI directly. No locking is needed.
- Code in a callback must return quickly. While it runs, the window cannot
  redraw or respond to input.

## Slow work goes on another goroutine

To do something slow, such as a network request or a large file read, start a
goroutine from the callback. When the work is done, hand the UI update back
to the main goroutine with
[`bellui.Do`](https://pkg.go.dev/github.com/belluikit/bellui#Do):

{{< example file="concepts/thread/main.go" region="work" >}}

The button's function runs on the main goroutine, so it disables the button
and sets the label directly. The goroutine sleeps without blocking anything.
Its result goes back through `bellui.Do`.

## Updating the UI from a goroutine

The same applies to any goroutine you start, such as a timer:

{{< example file="concepts/thread/main.go" region="clock" >}}

`clock.Text` is a plain struct field, and `Refresh` redraws the object. Both
must run on the main goroutine, so they are wrapped in `bellui.Do`.

The whole program:

{{< example "concepts/thread/main.go" >}}

## The functions

| Function | What it does |
| --- | --- |
| [`bellui.Do(fn)`](https://pkg.go.dev/github.com/belluikit/bellui#Do) | Queues `fn` to run on the main goroutine and returns at once. |
| [`bellui.DoAndWait(fn)`](https://pkg.go.dev/github.com/belluikit/bellui#DoAndWait) | Runs `fn` on the main goroutine and waits until it has finished. Use it when you need a result back, or need the change made before you go on. |
| [`bellui.RunOnMain(fn)`](https://pkg.go.dev/github.com/belluikit/bellui#RunOnMain) | Runs `fn` straight away if the caller is already on the main goroutine; otherwise works like `DoAndWait`. |
| [`bellui.IsMainGoroutine()`](https://pkg.go.dev/github.com/belluikit/bellui#IsMainGoroutine) | Reports whether the caller is on the main goroutine. |

`Do` and `DoAndWait` are for goroutines you created. Do not call `DoAndWait`
from the main goroutine, for example from inside a callback. By default it
then logs an error and runs your function on a new goroutine, not on the main
one. If a
function can be reached both from callbacks and from background goroutines,
use `RunOnMain`. It is safe from either.

## What is already safe

Some widget setters move their work to the main goroutine for you, using
`RunOnMain`. Examples are `Label.SetText`, `Container.Add`, and
`Enable` and `Disable` on widgets that can be disabled. You can call those
from any goroutine.

Do not rely on that for everything:

- Assigning to a field (`label.Text = "..."`, `rect.FillColor = ...`) is
  never marshalled. Neither is the `Refresh` you call afterwards.
- Methods on [canvas primitives](../canvas-objects/) are not marshalled.
- Not every widget method is marshalled. `Container.Remove` is not, for
  example.

When in doubt, wrap the change in `bellui.Do`.

{{< callout type="info" >}}
Callbacks that the toolkit calls for you, such as `OnTapped`, `OnChanged`
and `OnSubmitted` on widgets, and the event methods in
[Input events](../input-events/), already run on the main goroutine. They do
not need `bellui.Do`.
{{< /callout >}}
