---
title: Custom widgets
description: Build your own widgets from canvas primitives, or extend an existing widget with new behaviour.
weight: 9
---

When the standard widgets do not do what you need, you can make your own.
There are two ways:

- **Build a widget** from scratch: a struct that embeds
  `widget.BaseWidget`, and a renderer that draws it with
  [canvas primitives](/docs/concepts/canvas-objects/). You control every
  pixel and every event.
- **Extend a widget**: embed an existing widget, such as `widget.Entry`, and
  add or replace methods. You keep everything the widget already does.

Before writing a widget, check whether a [container](/docs/layout/) of
existing widgets does the job. A widget is worth it when you need a new look,
new input handling, or a reusable unit with its own state and API.

{{< cards >}}
  {{< card link="building/" title="Building a widget" subtitle="BaseWidget, the renderer, MinSize, Layout and Refresh, with a complete rating widget." >}}
  {{< card link="extending/" title="Extending a widget" subtitle="Add taps to an icon, filter an entry's input." >}}
{{< /cards >}}
