---
title: Labels and announcements
description: Give every control a name a screen reader can speak, and tell screen-reader users about changes they cannot see.
weight: 1
---

A screen reader speaks a control's **name**, then its role and value: "Save,
button" or "Volume, slider, 30". Most widgets take their name from their
text, so a button labelled "Save" is named "Save" with no extra work:

{{< example file="accessibility/labels/main.go" region="text" >}}

## Name controls that have no text

An icon-only button has nothing to speak, so a screen reader can only say
"button". Give it a name with
[`SetAccessibilityLabel`](https://pkg.go.dev/github.com/belluikit/bellui/widget#BaseWidget.SetAccessibilityLabel):

{{< example file="accessibility/labels/main.go" region="label" >}}

`SetAccessibilityLabel` is a method of `widget.BaseWidget`, so every
standard widget has it. A non-empty label you set replaces the widget's own
name. It only has an effect on widgets that describe themselves to screen
readers, which most standard widgets do. Translate the label like any other
text the user sees or hears.

Where you need a name:

- Icon-only buttons: `widget.NewIconButton`, toolbar actions, floating
  action buttons without text.
- Entries without a placeholder. An entry is named by its placeholder text.
- Entries in a form. The form shows a label next to each field, but it does
  not link the label to the field, so name the field yourself:

{{< example file="accessibility/labels/main.go" region="form" >}}

Bellui names the icon buttons inside its own widgets, such as the previous
and next buttons of the date picker.

## Announce changes

Some changes happen away from where the user is: a message is sent, a file
finishes downloading, a search returns nothing. A sighted user sees it; a
screen-reader user hears nothing unless you announce it:

{{< example file="accessibility/labels/main.go" region="announce" >}}

[`bellui.AnnounceToScreenReader`](https://pkg.go.dev/github.com/belluikit/bellui#AnnounceToScreenReader)
asks the screen reader to speak the text. The second argument is
`assertive`:

- `false` waits until the screen reader finishes what it is saying. Use it
  for most status messages.
- `true` interrupts it. Keep it for errors and warnings that need attention
  now.

Call it on the UI goroutine. Without the `accessibility` build tag, or when
no screen reader is running, it does nothing.

Toasts and snackbars announce their own text when they are shown. Danger
and warning toasts are announced assertively, the others politely.

The full program:

{{< example "accessibility/labels/main.go" >}}
