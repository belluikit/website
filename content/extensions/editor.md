---
title: Code editor
description: A code editor widget with syntax highlighting, a gutter for line numbers, breakpoints and diagnostics, folding, find, completion and hover cards.
weight: 3
---

`x/editor` provides
[`editor.Editor`](https://pkg.go.dev/github.com/belluikit/bellui/x/editor#Editor),
a text editor for source code. It adds to a multi-line text field what a
code editor needs: coloured syntax, a gutter beside the text, folding, a find
bar, and hooks for completion, hover documentation and error markers. Each
piece is optional and pluggable, so you can have line numbers without
highlighting, or highlighting without a gutter.

{{< screenshot name="code-editor" alt="The gallery's code editor page" >}}

## Install

```sh
go get github.com/belluikit/bellui/x/editor@latest
```

## A first editor

{{< example "extensions/editor/basic/main.go" >}}

- `editor.NewEditor` creates the widget with sensible defaults: automatic
  indentation, four-space indents, and at least 20 visible rows.
- `Highlighter` colours the text; here, the Go highlighter.
- `Gutter` draws to the left of the text. It stays in place when the text
  scrolls sideways.
- `SetText` replaces the whole document. It clears the undo history and does
  not call `OnChanged`, which reports only the user's edits. To change the text
  the way a user edit would, with undo, use `ApplyEdit` or `ApplyEdits`.
- `OnSave` and `OnRun` are called for Cmd+S and Cmd+R on macOS. On other
  systems these, and the find shortcuts below, use the Super (Windows) key in
  place of Cmd.

Keyboard shortcuts only reach the editor while it has focus: the user clicks
into it, or you call `Focus()` once the window is showing.

{{< callout type="warning" >}}
The editor scrolls itself. Do not put it inside a `container.Scroll`: the outer
scroll takes the pointer events, so clicks never reach the editor. Place it
directly in a layout that gives it space, such as the centre of a border
layout. The editor logs a warning if it finds itself inside a scrolling
container.
{{< /callout >}}

`SetReadOnly(true)` stops editing but keeps the syntax colours.
`MinRows` sets the minimum height in rows, and `AutoGrow` makes the editor grow
with its content, between `MinAutoGrowRows` and `MaxAutoGrowRows`.

## Syntax highlighting

Four highlighters ship with the module, each in its own package under
`highlight/`:

| Package | Language |
| --- | --- |
| [`highlight/golang`](https://pkg.go.dev/github.com/belluikit/bellui/x/editor/highlight/golang) | Go, including raw strings and block comments that span lines |
| [`highlight/starlark`](https://pkg.go.dev/github.com/belluikit/bellui/x/editor/highlight/starlark) | Starlark and Python. `AddBuiltins` registers extra built-in names, and `starlark.IndentProvider()` indents after a line ending in a colon and dedents `else:`, `elif`, `except` and `finally`. |
| [`highlight/proto3`](https://pkg.go.dev/github.com/belluikit/bellui/x/editor/highlight/proto3) | Protocol Buffers |
| [`highlight/pxf`](https://pkg.go.dev/github.com/belluikit/bellui/x/editor/highlight/pxf) | PXF, a text format for Protocol Buffer messages |

For another language, implement
[`Highlighter`](https://pkg.go.dev/github.com/belluikit/bellui/x/editor#Highlighter),
or wrap a function in `HighlighterFunc`. The editor calls it once per line. It
returns styled runs that together cover the line, and a state number that is
passed to the next line, so a construct like a block comment can continue
across lines. State 0 means "nothing open".

{{< example file="extensions/editor/highlighter/main.go" region="highlighter" >}}

A run names a colour with `ColorName` rather than a fixed colour, so it follows
the theme. The `ColorNameSyntax...` constants (keyword, string, comment,
number, builtin, function, operator, type, annotation) are looked up in your
app's theme first. A theme that does not define them gets built-in light and
dark defaults. A run can also set `ForegroundColor` directly, a
`BackgroundColor`, and bold, italic or underlined text.

## The gutter

A [`GutterProvider`](https://pkg.go.dev/github.com/belluikit/bellui/x/editor#GutterProvider)
draws one column of the gutter. Combine several with `NewCompositeGutter`:

{{< example file="extensions/editor/gutter/main.go" region="gutter" >}}

Whatever order you pass them in, the composite arranges the columns
conventionally: markers (breakpoints, diagnostics) on the outside, line
numbers next, and fold controls beside the text. Use
`NewCompositeGutterInOrder` to keep your own order.

| Provider | What it shows |
| --- | --- |
| `NewLineNumberGutter()` | Line numbers. `MinDigits` reserves width (default 3). |
| `NewBreakpointGutter(fn)` | A dot on lines the user clicks. `fn` receives the set of marked lines, numbered from 1. |
| `ed.FoldGutter()` | Arrows that fold and unfold regions. |
| `ed.DiagnosticGutter()` | A marker for each diagnostic, shaped by severity. |

The fold and diagnostic gutters belong to the editor, so you get them from it
rather than creating new ones. If you change the gutter after the editor is
shown, call `ApplyGutter`.

### Folding

The editor does not work out fold regions by itself. You tell it which line
ranges can fold, typically from your own parser or language server:

{{< example file="extensions/editor/gutter/main.go" region="folds" >}}

A region with `Collapsed: true` starts folded; `Summary` replaces the default
`...` placeholder. From code, `Fold`, `Unfold`, `FoldAll` and `UnfoldAll`
change the state, and `AutoFold` folds a range that was not registered.

### Diagnostics

A [`Diagnostic`](https://pkg.go.dev/github.com/belluikit/bellui/x/editor#Diagnostic)
marks a range of text with a severity and a message. The editor draws a wavy
underline under the range, a marker in the diagnostic gutter, and shows the
message when the pointer rests on the range.

{{< example file="extensions/editor/gutter/main.go" region="diagnostics" >}}

Severities are `SeverityError`, `SeverityWarning`, `SeverityInfo` and
`SeverityHint`. `SetDiagnostics` replaces the whole set each time, so you can
call it whenever your checker finishes; pass `nil` to clear them. Call it on the
UI thread.

## Find

Cmd+F opens a find bar above the text (the Super key on Linux and Windows).
Typing highlights every match; Enter or Cmd+G moves to the next and
Shift+Cmd+G to the previous; Escape closes it. A switch makes the search case
sensitive. `ShowFind`, `FindNext`, `FindPrev` and `HideFind` do the same from
code, for example from your own menu.

## Completion, hover and signatures

The editor asks providers you install for language help. Each is an interface
with one method, called on the UI thread, so it should answer quickly.

{{< example file="extensions/editor/assist/main.go" region="assist" >}}

- **Completion** shows a list of candidates while the user types, or on
  Ctrl+Space. A
  [`CompletionProvider`](https://pkg.go.dev/github.com/belluikit/bellui/x/editor#CompletionProvider)
  returns `CompletionItem`s with a label, the text to insert, an optional
  detail and documentation. `WordListCompletionProvider` offers a fixed list.
- **Hover** shows a card when the pointer rests on a word, or on Cmd+K
  Cmd+I (Ctrl+K Ctrl+I on Linux and Windows). A
  [`HoverProvider`](https://pkg.go.dev/github.com/belluikit/bellui/x/editor#HoverProvider)
  returns markdown or plain text. `WordHoverProvider` looks words up in a map.
- **Signature help** shows a function's parameters when the user types `(`
  or `,` in a call, or presses Cmd+Shift+Space (Ctrl+Shift+Space on Linux and
  Windows). Install a `SignatureHelpProvider`; `FuncSignatureProvider` reads
  signatures from a map.
- **Go to definition**: a `LinkProvider` marks words that can be followed.
  Cmd-click on macOS, or Ctrl-click elsewhere, or F12, calls the function you
  register with `OnRequestNavigation`.

The interfaces are shaped like the Language Server Protocol's requests, so a
bridge to a language server can implement them, but they carry no LSP types.
`ed.KeyMap` changes the completion, signature and go-to-definition shortcuts;
start from `editor.DefaultKeyMap()`.

## Decorations

Decorations draw over the text without changing it.

A **row decorator** draws across a whole line, for example the current
line in a debugger. A **span decorator** draws over a character range: an
underline (solid, wavy or dotted), a background, a box or a strike-through.

{{< example file="extensions/editor/assist/main.go" region="decorate" >}}

Note that `CurrentLineDecorator.Row` counts lines from 1, while cursor
positions and `TextRange` count from 0. After changing a decorator's fields,
call `RefreshDecorators`. After replacing the `Decorators` or
`SpanDecorators` slice, call `ApplyDecorators`.

Two related features protect parts of the text:

- `SetRedactions` hides ranges behind an opaque bar, for secrets such as API
  keys. The caret cannot enter a hidden range, and edits to it are refused.
- `SetProtectedRanges` makes ranges read-only while the rest stays editable.
  The ranges move with the text as the user edits above them.

## Large documents

The gutter and the decorations are drawn only for the visible lines, so long
files stay responsive. For documents of several megabytes, set `DeferTextSync`: the editor
then reports the text to `OnChanged` at natural pauses (a new line, a deletion,
moving to another line, losing focus) instead of on every keystroke.

## Licence

`x/editor` is BSD 3-Clause licensed, like Bellui.
