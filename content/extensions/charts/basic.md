---
title: Line, area, bar, scatter and pie
description: The everyday chart types in x/charts, with their main fields, examples and screenshots.
weight: 1
---

These are the series most charts are made of. Line, bar and scatter series
are drawn on the chart's axes and can be mixed in one chart. A pie needs no
axes. Every example on this page is a complete program; the parts shown are
the `Option` passed to `chart.New`.

## Line

A [`LineSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#LineSeries)
joins one value per category with straight segments. Use it for values that
change over an ordered axis, such as time.

{{< screenshot name="charts-line" alt="A line chart with three series: a solid line with circle markers, a dashed target line, and a stepped line with a gap in April" >}}

{{< example file="extensions/charts/series/line/main.go" region="line" >}}

The main fields:

- `Data` holds one value per category. `math.NaN()` marks a missing value: the
  line breaks there. With `ConnectNulls` set, the gap is bridged with a dashed
  segment instead.
- `Symbol` and `SymbolSize` put a marker on each point (default size 6).
  Markers are drawn only while the line has 200 points or fewer.
- `Step` draws a staircase instead of slopes: `StepStart`, `StepMiddle` or
  `StepEnd` says where in each interval the line jumps.
- `LineStyle` sets `Color`, `Width` (default 2) and `Dash`, a pattern of dash
  and gap lengths in pixels: `{6, 4}` draws 6 px dashes with 4 px gaps.
- `Stack` stacks the line on top of other series with the same stack name.
- `YAxisIndex` picks the Y axis when the chart has several.

A line with more points than the plot is pixels wide is downsampled before it
is drawn, with the LTTB algorithm, which keeps the shape of the curve.
`AppendValues`, or `AppendData` on the chart, adds points without rebuilding
the option (see [Live data](../#live-data)).

{{< callout type="warning" >}}
In v1.0.0 `Smooth` has no effect: the renderer never reads it, and lines are
always drawn with straight segments. `SymbolTriangle` and `SymbolDiamond` are
drawn as squares, the same as `SymbolRect`.
{{< /callout >}}

## Area

An area chart is a line series with an
[`AreaStyle`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#AreaStyle).
The fill shows volume under the line, such as load or traffic over a day.

{{< screenshot name="charts-area" alt="An area chart of CPU load over 24 hours; the fill under the line has a stepped top edge" >}}

{{< example file="extensions/charts/series/area/main.go" region="area" >}}

`AreaStyle.Opacity` sets how transparent the fill is (default 0.3), and
`AreaStyle.Color` overrides the line's colour for the fill.

{{< callout type="warning" >}}
The area fill does not render correctly in v1.0.0. It is drawn as one
rectangle per segment, from the higher of the segment's two points down to the
bottom of the plot. The top of the fill is therefore a staircase that pokes
above the line wherever the line slopes, as the screenshot shows. The fill
always reaches the bottom of the plot rather than zero, so stacked areas
overlap each other instead of sitting on top of one another. Stacked lines
without a fill are drawn correctly.
{{< /callout >}}

## Bar

A [`BarSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#BarSeries)
draws one bar per category. Use it to compare amounts between categories.

{{< screenshot name="charts-bar" alt="A bar chart of quarterly profit with one negative quarter, and an average line drawn over the bars" >}}

{{< example file="extensions/charts/series/bar/main.go" region="bar" >}}

The main fields:

- `Data` holds one value per category. Negative values grow down from zero.
- `BarWidth` is the bar's width as a fraction of the space per category
  (default 0.6).
- `CornerRadius` rounds the corners. `BorderColor` and `BorderWidth` add an
  outline. `Color` overrides the palette colour.
- `Stack` stacks bars with the same stack name, as in the
  [stacked example](../#series-types).

`AppendValues` adds bars in place, like the line series.

{{< callout type="warning" >}}
v1.0.0 has no grouped bars. Two bar series in one chart without a common
`Stack` are drawn at the same position, so the later one covers the earlier
one. `BarGap` is declared but has no effect.
{{< /callout >}}

## Scatter

A [`ScatterSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#ScatterSeries)
places a marker at each x, y pair. Use it to show how two measures relate.
Both axes are usually value axes.

{{< screenshot name="charts-scatter" alt="A scatter chart with two clusters of points, one drawn with circles and one with squares" >}}

{{< example file="extensions/charts/series/scatter/main.go" region="scatter" >}}

`Data` stores the pairs flat: `[x0, y0, x1, y1, ...]`. `Symbol` chooses the
marker shape, `SymbolSize` its size (default 8), and `Color` overrides the
palette. `AppendValues` takes further pairs in the same flat form.

## Pie

A [`PieSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#PieSeries)
shows how a whole divides into parts. It ignores the axes and is drawn in the
middle of the plot area.

{{< screenshot name="charts-pie" alt="A pie chart of storage by file type with five slices, small gaps between them, and a label next to each" >}}

{{< example file="extensions/charts/series/pie/main.go" region="pie" >}}

The main fields:

- `Data` is a list of
  [`PieDataItem`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#PieDataItem):
  a `Name`, a `Value` and an optional `Color` for that slice.
- `Radius` is a fraction of the available space (default 0.75).
- `InnerRadius` is a fraction of `Radius`. Any value above 0 cuts a hole in the
  middle and makes a doughnut.
- `StartAngle` is where the first slice starts, in degrees, with 0 at the top.
  Slices go clockwise.
- `PadAngle` leaves a gap of that many degrees between slices.
- `ShowLabels` writes each slice's name outside it, in the slice's colour.

Slices take their colours from a fixed built-in palette, not from
`Option.Color`. Set `PieDataItem.Color` to choose them.
