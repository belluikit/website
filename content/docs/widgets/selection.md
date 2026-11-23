---
title: Selection controls
description: Check boxes, radio groups, chip groups, drop-downs, switches, sliders and spin boxes for choosing among options or values.
weight: 3
---

Selection controls let the user choose instead of type. Each one reports the
choice through an `OnChanged` function and keeps it in an exported field you
can read at any time.

## Which one to use

| You want the user to | Use |
| --- | --- |
| Agree to one thing, or turn one option on | `Check` |
| Turn a setting on or off that applies at once | `Switch` |
| Pick any number from a short list | `CheckGroup` |
| Pick exactly one from a short list, all options visible | `RadioGroup` |
| Pick one or several options shown as chips that wrap onto new rows | `ChipGroup` |
| Pick one from a longer list, shown on demand | `Select` (or `SearchSelect` for long lists, see [Text input](../text-input/#choices-and-search)) |
| Pick one of 2 to 5 options shown as buttons | `SegmentedButton` (see [Buttons](../buttons/#button-groups-segmented-and-split-buttons)) |
| Pick a value along a range | `Slider` |
| Pick a low and high bound | `RangeSlider` |
| Set an exact number by stepping up and down | `SpinBox` |

## Check and CheckGroup

{{< example file="widgets/selection/main.go" region="check" >}}

[`Check`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Check) is a
check box with a label. `Checked` holds its state and `SetChecked` changes it
from code (safe from any goroutine). Setting `Partial` shows a dash, the
"some but not all" state of a parent check box. The user cannot set it; a tap
on a partial check box checks it. `NewCheckWithData` connects the check box
to a `binding.Bool`.

[`CheckGroup`](https://pkg.go.dev/github.com/belluikit/bellui/widget#CheckGroup)
shows one check box per option and passes the list of checked options to
`OnChanged`. `Horizontal` lays them out in a row. With `Required` set, the
user cannot uncheck the last checked option. `Append` and `Remove` change the
options later.

## RadioGroup

{{< example file="widgets/selection/main.go" region="radio" >}}

[`RadioGroup`](https://pkg.go.dev/github.com/belluikit/bellui/widget#RadioGroup)
keeps at most one option selected and passes its text to `OnChanged`.
Tapping the selected option again clears the selection, unless `Required` is
set. `SetSelected` picks an option from code and is safe to call from any
goroutine.

## ChipGroup

{{< example file="widgets/chips/main.go" region="group" >}}

[`ChipGroup`](https://pkg.go.dev/github.com/belluikit/bellui/widget#ChipGroup)
shows one [filter chip](../buttons/#chips) per option. Each chip is as wide
as its label, and the chips wrap onto the next row when a row is full.

- Without `Multiple`, one chip can be selected, and tapping another moves the
  selection to it. With `Multiple`, any number can be selected.
- Tapping a selected chip deselects it, unless `Required` is set and it is
  the last one selected.
- `OnChanged` receives the selected options, in the order they were
  selected. `Selected` holds them.
- `SetSelected` selects options from code, fires `OnChanged` and refreshes.
  `SetSelectedSilently` does the same without firing anything.
- `Alignment` places each row of chips. `bellui.TextAlignLeading` (the
  default) puts it at the start, `TextAlignCenter` in the middle and
  `TextAlignTrailing` at the end. A shorter last row is placed the same way.
  Start and end follow the layout direction, so under right-to-left the chips
  run from the right edge.

### Excluding options

{{< example file="widgets/chips/main.go" region="exclude" >}}

A filter often needs "not this one" as well as "only these". Set `Excludable`
on a group with `Multiple`, and each tap moves a chip from off to selected,
to excluded, and back to off. `Excludable` does nothing without `Multiple`.

- An excluded chip shows a cross where the check mark goes, an
  error-coloured outline, and its label struck through, so the state does not
  depend on colour alone. Material 3 has no excluded filter chip, so this
  look is BellUI's own. A screen reader reads it as "Rock, excluded".
- `Excluded` holds the excluded options, in the order they were excluded. An
  option is never in both `Selected` and `Excluded`.
- `OnStateChanged` receives each option whose state changed, with its new
  `ChipState`: `ChipStateOff`, `ChipStateSelected` or `ChipStateExcluded`.
  A tap changes one option. Both lists are already updated when it runs.
  The tap from selected to excluded also takes the option out of
  `Selected`, so `OnChanged` fires first with the smaller selection.
- `State` returns one option's state.
- `SetExcluded` excludes options from code and fires `OnStateChanged`. If it
  takes an option out of `Selected`, it fires `OnChanged` too.
  `SetExcludedSilently` fires nothing.
- To refuse a change, put the lists back from `OnStateChanged` with
  `SetSelectedSilently` and `SetExcludedSilently`.

`OnChanged` keeps its `func([]string)` signature and still reports only the
selection. Code written for an earlier version works unchanged.
`OnStateChanged` runs in every group, not only excludable ones, after
`OnChanged`. In a single-choice group, a tap that moves the selection reports
two options: the one that lost the selection and the one that gained it.

## Select

{{< example file="widgets/selection/main.go" region="select" >}}

[`Select`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Select)
shows the current option and opens the full list when tapped. Until the user
picks something it shows `PlaceHolder`.

- `Selected` holds the chosen text. `SelectedIndex()` returns its position,
  or -1.
- `SetSelected` and `SetSelectedIndex` choose from code. `ClearSelected`
  goes back to the placeholder.
- `SetOptions` replaces the list.
- `NewSelectWithData` connects the selection to a `binding.String`.

## Switch

{{< example file="widgets/selection/main.go" region="switch" >}}

[`Switch`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Switch) is
the MD3 toggle with a sliding thumb. `On` holds the state and `OnChanged`
receives each change. It has no label of its own, so put a `Label` next to
it. Space toggles a focused switch.

Use a `Switch` for a setting that takes effect immediately. Use a `Check` when
the choice is applied later, for example when a form is submitted.

## Slider

{{< example file="widgets/sliders/main.go" region="slider" >}}

[`Slider`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Slider)
moves between `Min` and `Max`. `Step` snaps the value; with `Step` above zero,
`ShowTickMarks` draws a mark at each step. `ShowValueIndicator` shows the
value above the thumb while dragging. Set `Orientation` to `widget.Vertical`
for an upright slider.

- `OnChanged` runs continuously while the user drags. `OnChangeEnded` runs
  when the change is finished, for example when the drag ends. Save or send
  the value in `OnChangeEnded`.
- `SetValue` clamps the value into range and fires both callbacks, as if the
  user had moved it. `SetValueSilently` moves the thumb without firing them,
  for example to show playback position. Both are safe from any goroutine.
- `Dragging()` reports whether the user is dragging right now. A media player
  can use it to stop updating the position while the user scrubs.
- `OnHover` and `OnHoverEnd` report the value under the pointer without
  moving the thumb, which suits a preview on a seek bar.

`NewSliderWithData` connects the value to a `binding.Float`.

## RangeSlider

{{< example file="widgets/sliders/main.go" region="range" >}}

[`RangeSlider`](https://pkg.go.dev/github.com/belluikit/bellui/widget#RangeSlider)
has two thumbs. `LowValue` and `HighValue` hold the range, and both callbacks
receive both values. `SetValues` clamps, snaps to `Step` and swaps the values
if they arrive the wrong way round; `SetValuesSilently` does the same without
firing the callbacks. From the keyboard, the arrows along the slider move the
focused thumb, and the arrows across it switch between thumbs.

## SpinBox

{{< example file="widgets/sliders/main.go" region="spin" >}}

[`SpinBox`](https://pkg.go.dev/github.com/belluikit/bellui/widget#SpinBox)
is a number field with increment and decrement buttons. Each press adds or
subtracts `Step`, and `SetValue` clamps to `Min` and `Max`. The up and down
arrow keys step the value when it has focus.
