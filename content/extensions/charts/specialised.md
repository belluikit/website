---
title: Specialised charts
description: Radar, gauge, funnel, parallel, pictorial bar, effect scatter and custom series in x/charts.
weight: 4
---

These series each answer one kind of question: how several measures compare,
where a single value stands, how many drop out at each step. Radar, gauge,
funnel, parallel and custom series lay themselves out in the plot area.
Pictorial and effect scatter series are drawn on the chart's axes.

## Radar

A [`RadarSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#RadarSeries)
puts several measures on spokes around a centre and joins each item's values
into a polygon. Use it to compare a few items across the same measures.

{{< screenshot name="charts-radar" alt="A radar chart comparing two laptops on battery, screen, weight, speed and price; each is a filled polygon" >}}

{{< example file="extensions/charts/series/radar/main.go" region="radar" >}}

`Indicators` lists the spokes, each a
[`RadarIndicator`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#RadarIndicator)
with a `Name` and a `Max`. With `Max` at 0 the spoke's maximum is the largest
value on it. `Data` has one row per item and one value per indicator; a row
with fewer values than indicators is skipped. `ShowArea` fills the polygons,
with `Opacity` as the fill's opacity (default 0.3).

The fill is drawn as a fan of lines from the centre, so it looks striped rather
than solid. Polygons take their colours from the built-in palette, and
`DataNames` is not used for drawing in v1.0.0.

## Gauge

A [`GaugeSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#GaugeSeries)
shows one value on a dial, with coloured ranges for normal, warning and
critical.

{{< screenshot name="charts-gauge" alt="A gauge reading 312 out of 500, with the gap at the bottom; the value arc runs from the bottom right over the top, followed by a yellow and a red range" >}}

{{< example file="extensions/charts/series/gauge/main.go" region="gauge" >}}

The main fields:

- `Value` is the reading. `Min` and `Max` set the scale; when `Max` is not
  above `Min`, the scale runs to 100.
- `Sectors` colour ranges of the scale, each a
  [`GaugeSector`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#GaugeSector)
  with `Min`, `Max` and `Color`.
- `ShowValue` writes the value, with the series name under it, in the middle.
- `Radius` is a fraction of the available space (default 0.85).
- `StartAngle` and `EndAngle`, in degrees with 0 at the top, set where the
  scale starts and ends.

{{< callout type="warning" >}}
The gauge does not draw as its documentation describes in v1.0.0. Values
always advance counter-clockwise from `StartAngle`. With the default angles
(225 and -45) the scale starts at the bottom left and runs through the bottom,
the right and the top, leaving the gap on the left side. The example sets
`StartAngle: 135, EndAngle: -135`, which puts the gap at the bottom, with the
minimum at the bottom right. The value arc is drawn over the sectors, so a
sector below the current value is hidden. The value text is a fixed light
colour, which is faint in the light theme.
{{< /callout >}}

## Funnel

A [`FunnelSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#FunnelSeries)
shows how many remain at each step of a process, such as a checkout.

{{< screenshot name="charts-funnel" alt="A funnel of five checkout stages, each a centred bar narrower than the one above, with its name inside" >}}

{{< example file="extensions/charts/series/funnel/main.go" region="funnel" >}}

Each [`FunnelData`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#FunnelData)
stage has a `Name`, a `Value` and an optional `Color`. Stages are centred bars
whose width is proportional to their value. All stages are the same height,
unless `DynamicHeight` is set, which makes each stage's height proportional to
its value.

Stages are drawn top to bottom in the order you give them. `Sort` is declared
but has no effect in v1.0.0, so put the stages in order yourself.

## Parallel coordinates

A [`ParallelSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#ParallelSeries)
draws one vertical axis per measure and one line per item across them. Use it
to spot patterns across many measures, such as heavier cars being slower and
more expensive.

{{< screenshot name="charts-parallel" alt="Parallel coordinates for six cars across power, weight, acceleration, fuel use and price, one coloured line per car" >}}

{{< example file="extensions/charts/series/parallel/main.go" region="parallel" >}}

`Axes` lists the
[`ParallelAxis`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#ParallelAxis)es,
each with a `Name`, `Min` and `Max`. When both `Min` and `Max` are 0, the range
comes from the data. `Data` holds one row per line and one value per axis.
`LineWidth` defaults to 1.5. Each line takes the next colour from the built-in
palette.

## Pictorial bar

A [`PictorialSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#PictorialSeries)
is a bar chart whose bars are stacks of symbols. It suits counts of things,
where a column of items reads better than a solid bar.

{{< screenshot name="charts-pictorial" alt="A pictorial bar chart of trees planted per district, each bar a column of circles" >}}

{{< example file="extensions/charts/series/pictorial/main.go" region="pictorial" >}}

`SymbolSize` is the size of one symbol (default 12), and the number of symbols
follows the bar's height. With `Symbol: series.SymbolCircle` the bar is a
column of circles; any other symbol draws a column of rounded bars instead.
`Color` overrides the palette colour.

## Effect scatter

An [`EffectScatterSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#EffectScatterSeries)
is a scatter series whose points have a ring around them. Use it to make a few
points stand out, here alerts over ordinary readings.

{{< screenshot name="charts-effectscatter" alt="A scatter of sensor readings with three alert points above them, each with a faint ring around it" >}}

{{< example file="extensions/charts/series/effectscatter/main.go" region="effectscatter" >}}

`Data` stores x, y pairs flat, like a scatter series. `SymbolSize` sets the
point size (default 8), and `RippleScale` the size of the ring relative to the
point (default 2.5). The ring grows during the chart's entrance animation and
then stays still: in v1.0.0 it does not pulse.

## Custom

A [`CustomSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#CustomSeries)
hands drawing to your own function. Use it for a chart the other types do not
cover, while keeping the chart's title, background and theme.

{{< screenshot name="charts-custom" alt="A custom series drawing four labelled progress bars for a release checklist" >}}

{{< example file="extensions/charts/series/custom/main.go" region="custom" >}}

`RenderFunc` receives a
[`RenderContext`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#RenderContext)
and returns the canvas objects to draw. The fields you will use most:

- `Viewport` is the plot area in pixels. Position everything inside it.
- `Color` is the series' colour from the palette.
- `Animate` and `AnimProgress` (0 to 1) tell you when the entrance animation is
  running, so you can animate too.

A custom series is laid out like a pie or a radar: the chart draws no axes for
it, and `Coord` maps data values to the same pixel values.
