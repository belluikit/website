---
title: About
description: Who makes Bellui, where it came from, and how it is licensed.
---

Bellui is a Go UI toolkit for building apps that run on desktop, mobile and the
web from one codebase. It is developed by TrendVidia and the Bellui developers in the open, in the
[belluikit](https://github.com/belluikit) GitHub organization.

## Where it came from

Bellui began as a fork of [Fyne](https://fyne.io), the Go toolkit
created by the Fyne.io developers. It has since diverged a long way and is not
intended to merge back. Among the larger changes:

- a Material Design 3 theme system and component set;
- a WebGPU renderer, now the default, alongside the original OpenGL one;
- a reworked package layout, with fundamental types in `geo`, `event` and `core`;
- extension modules for charts, data tables, a code editor, graphs and media.

We are grateful to the Fyne developers for the foundation Bellui is built on.

## License

Bellui is released under the
[BSD 3-Clause license](https://github.com/belluikit/bellui/blob/main/LICENSE).
Because it is a derivative of Fyne, its license file keeps Fyne's copyright
notice and conditions in full, as the BSD license requires, and Fyne's
contributors are listed in
[AUTHORS](https://github.com/belluikit/bellui/blob/main/AUTHORS). Third-party
code, fonts and icons are listed in
[THIRD_PARTY_NOTICES.md](https://github.com/belluikit/bellui/blob/main/THIRD_PARTY_NOTICES.md).

Several supporting libraries are forks too, each under its original license:

| Library | Forked from |
| --- | --- |
| [webgpu](https://github.com/belluikit/webgpu) | oliverbestmann/webgpu |
| [glfw](https://github.com/belluikit/glfw) | go-gl/glfw |
| [glfw-js](https://github.com/belluikit/glfw-js) | fyne-io/glfw-js |
| [gl-js](https://github.com/belluikit/gl-js) | fyne-io/gl-js |
| [systray](https://github.com/belluikit/systray) | fyne-io/systray |
| [fsnotify](https://github.com/belluikit/fsnotify) | fsnotify/fsnotify |
| [typesetting](https://github.com/belluikit/typesetting) | go-text/typesetting |

This website's content and examples are BSD 3-Clause licensed as well.
