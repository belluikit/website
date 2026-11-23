---
title: Containment and overlays
description: Cards, accordions, carousels and swipe areas that hold content, sheets, pop-ups, menus and tooltips that float above it, and inner windows the user moves around.
weight: 6
---

Some widgets exist to hold other content. Cards, accordions and carousels
group it in place. Sheets, pop-ups, menus and tooltips float above the rest
of the window for a moment. A swipe area reports swipes across what it holds.
Inner windows are small windows inside your app's window, which the user
drags, resizes and closes.

{{< screenshot name="containment" alt="The gallery's Containment page: a tree view on the left, elevated, filled and outlined cards, and a carousel showing its first slide" >}}

## Which one to use

| You want | Use |
| --- | --- |
| A titled box around related content | `Card` |
| Sections the user expands one at a time | `Accordion` |
| A row of items the user steps through one at a time | `Carousel` |
| Content the user swipes left, right, up or down, such as a card to keep or skip | `SwipeArea` |
| Extra actions or content sliding up from the bottom | `BottomSheet` |
| Supplementary details beside the main content | `SideSheet` |
| Any content floating above the window | `PopUp` |
| Several movable, resizable windows inside the app's window | `container.InnerWindow` in a `container.MultipleWindows` |
| A right-click menu on a widget | `ContextMenu` |
| A menu opened from a button | `ShowPopUpMenuAtRelativePosition` |
| A short hint on hover | `SetToolTipText`, or `Tooltip` for non-widget content |

