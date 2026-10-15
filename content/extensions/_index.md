---
title: Extensions
description: Optional modules for charts, data tables, a code editor, graphs and media.
weight: 3
---

The extensions are larger widgets that not every app needs: charts, a data
grid, a code editor, a graph editor and a media player. They live in the
Bellui repository under `x/`, but each is a separate Go module. An app pays
for one, in build time and binary size, only when it imports it.

Each extension is released together with the toolkit and carries the same
version number: `x/charts v1.0.0` goes with Bellui v1.0.0. Add one with
`go get`:

```sh
go get github.com/belluikit/bellui/x/charts@latest
```

{{< cards >}}
  {{< card link="charts" title="Charts" subtitle="Line, bar, pie, financial, hierarchical, 3D and map charts, with zoom, tooltips and live data." icon="chart-bar" >}}
  {{< card link="tables" title="Tables" subtitle="A data grid for large data sets: sorting, filtering, editing, pinned columns, grouping, sparklines and export." icon="table" >}}
  {{< card link="editor" title="Code editor" subtitle="Syntax highlighting, a gutter for line numbers, breakpoints and diagnostics, folding, find and completion." icon="code" >}}
  {{< card link="dag" title="Graphs" subtitle="A validated graph model and an interactive canvas for editing pipelines and workflows." icon="share" >}}
  {{< card link="mmedia" title="Media" subtitle="Video and audio playback, subtitles and cameras, decoded by FFmpeg." icon="film" >}}
{{< /cards >}}

All five are BSD 3-Clause licensed, like Bellui. The media module links FFmpeg,
which is licensed under the LGPL; see [its licensing notes](mmedia/#licensing)
before you distribute an app that uses it.
