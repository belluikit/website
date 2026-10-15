---
title: Scrolling
description: Show content that is larger than its space with container.NewScroll, NewVScroll and NewHScroll.
weight: 2
---

A scroll container shows part of a larger object and lets the user move
around it with the mouse wheel, the touchpad, the scroll bars or a finger.
Wrap the content in one of:

| Constructor | Scrolls | Minimum size |
| --- | --- | --- |
| [`container.NewVScroll`](https://pkg.go.dev/github.com/belluikit/bellui/container#NewVScroll) | Up and down | As wide as the content, but short |
| [`container.NewHScroll`](https://pkg.go.dev/github.com/belluikit/bellui/container#NewHScroll) | Left and right | As tall as the content, but narrow |
| [`container.NewScroll`](https://pkg.go.dev/github.com/belluikit/bellui/container#NewScroll) | Both ways | Small in both directions |

{{< example file="layout/scroll/main.go" region="scroll" >}}

## Why the minimum size matters

A scroll container's purpose is to have a **smaller** minimum size than its
content. That is what lets a long list fit in a small window. Two things
follow from it:

- In a VBox, a VScroll gets its own small minimum height, not the height of
  the content, so it shows only a sliver. Put the scroll where it is
  stretched instead: the middle of a
  [Border](../containers/#border), a split pane, a tab, or the window itself.
- If the scroll is still too small, give it a floor with
  `scroll.SetMinSize(size)`.

A VScroll keeps the content's minimum width, so the window cannot be made
narrower than the content. An HScroll does the same for height.

## Reacting to and controlling the position

The scroll position is the `Offset` field, a
[`geo.Position`](https://pkg.go.dev/github.com/belluikit/bellui/geo#Position):
how far the content is moved from its top-left corner.

{{< example file="layout/scroll/main.go" region="control" >}}

- `OnScrolled` is called when the user scrolls. Do not change `Offset` from
  inside it.
- `ScrollToTop()`, `ScrollToBottom()` and `ScrollToOffset(pos)` move the
  content from code.
- `AddScrollListener(fn)` registers any number of extra listeners and
  returns a function that removes yours.

{{< callout type="warning" >}}
`ScrollToTop` and `ScrollToOffset` move the content without calling
`OnScrolled` or the listeners. `ScrollToBottom` does call them. If you keep
something in step with the offset, update it yourself after `ScrollToTop`, as
the example does.
{{< /callout >}}

Two fields change how the container looks:

- `ReserveBarSpace` makes the content narrower (or shorter) by the width of
  the scroll bar, so the bar never covers it. The three constructors turn it
  on; set it to `false` to let the bar float over the content.
- `HideShadows = true` turns off the fading edge shown where content
  continues out of view.

The `Direction` field holds which way the container scrolls:
`container.ScrollBoth`, `ScrollHorizontalOnly`, `ScrollVerticalOnly` or
`ScrollNone`. The three constructors set it for you.

## Long lists

Everything inside a scroll container exists at once, including the parts out
of view. For long data, such as thousands of rows, use
[`widget.List`](https://pkg.go.dev/github.com/belluikit/bellui/widget#List) or
[`widget.Table`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Table)
instead. They scroll by themselves and reuse a pool of row (or cell) objects,
filling them in with your data as the user scrolls.

## The whole program

{{< example "layout/scroll/main.go" >}}
