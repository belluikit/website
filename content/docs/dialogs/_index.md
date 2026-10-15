---
title: Dialogs
description: Show messages, ask questions, collect input, and let people pick files, folders and colours.
weight: 6
---

The [`dialog`](https://pkg.go.dev/github.com/belluikit/bellui/dialog) package
has ready-made dialogs for the common cases: telling the user something, asking
them to confirm, collecting a few values, and choosing a file, a folder or a
colour. Every dialog is shown over a parent window that you pass in.

Most dialogs come in two forms:

- `ShowXxx(…)` builds the dialog and shows it at once.
- `NewXxx(…)` builds it and returns it, so you can change it before you call
  `Show()`.

Dialogs do not block. `Show` returns straight away, and the result arrives
later in the callback you passed. Write the code that depends on the answer
inside that callback.

{{< cards >}}
  {{< card link="messages/" title="Messages and confirmation" subtitle="Information, error and yes/no dialogs, and the options every dialog shares." >}}
  {{< card link="custom/" title="Custom and form dialogs" subtitle="Your own content and buttons, and forms with validation." >}}
  {{< card link="files/" title="File and folder dialogs" subtitle="Open, save, folder and multi-file pickers, native or drawn by Bellui." >}}
  {{< card link="color/" title="Colour picker" subtitle="Preset swatches, recent colours, and full RGB/HSL editing." >}}
  {{< card link="about/" title="About dialog" subtitle="Your app's icon, name, version and links, from its metadata." >}}
{{< /cards >}}
