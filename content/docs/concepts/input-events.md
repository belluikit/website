---
title: Input events
description: How taps, hover, drags, scrolling, keys and shortcuts reach your objects, and what differs between desktop and touch.
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

A [`PointEvent`](https://pkg.go.dev/github.com/belluikit/bellui/event#PointEvent)
has `Position`, relative to the object's top-left corner, and
`AbsolutePosition`, relative to the window. `DragEvent.Dragged` and
`ScrollEvent.Scrolled` are a [`geo.Delta`](https://pkg.go.dev/github.com/belluikit/bellui/geo#Delta):
how far the pointer moved since the last event. `ScrollEvent.Source` says
whether a scroll came from a mouse wheel, a touchpad or a touch gesture.

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
| Scroll | Mouse wheel or touchpad | Dragging a finger over scrollable content, such as a scroll container |

On touch, a widget that implements both `LongPressable` and
`SecondaryTappable` gets `LongPressed` while the finger is held, and then
`TappedSecondary` when it lifts.

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
