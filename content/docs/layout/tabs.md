---
title: Tabs
description: AppTabs for switching between the areas of an app, and DocTabs for documents the user opens and closes.
weight: 4
---

Tab containers show one of several pages, with a bar of tab buttons to
switch between them. Each page is a
[`container.TabItem`](https://pkg.go.dev/github.com/belluikit/bellui/container#TabItem):
a title, an optional icon, and the content to show.

| Container | Use it for |
| --- | --- |
| [`container.AppTabs`](https://pkg.go.dev/github.com/belluikit/bellui/container#AppTabs) | The main areas of your app, such as Home, Documents and Settings. A fixed set. |
| [`container.DocTabs`](https://pkg.go.dev/github.com/belluikit/bellui/container#DocTabs) | Open documents. Each tab has a close button, and the bar can have a button to add a tab. |

## AppTabs

{{< example file="layout/tabs/main.go" region="apptabs" >}}

- `container.NewTabItem(text, content)` and
  `container.NewTabItemWithIcon(text, icon, content)` create the pages.
- `SetTabLocation` moves the tab bar: `container.TabLocationTop` (the
  default), `TabLocationBottom`, or one of the sides with
  `TabLocationLeading` and `TabLocationTrailing`.
- `OnSelected` and `OnUnselected` are called when the current tab changes.

To change tabs from code, use `Select(item)` or `SelectIndex(i)`;
`Selected()` and `SelectedIndex()` tell you which one is showing. `Append`,
`Remove`, `RemoveIndex` and `SetItems` change the list of tabs. Removing the
selected tab selects the tab after it, or the one before when it was the
last. `OnUnselected` then receives the removed tab and `OnSelected` the new
one, as when the user picks a tab.
`DisableIndex(i)` and `DisableItem(item)` grey out a tab so it cannot be
chosen, and `EnableIndex` and `EnableItem` undo that.

### Hiding the bar, dragging, disabling

{{< example file="layout/tabs/main.go" region="options" >}}

- `DragToSelect` lets the user drag along the tab titles. The selection
  indicator follows the pointer, and letting go selects the nearest enabled
  tab. Tapping a title still works. It is off by default.
- `HideTabBar` hides the tab bar, its divider and the selection indicator, so
  the selected tab's content fills the whole container. The tabs and the
  selection stay as they are. Call `Refresh` after changing it.
- `Disable()` disables the whole container. It keeps its place, but the user
  cannot switch tabs or open the overflow menu. Every widget inside the tabs
  that can be disabled is disabled too, in every tab, including content you
  add while it is disabled. `Enable()` turns back on only the widgets that
  `Disable` turned off, so a widget you disabled yourself stays disabled.
  `Select` and the other calls from code still work.

{{< callout type="info" >}}
On mobile devices, the tab bar does not go along a long edge of the screen.
In portrait, `TabLocationLeading` becomes `TabLocationTop` and
`TabLocationTrailing` becomes `TabLocationBottom`. In landscape, top becomes
leading and bottom becomes trailing.
{{< /callout >}}

## DocTabs

{{< example file="layout/tabs/main.go" region="doctabs" >}}

DocTabs has the same methods as AppTabs, plus:

- `CreateTab`: when set, a **+** button appears in the tab bar. Your function
  returns the new item; the container appends it and selects it. Return
  `nil` to add nothing.
- `OnClosed`: called after the user closes a tab and it has been removed.
- `CloseIntercept`: when set, closing a tab calls this function **instead**.
  The tab stays open until you call `Remove(item)` yourself, and `OnClosed`
  is not called. Use it to ask about unsaved changes before closing.

Closing the selected tab reports the new selection through `OnUnselected`
and `OnSelected`, as `Remove` does. A disabled DocTabs works like a disabled
AppTabs: the user cannot select, close or create tabs, and the widgets inside
are disabled with it.

## The whole program

The example puts a DocTabs inside one page of an AppTabs.

{{< example "layout/tabs/main.go" >}}
