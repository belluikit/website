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
  {{< card link="selection/" title="Selection controls" subtitle="Check, CheckGroup, RadioGroup, ChipGroup, Select, Switch, Slider, RangeSlider, SpinBox." >}}
  {{< card link="pickers/" title="Pickers" subtitle="Date, time and date-time pickers and entries, wheel picker, font picker, theme selector." >}}
  {{< card link="navigation/" title="Navigation" subtitle="Top app bar, navigation bar, rail and drawer, bottom app bar." >}}
  {{< card link="containment/" title="Containment and overlays" subtitle="Cards, accordions, carousels, swipe areas, sheets, pop-ups, menus and tooltips." >}}
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
`Visible()`. Most interactive widgets also have `Disable()`, `Enable()` and
`Disabled()`; see [Disabled widgets](#disabled-widgets) below.
`SetToolTipText` adds a hover tooltip to any widget, and
`SetAccessibilityLabel` names a widget that has no visible text, such as an
icon button, for screen readers.

**Pressed state.** Buttons, chips, FABs, links, selects, check boxes, radio
buttons, switches, list, grid and tree rows, menu items, sortable table
headers, navigation items and tabs show the MD3 pressed state from the press
to the release. A touch that turns into a scroll ends it, so a fling over a
list leaves no row pressed.

**Data binding.** A few widgets have a `NewXWithData` constructor that
connects them to a value from the
[`data/binding`](https://pkg.go.dev/github.com/belluikit/bellui/data/binding)
package: Check, Entry, Label, ProgressBar, Select, Slider, List, GridWrap and
Tree. All of these except List, GridWrap and Tree also have a `Bind` method
for a widget you already created. Binding is covered in [Data binding](/docs/data/).

## Disabled widgets

Call `Disable()` to keep a widget on screen while the user cannot use it, and
`Enable()` to give it back. These can be disabled:

| Kind | Widgets |
| --- | --- |
| Buttons | `Button`, `IconButton`, `FAB`, `SplitButton`, `ButtonGroup`, `SegmentedButton`, `Chip`, `ChipGroup`, `AuthButton`, toolbar actions |
| Text input | `Entry` and every entry variant, `SelectEntry`, `SearchSelect`, `DateEntry`, `DateTimeEntry`, `RichEntry` |
| Selection | `Check`, `CheckGroup`, `RadioGroup`, `Select`, `Switch`, `Slider`, `RangeSlider`, `SpinBox` |
| Pickers | `DatePicker`, `TimePicker`, `DateTimePicker`, `WheelPicker`, `FontPicker`, `ThemeSelector` |
| Navigation | `TopAppBar`, `NavigationBar`, `NavigationRail`, `BottomAppBar` |
| Collections and containment | `List`, `Tree`, `TreeView`, `FileTree`, `Carousel`, `SwipeArea` |
| Containers and forms | `Form`, `container.AppTabs`, `container.DocTabs`, [`container.InnerWindow`](/docs/widgets/containment/#inner-windows) |

Some cannot, for example `Hyperlink`, `Table`, `GridWrap` and
`NavigationDrawer`. Put them inside one that can, such as a tab or an inner
window, and they are blocked with it.

A disabled widget:

- keeps its place and draws in the MD3 disabled colours;
- takes no tap, key, hover or pressed state, and Tab skips it;
- gives up the keyboard focus if it had it, and cannot take it again. The
  exceptions are an `Entry` and a selectable `Label` or `RichText`: a click
  can still focus them, so the user can select and copy their text;
- still does what your code asks: `Select`, `SetText`, `Next` and similar
  calls work as before.

Everything inside a disabled widget is blocked too, including widgets that
cannot be disabled themselves:

- A tap, double tap, long press, mouse press, hover or drag inside it does
  nothing. A link in a disabled tab does not open.
- Tab skips everything inside it.
- What has no disabled look of its own is drawn at 38% opacity. That covers
  a label, a link, a card, the pages of a `Carousel` and your own widgets.
  Widgets that have their own disabled look, such as a `Button` or an
  `Entry`, draw it once and are not dimmed a second time.
- The mouse wheel and a scrolling drag still reach the widget under them.
  That widget decides whether to scroll.

`Disable()` is safe to call before the widget is first shown. Some widgets
add their own rules:

- A disabled `List`, `Tree`, `TreeView` or `FileTree` does not scroll, by
  wheel, touch or scroll bar, and a wheel over it does not scroll the page
  around it either. Its selected row shows in the disabled colours.
- In a collection, a `Carousel` or a `SwipeArea`, the widgets you supply,
  such as row templates and pages, are not disabled. They take no input and
  are drawn dimmed, as above.
- A disabled `TopAppBar` keeps its title at full colour. Only its buttons
  look disabled.
- A `RichEntry` cannot be focused or selected while it is disabled.
- `ButtonGroup`, `BottomAppBar`, `AppTabs`, `DocTabs`, `InnerWindow` and
  `Form` disable what they hold. `Enable()` turns back on only what `Disable`
  turned off, so a widget you disabled yourself stays disabled. The tab
  containers and `InnerWindow` also disable content you add while they are
  disabled. A disabled `Form` also draws its labels, hints and validation
  messages in the disabled colour.
- A disabled `Entry` disables its `ActionItem`, such as a search field's
  clear button or a date entry's calendar button.

{{< example file="widgets/disable/main.go" region="disable" >}}

Tabs are containers, not widgets: see `container.NewAppTabs` and
`container.NewDocTabs` in [Layout](/docs/layout/). Dialogs are in the
[`dialog`](https://pkg.go.dev/github.com/belluikit/bellui/dialog) package.
