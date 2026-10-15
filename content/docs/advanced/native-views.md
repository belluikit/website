---
title: Native views
description: Place a view owned by the operating system, such as a map, a video surface or a web view, inside a Bellui layout.
weight: 5
---

Some content has to be drawn by the operating system: a platform map, a
camera preview, a hardware video surface. A
[`widget.NativeViewHost`](https://pkg.go.dev/github.com/belluikit/bellui/widget#NativeViewHost)
reserves a place in your layout and keeps an OS view positioned exactly over
it, through scrolling, resizing and layout changes.

You supply the OS view by implementing
[`widget.NativeView`](https://pkg.go.dev/github.com/belluikit/bellui/widget#NativeView).
That part is platform code (cgo, Objective-C, JNI and so on) and belongs in its
own package. The [`webview`](../webview/) package is a ready-made example.

## The contract

A `NativeView` has four methods. The example implements them with log
messages, so you can watch what the host asks for:

{{< example file="advanced/nativeview/main.go" region="view" >}}

| Method | When the host calls it | What a real view does |
| --- | --- | --- |
| `Attach(win)` | once, the first time the host is laid out in a window | creates the OS view inside the window |
| `SetGeometry(rect, clip, scale)` | when the host's place on screen changes | moves and sizes the OS view to `rect`, and masks it to `clip` |
| `SetVisible(visible)` | when the host is hidden or shown, or scrolls fully out of view and back | hides or shows the OS view |
| `Destroy()` | when the host leaves the window or the window closes | removes the OS view; must be safe to call twice, or without a successful `Attach` |

`rect` and `clip` are in Bellui's device-independent units, measured from the
window's content. Multiply by `scale` for pixels. `clip` is the part of `rect`
that is not cut off by a scroll container or other clipping parent. The host
never passes an empty clip; it hides the view instead.

To reach the platform window inside `Attach`, assert `win` to
[`driver.NativeWindow`](https://pkg.go.dev/github.com/belluikit/bellui/driver#NativeWindow)
and call `RunNative`. Its callback gets a platform context, such as
`driver.MacWindowContext` or `driver.WindowsWindowContext`.

If `Attach` returns an error, the host logs it and stays empty. It keeps its
space in the layout and draws nothing. It does not try again.

Every method is called on the UI thread.

## Host it

{{< example file="advanced/nativeview/main.go" region="host" >}}

The host has no natural size, because the OS owns the content. Its minimum
size is zero unless you call `SetMinSize`, so put it where the layout
stretches it or give it a floor.

## Limits

The OS view is not drawn by Bellui. It sits on top of Bellui's window, in the
space the host leaves empty. That has consequences you cannot work around:

- **Nothing Bellui draws can appear above it.** Pop-ups, dialogs, menus,
  tooltips, toasts and the text-selection handles on mobile are hidden where
  they overlap the native view.
- **Screenshots miss it.** `Canvas.Capture` returns whatever Bellui drew
  underneath, so the area looks empty in screenshots and image-based tests.
- **Bellui's visual effects do not reach it.** Group opacity, transforms,
  repaint caching and right-to-left mirroring apply only to what Bellui draws.
- **Clipping is up to the view.** If your `SetGeometry` ignores `clip`, the
  view spills out of a scroll container.

When Bellui can draw the content itself, prefer a normal widget.

## Keyboard focus

If your view can take keyboard input, also implement
[`widget.NativeViewFocuser`](https://pkg.go.dev/github.com/belluikit/bellui/widget#NativeViewFocuser):

- `TakeFocus()` gives the OS keyboard focus to the view.
- `ResignFocus()` gives it back to Bellui.
- `HasFocus()` reports whether the view has it. It is called every frame, so
  keep it cheap.

With it, the host keeps the two sides in step. Clicking into the view clears
Bellui's focus, so no text field keeps a blinking cursor. Focusing a Bellui
widget makes the view give up the keyboard. Without it, the two can disagree
about who receives keys.

The host is deliberately not in the Tab order. Once the OS view has focus,
Bellui never sees the Tab key, so it could not move focus back out again.
Users reach the view by clicking it, and your code can call
`host.FocusNativeView()`.
