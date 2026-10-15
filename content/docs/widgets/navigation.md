---
title: Navigation
description: The MD3 top app bar, navigation bar, navigation rail, navigation drawer and bottom app bar.
weight: 5
---

Navigation widgets move the user between the main screens of your app. They
show destinations and report which one was tapped. Swapping the content is up
to you.

{{< screenshot name="navigation" alt="The gallery's Navigation page: a navigation bar with Home, Search and Profile, and a navigation rail beside a content area" >}}

## Which one to use

| You want | Use |
| --- | --- |
| A title bar with a menu or back icon and actions | `TopAppBar` |
| 3 to 5 main destinations along the bottom, on a phone | `NavigationBar` |
| 3 to 7 destinations down the side, on a tablet or desktop | `NavigationRail` |
| Many destinations in a panel that slides in | `NavigationDrawer` |
| A bottom bar with actions and a FAB | `BottomAppBar` |
| Sections of one screen | tabs: `container.NewAppTabs` (see [Layout](/docs/layout/)) |

## Navigation bar and rail

[`NavigationBar`](https://pkg.go.dev/github.com/belluikit/bellui/widget#NavigationBar)
and [`NavigationRail`](https://pkg.go.dev/github.com/belluikit/bellui/widget#NavigationRail)
take a list of items, each with a `Label` and an `Icon`, and a function that
receives the index of the tapped item. A tap also updates `Selected`, which
highlights the item. Set `Selected` yourself (then call `Refresh`) when the
destination changes some other way. The rail shows labels under its icons by
default; set `ShowLabels` to false for icons only.

{{< example file="widgets/navigation/main.go" region="bars" >}}

Place a bar at the bottom of a `container.NewBorder` and a rail at its
leading side. A common pattern is the bar on narrow windows and the rail on
wide ones; this example shows both so you can compare them.

## Top app bar and navigation drawer

{{< example file="widgets/navigation/main.go" region="top" >}}

[`TopAppBar`](https://pkg.go.dev/github.com/belluikit/bellui/widget#TopAppBar)
shows a `Title`, a leading navigation icon that calls `OnNavigation`, and
trailing `Actions`, each a
[`TopAppBarAction`](https://pkg.go.dev/github.com/belluikit/bellui/widget#TopAppBarAction)
with an icon and a function.

[`NavigationDrawer`](https://pkg.go.dev/github.com/belluikit/bellui/widget#NavigationDrawer)
is a side panel that you open over the content with `Show(canvas)` or
`Toggle(canvas)`, and close with `Hide()`. It is modal by default: the rest
of the window is dimmed while it is open, and a tap outside the drawer closes
it. Tapping a destination calls `OnChanged` and closes the drawer.

- An item with `IsDivider: true` draws a separator line. Dividers count as
  items, so indexes passed to `OnChanged` include them.
- `Header` adds content, such as an account name, above the items.

## Bottom app bar

[`BottomAppBar`](https://pkg.go.dev/github.com/belluikit/bellui/widget#BottomAppBar)
is a bar of icon actions with an optional FAB on its right. Pass `nil`
for the FAB if you do not need one.

{{< example file="widgets/bottom-bar/main.go" region="bottom" >}}

Place it at the bottom of a `container.NewBorder`, like a navigation bar.
