---
title: Releasing
description: Build signed packages for the Mac App Store, Microsoft Store, Google Play and the App Store with bellui release.
weight: 3
---

`bellui release` builds the app in release mode, packages it, and then does the
signing and store packaging each platform needs. It takes most of the flags of
`bellui package` (not `--src`, `--exe` or `--release`; it always builds in
release mode), plus the signing details below. Each platform's own tools
must be installed; the command calls them for you.

{{< tabs >}}

{{< tab name="macOS" >}}
```sh
bellui release -os darwin --app-id com.example.notes --app-version 1.2.0 \
    --profile "Notes Distribution" --category productivity
```

| Flag | Needed | Notes |
| --- | --- | --- |
| `--profile` | Yes | The provisioning profile. |
| `--category` | Yes | A Mac App Store category, such as `productivity`. An unknown name is rejected with the list of valid ones. |
| `--certificate` | No | Defaults to `3rd Party Mac Developer Application`. The matching `Installer` certificate is used for the package. |

The command signs the app bundle with `codesign`, builds an installer package
with `productbuild`, and signs that with `productsign`. The result is
`Name.pkg`. You need Xcode's command line tools and both certificates in your
keychain.
{{< /tab >}}

{{< tab name="Windows" >}}
```powershell
bellui release -os windows --app-id com.example.notes --app-version 1.2.0 --app-build 3 `
    --developer "CN=Example Ltd, O=Example Ltd, L=City, S=State, C=GB" `
    --certificate cert.pfx --password ****
```

| Flag | Needed | Notes |
| --- | --- | --- |
| `--app-id`, `--app-version`, `--app-build` | Yes | |
| `--developer` | Yes | Your publisher identity from the Microsoft Partner Center, in the `CN=…, O=…` form. |
| `--certificate`, `--password` | Yes | The signing certificate file and its password. |

The command builds `Name.appx` with `makeappx.exe` and signs it with
`signtool.exe`. Both come with the Windows SDK; the command finds them in the
SDK's usual install folder, `C:\Program Files (x86)\Windows Kits`.
{{< /tab >}}

{{< tab name="Android" >}}
```sh
bellui release -os android --app-id com.example.notes --app-version 1.2.0 --app-build 3 \
    --keystore release.keystore --key-name upload
```

| Flag | Needed | Notes |
| --- | --- | --- |
| `--app-id`, `--app-version`, `--app-build` | Yes | |
| `--keystore` | Yes | Your `.keystore` file. |
| `--key-name` | Yes | The alias of the signing key. |
| `--keystore-pass`, `--key-pass` | No | Passwords. If you leave them out you are asked for them. |

A release builds an Android App Bundle, `Name.aab`, which is what Google Play
expects; `bellui package` builds an `.apk` instead. You need `ANDROID_HOME`
set, and `bundletool` on your `PATH` to make the bundle. The command aligns it
with `zipalign` from the SDK's build tools and signs it with `jarsigner` from
the JDK.
{{< /tab >}}

{{< tab name="iOS" >}}
```sh
bellui release -os ios --app-id com.example.notes --app-version 1.2.0 --app-build 3 \
    --profile "Notes App Store"
```

| Flag | Needed | Notes |
| --- | --- | --- |
| `--app-id`, `--app-version`, `--app-build` | Yes | |
| `--profile` | Yes | The App Store provisioning profile. |
| `--certificate` | No | Defaults to `Apple Distribution`. |

The command signs the app with `codesign`, using your team ID from the
certificate, and zips it into `Name.ipa` for upload.
{{< /tab >}}

{{< /tabs >}}

On Linux, BSD and the web, `bellui release` produces the same files as
`bellui package --release`; there is no store packaging step.

Like `bellui package`, a successful release adds one to the build number in
your [manifest](../metadata/).
