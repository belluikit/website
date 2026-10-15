---
title: Renderers
description: WebGPU is the default renderer on every platform; OpenGL is available with the gl build tag.
weight: 3
---

Bellui draws with **WebGPU** by default. On each platform that means the
native graphics API: Metal on Apple platforms, Vulkan on Linux and Android,
Direct3D 12 on Windows, and the browser's own WebGPU on the web. One set of
shaders drives all of them.

The older **OpenGL** renderer (OpenGL, OpenGL ES and WebGL) is still supported.
Choose it at build time with the `gl` tag:

| Build | Renderer | Command |
| --- | --- | --- |
| default | WebGPU | `go build .` |
| fallback | OpenGL / GLES / WebGL | `go build -tags gl .` |

There is no runtime switch; the tag decides what is compiled in.

## Where WebGPU runs

| Platform | Graphics API | Status |
| --- | --- | --- |
| macOS | Metal | Verified |
| Linux (Wayland and X11) | Vulkan | Verified on Wayland; X11 builds and renders off-screen |
| Windows | Direct3D 12 | Verified |
| Web (WebAssembly) | Browser WebGPU | Verified in Chromium |
| Android | Vulkan | Verified in the emulator |
| iOS Simulator | Metal | Verified |
| iOS device | Metal | Verified on iPad |

The OpenGL renderer supports the same desktop, mobile and web platforms.

## What WebGPU needs

**On the desktop**, the build links the native wgpu-native library through cgo.
The library comes as per-platform Go modules (`github.com/belluikit/webgpu/libs-*`),
so there is
nothing to install beyond the [prerequisites](../prerequisites/). Two things
to know:

- Stray C linker or compiler flags in your environment can break the link. If
  the link step fails, try again with them cleared:

  ```sh
  CGO_LDFLAGS='' CGO_CFLAGS='' go build .
  ```

- Running the app needs a working GPU driver: Metal, Vulkan (a real GPU or the
  `lavapipe` software driver) or Direct3D 12. Building does not.

**On the web**, the browser provides WebGPU, so there is no native library and
no cgo. The page must be served from a secure context (HTTPS or `localhost`),
and the browser must support WebGPU. See [Web](/docs/platforms/web/).

**On Android and iOS**, the `bellui` command packages the right native library
for the target. See [Mobile](/docs/platforms/mobile/).

## When to use the OpenGL renderer

Use `-tags gl` when the target has no usable WebGPU: an older GPU or driver
without Vulkan or Direct3D 12 support, a browser without WebGPU, or a build
environment where linking wgpu-native is not possible. The OpenGL build links
no wgpu-native at all.
