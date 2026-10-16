---
title: Custom and form dialogs
description: Put any content in a dialog, choose its buttons, and collect validated input with a form dialog.
weight: 2
---

## Your own content

[`dialog.ShowCustom`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#ShowCustom)
shows any canvas object with a title and one button that closes the dialog.
You choose the button's label. The dialog starts at the minimum size of the
content you pass.

{{< example file="dialogs/custom/main.go" region="custom" >}}

## Content with confirm and cancel

[`dialog.ShowCustomConfirm`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#ShowCustomConfirm)
adds a second button. The callback gets `true` for the confirm button and
`false` for the cancel button or any other way of closing. Read the values from
your widgets inside the callback.

{{< example file="dialogs/custom/main.go" region="customconfirm" >}}

`NewCustomConfirm` returns a
[`ConfirmDialog`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#ConfirmDialog),
the same type the yes/no dialog uses, so `SetConfirmImportance` and the other
methods work here too.

## Your own buttons

[`dialog.NewCustomWithoutButtons`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#NewCustomWithoutButtons)
creates a dialog with an empty button row.
[`SetButtons`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#CustomDialog.SetButtons)
then fills it with whatever you like, laid out in one row. Your buttons do not
close the dialog by themselves, so call `Hide()` from their handlers.

{{< example file="dialogs/custom/main.go" region="buttons" >}}

`CustomDialog` also has `SetIcon`, which shows an icon in the dialog's top
corner. Pass `nil` to remove it.

## Forms

[`dialog.ShowForm`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#ShowForm)
lays out a list of
[`widget.FormItem`](https://pkg.go.dev/github.com/belluikit/bellui/widget#FormItem)
values as a form, with confirm and cancel buttons under it.

{{< example file="dialogs/custom/main.go" region="form" >}}

The form dialog ties into validation. If any field has a validator (here the
email entry's `Validator`) and its value fails, the confirm button is disabled.
The button's state is checked when the dialog is built and again whenever a
field's validation result changes. So the callback only ever gets `true` with
valid values. To submit from code, call
[`Submit`](https://pkg.go.dev/github.com/belluikit/bellui/dialog#FormDialog.Submit)
on the dialog that `NewForm` returns; it does nothing while the confirm button
is disabled.

For more on validators and forms, see [Forms](/docs/widgets/forms/).

## The complete example

{{< example "dialogs/custom/main.go" >}}
