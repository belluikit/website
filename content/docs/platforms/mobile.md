---
title: Mobile
description: Simulate a phone on your computer, package for Android and iOS, and use orientation, lifecycle and keyboard APIs.
weight: 2
---

Bellui apps run on Android and iOS from the same code as on the desktop. You
build them with the [`bellui` command](/docs/packaging/), which produces an
APK or an iOS app bundle. Rendering uses WebGPU: Vulkan on Android and Metal
on iOS. The [OpenGL ES renderer](/develop/rendering/) is available with
`--tags gl`.

## Trying it without a phone

The `mobile` build tag runs your app on your computer with the mobile driver,
the same code that runs on a phone:

```sh
go run -tags mobile .
```

It behaves like a phone in the ways that matter for layout and flow:
`CurrentDevice().IsMobile()` is `true`, windows stack instead of opening side
by side, and dialogs stay inside the window. It does not emulate the operating
system, so test on a device or emulator before you ship.

## How mobile apps behave

- **Windows form a stack.** Only the top visible window is shown. Opening a
  window covers the current one, and closing it reveals the one below.
- **The OS runs the app's lifecycle.** On a phone `App.Quit` closes the app's
  windows but does not stop the app; the system decides when the app stops.
- **Dialogs stay inside the app window.** `SetWindowHosted` has no effect.
- **File dialogs use the system document pickers**, which are the only way to
  reach files outside the app's own storage. See
  [File and folder dialogs](/docs/dialogs/files/).

## Device, orientation and lifecycle

{{< example file="platforms/mobile/main.go" region="device" >}}

[`bellui.CurrentDevice`](https://pkg.go.dev/github.com/belluikit/bellui#CurrentDevice)
reports the orientation and whether a hardware keyboard is available.
`bellui.IsVertical` and `bellui.IsHorizontal` sort the four
[`DeviceOrientation`](https://pkg.go.dev/github.com/belluikit/bellui#DeviceOrientation)
values into portrait and landscape. A layout that should change with the
orientation usually does not need this: the window is resized when the device
rotates, and layouts that respond to size, such as those described under
[Layout](/docs/layout/), adapt by themselves.

[`App.Lifecycle`](https://pkg.go.dev/github.com/belluikit/bellui#Lifecycle)
tells you when the app starts and stops, and when it goes to the background
and comes back. On a phone, going to the background is the moment to save
state, because the system may stop the app afterwards without warning.

{{< example file="platforms/mobile/main.go" region="lifecycle" >}}

## The on-screen keyboard

The on-screen keyboard appears when a widget that accepts typing gets focus,
and hides when focus leaves it. To control it directly, assert the device to
[`mobile.Device`](https://pkg.go.dev/github.com/belluikit/bellui/driver/mobile#Device):

{{< example file="platforms/mobile/main.go" region="keyboard" >}}

`ShowVirtualKeyboardType` picks a layout: `mobile.DefaultKeyboard`,
`SingleLineKeyboard`, `NumberKeyboard` or `PasswordKeyboard`. A custom text
widget can choose its layout by implementing
[`mobile.Keyboardable`](https://pkg.go.dev/github.com/belluikit/bellui/driver/mobile#Keyboardable).

## Android

What you need:

- Go and a C compiler, as for any Bellui app;
- the Android SDK, with `ANDROID_HOME` pointing to it;
- the Android NDK, found through `ANDROID_NDK_HOME` or in
  `$ANDROID_HOME/ndk-bundle`;
- the SDK build tools, whose `apksigner` makes the APK installable on Android
  11 and later.

```sh
bellui package -os android --app-id com.example.notes
bellui install -os android --app-id com.example.notes   # with adb, to a connected device
```

`-os android` builds for all four architectures (arm, arm64, 386, amd64);
`android/arm64` and the other variants build one. The app ID must contain a
dot, and no part of it may contain `-` or start with a digit or `_`. For an
adaptive launcher icon, set `adaptive_icon` in the
[manifest](/docs/packaging/metadata/). For Google Play, use
[`bellui release`](/docs/packaging/release/), which builds a signed `.aab`.

## iOS

What you need:

- a Mac with Xcode;
- for a device, an Apple development certificate in your keychain.

```sh
# The simulator needs the iossimulator target and build tag.
bellui package -os iossimulator --tags iossimulator --app-id com.example.notes
bellui install -os iossimulator --tags iossimulator --app-id com.example.notes

# A device build.
bellui package -os ios --app-id com.example.notes
bellui install -os ios --app-id com.example.notes   # needs ios-deploy
```

For a device build, `--certificate` names the signing certificate and
`--profile` a provisioning profile. Leave `--profile` out to let Xcode sign
automatically for the certificate's team. `bellui install -os iossimulator`
installs into the booted simulator and starts the app.

If your app asks for a permission, such as the camera, iOS requires a
sentence explaining why. Add it under `ios_usage_descriptions` in the
[manifest](/docs/packaging/metadata/). For the App Store, use
[`bellui release`](/docs/packaging/release/), which builds a signed `.ipa`.

{{< callout type="info" >}}
WebGPU on iOS is verified in the simulator. Device builds compile and link but
have not yet been verified on hardware.
{{< /callout >}}

## The complete example

{{< example "platforms/mobile/main.go" >}}
