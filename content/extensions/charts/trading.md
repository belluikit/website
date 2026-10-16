---
title: Trading charts
description: A price pane and a volume pane that zoom and hover together, moving averages, a crosshair with axis tags, and drawing tools such as trend lines and Fibonacci retracements.
weight: 7
---

This page builds the kind of chart a trading app shows: candlesticks with
moving averages, a volume pane underneath that zooms and hovers together with
the prices, a crosshair with value tags on the axes, and tools the user draws
with, such as trend lines, horizontal levels and Fibonacci retracements. It
also shows how to create drawings from code and how to save and restore them.

{{< screenshot name="charts-trading" alt="A candlestick chart of 120 trading days with an orange 10-day and a blue 30-day moving average, a volume pane below it, a toolbar of drawing tools above, and a crosshair at Jun 18 with value tags on both axes of both panes" >}}

The sections below show the program in parts; the whole program is at the
[end of the page](#the-whole-program).

## Price and volume panes

A [`chart.ChartGroup`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/chart#ChartGroup)
stacks several charts on top of each other and keeps them in step. Each
[`Pane`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/chart#Pane)
holds a chart and a relative `Height`: with heights 3 and 1, the price pane
takes three quarters of the space.

{{< example file="extensions/charts/trading/main.go" region="panes" >}}

The group does two things:

- **Zoom and pan.** When the user zooms or pans one pane, with the wheel, a drag
  or the zoom buttons, the group sets the same zoom window on the other panes.
  A pane only follows if it has a `DataZoom`. Give it one with `Enabled: false`,
  as the volume pane does: it then follows the group but has no zoom buttons
  and ignores the wheel.
- **Hover.** When the pointer moves over one pane, the other panes show their
  crosshair at the same X position.

The panes line up only if you make them line up. Give every pane the same
number of categories and the same `Grid.Left` and `Grid.Right` margins, so the
plots have the same width and a category sits at the same X in each pane.

`SetSyncEnabled(false)` switches the syncing off. The group syncs only zoom
changes the user makes: when you call `SetZoomWindow` from code, call it on
every pane.

### Moving averages

A moving average is an ordinary `LineSeries` on the same axes as the candles.
Leave the first values, where the average is not defined yet, as `NaN`: the
line starts where the data does.

{{< example file="extensions/charts/trading/main.go" region="ma" >}}

`Option.Color` gives the averages their colours. The candles use their own
`UpColor` and `DownColor` (green and red by default), so the palette entry at
their index only colours their legend swatch.

### Crosshair and axis tags

The tooltip settings in the example are the ones that suit a price chart:

- `Trigger: TooltipTriggerAxis` finds the data at the pointer's X position.
- `Crosshair: CrosshairBoth` draws a vertical and a horizontal line through it.
- `ShowAxisTags` puts a tag with the X and Y values where the lines meet the
  axes, and `AxisTagFormatter` writes those values. Here it turns the category
  index into a date.
- `CrosshairOnly` hides the text box while hovering. A click pins the box,
  filled in by `Formatter`; a second click, or 15 seconds (`PinDuration`),
  removes it.

{{< callout type="warning" >}}
In v1.0.0 an axis tooltip only finds line and scatter series. It never finds
candlestick or bar series, so a pane that holds only candles or bars shows no
crosshair at all, and the group cannot show one in it either. The example
works around this by adding a line series with a transparent colour through
the closing prices, and another through the volumes. An unnamed series stays
out of the legend.
{{< /callout >}}

### Zooming

With `DataZoom` enabled, the wheel zooms around the pointer, a drag pans once
you are zoomed in, and the **+**, **−** and **Reset** buttons sit at the top
right of the pane. [Zoom](../interactive/#zoom) has the details. Three things
matter for price charts in v1.0.0:

- The Y axis keeps the range of all the data. Zooming into a quiet week does
  not stretch it to fill the pane.
- Candles and bars keep the width they have when all the data is shown, so
  they look thin when you zoom in far.
- A category axis writes a label under every category. With 120 days that is
  unreadable, so the example passes a label for every tenth day and empty
  strings for the rest.

To add candles as they arrive, see [Live data](../#live-data). In a group,
append to every pane, including the hidden lines, so the panes keep the same
number of points.

## Drawing tools

The user draws on a chart after you switch a tool on with
[`SetDrawingTool`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/chart#Chart.SetDrawingTool).
There is no built-in toolbar: you make one from your own buttons.

{{< example file="extensions/charts/trading/main.go" region="tools" >}}

These are all the tools in v1.0.0, the
[`interact.DrawingToolMode`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/interact#DrawingToolMode)
values:

| Tool | How the user draws it | Stored as |
| --- | --- | --- |
| `DrawingToolHLine` | One click; a horizontal line across the plot at that price | `HLineDrawing` |
| `DrawingToolVLine` | One click; a vertical line at that point in time | `VLineDrawing` |
| `DrawingToolTrendLine` | Two clicks; a segment between them | `TrendLineDrawing` |
| `DrawingToolRay` | Two clicks; a line from the first through the second, to the edge of the plot | `RayDrawing` |
| `DrawingToolRectangle` | Two clicks on opposite corners; a shaded box | `RectangleDrawing` |
| `DrawingToolChannel` | Three clicks: two for a line, a third for the parallel line | `ChannelDrawing` |
| `DrawingToolFibRetracement` | Two clicks, at the high and the low; dashed lines at 0, 23.6, 38.2, 50, 61.8, 78.6 and 100 % with their prices | `FibRetracementDrawing` |
| `DrawingToolTextLabel` | One click, then type; Enter or another click finishes | `TextLabelDrawing` |
| `DrawingToolMeasure` | Drag; shows the change in X, in price and in percent while you drag | nothing: it disappears when you let go |
| `DrawingToolNone` | Switches drawing off | |

While a two- or three-click drawing is in progress, a preview follows the
pointer. A new drawing takes its colour from a fixed set of six, picked by how
many drawings the chart already has.

A tool stays active after a drawing is finished, so the user can draw several
in a row. The example switches back to selecting after each one.

### Selecting, moving and deleting

With no tool active, a click within 10 pixels of a drawing selects it: it is
drawn thicker, with square handles at its points. Dragging a selected drawing
moves it. Delete or Backspace removes it, and Escape deselects it, or cancels a
drawing that is still in progress. The chart takes keyboard focus when it is
clicked, so these keys work right after a click.

{{< callout type="warning" >}}
A click on a chart that is showing a tooltip pins the tooltip instead of
selecting a drawing. With an axis tooltip, which is showing almost everywhere
over the plot, the user cannot select drawings at all. The example has a
**Select** mode that sets `TooltipTriggerNone` while the user edits drawings,
and puts the crosshair back for the drawing tools. Clicks with a drawing tool
active are not affected.
{{< /callout >}}

### Events

{{< example file="extensions/charts/trading/main.go" region="events" >}}

Each handler receives an
[`interact.DrawingEvent`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/interact#DrawingEvent)
with the `DrawingID` and, for `EventDrawingCreated`, the `Drawing` itself:

- `EventDrawingCreated` fires when the user finishes a drawing. It arrives on
  the UI thread, so the handler can update widgets.
- `EventDrawingSelected` fires when the user clicks a drawing.
- `EventDrawingDeleted` fires when the user presses Delete and when you call
  `RemoveDrawing`.

Nothing fires when a drawing is moved, when you call `AddDrawing`, or when
`ClearDrawings` removes everything. To keep a saved copy current, save again
when the user is done, for example when the window closes.

## Drawings from code

[`AddDrawing`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/chart#Chart.AddDrawing)
puts a drawing on the chart without the mouse, for example to mark levels
your app computes. X is the category index, where 0 is the first candle, and Y
is the price.

{{< screenshot name="charts-drawings" alt="A candlestick chart with drawings added from code: a red resistance line, an orange trend line along the highs, a blue ray along the lows, a blue box around a trading range, red dashed Fibonacci levels with their percentages, and an orange text label reading Lower high" >}}

{{< example file="extensions/charts/drawings/main.go" region="add" >}}

Every drawing embeds a
[`DrawingBase`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/interact#DrawingBase)
with an `ID_`, a `Color_` and a `Label_`. A nil colour uses the default
overlay colour of the theme. The label is written next to the drawing's
first point. The rest of the fields depend on the type:

| Type | Fields |
| --- | --- |
| `HLineDrawing` | `Y` |
| `VLineDrawing` | `X` |
| `TrendLineDrawing` | `Start`, `End` |
| `RayDrawing` | `Start`, and `Direction`, a second point the ray passes through |
| `RectangleDrawing` | `Corner1`, `Corner2` |
| `ChannelDrawing` | `Line1Start`, `Line1End`, and `OffsetY`, the price distance to the parallel line |
| `FibRetracementDrawing` | `HighY`, `LowY`, and `X` |
| `TextLabelDrawing` | `Pos`, `Text` |

The other methods on the chart are `RemoveDrawing(id)`, `ClearDrawings()`,
`Drawings()`, `SelectDrawing(id)`, `SelectedDrawing()` and `DrawingTool()`.

Drawings are measured against the first Y axis. A text label always has a dark
background, in both themes.

## Saving and restoring drawings

The chart does not save drawings for you. Every drawing can describe itself,
though: `Type()` says which kind it is, and `DataPoints()` returns its points
in data coordinates. That is enough to write it out and read it back.

{{< example file="extensions/charts/drawings/main.go" region="save" >}}

To restore, make an empty drawing of the saved type, then hand it the points
with `SetDataPoints`, which reads them in the same order `DataPoints` wrote
them:

{{< example file="extensions/charts/drawings/main.go" region="restore" >}}

{{< example file="extensions/charts/drawings/main.go" region="buttons" >}}

Things to know when you save drawings:

- **Choose your own IDs.** The chart names the drawings the user makes
  `drawing-1`, `drawing-2` and so on, counting from 1 for each chart, and does
  not look at IDs already in use. A restored drawing called `drawing-1` would
  share its ID with the next drawing the user makes, and `RemoveDrawing` would
  then remove the wrong one.
- **Copy before you keep.** `Drawings()` returns the chart's own list, and a
  drag moves a drawing by changing it in place.
- **X is an index.** On a category axis, a saved X is a position in your data.
  If you later load the data starting from a different day, shift the X values
  by the same number of days.
- The text of a text label is not one of its points; save `Text` separately,
  as the example does.

## The whole program

{{% details title="charts/trading/main.go" closed="true" %}}
{{< example "extensions/charts/trading/main.go" >}}
{{% /details %}}
