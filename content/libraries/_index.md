---
title: Libraries
description: The Go libraries Bellui is built on. Each is a separate module you can use on its own.
---

Bellui depends on a set of libraries maintained in the
[belluikit](https://github.com/belluikit) organization. Several are forks of
established projects, carrying fixes and platform support that Bellui needs;
each keeps its original license. Any of them can be used on its own.

How you get them differs. The graphics, windowing and desktop libraries are
dependencies of Bellui itself and arrive with `go get`. goffmpeg and its
prebuilt libraries come with [x/mmedia](/extensions/mmedia/). gocef-libs is
fetched by the `bellui` command when it packages a Linux app with a web view.

Uploader and Clair are separate tools, not libraries Bellui depends on.
They are described under [Tools](/tools/).

## Graphics and windowing

| Library | What it does | License |
| --- | --- | --- |
| [webgpu](https://github.com/belluikit/webgpu) | Go bindings for wgpu-native (WebGPU) on desktop and mobile, and for the browser's WebGPU on the web. Fork of oliverbestmann/webgpu. | Apache-2.0 |
| [glfw](https://github.com/belluikit/glfw) | Go bindings for GLFW 3.4, with Wayland drag-and-drop, touch and input-method extensions. Fork of go-gl/glfw. | BSD-3-Clause |
| [glfw-js](https://github.com/belluikit/glfw-js) | GLFW-style windowing for Go on the desktop and in the browser (WebAssembly). Fork of fyne-io/glfw-js. | MIT |
| [gl-js](https://github.com/belluikit/gl-js) | Go bindings for OpenGL, OpenGL ES and WebGL. Fork of fyne-io/gl-js. | BSD-3-Clause |
| [typesetting](https://github.com/belluikit/typesetting) | Text shaping and layout in pure Go, a port of HarfBuzz. Fork of go-text/typesetting. | Unlicense or BSD-3-Clause |

## Desktop integration

| Library | What it does | License |
| --- | --- | --- |
| [systray](https://github.com/belluikit/systray) | An icon and menu in the system tray on macOS, Windows and Linux. Fork of fyne-io/systray. | Apache-2.0 |
| [fsnotify](https://github.com/belluikit/fsnotify) | Cross-platform file system change notifications. Fork of fsnotify/fsnotify. | BSD-3-Clause |
| [gocef](https://github.com/belluikit/gocef) | Go bindings for the Chromium Embedded Framework, loaded at run time. Powers the web view on Linux. | BSD-3-Clause |
| [gocef-libs](https://github.com/belluikit/gocef-libs) | Pinned, checksummed CEF runtime payloads for gocef. | BSD-3-Clause |

## Media

| Library | What it does | License |
| --- | --- | --- |
| [goffmpeg](https://github.com/belluikit/goffmpeg) | Go bindings for FFmpeg (libavformat, libavcodec) with prebuilt libraries for each platform. Used by [x/mmedia](/extensions/mmedia/). | BSD-3-Clause |
| [goffmpeg-libs](https://github.com/belluikit/goffmpeg-libs) | The prebuilt FFmpeg libraries for goffmpeg, one Go module per platform. | LGPL-2.1 |
