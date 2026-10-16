---
title: Desktop
description: Main menus, the system tray, notifications, several windows, keyboard shortcuts and other desktop extras.
weight: 1
---

On macOS, Windows, Linux and BSD, Bellui draws its windows with the system's
window manager and adds a set of desktop features. This page uses two
examples: one for menus, the tray, notifications and windows, and one for the
keyboard.

## The main menu

[`Window.SetMainMenu`](https://pkg.go.dev/github.com/belluikit/bellui#Window)
sets the menu bar. Build it from
[`bellui.NewMainMenu`](https://pkg.go.dev/github.com/belluikit/bellui#NewMainMenu),
[`bellui.NewMenu`](https://pkg.go.dev/github.com/belluikit/bellui#NewMenu) and
[`bellui.NewMenuItem`](https://pkg.go.dev/github.com/belluikit/bellui#NewMenuItem).

{{< example file="platforms/desktop/main.go" region="menu" >}}

- On macOS the menu goes into the system menu bar at the top of the screen.
  Elsewhere it is drawn at the top of the window.
- A menu item's `Shortcut` is shown next to its label, and pressing the keys
  runs the item's action while the window has focus.
  [`event.KeyModifierShortcutDefault`](https://pkg.go.dev/github.com/belluikit/bellui/event#KeyModifierShortcutDefault)
  is Command on macOS and Control elsewhere.
- Outside macOS, the first menu always ends with a **Quit** item. If you do not
  add one, Bellui adds it; a last item labelled "Quit" is treated as that item.
- `MenuItem` also has `Checked`, `Disabled`, `Icon`, `ChildMenu` for sub-menus,
  and `IsSeparator`. After changing items, call `Refresh()` on the menu.

## The system tray

The system tray (the menu bar extras on macOS, the notification area on
Windows) can hold an icon for your app with its own menu.

{{< example file="platforms/desktop/main.go" region="tray" >}}

| Method | What it does |
| --- | --- |
| `SetSystemTrayMenu` | Shows the tray icon with this menu. |
| `SetSystemTrayIcon` | Sets the tray icon. Without it, the app icon is used. A themed resource adapts to light and dark menu bars where the system does that. |
| `SetSystemTrayTooltip` | Sets the hover text, where the system shows one. |
| `SetSystemTrayWindow` | Makes a left click on the icon show this window. It also makes closing the window hide it instead, so the app keeps running. |
| `SetSystemTrayOnTapped` | Runs your own function on a left click instead. This and `SetSystemTrayWindow` share one slot; the last one set wins. |

Because closing the tray window only hides it, the app keeps running in the
tray. Give the tray menu an item that calls `App.Quit` so users can still leave.

{{< callout type="warning" >}}
On Linux the tray icon appears only if the desktop runs a StatusNotifierItem
host. KDE Plasma and most Wayland compositors do; stock GNOME needs the
AppIndicator extension, and older trays that only support the XEmbed
protocol are not supported. Never make the tray the only way to reach a
feature.
{{< /callout >}}

{{< callout type="info" >}}
In this release, assert `desktop.App` on `a.Driver()`, as above. The app value
returned by `app.New` does not implement the tooltip and tap methods, so
`a.(desktop.App)` fails.
{{< /callout >}}

## Notifications

{{< example file="platforms/desktop/main.go" region="notify" >}}

[`bellui.SendNotification`](https://pkg.go.dev/github.com/belluikit/bellui#SendNotification)
shows a title and text in the system's notification area.
[`bellui.SendNotificationWithOptions`](https://pkg.go.dev/github.com/belluikit/bellui#SendNotificationWithOptions)
adds a function to run when the user clicks the notification, action buttons,
and a per-notification icon. Your functions run on the UI thread, but only
while the app is still running.

Where actions or click handling are not available, the notification still
appears with its title and text. On macOS the rich form needs a packaged app
(a `.app` bundle with an app ID); a program started with `go run` shows plain
notifications.

## More than one window

{{< example file="platforms/desktop/main.go" region="window" >}}

On the desktop every window from `App.NewWindow` is a separate OS window, and
any number can be open. The app quits when its last window closes. To make
one window the main one, so that closing it quits the app even while others
are open, call `SetMaster` on it (the
[`bellui.MasterWindow`](https://pkg.go.dev/github.com/belluikit/bellui#MasterWindow)
interface). To ask before a window closes, use `SetCloseIntercept`.

The [`driver/desktop`](https://pkg.go.dev/github.com/belluikit/bellui/driver/desktop)
package has more window control. Assert the window to:

- [`desktop.Window`](https://pkg.go.dev/github.com/belluikit/bellui/driver/desktop#Window)
  to request a position, keep the window on top, or go full screen on a second
  monitor;
- [`desktop.PositionableWindow`](https://pkg.go.dev/github.com/belluikit/bellui/driver/desktop#PositionableWindow)
  to read the position back and be told when the window moves, which is how
  you save and restore window placement.

`a.Driver().(desktop.ScreensDriver)` lists the connected monitors with their
size, usable area, scale and which one is primary. Positions are ignored on
Wayland, which does not let apps place windows.

## Keyboard shortcuts

A shortcut added to a window's canvas runs while that window has focus:

{{< example file="platforms/shortcuts/main.go" region="canvas" >}}

[`desktop.CustomShortcut`](https://pkg.go.dev/github.com/belluikit/bellui/driver/desktop#CustomShortcut)
combines a key with modifiers such as `event.KeyModifierShift`,
`KeyModifierControl`, `KeyModifierAlt` and `KeyModifierSuper`. A focused
widget that handles shortcuts itself, such as an entry handling copy and
paste, gets them first; menu shortcuts are checked before both.

### Key down and key up

The usual key callbacks report keys as they are typed. For games and other
apps that need to know when a key goes down and when it comes up, assert the
canvas to
[`desktop.Canvas`](https://pkg.go.dev/github.com/belluikit/bellui/driver/desktop#Canvas):

{{< example file="platforms/shortcuts/main.go" region="keys" >}}

### Global shortcuts

A global shortcut works even when your app is in the background:

{{< example file="platforms/shortcuts/main.go" region="global" >}}

Registration can fail, for example because another app already owns the keys,
so always check the error. On Wayland the request goes through the desktop
portal, which may ask the user for permission; a `nil` error there means the
request was accepted, not that the shortcut is already active.

## Other desktop features

| Feature | How |
| --- | --- |
| Badge on the dock or taskbar icon | `a.(bellui.BadgeApp).SetBadge(n)`; `0` clears it. |
| Start at login | `a.(bellui.LaunchAtLoginApp).SetLaunchAtLogin(true)`. Meant for installed apps; on macOS it needs a `.app` bundle. |
| One running copy | `a.(bellui.SingleInstanceApp).SetSingleInstance(func(args []string) {…})`, called before `Run`. A second launch passes its arguments to the first and should exit. |
| Files and links opened with the app | `a.(bellui.OpenHandlerApp)` with `SetOnFileOpened` and `SetOnURLOpened`, together with [file associations and URL schemes](/docs/packaging/metadata/) in the manifest. |
| Mouse buttons and cursors | Widgets can implement [`desktop.Mouseable`](https://pkg.go.dev/github.com/belluikit/bellui/driver/desktop#Mouseable) and [`desktop.Cursorable`](https://pkg.go.dev/github.com/belluikit/bellui/driver/desktop#Cursorable). |

## The complete examples

{{< example "platforms/desktop/main.go" >}}

{{< example "platforms/shortcuts/main.go" >}}
