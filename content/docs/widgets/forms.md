---
title: Forms
description: The Form widget lays out labelled inputs with hints, checks their validation and shows submit and cancel buttons.
weight: 10
---

[`Form`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Form) turns a
list of inputs into a form: a label beside each input, a hint line under it,
and submit and cancel buttons at the bottom. It also watches the inputs'
validation and only lets the user submit when every field is valid.

## Which one to use

- Use a `Form` for a set of fields that are submitted together, such as a
  sign-up or a settings dialog.
- For a single field, use an [`Entry`](../text-input/) with its own
  validator.
- For a form inside a dialog, use `dialog.ShowForm` from the
  [`dialog`](https://pkg.go.dev/github.com/belluikit/bellui/dialog) package,
  which takes the same `FormItem`s.

## Building a form

{{< example file="widgets/form/main.go" region="form" >}}

Each row is a
[`FormItem`](https://pkg.go.dev/github.com/belluikit/bellui/widget#FormItem)
with a `Text` label, a `Widget` and an optional `HintText` shown under it.
`NewFormItem` sets the first two. `Append(text, widget)` and `AppendItem`
add rows after the form is created. Any widget can be a row, not only
entries.

The buttons appear only when you set their functions:

- `OnSubmit` shows the submit button, labelled `SubmitText` ("Submit" when
  empty). It is the form's high-importance button.
- `OnCancel` shows the cancel button, labelled `CancelText` ("Cancel" when
  empty).

If you set or clear `OnSubmit` or `OnCancel` after the form is on screen,
call `Refresh()` so the buttons update. Setting both to `nil` removes the
button row.

## Validation

A row takes part in validation when its widget can be validated, such as an
`Entry` with a `Validator` (see [Validation](../text-input/#validation)).

- While any such row is invalid, the submit button is disabled.
- For a plain `Entry`, the row's hint line shows `HintText` until the user
  has edited the field and left it. After that, an invalid row shows the
  error message in the error colour in place of the hint. Other widgets that
  can be validated, including `NumericEntry` and `DateEntry`, show the error
  as soon as there is one.
- `Validate()` checks every row and returns the first error, for example
  before you save in code.
- `Disable()` disables both the submit and the cancel button, regardless of
  validation. `Enable()` enables them again; the submit button stays
  disabled while a row is invalid.

## Layout

By default each row has its label on the left and the input on the right.
Set `Orientation` to `widget.Vertical` to put each label above its input, or
`widget.Adaptive` to switch between the two with a mobile device's
orientation.
