---
title: Text input
description: Entry and its variants for passwords, multi-line text, numbers, search, choices and completion, plus validation and the RichEntry editor.
weight: 2
---

[`widget.Entry`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Entry)
is the text field. Most of the other text inputs on this page build on it, so
what you learn about `Entry` (placeholder, callbacks, validation, styles)
applies to them too. `SearchSelect` and `RichEntry` are separate widgets.

{{< screenshot name="inputs" alt="The gallery's Inputs page: a filled and an outlined text field, check boxes and radio buttons" >}}

## Which one to use

| You want | Use |
| --- | --- |
| One line of free text | `NewEntry` |
| A secret, with a button to reveal it | `NewPasswordEntry` |
| Several lines, such as notes or a message | `NewMultiLineEntry` |
| A number, with typing limited to digits | `NewNumericEntry`, `NewIntegerEntry` |
| Free text with a list of common values to pick from | `NewSelectEntry` |
| A search box that filters as the user types | `NewSearchField` |
| One value from a long list, found by typing | `NewSearchSelect` |
| Suggestions under the caret while the user types | `Entry.ShowCompletion` |
| Text with bold, italic and other styled runs | `NewRichEntry` |
| A date or a date and time | `DateEntry`, `DateTimeEntry` (see [Pickers](../pickers/)) |

