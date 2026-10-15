---
title: Accessibility
description: What Bellui exposes to screen readers on each platform, how to turn it on, and how to make your own widgets accessible.
weight: 8
---

Bellui draws its widgets itself, so the operating system cannot see them
the way it sees native controls. To make an app usable with a screen reader,
Bellui describes each widget to the platform's accessibility API: what it
is (a button, a checkbox, a slider), its name, its value and state, and what
can be done with it.

{{< cards >}}
  {{< card link="labels/" title="Labels and announcements" subtitle="Name icon-only controls and announce status changes." >}}
  {{< card link="custom-widgets/" title="Accessible custom widgets" subtitle="Give your own widgets a role, a name, a value and actions." >}}
  {{< card link="keyboard/" title="Keyboard and focus" subtitle="Tab order, initial focus and keyboard-operable widgets." >}}
{{< /cards >}}

## Turn it on

Screen-reader support is compiled in only when you build with the
`accessibility` tag:

```sh
go build -tags accessibility .
```

Without the tag, the description of each widget still exists in your code,
but nothing passes it to the operating system, and the app is invisible to
screen readers. Make sure the tag reaches every build you ship.

## What each platform gets

| Platform | Accessibility API | Notes |
| --- | --- | --- |
| macOS | NSAccessibility, used by VoiceOver | Text fields support caret and selection. Headings feed VoiceOver's heading navigation. Extra actions appear in VoiceOver's actions menu. |
| Windows | UI Automation, used by Narrator and NVDA | Buttons, toggles, expand and collapse, selection, values and ranges, text, and grids. Announcements need Windows 10 version 1709 or later. Rename and delete actions are not exposed. |
| Linux | AT-SPI2 over D-Bus, used by Orca | Written in Go, with no C library to install. When the desktop has no accessibility bus, it does nothing. |
| Android | Android accessibility, used by TalkBack | Extra actions appear in TalkBack's actions menu. The tree is only built while an accessibility service is running. |
| iOS | UIKit accessibility, used by VoiceOver | Supports activating, adjusting values, dismissing with the escape gesture, and editing text. |
| Web | none | There is no screen-reader support in the browser build. |

On every platform the screen reader gets the same information from the same
widget code:

- **Role**: one of about 45 kinds, from button and checkbox to table cell
  and heading.
- **Name**: the button text, the label text, the entry's placeholder, or a
  name you set.
- **Value**: the text in an entry, a slider's position, a date picker's date.
- **State**: checked, selected, expanded, disabled, focused, read-only,
  required, modal.
- **Actions**: press, toggle, increment, decrement, expand, collapse,
  select, dismiss, and for some widgets rename, delete, undo and redo.

Most standard widgets describe themselves: buttons, checks, radio groups,
switches, entries, selects, sliders, progress bars, lists, tables, trees,
tabs, menus, navigation bars and rails, date and time pickers, toasts,
dialogs and more. A widget that can be disabled is reported as disabled
automatically.

Elements are listed in the order they appear in the widget tree, which is
the order a screen reader reads them. Under a right-to-left layout the order
stays the same and only the positions are mirrored.

Things that are not described: decorative images and icons (`canvas.Image`,
`widget.Icon`), and anything you draw yourself with canvas objects. Wrap them
in a widget of your own if they carry meaning. See
[Accessible custom widgets](custom-widgets/).
