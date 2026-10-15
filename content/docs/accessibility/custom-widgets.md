---
title: Accessible custom widgets
description: Describe your own widgets to screen readers with a role, a name, a value, a state and actions, and test the result.
weight: 2
---

A widget you write is invisible to screen readers until it describes
itself. You do that by implementing interfaces from the
[`core`](https://pkg.go.dev/github.com/belluikit/bellui/core) package. The
platform bridges look for them; there is nothing to register.

This page builds a five-star rating control that screen readers announce as
a slider: "Food, slider, 3 of 5 stars".

## Role and name

Every accessible widget implements
[`core.Accessible`](https://pkg.go.dev/github.com/belluikit/bellui/core#Accessible):

- `AccessibleRole()` returns a role constant such as `core.RoleButton`,
  `core.RoleCheckbox` or `core.RoleSlider`. Pick the role whose behaviour
  matches your widget, so users know how to operate it.
- `AccessibleLabel()` returns the name.

If your widget embeds `widget.BaseWidget`, apps can also override the name
with `SetAccessibilityLabel`, as for standard widgets.

## Value, state and actions

Optional interfaces add more:

| Interface | Methods | Use it for |
| --- | --- | --- |
| [`core.AccessibleValuer`](https://pkg.go.dev/github.com/belluikit/bellui/core#AccessibleValuer) | `AccessibleValue()`, `AccessibleRange()` | A value to speak, and a numeric range for adjustable controls. Return `0, 0, 0` from `AccessibleRange` when there is no range. |
| [`core.AccessibleStateful`](https://pkg.go.dev/github.com/belluikit/bellui/core#AccessibleStateful) | `AccessibleState()` | Flags such as `core.StateChecked`, `StateSelected` or `StateExpanded`, combined with `\|`. |
| [`core.AccessibleActioner`](https://pkg.go.dev/github.com/belluikit/bellui/core#AccessibleActioner) | `AccessibleActions()`, `PerformAccessibleAction(name)` | What a screen reader can do with the widget, such as `core.ActionPress`, `ActionToggle`, `ActionIncrement`. |
| [`core.AccessibleDescriber`](https://pkg.go.dev/github.com/belluikit/bellui/core#AccessibleDescriber) | `AccessibleDescription()` | Extra help text, beyond the name. |

The rating widget uses a value and two actions:

{{< example file="accessibility/customwidget/main.go" region="semantics" >}}

How the platforms use these:

- **Standard actions** map to the platform's own commands where one exists.
  On macOS and iOS, VoiceOver's increment and decrement commands call
  `increment` and `decrement`. On Windows, setting a slider's value steps
  through them, so your widget keeps control of the step size. On macOS,
  VoiceOver's press command runs `press`, `toggle`, `select` or `dismiss`.
- **Other actions**, such as `rename` or `delete`, appear in the screen
  reader's extra actions menu on macOS, Linux and Android. Windows and iOS
  do not show them.
- **Disabled** does not need reporting. If your widget implements
  `core.Disableable` (it has `Disable`, `Enable` and `Disabled` methods),
  Bellui adds `StateDisabled` when it is disabled.

## Children

For some roles, a screen reader treats the widget as one unit and ignores
the accessible widgets inside it. A slider, button, checkbox, switch, tab,
heading, link or image is announced once, not once for itself and again for
the label that draws its text. The rating widget draws its stars with a
`widget.Label`, and that label does not appear separately.

Containers and groups are different: their accessible children stay in the
tree. If you build a composite widget from standard widgets, such as a
search box with a button, give the outer widget a grouping role or no role
at all, and let the inner widgets describe themselves.

## Test it

[`core.SnapshotAccessible`](https://pkg.go.dev/github.com/belluikit/bellui/core#SnapshotAccessible)
returns what a screen reader would receive for one widget, so you can check
it in an ordinary Go test, without the build tag and without a screen
reader:

{{< example file="accessibility/customwidget/main_test.go" region="test" >}}

The snapshot applies the same rules as the platform bridges: a label set
with `SetAccessibilityLabel` wins over `AccessibleLabel`, and a disabled
widget gets `StateDisabled`.
[`core.WalkAccessibleTree`](https://pkg.go.dev/github.com/belluikit/bellui/core#WalkAccessibleTree)
visits every accessible object under a root, if you want to check a whole
screen.

The rating widget is also operable from the keyboard. See
[Keyboard and focus](../keyboard/#make-a-custom-widget-focusable).

The full program:

{{< example "accessibility/customwidget/main.go" >}}
