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

The package targets API level 36 and installs on Android 6 (API 23) and
later. If a library you ship needs a newer Android, raise the minimum with
`--android-min-sdk`, which `bellui package` and `bellui release` both take:

```sh
bellui package -os android --app-id com.example.notes --android-min-sdk 26
```

The manifest, the resource and dex tools and the native build all use that
level, so the native code never needs a newer Android than the manifest
promises. The lowest level you can set is 21.

### The android folder

An `android` folder next to your `main` package adds what the manifest has
no field for. Each part is optional, and each needs `ANDROID_HOME` set.

| Path | What `bellui package` does with it |
| --- | --- |
| `android/*.java` | Compiles it against the SDK's `android.jar` and adds it to the APK, for small pieces of Java such as a JNI bridge. |
| `android/libs/*.jar` | Puts it on the classpath of that Java, and adds it to the APK. |
| `android/libs/*.aar` | Adds the Android library: its classes, its native libraries for each architecture you build, its assets and its resources. Its manifest's permissions, features, queries and application components are merged into the app's. |
| `android/res/` | The app's own resources, laid out like the `res` folder of a Gradle module. They are linked after the libraries' resources, so yours win a name clash. |

An `AndroidManifest.xml` next to your `main` package (not in `android`)
replaces the generated manifest. It must declare the
`org.golang.app.GoNativeActivity` activity. `--app-version` and `--app-build`
still replace its version numbers, with a warning when they differ.

### Android libraries

With jars and AARs in `android/libs`, your Java can use AndroidX and Google
Play services libraries. The command does not download anything: the folder
must hold the library and every library it depends on, as listed in each
one's Maven POM file.

The manifest merge is simpler than Gradle's:

- An element whose `android:name` your manifest already declares is left out.
  Your manifest wins.
- An element two libraries declare is kept once.
- A library's `<application>` attributes are not merged. Put them in your
  own manifest.
- One merge rule is read, `tools:node="remove"`. An element in your own
  manifest marked with it keeps every library's element of the same name and
  `android:name` out of the package, and is itself left out. Other `tools:`
  attributes are dropped.

For example, an app that bundles an ads library but does not ask for
tracking removes the advertising ID permission that library brings:

```xml
<manifest xmlns:android="http://schemas.android.com/apk/res/android"
    xmlns:tools="http://schemas.android.com/tools">
  <uses-permission android:name="com.google.android.gms.permission.AD_ID"
      tools:node="remove" />
  ...
</manifest>
```

An AAR with native libraries for some architectures but not for one you
build fails the build. Build only the architectures it has, for example
`-os android/arm64`.

Native libraries are loaded straight from the APK, not unpacked at install
(the generated manifest sets `android:extractNativeLibs="false"`). A library
loaded with `System.loadLibrary` works. One that needs its `.so` as a file on
disk does not: give it your own `AndroidManifest.xml` with that attribute set
to `"true"`.

The shared libraries your Go code loads, such as FFmpeg's for
[`x/mmedia`](/extensions/mmedia/), go into the APK by themselves. See
[Shared libraries the app loads](/docs/packaging/package/#shared-libraries-the-app-loads).

## iOS

What you need:

- a Mac with Xcode;
- for a device, an Apple development certificate in your keychain.

Apps packaged with `bellui` v1.1.0 or later run on iOS 15 and later.

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

### The ios folder

An `ios` folder next to your `main` package adds what the manifest has no
field for. Each part is optional.

| Path | What `bellui package -os ios` does with it |
| --- | --- |
| `ios/Info.plist` | Merges the `<dict>` into the generated `Info.plist`. Values of any type work. |
| `ios/Entitlements.plist` | Signs the app with these entitlements as well as the ones `bellui` sets. `bellui release` uses the same ones. |
| `ios/PrivacyInfo.xcprivacy` | Merges your declarations into the app's privacy manifest. |
| `ios/Frameworks/*.xcframework`, `*.framework` | Links the slice for each architecture you build. A dynamic framework is also copied into the app and signed. |
| `ios/Frameworks/*.a` | Links a plain static library into every architecture it holds. |

**Info.plist values.** For example, an app that keeps playing audio in the
background, as [`x/mmedia`](/extensions/mmedia/) can, declares the audio
background mode:

```xml
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>UIBackgroundModes</key>
	<array>
		<string>audio</string>
	</array>
</dict>
</plist>
```

A key the command owns is an error: `CFBundleIdentifier`,
`CFBundleExecutable`, `CFBundleName`, `CFBundleShortVersionString`,
`CFBundleVersion` and `UIApplicationSceneManifest`. So is a key the manifest
also writes, such as a usage description from `ios_usage_descriptions` or the
URL types from `url_schemes`. Any other generated key, such as
`UISupportedInterfaceOrientations`, takes your value.

**Entitlements.** Setting `com.apple.developer.associated-domains` here and
`associated_domains` in the manifest is an error, and so is
`application-identifier`, which signing sets. `$(AppIdentifierPrefix)` and
the other build settings Xcode knows are expanded. With automatic signing,
the capability an entitlement needs, such as HealthKit, is registered on the
app's ID. [Sign in with Apple](/docs/advanced/auth/#on-phones) needs one.

**Privacy manifest.** Every iOS app `bellui` packages ships a privacy
manifest, which App Store Connect requires. It declares the two "required
reason" APIs every Go program calls: file timestamps, for files in the app's
own container (`C617.1`), and the system boot time, for measuring elapsed
time (`35F9.1`). Declare what your own code and libraries call in
`ios/PrivacyInfo.xcprivacy`; its API reasons are merged with these, category
by category. Tracking and collected data are yours alone to declare.

**Frameworks.** The command adds the flags that link a framework or library
to each architecture's cgo build, so your cgo code needs only the
`#include`. It does not compile Swift. Build Swift code into an xcframework in
Xcode, or into a static library with `swiftc -emit-library -static`, and
expose a C or `@objc` API. Link such a library from `ios/Frameworks`, not
from a `#cgo LDFLAGS: -L` line: from Go 1.27.4, the `go` command refuses the
objects Xcode 27.2's `swiftc` writes in a folder a `#cgo` line names. A
static library does not say whether it was built for a device or the
simulator, so for an app built for both, ship an xcframework. Anything else
in `ios/Frameworks` is an error, apart from files whose name starts with a
dot.

The shared libraries your Go code loads, such as FFmpeg's, are embedded and
signed by themselves. See
[Shared libraries the app loads](/docs/packaging/package/#shared-libraries-the-app-loads).

## The complete example

{{< example "platforms/mobile/main.go" >}}
