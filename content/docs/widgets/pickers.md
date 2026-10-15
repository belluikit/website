---
title: Pickers
description: MD3 date, time and date-time pickers, inline or as dialogs, date entries that accept typed dates, the font picker and the theme selector.
weight: 4
---

Pickers let the user choose a structured value, such as a date, a time or a
font, without typing it in a particular format.

## Which one to use

| You want | Use |
| --- | --- |
| A date chosen on a calendar, shown in place | `NewDatePicker` |
| A date or time chosen in a dialog from a button | `ShowDatePicker`, `ShowTimePicker`, `ShowDateTimePicker` |
| A form field where the user can type a date or open a calendar | `NewDateEntry` |
| The same for a date and a time | `NewDateTimeEntry` |
| A time on a clock dial, shown in place | `NewTimePicker` |
| A font family, size, weight and style | `NewFontPicker`, `NewFontPickerWithSystemFonts` |
| A light, dark or system appearance setting with a colour | `NewThemeSelector` |

## Date and time pickers

[`DatePicker`](https://pkg.go.dev/github.com/belluikit/bellui/widget#DatePicker)
is a month calendar. `Selected` holds the chosen date and `OnChanged` receives
each new one.

{{< example file="widgets/pickers/main.go" region="inline" >}}

With the keyboard, the arrow keys move between days (and into the next or
previous month), Home and End jump to the edges of the week, Page Up and Page
Down move a month, and Enter or Space selects.

[`TimePicker`](https://pkg.go.dev/github.com/belluikit/bellui/widget#TimePicker)
is the MD3 clock dial. `Hour` and `Minute` hold the time, and `Use24Hour`
switches between the 24-hour dial and AM/PM. The user can also type digits
into the hour and minute fields. [`DateTimePicker`](https://pkg.go.dev/github.com/belluikit/bellui/widget#DateTimePicker)
puts a date picker and a time picker on two tabs.

Each picker has a `Show` function that opens it in a modal dialog with OK
and Cancel buttons. Your function runs only when the user presses OK:

{{< example file="widgets/pickers/main.go" region="modal" >}}

| Picker | In place | In a dialog |
| --- | --- | --- |
| Date | `NewDatePicker` | `ShowDatePicker` |
| Time | `NewTimePicker`, `NewAutoTimePicker` | `ShowTimePicker`, `ShowAutoTimePicker` |
| Date and time | `NewDateTimePicker`, `NewAutoDateTimePicker` | `ShowDateTimePicker`, `ShowAutoDateTimePicker` |

The `Auto` variants choose 12-hour or 24-hour mode from the user's locale.
The others take a `use24h` argument.

## Date entries

[`DateEntry`](https://pkg.go.dev/github.com/belluikit/bellui/widget#DateEntry)
and [`DateTimeEntry`](https://pkg.go.dev/github.com/belluikit/bellui/widget#DateTimeEntry)
are entries for forms. The user types a date in the locale's format, or taps
the button at the end of the field to open a calendar. Characters that cannot
be part of a date are ignored, including in pasted text.

{{< example file="widgets/picker-entries/main.go" region="entries" >}}

- Use the entry's own `OnChanged`, which receives a `time.Time`. The embedded
  `Entry.OnChanged` receives the raw text instead.
- `Selected` holds the last valid value. `Date()` and `DateTime()` parse the
  current text and return an error if it is not valid.
- `SetDate` and `SetDateTime` write a value in the locale's format.
- `DateTimeEntry.Use24Hour` is set from the locale and controls the time
  picker it opens.

## Font picker

[`FontPicker`](https://pkg.go.dev/github.com/belluikit/bellui/widget#FontPicker)
combines controls for a font family, size, weight and style. It reports a
[`FontPickerConfig`](https://pkg.go.dev/github.com/belluikit/bellui/widget#FontPickerConfig)
with `Family`, `Size`, `Weight`, `Bold`, `Italic` and `Underline`.
`TextStyle()` turns that into a `bellui.TextStyle` for a label or rich text.

{{< example file="widgets/font-theme/main.go" region="font" >}}

`NewFontPickerWithSystemFonts` lists the fonts installed on the system,
loading the list in the background. `NewFontPicker` takes your own
[`FontFamily`](https://pkg.go.dev/github.com/belluikit/bellui/widget#FontFamily)
list, each with a name and the font file as a resource. `ShowFontPickerPopup`
opens a picker in a modal dialog with OK and Cancel buttons.

## Theme selector

[`ThemeSelector`](https://pkg.go.dev/github.com/belluikit/bellui/widget#ThemeSelector)
is a ready-made appearance setting: System, Light or Dark above a row of
primary colour swatches. It only reports the user's choice. Call
`ThemeSelection.Apply` to re-theme the app and `SaveThemeSelection` to
remember it:

{{< example file="widgets/font-theme/main.go" region="theme" >}}

[Theme selector](/docs/theming/theme-selector/) in the Theming section covers
it in more detail.