If the user must pick from a fixed list and may not type their own value, use
[`Select`](../selection/#select) instead.

## Entry

{{< example file="widgets/entry/main.go" region="basic" >}}

- `OnChanged` runs after every edit with the whole text.
- `OnSubmitted` runs when the user presses Enter or Return in a single-line
  entry. In a multi-line entry, Shift+Enter submits.
- `OnEscape`, when set, runs when the user presses Escape.
- `SetText` replaces the text and clears the undo history. `Text` holds the
  current text. `SetText` and `SetPlaceHolder` are safe to call from any
  goroutine.
- `Undo()` and `Redo()` step through the edit history; the usual keyboard
  shortcuts work too.

### Variants

{{< example file="widgets/entry/main.go" region="variants" >}}

- **Style.** `EntryStyleFilled` (the default) has a filled background.
  `EntryStyleOutlined` has a border and no fill.
- **Icon.** `SetIcon` shows an icon at the leading edge. It is decoration
  only and cannot be tapped. `ActionItem` holds an object for the trailing
  edge; the password entry puts its reveal button there.
- **Password.** `NewPasswordEntry` hides the characters and adds a button that
  reveals them. Screen readers are never given the text.
- **Multi-line.** `NewMultiLineEntry` accepts new lines. Set `Wrapping` to
  `bellui.TextWrapWord` to wrap long lines, and `SetMinRowsVisible(n)` to
  make it at least `n` rows tall (3 by default).
- **Read-only.** `ReadOnly = true` stops editing but keeps the normal look,
  and the user can still focus, select and copy the text. `Disable()` instead
  blocks all interaction and draws the entry in the disabled style.
- **Alignment and style.** `Alignment` aligns the text, and `TextStyle` sets
  bold, italic or monospace for the whole entry.

## Validation

Set `Validator` to a function that returns an error for invalid text. A
[`bellui.StringValidator`](https://pkg.go.dev/github.com/belluikit/bellui#StringValidator)
is a `func(string) error`. The
[`data/validation`](https://pkg.go.dev/github.com/belluikit/bellui/data/validation)
package has ready-made ones: `NewRegexp`, `NewTime` and `NewAllStrings`,
which combines several.

{{< example file="widgets/validation/main.go" region="validate" >}}

The entry shows its state at the trailing edge:

- valid, non-empty text gets a check mark;
- invalid text gets an error mark, but only once the user has edited the
  entry and moved focus away, so a half-typed value is not flagged.

Set `AlwaysShowValidationError` to show an error before the user has typed
anything. Set `ValidateWithBorder` to show errors as a thicker border in the
error colour instead of the icon.

`Validate()` runs the validator now and returns its error, which is useful
before you save:

{{< example file="widgets/validation/main.go" region="check" >}}

`SetValidationError` sets the error shown until the next edit, for example an
error your server returned. Inside a [Form](../forms/), the form collects these
errors and disables its submit button while any field is invalid.

## Numbers

{{< example file="widgets/special-entry/main.go" region="numeric" >}}

[`NumericEntry`](https://pkg.go.dev/github.com/belluikit/bellui/widget#NumericEntry)
right-aligns its text and ignores keys that cannot be part of a number,
including in pasted text. `NewNumericEntry` accepts decimals and
`NewIntegerEntry` accepts whole numbers.

- `OnValueChanged` runs with the parsed number when the text changes to a
  valid number. Do not assign `OnChanged` on a `NumericEntry`: it replaces
  the handler that drives `OnValueChanged`.
- `Value()` returns the current number, or an error if the text is not one.
  `SetValue` writes a number.
- `SetBounds(min, max)` sets `Min` and `Max`. A value outside them is a
  validation error such as "minimum is 1". Pass `math.Inf(-1)` or
  `math.Inf(1)` to remove a bound.
- `AllowNegative` and `AllowFloat` control whether a minus sign and a
  decimal point are accepted. `NewNumericEntry` sets both to true.
  `NewIntegerEntry` sets `AllowNegative` to true and `AllowFloat` to false. `LocaleFormat` uses the
  locale's decimal separator instead of a dot.

For a number the user nudges up and down, see [SpinBox](../selection/#spinbox).

## Choices and search

{{< example file="widgets/special-entry/main.go" region="choice" >}}

- [`SelectEntry`](https://pkg.go.dev/github.com/belluikit/bellui/widget#SelectEntry)
  is an entry with a drop-down button listing `options`. The user can pick an
  option or type anything else. `OnOptionSelected` runs only for a pick, never
  for typing. `SetOptions` changes the list.
- [`SearchSelect`](https://pkg.go.dev/github.com/belluikit/bellui/widget#SearchSelect)
  looks like a `Select`. Tapping it opens a pop-up with a search box that
  filters the options. `Preferred` options are listed first, above a divider.
  `ClearSelected` empties the choice.

{{< example file="widgets/special-entry/main.go" region="search" >}}

[`SearchField`](https://pkg.go.dev/github.com/belluikit/bellui/widget#SearchField)
is an entry with a search icon, a clear button and the placeholder
"Search...". The function you pass to `NewSearchField` is its own `OnChanged`
and runs on every keystroke, which suits live filtering. `OnSearched` runs
when the user presses Enter.

## Completion

`Entry.ShowCompletion(title, items)` opens a suggestion list just below the
caret, or above it when there is no room. Your code decides what to suggest,
usually from `OnChanged`. Passing no items hides the list. When the user
accepts an item with Enter, Tab or a tap, `OnCompletionSelected` receives
it. Space types a space into the entry.

{{< example file="widgets/special-entry/main.go" region="completion" >}}

While the list is open the user keeps typing into the entry, and the arrow
keys move through the suggestions. A
[`CompletionItem`](https://pkg.go.dev/github.com/belluikit/bellui/widget#CompletionItem)
can also carry an `Icon`, a short `Detail` and longer `Documentation`
(optionally Markdown), which is shown in a panel beside the list.
`CompletionPopup` is the same list as a standalone widget, for editors that
manage their own caret.

## RichEntry

[`RichEntry`](https://pkg.go.dev/github.com/belluikit/bellui/widget#RichEntry)
is an editor whose content is a list of styled
[`RichText`](../text/#richtext) segments instead of one string. It supports
typing, selection, undo and redo, input methods, images, and word wrapping
through `Wrapping`.

{{< example file="widgets/richentry/main.go" region="rich" >}}

- `SetSegments` replaces the content and `Segments()` returns it. Treat the
  returned slice as read-only.
- `ApplyStyle` styles the current selection. With nothing selected it sets
  `TypingStyle`, the style for the next characters typed.
- `String()` returns the plain text. `OnChanged` runs after each edit, and
  `OnContentChanged` receives the segments.
- `InsertImage` and `InsertInlineImage` add an image as a block or within a
  line of text.

`RichEntry` is newer than `Entry` and has fewer features. It has no
validator, placeholder or completion, and images added with `SetSegments` are
shown but the caret cannot stop on them yet. For plain text, use `Entry`.
