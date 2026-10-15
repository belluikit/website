---
title: Buttons and actions
description: Buttons in the five MD3 styles, floating action buttons, icon, segmented and split buttons, chips and toolbars.
weight: 1
---

Actions are what the user taps to make something happen. Bellui has one
general [`Button`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Button)
and several specialised buttons for the MD3 patterns.

## Which one to use

| You want | Use |
| --- | --- |
| A labelled action, the default choice | `Button` |
| The single most important action on a screen, floating above content | `FAB` (`NewFAB`, `NewExtendedFAB`) |
| A floating button that opens several related actions | `FABMenu` |
| A compact action shown only as an icon | `IconButton` |
| A choice between 2 to 5 mutually exclusive options | `SegmentedButton` |
| A main action with alternatives in a drop-down | `SplitButton` |
| Several related actions joined into one shape | `ButtonGroup` |
| Small actions, filters or tags among content | `Chip` |
| A row of icon actions along the top of a window | `Toolbar` |

## Button

[`widget.NewButton`](https://pkg.go.dev/github.com/belluikit/bellui/widget#NewButton)
takes a label and a function that runs when the button is tapped.
`NewButtonWithIcon` adds a leading icon.

{{< example file="widgets/buttons/main.go" region="basic" >}}

Space activates a focused button, as a tap does. `SetText` and `SetIcon`
change the button later and are safe to call from any goroutine. Pass `nil`
to `SetIcon` to remove the icon.

### Styles

The `Style` field picks one of the MD3 button variants:

{{< example file="widgets/buttons/main.go" region="styles" >}}

| Style | Looks like |
| --- | --- |
| `ButtonStyleFilled` | Primary colour fill, the strongest emphasis |
| `ButtonStyleFilledTonal` | Secondary container fill |
| `ButtonStyleElevated` | Low surface fill with primary-coloured text |
| `ButtonStyleOutlined` | A border and primary-coloured text |
| `ButtonStyleText` | Primary-coloured text only, the lightest emphasis |

### Importance

When `Style` is left at `ButtonStyleDefault`, the `Importance` field chooses
the colours instead:

- `HighImportance` draws the button as `ButtonStyleFilled`.
- `MediumImportance` (the default) and `LowImportance` draw it as
  `ButtonStyleFilledTonal`.
- `DangerImportance`, `WarningImportance` and `SuccessImportance` fill it with
  the theme's error, warning and success colours.

A `Style` other than the default wins over `Importance`.

{{< example file="widgets/buttons/main.go" region="importance" >}}

`Alignment` moves the icon and text to the leading or trailing edge
(`ButtonAlignLeading`, `ButtonAlignTrailing`), and `IconPlacement` puts the
icon after the text with `ButtonIconTrailingText`.

## Icon buttons

[`IconButton`](https://pkg.go.dev/github.com/belluikit/bellui/widget#IconButton)
shows only an icon. Its `Style` is `IconButtonStyleStandard` (no background),
`IconButtonStyleFilled`, `IconButtonStyleFilledTonal` or
`IconButtonStyleOutlined`.

{{< example file="widgets/buttons/main.go" region="icon" >}}

An icon button has no text, so screen readers have nothing to announce. Give
it a name with `SetAccessibilityLabel`, and a hover hint with
`SetToolTipText`.

## Button groups, segmented and split buttons

{{< example file="widgets/buttons/main.go" region="group" >}}

- [`ButtonGroup`](https://pkg.go.dev/github.com/belluikit/bellui/widget#ButtonGroup)
  joins ordinary buttons into one connected shape. Each button keeps its own
  `OnTapped`.
- [`SegmentedButton`](https://pkg.go.dev/github.com/belluikit/bellui/widget#SegmentedButton)
  shows a row of labels and keeps exactly one selected. `OnChanged` receives
  the index of the new selection, and `Selected` holds it (index 0 after
  `NewSegmentedButton`; -1 selects nothing). When it has focus, the left and
  right arrow keys move the selection.
- [`SplitButton`](https://pkg.go.dev/github.com/belluikit/bellui/widget#SplitButton)
  has a main action and an arrow that opens a menu of
  [`bellui.MenuItem`](https://pkg.go.dev/github.com/belluikit/bellui#MenuItem)
  alternatives.

## Floating action buttons

A [`FAB`](https://pkg.go.dev/github.com/belluikit/bellui/widget#FAB) is the
floating button for a screen's main action. `NewFAB` makes an icon-only
FAB. `NewExtendedFAB` adds a label. `Style` picks the colour:
`FABStylePrimary` (the default), `FABStyleSurface`, `FABStyleSecondary` or
`FABStyleTertiary`.

{{< example file="widgets/fab/main.go" region="fab" >}}

A FAB is an ordinary widget, so you place it with a layout like any other.
To keep it in a corner, put it at the end of a box with a spacer before it.

### FAB menu

[`FABMenu`](https://pkg.go.dev/github.com/belluikit/bellui/widget#FABMenu)
is the MD3 speed dial: a list of labelled mini FABs that opens over the
content. It is not a widget you place in a layout. You open it with
`Show`, `Toggle` or close it with `Hide`, giving the canvas and a position.
The menu opens above and to the left of that position. Tapping an item
closes the menu and calls its `OnTapped`.

{{< example file="widgets/fab/main.go" region="menu" >}}

## Chips

A [`Chip`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Chip) is a
small rounded label. Its `Style` decides what it is for:

- `ChipStyleAssist` (from `NewChip`): an action, with an optional leading
  `Icon`.
- `ChipStyleFilter` (from `NewFilterChip`): a toggle. Each tap flips
  `Selected`, fills the chip with the secondary container colour while it is
  selected, and calls your function with the new state.
- `ChipStyleInput`: a tag the user can remove. Set `OnDeleted`; a tap on the
  trailing close mark calls it, and a tap anywhere else calls `OnTapped`.
- `ChipStyleSuggestion`: an outlined label, drawn the same way as an assist
  chip.

There is no constructor for input and suggestion chips, so build them as a
struct literal.

{{< example file="widgets/chips/main.go" region="chips" >}}

## Toolbar

[`Toolbar`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Toolbar)
lays out a row of toolbar items:

- `NewToolbarAction(icon, func)`: an icon button. It has `Disable`,
  `Enable` and `SetIcon`.
- `NewTwoStateToolbarAction(offIcon, onIcon, func(on bool))`: a toggle that
  swaps icons, such as play and pause. `SetOn` changes the state from code.
- `NewToolbarSeparator()`: a thin divider between groups.
- `NewToolbarSpacer()`: empty space that stretches, which pushes the items
  after it to the far end. Several spacers share the space equally.

{{< example file="widgets/toolbar/main.go" region="toolbar" >}}

`Append` and `Prepend` add items after the toolbar is created.
