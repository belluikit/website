---
title: Concepts
description: The ideas every Bellui app is built on - the app and its windows, canvas objects, the UI thread, input events and resources.
weight: 1
---

A Bellui program creates one **app**, opens one or more **windows**, and puts
**canvas objects** in them. Widgets, containers and drawing primitives are all
canvas objects. The toolkit calls your code on one goroutine, the **UI
thread**, when the user does something. These pages explain each of those
pieces. Read them before the rest of the documentation.

The fundamental types live in small packages:

| Package | What it holds |
| --- | --- |
| [`bellui`](https://pkg.go.dev/github.com/belluikit/bellui) | `App`, `Window`, `Resource`, `ShowAndRun`, `Do` and the other app-level types |
| [`core`](https://pkg.go.dev/github.com/belluikit/bellui/core) | `CanvasObject`, `Widget`, `WidgetRenderer`, `Layout`, `Canvas` |
| [`geo`](https://pkg.go.dev/github.com/belluikit/bellui/geo) | `Size`, `Position`, `Delta` |
| [`event`](https://pkg.go.dev/github.com/belluikit/bellui/event) | Pointer and key events, and the interfaces that receive them |
| [`canvas`](https://pkg.go.dev/github.com/belluikit/bellui/canvas) | Drawing primitives: text, shapes, images, gradients |

{{< cards >}}
  {{< card link="app-and-windows/" title="The app and its windows" subtitle="Create the app, open windows, and control when the app quits." >}}
  {{< card link="canvas-objects/" title="Canvas objects" subtitle="The one interface everything on screen implements, and the drawing primitives." >}}
  {{< card link="ui-thread/" title="The UI thread" subtitle="Where your callbacks run, and how to update the UI from other goroutines." >}}
  {{< card link="input-events/" title="Input events" subtitle="Taps, hover, drags, scrolling, keys and shortcuts." >}}
  {{< card link="resources/" title="Resources and images" subtitle="Embed files in your app, use theme icons, load images." >}}
{{< /cards >}}
