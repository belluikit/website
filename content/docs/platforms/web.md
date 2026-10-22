---
title: Web
description: Build your app to WebAssembly, serve it, and choose between WebGPU and WebGL in the browser.
weight: 3
---

A Bellui app can run in a web browser. Go compiles it to WebAssembly (a binary
format browsers can run), and the app draws into a canvas element on the page.
Rendering uses the browser's WebGPU by default, or WebGL with the `gl` build
tag.

## The quickest way: bellui serve

In your app's folder, run:

```sh
bellui serve
```

This builds the web version, writes it to a `wasm` folder, and serves that
folder at `http://localhost:8080`. Open the address in a browser that supports
WebGPU. `--http-port` picks another port, and `--debug` builds a debug version;
by default `serve` builds in release mode.

## Packaging for a web server

```sh
bellui package -os web
```

The `wasm` folder it creates is a complete static site:

| File | Purpose |
| --- | --- |
| `index.html` | The page. It shows your icon and a loading message, checks that the browser supports WebGPU (or WebGL, for a build with `-tags gl`) and says so if it does not, loads the app, and hides the loading screen when the app creates its canvas. It also credits the gopher artwork on the loading screen, which its licence requires. |
| `<Name>.wasm` | Your compiled app. |
| `wasm_exec.js` | Go's WebAssembly loader, copied from your Go installation. |
| `icon.png`, `light.css`, `dark.css`, spinners | Page icon and loading-screen styling, in the user's light or dark preference. |
| `bellui-sw.js` | A service worker used for notification action buttons. |

Copy the folder to any static web host. Add `--release` for a smaller,
optimised build.

## Building it yourself

You can also build with the Go tool and write your own page:

```sh
GOOS=js GOARCH=wasm CGO_ENABLED=0 go build -o app.wasm .              # WebGPU
GOOS=js GOARCH=wasm CGO_ENABLED=0 go build -tags gl -o app.wasm .     # WebGL
```

The web build uses no C code. If your environment sets `CGO_ENABLED=1`, the
build fails in `runtime/cgo`, which is why the commands above turn it off.

Your page then needs Go's `wasm_exec.js` (in `$(go env GOROOT)/lib/wasm/`) and
a script that starts the app. One thing is specific to Bellui: before you
call `go.run`, store the instantiated module in `window.wasm`. The WebGPU
renderer reads the WebAssembly memory through it.

```html
<script src="wasm_exec.js"></script>
<script>
  const go = new Go();
  WebAssembly.instantiateStreaming(fetch("app.wasm"), go.importObject)
    .then((result) => {
      window.wasm = result;      // required by the WebGPU renderer
      go.run(result.instance);
    });
</script>
```

The `index.html` that `bellui package -os web` writes does this already.

## WebGPU or WebGL

| | WebGPU (default) | WebGL (`-tags gl`) |
| --- | --- | --- |
| Browser support | Browsers with WebGPU enabled | Browsers with WebGL |
| Page requirement | A secure context: HTTPS, or `http://localhost` | None beyond WebGL |
| Status | Verified in Chromium | The long-standing renderer |

WebGPU is only available to pages in a **secure context**. `bellui serve` on
`localhost` counts as secure; a page served over plain HTTP from another
address does not, and WebGPU is not available to it. Serve production builds
over HTTPS, or use the WebGL build if you cannot.

There is no automatic fallback at run time: the renderer is chosen when you
build. To reach browsers without WebGPU, publish the WebGL build.

## How the browser changes things

{{< example file="platforms/web/main.go" region="web" >}}

- `bellui.CurrentDevice().IsBrowser()` is `true`.
- **One canvas, several windows.** The first window fills the page. Further
  windows open as windows inside it that the user can drag, but they are not
  separate browser windows.
- **Files.** Opening a file uses the browser's file picker. Saving offers the
  file as a download when you close the writer. See
  [File and folder dialogs](/docs/dialogs/files/#native-pickers-or-belluis-dialog).
- **Storage.** [Preferences](/docs/storage/preferences/) are kept in the
  browser's `localStorage`, and [app storage](/docs/storage/app-storage/) in
  IndexedDB.
- **No desktop extras.** There is no system tray and there are no global
  shortcuts.

## The complete example

{{< example "platforms/web/main.go" >}}
