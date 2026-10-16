---
title: Web views
description: Embed web content in a Bellui layout, run JavaScript in it, and exchange messages with it.
weight: 3
---

The [`webview`](https://pkg.go.dev/github.com/belluikit/bellui/webview)
package puts a web engine inside your layout. You can load a URL or HTML,
run JavaScript in the page, and pass JSON messages both ways through a
bridge.

The package is optional. The web engine is linked only into apps that import
it.

## Which engine runs where

| Platform | Engine | Shipped by |
| --- | --- | --- |
| macOS | WKWebView | the OS |
| iOS | WKWebView | the OS |
| Android | `android.webkit.WebView` | the OS |
| Windows | WebView2 | a separate Microsoft runtime (see below) |
| Web (WebAssembly) | an `<iframe>` over the canvas | the browser |
| Linux | CEF (Chromium Embedded Framework) | **your app** |

Everywhere except Linux the engine belongs to the operating system, so its
security updates come with the OS. Linux has no system engine an app can rely
on, so a Linux app bundles Chromium itself. [Shipping on Linux](#shipping-on-linux)
explains what that means for you.

The constant
[`webview.Supported`](https://pkg.go.dev/github.com/belluikit/bellui/webview#Supported)
is false on platforms without an engine: the BSDs, and Linux builds with
`CGO_ENABLED=0`. There the code still compiles and runs. The view takes up its
space and draws nothing.

## Show a page

{{< example file="advanced/webview/main.go" region="host" >}}

[`NewHost`](https://pkg.go.dev/github.com/belluikit/bellui/webview#NewHost)
returns two values. The
[`Host`](https://pkg.go.dev/github.com/belluikit/bellui/webview#Host) is the
widget you put in your layout. The
[`WebView`](https://pkg.go.dev/github.com/belluikit/bellui/webview#WebView)
is what you drive. The engine is created when the host is first laid out in a
window, and the URL loads then.

A web page has no natural size, so the host's minimum size is zero. Give it a
layout that stretches it, such as the center of a `container.NewBorder`, or
set a floor with `SetMinSize`.

To change pages later, call `LoadURL`, `Reload`, `StopLoading`, `GoBack` or
`GoForward`. `LoadHTML(html, baseURL)` shows HTML you pass in, but only once
the view is attached; before that it does nothing.

## Overlays and the two hosting models

On macOS, iOS, Android, Windows and the web, the page is a separate OS view
placed over Bellui's window. Bellui cannot draw on top of it. Pop-ups, menus,
dialogs and tooltips that overlap the page are hidden behind it, and
screenshots taken with `Canvas.Capture` show a blank area there. Lay out your
UI so nothing needs to cover the page. See [Native views](../native-views/)
for the full list of limits.

On Linux, Chromium renders off-screen and Bellui draws the result like any
other widget. Overlays appear above the page, `Capture` includes it, and
clipping and opacity apply. The page also takes part in Tab navigation. While
the page has focus, Tab moves within the page, and Ctrl+Tab or Ctrl+Shift+Tab
moves back to Bellui's widgets.

## Talk to the page

The bridge carries JSON between Go and the page's JavaScript. It is **off**
until you enable it, and it only accepts messages from origins you trust:

{{< example file="advanced/webview/main.go" region="bridge" >}}

Call `EnableBridge` before the view is attached, which in practice means
before the window is shown. Later calls are ignored and logged. The policy is
checked for every incoming message, including those from frames inside the
page:

- [`TrustOrigins(...)`](https://pkg.go.dev/github.com/belluikit/bellui/webview#TrustOrigins)
  accepts exact origins, such as `"https://example.com"`. Use `OriginOpaque`
  for `data:` URLs and `about:blank`, and `"file://"` for files loaded from
  disk.
- `TrustAllOrigins()` accepts anything. Use it only when the view can never
  show content you did not write.
- A nil policy accepts nothing.

The page gets a `window.bellui` object:

| In the page | In Go |
| --- | --- |
| `await window.bellui.call(name, arg)` | answered by the function registered with `view.Handle(name, ...)` |
| `window.bellui.post(value)` | delivered to `view.OnMessage` |
| `window.bellui.handle(name, fn)` | called by `view.Call(name, arg, onResult)` |
| `window.bellui.onMessage = fn` | receives `view.Post(value)` |

Here is the page side of the example:

{{< example file="advanced/webview/main.go" region="page" >}}

A Go handler's return value resolves the page's promise. An error rejects it
with the error's message. A call to a name with no handler is rejected too,
so the page never waits forever.

Going the other way, `view.Call` runs a function the page registered:

{{< example file="advanced/webview/main.go" region="call" >}}

Its callback runs exactly once: with the result, or with an error such as
`ErrNotAttached` or `ErrBridgeDisabled`.

## Run JavaScript

| Method | What it does |
| --- | --- |
| `Eval(js)` | runs a script and ignores the result |
| `EvalResult(js, onResult)` | runs a script and reports the value of its last expression as JSON |
| `EvalFunc(body, args)` | runs `body` as a function whose parameters are the keys of `args` |
| `EvalFuncResult(body, args, onResult)` | the same, reporting the return value |

`Eval` takes source code. If you build it by joining strings that came from
outside your app, a quote in that text becomes code. `EvalFunc` passes values
as JSON data instead, so prefer it for anything that is not a constant:

{{< example file="advanced/webview/main.go" region="evalfunc" >}}

## Control navigation

{{< example file="advanced/webview/main.go" region="navigate" >}}

`OnNavigate` is asked before each navigation and blocks it by returning false.
The [`NavigationRequest`](https://pkg.go.dev/github.com/belluikit/bellui/webview#NavigationRequest)
tells you the URL, what started it (a link, a form, history, a reload), whether
it is the main frame, and whether the page asked for a new window. There is no
second window, so an allowed new-window request opens in the same view.

`OnLoadChanged` reports each change of the main frame's state: loading,
complete or failed. A main-frame navigation you blocked, or one you stopped
with `StopLoading`, is reported as failed with `ErrNavigationCancelled`.

The bridge policy decides who may *talk* to your app. `OnNavigate` decides what
may be *shown*. A view that can reach the open web needs both.

## Threading

Call every method on the UI thread. Every callback, from messages to load
events, also runs on the UI thread. A bridge handler holds up the engine while
it runs, so hand slow work to a goroutine and send the answer later with
`Post`.

## Platform differences

- **Web:** the bridge, `Eval` and `Call` reach same-origin content only.
  `OnNavigate` sees only navigations the app starts, and `CanGoBack` is
  always false.
- **Web and Android:** the bridge is added when the page loads, so the page's
  first lines may run before `window.bellui` exists.
- **Android and Windows:** a script that throws reports a null result, not an
  error. For the same reason, a `view.Call` to a page that has no bridge
  waits until the view is destroyed instead of failing at once.
- **Windows and Linux:** `LoadHTML` ignores its base URL. Relative links in the
  HTML do not resolve, and its origin is always `OriginOpaque`.
- **Windows:** every navigation the user starts, including a form submission,
  is reported as a link click.

## Shipping on Windows

WebView2 needs the Microsoft Edge WebView2 Runtime. Windows 11 and Microsoft
Edge include it, but Windows Server and LTSC images may not. Without it the
view stays blank and attaching fails with
[`ErrRuntimeMissing`](https://pkg.go.dev/github.com/belluikit/bellui/webview#ErrRuntimeMissing).
If your app must work everywhere, ship Microsoft's Evergreen Bootstrapper with
it.

## Shipping on Linux

On Linux your app carries Chromium. `bellui package` places the pinned CEF
files next to your binary, and the app finds them there at run time. The
`BELLUI_CEF_DIR` environment variable points it somewhere else. A binary
started without the files leaves the view blank and logs `ErrRuntimeMissing`.

{{< callout type="warning" >}}
**You are responsible for Chromium security updates on Linux.** The Chromium in
a Linux app stays at the version it was packaged with. When a Chromium security
update lands, rebuild and re-package your app against a current Bellui release
and ship it to your users. An app packaged months ago still carries the
Chromium of that time. The Bellui project's update policy for the pinned
version is in
[CEF-MAINTENANCE.md](https://github.com/belluikit/bellui/blob/v1.0.0/docs/CEF-MAINTENANCE.md).
{{< /callout >}}

Browsing data (cookies, local storage and the HTTP cache) is kept between runs,
in a `webview-cef` folder in the app's storage, like the system engines on
other platforms. Two functions change that. Both must be called before the
first web view is attached, and both do nothing on other platforms:

- [`SetCEFProfileDir(dir)`](https://pkg.go.dev/github.com/belluikit/bellui/webview#SetCEFProfileDir)
  moves the data to `dir`.
- [`SetCEFProfileInMemory()`](https://pkg.go.dev/github.com/belluikit/bellui/webview#SetCEFProfileInMemory)
  keeps it in memory only, like a private browsing window.

Chromium locks its data folder. If a second copy of your app starts while the
first is running, the second one's web views stay blank. Give each copy its
own folder, keep the data in memory, or allow only one running copy.
