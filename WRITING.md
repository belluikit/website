# Writing pages for bellui.com

## Accuracy comes first

- **Describe the code, not your memory of Fyne.** Bellui began as a fork of
  Fyne but its API has diverged: the fundamental types live in `geo`, `event`
  and `core`, `ShowAndRun` is `bellui.ShowAndRun(w)`, preferences are generic
  free functions, and much more. Read the bellui source at the version pinned
  in `examples/go.mod` before describing anything.
- **Never copy or paraphrase Fyne's documentation** (docs.fyne.io,
  developer.fyne.io). It has no open license. Write every page from the bellui
  source and its doc comments.
- **Every Go snippet is a compiled file** in `examples/`, included with the
  `example` shortcode (see the README). If a snippet cannot compile on its own,
  put it in a complete program and show only a region.
- If the source and a doc comment disagree, the source wins; note the
  disagreement in your PR so the doc comment gets fixed in bellui.
- Do not document something you could not confirm in the source. Leave it out.

## Examples

- One directory per example, `examples/<section>/<name>/main.go`, each a
  `package main` that builds on its own.
- Keep them small and runnable: an app that opens a window and shows the thing
  the page is about.
- Use `//docs:start <name>` / `//docs:end <name>` markers to show part of a file.
- Run `scripts/check-examples.sh` before committing. It must pass.

## Style

- Plain, direct English. Short sentences. One idea per sentence.
- Lead with what the reader can do, then how it works.
- Second person ("you"), present tense.
- Explain a term the first time it appears, or use the plain word instead.
- Link API names to pkg.go.dev when that helps:
  `[widget.Entry](https://pkg.go.dev/github.com/belluikit/bellui/widget#Entry)`.
- No marketing adjectives ("blazing", "powerful", "seamless").
- No internal issue numbers or project history. Pages are for app developers.

## Front matter

```yaml
---
title: Layouts
description: One sentence, shown in search results and link previews.
weight: 3          # order within the section
---
```

Each directory has an `_index.md` that introduces the section.
