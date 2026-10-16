---
title: Selection controls
description: Check boxes, radio groups, drop-downs, switches, sliders and spin boxes for choosing among options or values.
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
