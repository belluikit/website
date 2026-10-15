---
title: Widgets
description: The Material Design 3 widget set in the widget package, grouped by what you use each widget for.
weight: 3
---

The [`widget`](https://pkg.go.dev/github.com/belluikit/bellui/widget) package
holds every ready-made control in Bellui: buttons, text fields, pickers,
navigation bars, lists, tables and more. They follow Material Design 3 (MD3),
so most widgets have the MD3 variants as a `Style` field: filled, tonal,
outlined and so on. Colours come from the [theme](/docs/theming/), so every
widget follows light and dark mode without extra code.

{{< screenshot name="actions" alt="The gallery's Actions page: filled, tonal, elevated, outlined and text buttons, standard and extended floating action buttons, and four icon button styles" >}}

The [gallery app](https://github.com/belluikit/gallery) shows most of these
widgets in one window. Clone it and run `go run .` to try them.

{{< cards >}}
  {{< card link="buttons/" title="Buttons and actions" subtitle="Button styles, FAB and FAB menu, icon, segmented and split buttons, chips, toolbars." >}}
  {{< card link="text-input/" title="Text input" subtitle="Entry and its variants: password, multi-line, numbers, search, completion, rich text." >}}
  {{< card link="selection/" title="Selection controls" subtitle="Check, CheckGroup, RadioGroup, Select, Switch, Slider, RangeSlider, SpinBox." >}}
  {{< card link="pickers/" title="Pickers" subtitle="Date, time and date-time pickers and entries, font picker, theme selector." >}}
  {{< card link="navigation/" title="Navigation" subtitle="Top app bar, navigation bar, rail and drawer, bottom app bar." >}}
  {{< card link="containment/" title="Containment and overlays" subtitle="Cards, accordions, carousels, sheets, pop-ups, menus and tooltips." >}}
  {{< card link="feedback/" title="Feedback" subtitle="Progress bars, activity and loading indicators, badges, snackbars, toasts." >}}
  {{< card link="text/" title="Displaying text" subtitle="Label, Hyperlink, RichText, Markdown, HTML, TextGrid, icons and separators." >}}
  {{< card link="collections/" title="Lists, tables and trees" subtitle="List, Table, GridWrap, Tree, TreeView and FileTree for large data sets." >}}
  {{< card link="forms/" title="Forms" subtitle="Labelled rows, hints, validation and submit and cancel buttons." >}}
{{< /cards >}}

## Things every widget shares

**Create widgets with their constructor.** Every widget has a `NewX`
function, for example `widget.NewButton` or `widget.NewSlider`. Set extra
fields on the result before you show it. You can also build a widget as a
struct literal, such as `&widget.Chip{Text: "go", Style: widget.ChipStyleInput}`.
The constructors call `ExtendBaseWidget` for you. When you build a widget as a
struct literal, call `w.ExtendBaseWidget(w)` on it before you use its methods.

**Change fields, then refresh.** Exported fields such as `Text`, `Style` or
`Selected` are read when the widget draws. If you change a field after the
widget is on screen, call `Refresh()` so it redraws. Setter methods such as
`Label.SetText`, `Button.SetText` or `ProgressBar.SetValue` refresh for you,
and their documentation says when they are safe to call from any goroutine.
Callbacks such as `OnTapped` and `OnChanged` run on the UI thread, so they can
update other widgets directly.

**Hide, disable and describe.** Every widget has `Show()`, `Hide()` and
`Visible()`. Most input widgets, such as buttons, entries, check boxes,
switches and sliders, also have `Disable()`, `Enable()` and `Disabled()`.
Some interactive widgets do not, for example `Chip`, `SegmentedButton`,
`Hyperlink` and `RichEntry`. `SetToolTipText` adds a hover tooltip to any widget, and
`SetAccessibilityLabel` names a widget that has no visible text, such as an
icon button, for screen readers.

**Data binding.** A few widgets have a `NewXWithData` constructor that
connects them to a value from the
[`data/binding`](https://pkg.go.dev/github.com/belluikit/bellui/data/binding)
package: Check, Entry, Label, ProgressBar, Select, Slider, List, GridWrap and
Tree. All of these except List, GridWrap and Tree also have a `Bind` method
for a widget you already created. Binding is covered in [Data binding](/docs/data/).

Tabs are containers, not widgets: see `container.NewAppTabs` and
`container.NewDocTabs` in [Layout](/docs/layout/). Dialogs are in the
[`dialog`](https://pkg.go.dev/github.com/belluikit/bellui/dialog) package.
