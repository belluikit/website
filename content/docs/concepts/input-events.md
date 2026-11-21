---
title: Input events
description: How taps, hover, drags, scrolling, keys and shortcuts reach your objects, what differs between desktop and touch, and how to tell when the user is idle.
weight: 4
---

Bellui delivers input to an object by checking which interfaces it
implements. There is no event registration: if the object under the pointer
has a `Tapped` method with the right signature, a click calls it. The
interfaces are in the [`event`](https://pkg.go.dev/github.com/belluikit/bellui/event)
package.

Ready-made widgets already implement the interfaces they need; a button is
`Tappable`, an entry is `Focusable`. You implement them yourself when you
[build a widget](/docs/custom-widgets/). The canvas primitives implement
none of them, by design.

## A widget that receives everything

This widget implements every pointer interface except `DoubleTappable` (see
the warning below) and reports what it gets:

{{< example file="concepts/events/main.go" region="pad" >}}

`widget.BaseWidget`, `ExtendBaseWidget` and `CreateRenderer` are explained
in [Custom widgets](/docs/custom-widgets/). Every event method runs on the
[main goroutine](../ui-thread/).

## Pointer interfaces

| Interface | Methods | Delivered when |
| --- | --- | --- |
| [`Tappable`](https://pkg.go.dev/github.com/belluikit/bellui/event#Tappable) | `Tapped(*PointEvent)` | A click or tap is pressed and released on the object |
| [`DoubleTappable`](https://pkg.go.dev/github.com/belluikit/bellui/event#DoubleTappable) | `DoubleTapped(*PointEvent)` | Two taps in quick succession |
| [`SecondaryTappable`](https://pkg.go.dev/github.com/belluikit/bellui/event#SecondaryTappable) | `TappedSecondary(*PointEvent)` | A right-click on desktop; on touch, see below |
| [`LongPressable`](https://pkg.go.dev/github.com/belluikit/bellui/event#LongPressable) | `LongPressed(*PointEvent)` | The pointer is held down without moving |
| [`Hoverable`](https://pkg.go.dev/github.com/belluikit/bellui/event#Hoverable) | `HoverIn`, `HoverMoved`, `HoverOut` | A mouse pointer enters, moves over, or leaves the object |
| [`Draggable`](https://pkg.go.dev/github.com/belluikit/bellui/event#Draggable) | `Dragged(*DragEvent)`, `DragEnd()` | The pointer moves while pressed, and then is released |
| [`Scrollable`](https://pkg.go.dev/github.com/belluikit/bellui/event#Scrollable) | `Scrolled(*ScrollEvent)` | A mouse wheel or touchpad scrolls over the object |
| [`Pinchable`](https://pkg.go.dev/github.com/belluikit/bellui/event#Pinchable) | `Pinched(*PinchEvent)`, `PinchEnded()` | Two fingers on the object move apart or together; `PinchEnded` once when either lifts |

A [`PointEvent`](https://pkg.go.dev/github.com/belluikit/bellui/event#PointEvent)
has `Position`, relative to the object's top-left corner, and
`AbsolutePosition`, relative to the window. `DragEvent.Dragged` and
`ScrollEvent.Scrolled` are a [`geo.Delta`](https://pkg.go.dev/github.com/belluikit/bellui/geo#Delta):
how far the pointer moved since the last event. `ScrollEvent.Source` says
whether a scroll came from a mouse wheel, a touchpad or a touch gesture.
A [`PinchEvent`](https://pkg.go.dev/github.com/belluikit/bellui/event#PinchEvent)
is a `PointEvent` at the point between the fingers, plus `Scale`: the change
since the previous event, above 1 when the fingers moved apart. Multiply
successive values for the total zoom.

{{< callout type="warning" >}}
If an object implements `DoubleTappable`, a single tap is not delivered
straight away. The toolkit waits to see whether a second tap follows, and
calls `Tapped` only after that delay has passed. Implement `DoubleTappable`
only when you need it.
{{< /callout >}}

## Desktop and touch

The same interfaces work on every platform, but some gestures map
differently:

| Gesture | Desktop and web | Touch |
| --- | --- | --- |
| Tap | Primary button click | Finger tap |
| Secondary tap | Right-click | Lifting the finger after a press and hold |
| Long press | Holding the primary button down; it fires while the button is still down, and the release does not also deliver `Tapped` | Holding a finger still; it fires while the finger is still down |
| Hover | Mouse movement | Not delivered: there is no pointer between touches |
| Scroll | Mouse wheel or touchpad | Dragging a finger over scrollable content, such as a scroll container, or two fingers moving together anywhere |
| Pinch | Not delivered from a trackpad yet | Two fingers moving apart or together, when both are on a `Pinchable` |

On touch, a widget that implements both `LongPressable` and
`SecondaryTappable` gets `LongPressed` while the finger is held, and then
`TappedSecondary` when it lifts.

A pinch starts only when both fingers land on the same `Pinchable`. Two
fingers on widgets either side of it, such as a slider and a button, keep
their own drag and tap. Neither finger of a pinch or a two-finger scroll
produces a tap or a drag. Two fingers moving together scroll the content
under them, so a page of charts or 3D views, which keep one-finger drags for
themselves, still scrolls.

## Drags on touch screens

Two optional interfaces refine how a `Draggable` behaves under a finger. They
change nothing for a mouse drag.

{{< example file="concepts/drag/main.go" region="strip" >}}

- [`event.DragClaimer`](https://pkg.go.dev/github.com/belluikit/bellui/event#DragClaimer)
  lets a widget take a touch drag in some directions only. When a finger
  starts to drag, Bellui asks the `Draggable` widgets under it, innermost
  first, with `ClaimsDrag(delta)`, where `delta` is the first movement. The
  first one that returns true gets the drag until the finger lifts. A drag
  that nobody claims scrolls whatever can scroll there. A `Draggable` without
  `ClaimsDrag` takes every drag that starts on it. `ClaimsDrag` must not
  change any state, because it is asked while the target is still being
  chosen. A drag that follows a long press is not offered: it goes to the
  innermost `Draggable`.
- [`event.MomentumDraggable`](https://pkg.go.dev/github.com/belluikit/bellui/event#MomentumDraggable)
  lets a drag keep moving after a fast release, on Android, iOS and embedded
  devices. If `DragMomentum()` returns true, the widget gets more `Dragged`
  events that slow down, then `DragEnd`. Every other `Draggable` gets
  `DragEnd` where the finger lifted. Use it for widgets that pan content,
  such as a map or a canvas. Do not use it for something the user moves
  directly, such as a slider thumb or a divider: what the user let go of has
  to stay there. Scroll containers already use it.

In the example, an upward swipe that starts on the strip scrolls the page,
and a sideways one pans the strip:

{{< example "concepts/drag/main.go" >}}

## Platform-specific interfaces

For desktop-only details there are extra interfaces in
[`driver/desktop`](https://pkg.go.dev/github.com/belluikit/bellui/driver/desktop):
`Mouseable` (`MouseDown` and `MouseUp`, with the button and modifier keys),
`Keyable` (`KeyDown` and `KeyUp`) and `Cursorable` (the mouse cursor to show
over the object). The [`driver/mobile`](https://pkg.go.dev/github.com/belluikit/bellui/driver/mobile)
package has `Touchable` (`TouchDown`, `TouchUp`, `TouchCancel`) for raw touch
events.

## Keyboard input

Key events go to the **focused** object. An object can take focus if it
implements [`Focusable`](https://pkg.go.dev/github.com/belluikit/bellui/event#Focusable):

| Method | Called when |
| --- | --- |
| `FocusGained()`, `FocusLost()` | The object gains or loses focus |
| `TypedRune(rune)` | A character is typed |
| `TypedKey(*KeyEvent)` | A key without a character is pressed, such as an arrow key or Return |

`KeyEvent.Name` is a [`KeyName`](https://pkg.go.dev/github.com/belluikit/bellui/event#KeyName)
such as `event.KeyReturn` or `event.KeyEscape`. `KeyEvent.Modifier` holds the
modifier keys that were down. On desktop, Tab moves focus to the next
focusable object. Your code can move focus with `w.Canvas().Focus(obj)`.

Some widgets keep Tab for themselves: a multi-line `Entry` types a tab
character, and the DAG canvas moves between nodes. A custom widget does the
same by implementing
[`event.Tabbable`](https://pkg.go.dev/github.com/belluikit/bellui/event#Tabbable)
with `AcceptsTab` returning true. So that keyboard users are never stuck in
such a widget, two key sequences still move focus out of it:

- **Escape, then Tab** (or Shift+Tab to go back). This works everywhere,
  including the web. An Escape that closes a pop-up or dialog does not count;
  press Escape once more.
- **Ctrl+Tab** (or Ctrl+Shift+Tab), on desktop and mobile. Browsers keep
  Ctrl+Tab for switching their own tabs, so it does not work on the web.

On desktop, a key pressed with Control, Alt or Super (Command on macOS) goes
to the shortcut system, not to `TypedKey`. Only unmodified and
Shift-modified keys reach `TypedKey`.

When no object has focus, key events go to the window's canvas instead.
Handle them there with `SetOnTypedKey` and `SetOnTypedRune`, as the example
below does.

## Shortcuts

A shortcut is a key combination with a name. The
[`bellui`](https://pkg.go.dev/github.com/belluikit/bellui) package defines
the standard ones (`ShortcutCopy`, `ShortcutPaste`, `ShortcutCut`,
`ShortcutSelectAll`, `ShortcutUndo`, `ShortcutRedo` and more). For any other
combination, use
[`desktop.CustomShortcut`](https://pkg.go.dev/github.com/belluikit/bellui/driver/desktop#CustomShortcut).
Add a window-wide shortcut on the window's canvas:

{{< example file="concepts/events/main.go" region="keys" >}}

`event.KeyModifierShortcutDefault` is Command in an app built for macOS or
iOS, and Control in every other build, so one shortcut follows each
platform's convention. It is fixed when you compile: a web build uses
Control, even in a browser on a Mac.

When a shortcut is pressed, the toolkit tries, in order:

1. a matching shortcut in the window's main menu;
2. the focused object, if it implements
   [`bellui.Shortcutable`](https://pkg.go.dev/github.com/belluikit/bellui#Shortcutable)
   (`TypedShortcut`);
3. the shortcuts added to the canvas with `AddShortcut`.

So a focused text entry receives Ctrl+C before a canvas shortcut for the same
keys would. To give your own widget shortcuts, embed a
[`bellui.ShortcutHandler`](https://pkg.go.dev/github.com/belluikit/bellui#ShortcutHandler),
which provides `TypedShortcut`, `AddShortcut` and `RemoveShortcut`.

The complete example:

{{< example "concepts/events/main.go" >}}

## Detecting an idle user

[`bellui.LastUserActivity()`](https://pkg.go.dev/github.com/belluikit/bellui#LastUserActivity)
returns when the user last gave the app input, in any of its windows. Use it
for an idle timer: go back to a home screen, dim, or lock after a while
without input. A timer reads it and compares it with `time.Now`:

{{< example file="concepts/idle/main.go" region="idle" >}}

It counts input as it arrives, before any widget sees it, so a scroll inside
a list counts even though the list reports nothing. What counts:

- a mouse button pressed or released, and the mouse moved while a button is
  held;
- a touch going down, moving or lifting;
- a scroll from a wheel or a trackpad;
- a key pressed or released, typed text, and input from an on-screen
  keyboard;
- files dropped into a window.

Moving the mouse with no button held does not count. Neither does anything
the app does itself: an animation, a timer, content that keeps scrolling after
the finger lifts, or a value set from code. Actions a screen reader performs
through the accessibility API, such as activating a button, are not counted
yet. Before the first input it returns the zero time. You can call it from
any goroutine.

The whole program:

{{< example "concepts/idle/main.go" >}}
