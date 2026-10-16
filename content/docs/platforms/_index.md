---
title: Platforms
description: What is specific to desktop, mobile and web apps, and how to use each platform's extras.
weight: 13
---

The same Bellui code runs on macOS, Windows, Linux and BSD, Android, iOS and
in web browsers. Most of an app never needs to know which. These pages cover
the parts that differ: features only one kind of platform has, and how
building and running work there.

{{< cards >}}
  {{< card link="desktop/" title="Desktop" subtitle="Main menus, the system tray, notifications, several windows and keyboard shortcuts." >}}
  {{< card link="mobile/" title="Mobile" subtitle="Simulating a phone, packaging for Android and iOS, orientation and lifecycle." >}}
  {{< card link="web/" title="Web" subtitle="Building to WebAssembly, serving the app, and WebGPU or WebGL." >}}
{{< /cards >}}

## Finding out where you are

[`bellui.CurrentDevice()`](https://pkg.go.dev/github.com/belluikit/bellui#CurrentDevice)
describes the device the app runs on. `IsMobile()` reports a phone or tablet,
and `IsBrowser()` reports a web page. Prefer these to checking
`runtime.GOOS`: they are also correct in the mobile simulation described on
the [Mobile](mobile/) page.

Platform-specific features are offered as optional interfaces. You check for
them with a type assertion and skip the feature when it is missing:

| Feature | Assert | Available on |
| --- | --- | --- |
| System tray | `app.Driver().(desktop.App)` | Desktop |
| Global shortcuts | `app.Driver().(desktop.GlobalShortcutApp)` | Desktop |
| Raw key down and up events | `window.Canvas().(desktop.Canvas)` | Desktop, web |
| Window position | `window.(desktop.PositionableWindow)` | Desktop |
| Badge count on the app icon | `app.(bellui.BadgeApp)` | Desktop, iOS |
| Start at login | `app.(bellui.LaunchAtLoginApp)` | Desktop |
| Single running instance | `app.(bellui.SingleInstanceApp)` | Desktop (a no-op elsewhere) |
| Virtual keyboard control | `bellui.CurrentDevice().(mobile.Device)` | Mobile |
