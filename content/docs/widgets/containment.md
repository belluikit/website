---
title: Containment and overlays
description: Cards, accordions and carousels that group content, and sheets, pop-ups, menus and tooltips that float above it.
weight: 6
---

Some widgets exist to hold other content. Cards, accordions and carousels
group it in place. Sheets, pop-ups, menus and tooltips float above the rest
of the window for a moment.

{{< screenshot name="containment" alt="The gallery's Containment page: a tree view on the left, elevated, filled and outlined cards, and a carousel showing its first slide" >}}

## Which one to use

| You want | Use |
| --- | --- |
| A titled box around related content | `Card` |
| Sections the user expands one at a time | `Accordion` |
| A row of items the user steps through one at a time | `Carousel` |
| Extra actions or content sliding up from the bottom | `BottomSheet` |
| Supplementary details beside the main content | `SideSheet` |
| Any content floating above the window | `PopUp` |
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
of dots below that marks the current item. Dragging does not move it.

{{< example file="widgets/containment/main.go" region="carousel" >}}

`CurrentItem` is the index shown and `OnChanged` receives it after each move.
`Next()` and `Previous()` move from code.

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
- Set `OnDismissed` to run your own close logic instead of the plain hide
  when the user presses Escape. A tap outside a normal pop-up still hides it
  directly.

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