For a dialog with a title and buttons, such as a confirmation, use the
[`dialog`](https://pkg.go.dev/github.com/belluikit/bellui/dialog) package.

## Card

[`Card`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Card) shows a
`Title`, a `Subtitle`, optional `Content` below them and an optional `Image`
above them. Any of these can be left empty.

{{< example file="widgets/containment/main.go" region="card" >}}

`Style` is one of the three MD3 card variants: `CardStyleElevated` (the
default, with a shadow), `CardStyleFilled` (a filled surface, no shadow) and
`CardStyleOutlined` (a border, no shadow). `SetTitle`, `SetSubTitle`,
`SetContent` and `SetImage` change a card that is already shown.

## Accordion

[`Accordion`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Accordion)
stacks items, each a title that expands to show its `Detail` when tapped.

{{< example file="widgets/containment/main.go" region="accordion" >}}

By default opening one item closes the others. Set `MultiOpen` to let several
stay open. `Open(i)`, `Close(i)`, `OpenAll()` and `CloseAll()` work from
code; without `MultiOpen`, `OpenAll` opens only the first item. `Append`,
`Prepend`, `Remove` and `RemoveIndex` change the items.

## Carousel

[`Carousel`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Carousel)
shows one item at a time, with previous and next buttons beside it and a row
of dots below that marks the current item. The user can also swipe
sideways: the current item follows the finger and the next one slides in
beside it.

{{< example file="widgets/containment/main.go" region="carousel" >}}

When the user lets go, the carousel moves to the next item if the swipe went
more than a fifth of its width or was a quick flick. Otherwise it goes back to
the current item. Past the first or last item the drag meets resistance and
always goes back. In a right-to-left layout the order, buttons and swipe
direction are mirrored.

`CurrentItem` is the index shown and `OnChanged` receives it after each move.
`Next()` and `Previous()` move from code.

## SwipeArea

[`SwipeArea`](https://pkg.go.dev/github.com/belluikit/bellui/widget#SwipeArea)
holds one object and tells you when the user swipes across it. A swipe is a
drag that goes far enough and mostly along one axis. `OnSwiped` receives the
direction the finger moved: `SwipeLeft`, `SwipeRight`, `SwipeUp` or
`SwipeDown`. Taps still reach the content.

{{< example file="widgets/swipe/main.go" region="swipe" >}}

Set `Follow` to make the content a card that moves with the finger and tilts
a little as it goes:

{{< example file="widgets/swipe/main.go" region="card" >}}

- A release a fifth of the way across, or a flick, throws the card off that
  side. `OnSwiped` runs and the content is back in place, ready for the next
  item. A shorter drag springs it back.
- `OnSwipeProgress` reports how close the card is to a swipe, from 0 at rest
  to 1 where a release would swipe. Use it to show a hint as the card leaves.
- `Directions` sets what a swipe does in each direction. `SwipeThrow` (the
  default for any direction not in the map) throws the card. `SwipeStay`
  lets the card follow the finger but springs it back, and runs `OnSwiped`
  as it starts back. `SwipeIgnore` keeps the card still in that direction and
  never runs `OnSwiped` for it.

With animations off in the app settings, the card moves at once.

## Bottom and side sheets

{{< example file="widgets/sheets/main.go" region="bottom" >}}

[`BottomSheet`](https://pkg.go.dev/github.com/belluikit/bellui/widget#BottomSheet)
slides up from the bottom edge of the canvas you pass to `Show`. It is at
most 640 units wide and centred. `Hide()` closes it.

{{< example file="widgets/sheets/main.go" region="side" >}}

[`SideSheet`](https://pkg.go.dev/github.com/belluikit/bellui/widget#SideSheet)
opens at the right side of the canvas. `ShowBottomSheet` and `ShowSideSheet`
create and show a sheet in one call.

Sheets are not modal by default, so a tap outside closes them. Set `Modal`
before `Show` to dim the window and block it until your code calls `Hide`.
Escape closes either kind.

## PopUp

[`PopUp`](https://pkg.go.dev/github.com/belluikit/bellui/widget#PopUp)
floats any content above a canvas, with padding and a shadow. Sheets and
dialogs are built on it.

{{< example file="widgets/sheets/main.go" region="popup" >}}

- `NewPopUp` and `NewModalPopUp` create a pop-up without showing it. Show it
  with `Show()`, `ShowAtPosition(pos)` or `ShowAtRelativePosition(pos, obj)`,
  and close it with `Hide()`.
- `ShowPopUp`, `ShowModalPopUp`, `ShowPopUpAtPosition` and
  `ShowPopUpAtRelativePosition` create and show one in a single call.
- A normal pop-up closes when the user taps outside it. A modal one covers the
  window with a dimmed overlay, ignores taps outside, and stays until your
  code hides it. Escape closes both.
- Set `OnDismissed` to run your own close logic whenever the user dismisses
  the pop-up: Escape, a tap or right-click outside a normal pop-up, or the
  screen reader's dismiss action. When it is set, it replaces the plain hide,
  so your function must call `Hide()` itself. Calling `Hide()` from code does
  not run it.

## Inner windows

[`container.InnerWindow`](https://pkg.go.dev/github.com/belluikit/bellui/container#InnerWindow)
draws a window inside your app's window: a title bar with close, minimise and
maximise buttons, your content below it, and a corner handle at the bottom
right for resizing. Put inner windows in a
[`container.MultipleWindows`](https://pkg.go.dev/github.com/belluikit/bellui/container#MultipleWindows),
which lets the user move them around.

{{< screenshot name="innerwindow" alt="Two inner windows on a plain background. The Notes window has a document icon at the right of its title bar and an empty text area. The About window overlaps it, with a short hint about dragging and resizing." >}}

{{< example file="widgets/innerwindow/main.go" region="windows" >}}

`MultipleWindows` gives each window its behaviour:

- dragging the title bar moves the window;
- dragging the bottom-right corner resizes it, never below its minimum size;
- tapping the title bar brings the window to the front.

It does this by setting each window's `OnDragged`, `OnResized` and
`OnTappedBar`, so any functions you set there yourself are replaced. A window
outside a `MultipleWindows` does not move by itself: set those three to move,
resize and restack it as you like.

`MultipleWindows` does not place windows for you. Each one starts at the
position you give it with `Move`, at its minimum size unless you call
`Resize`. `Add` puts another window on top of the others. You can also change
the `Windows` slice and call `Refresh`; the last window in it is drawn on top.

### Window buttons

The close button hides the window. To ask first, set `CloseIntercept`. Your
function then runs instead, and calls `Close` when the window should go:

{{< example file="widgets/innerwindow/main.go" region="close" >}}

Minimise and maximise do nothing on their own. Each button works only once
you set its function, `OnMinimized` or `OnMaximized`; until then, it is shown
disabled on desktop and hidden on phones and tablets. Your function decides
what the button does. `SetMaximized` only swaps the maximise icon for a
restore icon, so resize the window yourself:

{{< example file="widgets/innerwindow/main.go" region="maximise" >}}

`Icon` adds an icon to the title bar. It is drawn disabled unless you set
`OnTappedIcon`. `SetTitle` changes the title, and `SetPadded(false)` removes
the padding around your content.

The buttons sit at the left of the title bar on macOS, and at the right on
Windows, Linux and BSD. Set `Alignment` to `widget.ButtonAlignLeading` or
`widget.ButtonAlignTrailing` to choose for every platform.

### Disabling an inner window

`Disable()` stops the user from closing, moving or resizing the window, and
disables every widget inside it that can be disabled, as described in
[Disabled widgets](/docs/widgets/#disabled-widgets). Your code can still
call `Close`, `SetContent` and `SetMaximized`.

## Menus

{{< example file="widgets/menus/main.go" region="context" >}}

[`ContextMenu`](https://pkg.go.dev/github.com/belluikit/bellui/widget#ContextMenu)
wraps any object and shows a menu of
[`bellui.MenuItem`](https://pkg.go.dev/github.com/belluikit/bellui#MenuItem)s
on right-click (a secondary tap). Taps, hovers and drags still reach the
wrapped object. `NewContextMenuWithMenu` takes a ready-made `*bellui.Menu`
instead of items.

To open a menu from a button, pass a `*bellui.Menu` to
`ShowPopUpMenuAtRelativePosition`, with a position relative to the button:

{{< example file="widgets/menus/main.go" region="popup" >}}

`ShowPopUpMenuAtPosition` takes an absolute position instead, and
`NewPopUpMenu` creates a menu you can show again and again.

## Tooltips

Every widget has `SetToolTipText`. On desktop, the text appears in a small
bubble after the pointer rests on the widget for a moment, and disappears
when it leaves. An empty string removes the tooltip.

{{< example file="widgets/menus/main.go" region="tooltip" >}}

[`Tooltip`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Tooltip)
does the same as a wrapper, which is useful for content that is not a widget,
such as a canvas object.
