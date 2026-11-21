---
title: Overlays, zoom and drill-down
description: Reference lines and annotations, zoom with buttons and a slider, the toolbox, brush selection, drill-down with a breadcrumb, timelines and value-based colours in x/charts.
weight: 8
---

These features sit on top of any chart's series. Overlays mark values on the
plot. Zoom, the toolbox, the brush and drill-down let the user explore the
data. A timeline steps through a sequence of charts. Each section has a
compiled example. The [overview](../#interaction) covers tooltips, the legend
and the event bus, and [Trading charts](../trading/) covers multi-pane groups
and drawing tools.

## Overlays

An overlay marks a value on the plot, in data coordinates, so it moves with
the data when the user zooms. There are three kinds:

- a [`ReferenceLine`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/chart#ReferenceLine)
  at one value, such as a target or a limit;
- a [`ReferenceBand`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/chart#ReferenceBand)
  shading the range between `From` and `To`;
- an [`Annotation`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/chart#Annotation),
  a text label above an X, Y point, with an optional arrow to it.

{{< screenshot name="charts-overlays" alt="A line chart of API latency over 48 hours with a dashed red SLO line at 250 ms, a blue Normal band from 80 to 160, an orange Maintenance band, a red vertical Deploy line, and a Cache miss storm label with an arrow pointing at a spike" >}}

{{< example file="extensions/charts/overlays/main.go" region="overlays" >}}

Pass the list as `Option.Overlays`. How they behave:

- `Axis` chooses the direction. `AxisHorizontal` marks a Y value across the
  whole width; `AxisVertical` marks an X value across the whole height. On a
  category axis, X is the category index.
- `YAxisIndex` picks the Y axis a horizontal line, band or annotation is
  measured against, when the chart has several.
- `Width` is the line width (default 1.5). `Dashed` dashes the line, 6 pixels
  on and 4 off unless `DashPattern` says otherwise.
- With no `Color`, lines and bands use the theme's overlay colours, and
  annotations the title colour.
- A line's label sits at the right end of a horizontal line, or at the top of
  a vertical one. A band's label sits in its top-left corner.
- An annotation whose point is outside the plot, for example after zooming,
  is not drawn.

Overlays need axes: on a chart with only pie, gauge or other self-laid-out
series, they are not drawn.

## Zoom

Set `DataZoom` with `Enabled: true` and the user can zoom along the X axis:

- the wheel zooms in and out around the pointer;
- once zoomed in, a drag pans;
- the **+**, **−** and **Reset** buttons in the top-right corner zoom around
  the middle, or show all the data again.

{{< screenshot name="charts-zoom" alt="A line chart zoomed to days 201 to 310 of a year, with the zoom buttons at the top right and a slider below whose highlighted window covers the same part of the year" >}}

{{< example file="extensions/charts/zoomslider/main.go" region="zoom" >}}

[`ZoomConfig`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/interact#ZoomConfig)
holds the limits. `MinSpan` and `MaxSpan` are the smallest and largest part of
the data the user can show, as fractions from 0 to 1. `ZoomSensitivity`
scales each wheel step.

The visible part is a
[`DataWindow`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/interact#DataWindow):
`Start` and `End`, also fractions of the full range. Read it with
`ZoomWindow()`, set it with `SetZoomWindow`, and listen for
`EventZoomChange` to hear about changes the user makes.

### A zoom slider

[`component.DataZoomSlider`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/component#DataZoomSlider)
is a separate widget: a bar with a highlighted window that the user drags by
its body or by either edge. It is not connected to a chart, so you connect it
both ways:

{{< example file="extensions/charts/zoomslider/main.go" region="slider" >}}

`SetZoomWindow` does not send `EventZoomChange`, so the two handlers cannot
call each other forever. The same goes for buttons that jump to a fixed
range:

{{< example file="extensions/charts/zoomslider/main.go" region="code" >}}

{{< callout type="info" >}}
Some things to know about zoom:

- Only the X axis zooms. The Y axis keeps the range of all the data, as the
  screenshot shows: the visible values stay between 1000 and 1400 while the
  axis runs from 800 to 2000.
- Bars, candlesticks, box plots and pictorial bars get wider as a zoom shows
  fewer of them.
- `SetZoomWindow` does nothing on a chart without `DataZoom`.
- The slider takes its colours from the theme.
{{< /callout >}}

`DataZoom` with `Enabled: false` still lets `SetZoomWindow` move the window,
but the user cannot zoom that chart and it has no buttons. That is how a pane
follows the others in a [chart group](../trading/#price-and-volume-panes).

## Toolbox

The toolbox adds small buttons at the top right of the chart, under the zoom
buttons if there are any.

{{< screenshot name="charts-toolbox" alt="A bar chart of signups per weekday with line and bar buttons at its top right, and a Save PNG button above the chart" >}}

{{< example file="extensions/charts/toolbox/main.go" region="toolbox" >}}

`MagicType` adds one button per chart type. `"line"` turns every bar series
into a line series, and `"bar"` turns every line series into bars. Name, data,
Y axis and stack are kept; other settings of the series, such as colours,
dashes and area fills, are not. Button labels are cut to four letters.

`SaveAsImage: true` adds a **Save** button that asks for a file name and
writes a PNG of the chart. The chart's `SaveAsImage` method does the same from
your own button:

{{< example file="extensions/charts/toolbox/main.go" region="save" >}}

The image holds the title, series, axes, legend, overlays and drawings, at
the screen's pixel density. It leaves out the toolbox and the zoom buttons, and
the tooltip and crosshair if they are showing.

`DrawingTools` adds a row of drawing-tool buttons above the plot; see
[Drawing tools](../trading/#drawing-tools).

## Brush selection

With `Brush` enabled, a drag over the plot draws a selection rectangle
instead of panning. When the user lets go, the chart sends
`EventBrushSelect` with the points inside the rectangle.

{{< screenshot name="charts-brush" alt="A scatter chart of height against weight with a blue selection rectangle over part of the points, and a label above reading 56 people selected, mean weight 75.5 kg" >}}

{{< example file="extensions/charts/brush/main.go" region="brush" >}}

Each
[`BrushSelectedPoint`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/interact#BrushSelectedPoint)
has the `SeriesIndex` and `DataIndex` of the point and its `Value`. The
rectangle disappears when the drag ends, and the chart does not mark the
selected points: to highlight them, change the option yourself.

The brush selects points of line and scatter series, and the bars and
candlesticks it covers. `BrushSelectEvent` also reports the rectangle in data
coordinates, as `DataMinX`, `DataMaxX`, `DataMinY` and `DataMaxY`.

## Drill-down

Drill-down replaces the chart with a more detailed one when the user clicks a
data point, and shows a breadcrumb trail to go back.

{{< screenshot name="charts-drilldown" alt="A bar chart of Europe's revenue by quarter, reached by clicking the Europe slice of a pie chart, with a breadcrumb reading Root > Europe and a back button under the title" >}}

Set `Drilldown.DataProvider` to a function that returns the child option for
a clicked point, or nil when that point has nothing below it:

{{< example file="extensions/charts/drilldown/main.go" region="provider" >}}

{{< example file="extensions/charts/drilldown/main.go" region="chart" >}}

The function receives the clicked
[`HitResult`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/interact#HitResult)
and the current path. It runs for every click on a data point, at every level:
a child option without its own `Drilldown` keeps the parent's. Use the path,
as the example does, to stop at the last level.

The breadcrumb appears under the title once the user is below the top level:

- **<** goes up one level;
- **Root** goes back to the top;
- every level between them can be clicked; the last one, where the user is,
  cannot.

Each level's name in the trail is the clicked point's label. Pie slices have
one, their `Name`. For other series the label is empty and the series name is
used instead, so give the series a name that makes sense in the trail.

From code, `DrillInto(label, option)` goes one level down, `DrillUp()` one
level up, and `DrillPath()` returns the names of the levels. Going down or up
resets the legend and the zoom.

`EventDrilldown` fires when the chart goes down a level, and `EventDrillUp`
when it goes up, whether by a click in the trail or a call to `DrillUp`. The
`DrilldownEvent` carries the new `Path`. For a click on a data point, `Hit` is
the point clicked; it is nil when you call `DrillInto` yourself, and on the
way up.

## Timeline

A timeline steps through a sequence of charts, one per year for example.
[`component.TimelineState`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/component#TimelineState)
keeps the current step and plays through the steps on a timer. It has no user
interface of its own and is not attached to a chart: it calls your function
with each step's index, and your function sets the chart.

{{< example file="extensions/charts/timeline/main.go" region="timeline" >}}

How it behaves:

- `Next`, `Prev` and `GoTo` change the step and call your function. They do
  not go past the first or the last step, except that with `Loop`, `Next`
  goes from the last step back to the first.
- `Play` advances one step every `Interval` (default 2 seconds). Without
  `Loop` it stops at the last step; with `Loop` it starts again at the first.
  `Stop` ends playback.
- Every step, and your function, runs on the UI thread, during playback too,
  so your function can update the chart and widgets directly. Call `Play`,
  `Stop` and the step methods from the UI thread as well.
- Stop playback before stepping by hand, as the example does.
- `TimelineConfig.AutoPlay` has no effect yet: call `Play` yourself.

With `Animation: true` in each option, `SetOption` animates from one step to
the next.

## Colours by value

`Option.VisualMap` takes a list of
[`component.VisualMapConfig`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/component#VisualMapConfig),
which describe a colour scale from `Min` to `Max`. The chart does not use
them yet: setting `Option.VisualMap`, or calling `VisualMap` on the builder,
changes nothing on screen.

The config's own methods do work. `MapColor` returns the colour for a value,
blending between the colours in `InRange.Colors`, and `MapSize` returns a size
between `InRange.SymbolSizeMin` and `SymbolSizeMax`. Values outside `Min` and
`Max` get the colour or size at the nearest end. Use them where a series
takes a colour per item, such as pie slices:

{{< screenshot name="charts-visualmap" alt="A pie chart of errors by service whose slices run from pale orange for few errors to deep orange for the most, billing" >}}

{{< example file="extensions/charts/visualmap/main.go" region="visualmap" >}}

A heatmap has a colour scale of its own, `ColorStops`; see
[Heatmap](../statistical/#heatmap).
