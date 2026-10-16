---
title: Prerequisites
description: What to install before building a Bellui app on macOS, Windows or Linux.
weight: 1
---

Every platform needs:

- **Go 1.25 or later.** Check with `go version`.
- **A C compiler**, because Bellui uses cgo for the native window system and
  GPU code. `CGO_ENABLED` must not be `0`.

You do not need to install a GPU library. The prebuilt WebGPU
([wgpu-native](https://github.com/gfx-rs/wgpu-native)) libraries ship inside
per-platform Go modules (`github.com/belluikit/webgpu/libs-darwin`, `libs-linux`,
`libs-windows` and so on) and are downloaded by `go build`
like any other dependency.

{{< tabs >}}

{{< tab name="macOS" >}}
Install the Xcode command line tools, which include clang:

```sh
xcode-select --install
```

Rendering uses Metal, which every supported Mac has.
{{< /tab >}}

{{< tab name="Windows" >}}
Install a MinGW-w64 C toolchain and put its `bin` directory on your `PATH`.
Either of these works:

- [MSYS2](https://www.msys2.org/), then `pacman -S mingw-w64-ucrt-x86_64-gcc`
  in the UCRT64 shell;
- [llvm-mingw](https://github.com/mstorsjo/llvm-mingw), which is also the
  toolchain to use on Windows on Arm.

Check that Go finds it: `go env CC` should name the compiler, and `gcc --version`
(or `clang --version`) should run.

Rendering uses Direct3D 12.

{{< callout type="info" >}}
The first build on Windows compiles a lot of C and can take several minutes.
Later builds reuse Go's build cache and are fast.
{{< /callout >}}
{{< /tab >}}

{{< tab name="Linux" >}}
Install a C compiler and the X11, Wayland and OpenGL development headers. On
Debian and Ubuntu:

```sh
sudo apt-get install gcc libgl1-mesa-dev libwayland-dev libx11-dev \
    libxkbcommon-dev xorg-dev
```

To run apps with the default WebGPU renderer you also need a Vulkan driver. A
desktop with a working GPU driver normally has one; on a minimal system or a
VM, install the Vulkan loader and Mesa's drivers (which include the `lavapipe`
software renderer):

```sh
sudo apt-get install libvulkan1 mesa-vulkan-drivers
```

Other distributions need the same pieces under their own package names: a C
compiler, the X11 (including Xcursor, Xrandr, Xinerama, Xi and Xxf86vm),
Wayland, xkbcommon and OpenGL development packages, and a Vulkan driver.
{{< /tab >}}

{{< /tabs >}}

### Optional packages

- **[x/mmedia](/extensions/mmedia/)** plays audio through ALSA on Linux and needs
  `libasound2-dev` (Debian/Ubuntu) to build.
- **The `bellui` command** packages, installs and releases apps, including for
  Android and iOS. Install it with:

  ```sh
  go install github.com/belluikit/bellui/cmd/bellui@latest
  ```

  See [Packaging and distribution](/docs/packaging/) for what each platform needs.

When the tools are in place, [write your first app](../first-app/).
