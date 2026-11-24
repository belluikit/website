---
title: Charts
description: Line, bar, pie, financial, hierarchical, 3D and map charts, drawn with Bellui canvas objects and configured with one declarative Option.
weight: 1
---

`x/charts` draws charts inside a Bellui window. You describe a chart with one
[`chart.Option`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/chart#Option)
value: a title, axes, a list of series, and the interactions you want. The
[`chart.Chart`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/chart#Chart)
widget renders it with ordinary Bellui canvas objects, so there is no web view
and no JavaScript.

The model follows [Apache ECharts](https://echarts.apache.org/): if you know
ECharts' `option` object, the field names will look familiar.

{{< screenshot name="charts" alt="The gallery's chart page" >}}

## Install

```sh
go get github.com/belluikit/bellui/x/charts@v1.3.0
```

The module is split into packages. You import the ones you use:

| Package | What it holds |
| --- | --- |
| [`chart`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/chart) | The `Chart` widget, `Option`, overlays, multi-pane groups |
| [`series`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series) | One struct per series type |
| [`axis`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/axis) | Axis configuration and scales |
| [`component`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/component) | Legend, tooltip, toolbox and visual-map settings |
| [`interact`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/interact) | Zoom and brush settings, events |
| [`dataset`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/dataset) | Column-oriented data with filter and sort steps |
| [`builder`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/builder) | A fluent alternative to writing an `Option` by hand |

## A first chart

{{< example "extensions/charts/basic/main.go" >}}

`chart.New` returns a widget, so you place it like any other: here it fills the
window. The X axis is a category axis, one label per data point. The Y axis is
a value axis whose range is computed from the data. `Animation: true` plays a
short entrance animation the first time the chart is drawn.

To change the chart later, call
[`SetOption`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/chart#Chart.SetOption)
with a new `Option`. When both the old and the new option have `Animation`
set, the chart animates from the old state to the new one.
[`Option()`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/chart#Chart.Option)
returns the current one, so you can modify it and set it back.

## Series types

Every series is a struct in the `series` package that you add to
`Option.Series`. `x/charts` has 27 of them. The pages below describe
each one, with its main fields, a compiled example and a screenshot:

{{< cards >}}
  {{< card link="basic" title="Line, area, bar, scatter and pie" subtitle="The everyday charts, on category or value axes." icon="chart-bar" >}}
  {{< card link="statistical" title="Statistical and financial" subtitle="Box plots, heatmaps and candlesticks." icon="presentation-chart-line" >}}
  {{< card link="hierarchical" title="Hierarchies and flows" subtitle="Trees, treemaps, sunbursts, sankey diagrams and node-link graphs." icon="share" >}}
  {{< card link="specialised" title="Specialised charts" subtitle="Radar, gauge, funnel, parallel coordinates, pictorial bars, effect scatter and custom series." icon="adjustments" >}}
  {{< card link="3d" title="3D charts" subtitle="Bars, scatter points, surfaces and lines in a 3D scene you can orbit." icon="cube" >}}
  {{< card link="maps" title="Maps" subtitle="Tile maps with points and routes, and choropleths from GeoJSON." icon="globe" >}}
{{< /cards >}}

Series fall into three groups, and the group decides how the chart lays them
out:

- **Cartesian series** (line, bar, scatter, effect scatter, pictorial,
  candlestick and boxplot) are drawn on the chart's X and Y axes. Any number of
  them can share one chart, such as bars and a line. Bars with the same `Stack`
  name are stacked:

  {{< example file="extensions/charts/types/main.go" region="stacked" >}}

- **Self-laid-out series** (pie, gauge, radar, heatmap, funnel, treemap,
  sunburst, tree, sankey, graph, parallel, custom and the map series) ignore
  the axes and draw themselves in the whole plot area.
- **3D series** draw into a 3D scene that the user can orbit.

{{< callout type="info" >}}
Pie, radar, funnel, treemap, sunburst, sankey, graph and parallel series
colour their parts from a fixed built-in palette, not from `Option.Color`.
Labels a series writes on the plot's background take the theme's text colour,
except pie labels, which take their slice's colour.
Labels on a coloured tile, such as a funnel stage or a treemap rectangle, are
black or white, whichever stands out more.
{{< /callout >}}

## Axes

Axes are lists, `Option.XAxis` and `Option.YAxis`, of
[`axis.Config`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/axis#Config).
A config chooses the scale type and its range:

- `axis.TypeCategory` labels each data index with a string from `Data`.
- `axis.TypeValue` is a linear scale. Its range is computed from the data
  unless you set `Min` or `Max` (both are pointers, so nil means automatic).
- `axis.TypeLog` is a base-10 logarithmic scale.
- `axis.TypeTime` formats tick labels with `TimeFormat`, a Go time layout.
- `SplitCount` is the number of tick intervals you would like (default 5); the
  scale rounds to readable values.

With no axes configured, a chart that needs them gets a category X axis and a
value Y axis.

A chart has **one X axis**, the first entry of `XAxis`. It can have **several Y
axes**: the first is drawn on the left, the others on the right. A series picks
its Y axis with `YAxisIndex`:

{{< example file="extensions/charts/axes/main.go" region="axes" >}}

Each `axis.Config` also sets how its axis looks:

- `Name` is the axis title. An X axis writes it under its labels, a Y axis
  above its top.
- `Position` moves the X axis to the top (`axis.PositionTop`), or puts a Y
  axis on the left or right (`axis.PositionLeft`, `axis.PositionRight`).
- `Inverse` runs the axis backwards. Series, tooltips, zoom and pan follow it.
- `LineColor` colours the axis line and ticks, and `LabelColor` the labels and
  the name. Without them the axis takes the theme's colours.
- `Hide` hides parts of the axis: an
  [`axis.Parts`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/axis#Parts)
  with `Line`, `Ticks` and `Labels`. Its zero value hides nothing.
- `SplitLine` draws grid lines across the plot. `axis.SplitLineAuto`, the
  default, draws them for the first Y axis only. `SplitLineShow` and
  `SplitLineHide` turn them on or off for one axis. A value, log or time axis
  draws them at its ticks. A category axis draws them between its
  categories, so in a bar chart they separate the bars.
- `Offset` moves the axis that many pixels outward from the plot. Use it to
  put two Y axes side by side on one side; the chart makes room for them.

{{< example file="extensions/charts/axisparts/main.go" region="axisparts" >}}

The plot shrinks to leave room for axis names, for labels on the right or at
the top, and for offset axes.

{{< callout type="info" >}}
`axis.Config`'s `ShowLine`, `ShowTicks`, `ShowLabels` and `ShowSplitLine` are
deprecated and have no effect. Use `Hide` and `SplitLine` instead.
`Option.VisualMap` and `Option.Decal` are not applied by the chart either.
{{< /callout >}}

**Overlays** mark values on the plot in data coordinates: a reference line, a
shaded band between two values, or a text annotation at a point, optionally with
an arrow.

{{< example file="extensions/charts/axes/main.go" region="overlays" >}}

`Grid` sets the margins around the plot area, in pixels. The chart adds the
room its axes need inside these margins.

## Interaction

**Tooltips** follow the pointer. With no `Tooltip` set, the chart shows the
nearest data point. A
[`TooltipConfig`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/component#TooltipConfig)
changes that: `TooltipTriggerAxis` lists the line, bar, scatter and
candlestick series at the pointer's X position, `TooltipTriggerNone` turns
tooltips off, `Crosshair` chooses
vertical, horizontal, both or no crosshair lines, and `Formatter` replaces the
text. Clicking pins the tooltip.

**Zoom** is enabled by setting `DataZoom` with `Enabled: true`. The scroll wheel zooms the X range
around the pointer, + and − buttons appear on the chart, and once zoomed you
drag to pan. `MinSpan` and `MaxSpan` limit the visible fraction of the data.

{{< example file="extensions/charts/interact/main.go" region="zoom" >}}

**Events** arrive on the chart's `Events` bus. Each handler receives a value
you type-assert to the event's struct:

{{< example file="extensions/charts/interact/main.go" region="events" >}}

Other event types include `EventHover`, `EventLegendToggle`, `EventDrilldown`
and `EventKeyPress`; see
[`interact.EventType`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/interact#EventType).
`On` returns an ID that you pass to `Off` to unsubscribe. To move the zoom from
code, call `SetZoomWindow`;
[`interact.DataWindow`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/interact#DataWindow)
expresses the window as fractions from 0 to 1 of the full data range.

**Brush selection** lets the user drag a rectangle over the plot. When the drag
ends, `EventBrushSelect` reports the points inside it. While the brush is
enabled, dragging selects instead of panning.

{{< example file="extensions/charts/interact/main.go" region="brush" >}}

More interactions, each switched on in the `Option`:

- **Legend.** Clicking a legend entry hides or shows that series. The legend
  appears when you set `Legend` or when the chart has more than one series.
- **Drilldown.** Set `Drilldown.DataProvider` to a function that returns a
  child `Option` for a clicked point. The chart shows the child with a
  breadcrumb trail back to the parent.
- **Toolbox.** `Toolbox: &component.ToolboxConfig{SaveAsImage: true}` adds a
  button that opens a save dialog and writes a PNG of the chart.
  `MagicType: []string{"line", "bar"}` adds buttons that switch the series
  between line and bar. `DrawingTools` adds buttons for the drawing tools (see
  [Trading charts](trading/#drawing-tools)).
- **Multi-pane layouts.** [`chart.NewChartGroup`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/chart#NewChartGroup)
  stacks several charts, such as price above volume, and keeps their zoom, pan
  and crosshair in step.

Two pages go further, with examples and screenshots:

{{< cards >}}
  {{< card link="trading" title="Trading charts" subtitle="Price and volume panes, moving averages, a crosshair with axis tags, and drawing tools such as trend lines and Fibonacci retracements." icon="trending-up" >}}
  {{< card link="interactive" title="Overlays, zoom and drill-down" subtitle="Reference lines and annotations, zoom buttons and a slider, the toolbox, brush selection, drill-down, timelines and colours by value." icon="cursor-click" >}}
{{< /cards >}}

## Live data

[`AppendData`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/chart#Chart.AppendData)
adds values to a line, bar or scatter series in place, without rebuilding the
`Option`. It is safe to call from any goroutine; the chart schedules its own
redraw on the UI thread. `AppendCandlestick` does the same for candlesticks.

{{< example file="extensions/charts/stream/main.go" region="stream" >}}

`SetFollowLatest(true)` keeps the zoom window on the newest data as it arrives.
It needs `DataZoom` to be set. If the user pans away, following stops until you
call it again.

## Datasets

The `dataset` package stores tabular data by column, with float, string and
time columns. A
[`Pipeline`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/dataset#Pipeline)
runs filter and sort steps over it, each producing a new dataset. `ExtractXY`
pulls two float columns out for plotting:

{{< example file="extensions/charts/dataset/main.go" region="dataset" >}}

A dataset is not attached to the chart: you extract the slices you need and
put them in a series.

## The builder

The `builder` package writes the same `Option` as a chain of calls. Each
`Add...` call starts a series; `Done()` returns to the chart. `Widget()` builds
the chart, and `Build()` returns the `Option` instead.

{{< example file="extensions/charts/builder/main.go" region="builder" >}}

`builder.LineChart()`, `BarChart()`, `ScatterChart()` and `PieChart()` start
with a suitable tooltip already set.

## Colours and themes

Charts take their colours from the current Bellui theme and follow it between
light and dark: the background, text, axis and grid lines come from the theme.
Series colours come from a palette that starts with the theme's primary colour.
To choose the series colours yourself, set `Option.Color` (or call
`Palette` on the builder); series use the colours in order and wrap around.
A colour set directly on a series, such as `LineStyle.Color`, wins over the
palette, and the legend shows it too. The series types that take one colour
for the whole series implement
[`series.ColoredSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#ColoredSeries),
whose `SeriesColor` returns that colour, or nil for the palette colour. Line,
bar, scatter, effect scatter, box plot, pictorial, geo scatter, geo lines, 3D
scatter and 3D line series implement it.

## Animation

With `Animation: true`, the chart plays a 400 ms entrance animation the first
time it is drawn, and a 300 ms transition when `SetOption` replaces an option
that also had `Animation` set.

## Licence

`x/charts` is BSD 3-Clause licensed, like Bellui. Its configuration model
follows Apache ECharts' `option` object; the module itself is written in Go
and has no JavaScript dependency.
