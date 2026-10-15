---
title: Your first app
description: Create a module, write a window with a label and a button, and run it.
weight: 2
---

{{% steps %}}

### Create a module

```sh
mkdir hello && cd hello
go mod init example.com/hello
```

### Write the app

Save this as `main.go`:

{{< example "hello/main.go" >}}

Reading it from the top:

- `app.New()` creates the application. There is one per program.
- `a.NewWindow("Hello")` creates a window. Nothing is shown yet.
- `widget.NewLabel` and `widget.NewButton` create two widgets. The button's
  function runs on the UI thread each time it is tapped, so it can update the
  label directly.
- `container.NewVBox` stacks them vertically, and `w.SetContent` makes the
  stack the window's content.
- `bellui.ShowAndRun(w)` shows the window and runs the app's event loop. It
  returns when the app quits, for example when you close the window. It is a
  function in the `bellui` package that takes the window, not a method on it.

### Fetch the dependencies

```sh
go mod tidy
```

Run this after writing the code, not before. `go mod tidy` adds exactly the
modules your imports need, and on an empty module it has nothing to go on.

### Run it

```sh
go run .
```

A window opens with the label and the button. Tap the button and the label
changes.

{{% /steps %}}

<p align="center">
  <img src="/images/hello-light.png" alt="The hello app in the light theme" width="207" />
  <img src="/images/hello-dark.png" alt="The hello app in the dark theme" width="207" />
</p>

## Try the mobile simulation

The `mobile` build tag runs the same program in a phone-sized window with
mobile behaviour, without a device or emulator:

```sh
go run -tags mobile .
```

To run on a real phone, package the app with the `bellui` command; see
[Packaging and distribution](/docs/packaging/).

## Where to go next

- [Renderers](../rendering/): what WebGPU and the OpenGL fallback each need.
- [Concepts](/docs/concepts/): apps, windows, canvas objects and the UI thread.
- [The gallery](https://github.com/belluikit/gallery): every widget in one app.
  Clone it and run `go run .` to explore.
