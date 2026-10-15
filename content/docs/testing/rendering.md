---
title: Golden files
description: Check that a widget still renders the way it did, by comparing it with saved markup or a saved image.
weight: 2
---

A golden-file test renders part of your UI and compares the result with a
file saved earlier, called the **master**. If anything changes, the test
fails and saves what it rendered, so you can see what changed. The `test` package can compare two kinds of
output:

- **Markup**: a text description of the rendered object tree, with each
  object's type, position, size and text. It is easy to read in a diff.
- **Images**: the rendered pixels, drawn by Bellui's software renderer, so no
  GPU is needed.

## A UI to test

{{< example file="testing/golden/main.go" region="ui" >}}

## Comparing markup

{{< example file="testing/golden/golden_test.go" region="markup" >}}

[`test.AssertObjectRendersToMarkup`](https://pkg.go.dev/github.com/belluikit/bellui/test#AssertObjectRendersToMarkup)
renders the object at its minimum size and compares the markup with
`testdata/status_card.xml`, relative to the test's package folder. The master
for this example looks like this:

{{< example file="testing/golden/testdata/status_card.xml" lang="xml" >}}

## Comparing images

{{< example file="testing/golden/golden_test.go" region="image" >}}

[`test.AssertObjectRendersToImage`](https://pkg.go.dev/github.com/belluikit/bellui/test#AssertObjectRendersToImage)
does the same with a PNG master. The comparison allows small differences, so
minor anti-aliasing changes between machines do not fail the test.

## Creating and updating masters

The first time a test runs there is no master. The test fails, and writes what
it rendered to `testdata/failed/`. To create or refresh the masters instead,
run the tests with `UPDATE_TESTDATA` set:

```sh
UPDATE_TESTDATA=1 go test ./...
```

Each master is written to `testdata/` and the test passes. Look at the new
files before you commit them: a master is only useful if what it shows is
right.

When a comparison fails later, the actual output is written to
`testdata/failed/` under the same name, and the failure message gives its
path. Compare it with the master, then either fix the code or rerun with
`UPDATE_TESTDATA=1` to accept the change. Add `testdata/failed/` to your
`.gitignore`.

## Rendering whole windows

The `…Object…` functions size a fresh canvas to the object. To check a whole
window, including anything drawn outside one object's bounds, render its
canvas with [`test.AssertRendersToMarkup`](https://pkg.go.dev/github.com/belluikit/bellui/test#AssertRendersToMarkup)
or [`test.AssertRendersToImage`](https://pkg.go.dev/github.com/belluikit/bellui/test#AssertRendersToImage),
passing `w.Canvas()` of a test window.

{{< callout type="info" >}}
A canvas created with `test.NewCanvas` has no renderer, so its images come out
blank. The `AssertObjectRendersToImage` helper sets up a software renderer for
you; for your own canvas use `test.NewCanvasWithPainter`.
{{< /callout >}}

When the tests run with the `mobile` build tag, the helpers first look for a
master under `testdata/mobile/` and fall back to the shared one. This lets
widgets that look different on touch screens keep separate masters.
