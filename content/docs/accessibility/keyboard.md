---
title: Keyboard and focus
description: How Tab moves focus through a window, how to set the first focused widget, and how to make your own widgets work from the keyboard.
weight: 3
---

Many people use an app from the keyboard alone, and screen-reader users
mostly do. Every control they need should be reachable with Tab and operable
with keys.

## The focus order

On desktop, Tab moves focus to the next widget that accepts it, and
Shift+Tab to the previous one. At either end, focus wraps around.

{{< example file="accessibility/keyboard/main.go" region="order" >}}

- The order is the order of the widget tree: the order of objects in their
  containers, depth first. To change the order, change the layout.
- Disabled widgets are skipped.
- Under a [right-to-left layout](/docs/theming/right-to-left/), Bellui reorders
  focus to follow the screen: rows from top to bottom, and each row from
  right to left. A single column keeps its order.
- Escape closes the topmost pop-up, menu or dialog, unless the focused
  widget uses Escape itself.

Buttons are activated with Space when they have focus.

## Set the first focused widget

For a window whose purpose is to type something, focus the first field
when you build it, so the user can start typing straight away:

{{< example file="accessibility/keyboard/main.go" region="initial" >}}

The canvas also has `FocusNext`, `FocusPrevious`, `Unfocus` and `Focused`.

## Widgets that keep Tab

A widget that implements
[`event.Tabbable`](https://pkg.go.dev/github.com/belluikit/bellui/event#Tabbable)
and returns `true` from `AcceptsTab` receives Tab itself, instead of
moving focus. A multi-line entry does this so that Tab inserts a tab
character. It also keeps Shift+Tab, and Bellui gives it no other key that
moves focus out, so a keyboard user who reaches one cannot Tab past it.

If you write such a widget, give users another key that leaves it, and move
focus on yourself with the canvas's `FocusNext`. A widget that keeps Tab
with no way out traps keyboard users.

## Make a custom widget focusable

A widget receives focus and keys by implementing
[`event.Focusable`](https://pkg.go.dev/github.com/belluikit/bellui/event#Focusable):

{{< example file="accessibility/customwidget/main.go" region="keyboard" >}}

- `FocusGained` and `FocusLost` tell the widget when it gains and loses
  focus. Show it: users must be able to see where focus is. The rating
  widget puts brackets around its stars.
- `TypedKey` receives keys such as the arrows, Home and Space.
- `TypedRune` receives typed characters.

Give keyboard users the same control a pointer user has. The rating widget
uses the arrow keys for what a tap does.

A focusable widget that a key activates, as Space activates a button, can
also implement
[`event.Activatable`](https://pkg.go.dev/github.com/belluikit/bellui/event#Activatable).
`ActivatedByKey` reports which keys activate it. Bellui then reports that
key press as an activation, the same signal a tap sends. The widget still
handles the key itself in `TypedKey`.

The full program:

{{< example "accessibility/keyboard/main.go" >}}
