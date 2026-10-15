---
title: Packaging and installing
description: What bellui package produces for each platform, and how bellui install puts it on a computer or phone.
weight: 2
---

Run `bellui package` in the folder of your app's `main` package. It builds the
app and wraps it in the target platform's format:

```sh
bellui package -os darwin
```

`-os` (also spelled `--target`) chooses the platform. Without it, the command
packages for the system it runs on. The app name comes from `--name`, then the
manifest's `name`, then the last element of your module path.

| `-os` | Output |
| --- | --- |
| `darwin` | `Name.app`, a macOS application bundle. |
| `windows` | `Name.exe` with the icon, version information and an application manifest built in. |
| `linux`, `freebsd`, `openbsd`, `netbsd` | `Name.tar.xz` (`.tar.gz` on OpenBSD) holding the binary, a desktop entry, the icon and a `Makefile`. |
| `android`, `android/arm64`, `android/arm`, `android/amd64`, `android/386` | `Name.apk`. Plain `android` includes all four architectures. |
| `ios`, `iossimulator` | `Name.app`, for a device or for the iOS Simulator. |
| `web` (or `wasm`) | A `wasm` folder holding everything needed to serve the app. |

Useful flags:

- `--app-id`, `--name`, `--app-version`, `--app-build` and `--icon` override
  the [manifest](../metadata/).
- `--release` builds an optimised, non-debuggable package. It is still signed
  with a development key on Android, so you can install it directly. For a
  store submission use [`bellui release`](../release/).
- `--tags` passes Go build tags, for example `--tags gl` for the
  [OpenGL renderer](/develop/rendering/).
- `--exe` wraps an executable you built yourself instead of building one
  (desktop and web only).
- `--src` packages a different folder than the current one (not for mobile).

{{< callout type="info" >}}
Desktop packages are built with cgo, which needs a C compiler for the target
system. The simplest way is to package each desktop platform on that platform.
{{< /callout >}}

## Platform notes

{{< tabs >}}

{{< tab name="macOS" >}}
The bundle holds the executable in `Contents/MacOS`, an `icon.icns` made from
your icon, and an `Info.plist` with the app ID, name, version and build
number. File associations, URL schemes and usage descriptions from the
manifest are written into `Info.plist` too.

Without an app ID the bundle gets `com.example.<name>`.
{{< /tab >}}

{{< tab name="Windows" >}}
The executable is built as a GUI program, so no console window opens with it.
The icon and version details are embedded as Windows resources. If your source
folder has a file named `<name>.exe.manifest`, matching the executable's name,
it is used instead of the generated application manifest.
{{< /tab >}}

{{< tab name="Linux and BSD" >}}
Unpack the archive and use its `Makefile`:

```sh
tar -xf Name.tar.xz
sudo make install        # into /usr/local (or /usr)
make user-install        # into ~/.local, no root needed
```

`uninstall` and `user-uninstall` remove the files again. The desktop entry uses
the manifest's `linux_and_bsd` details and lists the MIME types from your file
associations and URL schemes.
{{< /tab >}}

{{< tab name="Android" >}}
You need the Android SDK and NDK. Set `ANDROID_HOME` to the SDK, and
`ANDROID_NDK_HOME` to the NDK if it is not in `$ANDROID_HOME/ndk-bundle`. An
app ID is required. It must contain a dot, and no part of it may contain `-`
or start with a digit or `_`.

```sh
bellui package -os android/arm64 --app-id com.example.notes
```

The APK is re-signed with `apksigner` from the SDK's build tools, which modern
Android versions require. If the build tools are missing, the command warns
that the APK will not install on Android 11 or later. See
[Mobile](/docs/platforms/mobile/) for more.
{{< /tab >}}

{{< tab name="iOS" >}}
iOS packages need a Mac with Xcode. An app ID is required.

```sh
bellui package -os ios --app-id com.example.notes
bellui package -os iossimulator --tags iossimulator --app-id com.example.notes
```

A simulator build needs the `iossimulator` build tag as well as the
`iossimulator` target. See [Mobile](/docs/platforms/mobile/) for signing and
the other requirements.
{{< /tab >}}

{{< tab name="Web" >}}
The `wasm` folder holds `index.html`, the app's `.wasm` file, Go's
`wasm_exec.js` loader, the icon, and a few support files. Serve the folder
with any static web server. See [Web](/docs/platforms/web/).
{{< /tab >}}

{{< /tabs >}}

## Installing on this computer

`bellui install` packages the app and installs it in the usual place:

| System | Installed to |
| --- | --- |
| macOS | `/Applications` |
| Windows | `Program Files\<name>`, after an administrator prompt |
| Linux and BSD | `/usr/local` (or `/usr` if there is no `/usr/local`) |

`--dst` chooses another folder. On Linux, installing to the system folders
needs root.

`bellui install` can also fetch an app from its Git repository, build it in
release mode and install it:

```sh
bellui install github.com/example/notes@latest
```

`@<branch-or-tag>` installs that branch or tag, and with no `@` the default
branch is used. `@latest` takes the last tag starting with `v` in the
repository's tag list.

## Installing on a phone or simulator

With `-os`, `bellui install` packages for a mobile target and installs there:

| `-os` | Uses |
| --- | --- |
| `android` and the `android/…` variants | `adb install` to the connected device or emulator |
| `ios` | `ios-deploy` to the connected device |
| `iossimulator` | `xcrun simctl install` to the booted simulator, then starts the app |

The tools in the right-hand column must be installed and on your `PATH`.

## Building without packaging

`bellui build` builds only the executable, with the same `-os`, `--tags`,
`--release` and `-o` flags. `--pprof` adds Go's profiling server, on port
6060 unless you set `--pprof-port`.
