---
title: Statistical and financial
description: Box plots, heatmaps and candlestick charts in x/charts.
weight: 2
---

These series summarise many values at once: the spread of a measurement, a
grid of intensities, or a day of trading.

## Boxplot

A [`BoxplotSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#BoxplotSeries)
draws one box and whisker per category. Use it to compare the spread of a
measurement between groups, such as latency per region.

{{< screenshot name="charts-boxplot" alt="A box plot of request latency for five regions; each box has a white median line and whiskers to the minimum and maximum" >}}

{{< example file="extensions/charts/series/boxplot/main.go" region="boxplot" >}}

Each
[`BoxplotData`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#BoxplotData)
is a five-number summary: `Min`, `Q1`, `Median`, `Q3` and `Max`. The series
does not compute it from raw samples, so you work out the quartiles yourself.
The box runs from `Q1` to `Q3` with a line at the median, and the whiskers
reach `Min` and `Max`.

`BoxWidth` is the box's width as a fraction of the space per category
(default 0.5). `Color` sets the box colour. The median line is always white
and the whiskers always light grey.

## Heatmap

A [`HeatmapSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#HeatmapSeries)
colours each cell of a grid by its value. Use it for values that depend on two
categories at once, such as activity by weekday and hour.

{{< screenshot name="charts-heatmap" alt="A heatmap of commits by weekday and hour; the weekday rows are darkest around midday" >}}

{{< example file="extensions/charts/series/heatmap/main.go" region="heatmap" >}}

The main fields:

- `Rows` and `Cols` give the grid size, and `Data` holds the values row by row:
  the value for row `r` and column `c` is `Data[r*Cols+c]`. `Value(row, col)`
  reads one back.
- `Min` and `Max` set the range that the colours span. When both are zero they
  are taken from the data.
- `ColorStops` is the colour scale: each
  [`HeatmapColorStop`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#HeatmapColorStop)
  maps a position from 0 to 1 in that range to a colour, and colours in
  between are blended. Without stops the scale runs from deep blue through
  yellow to deep red.

The heatmap fills the whole plot area and draws no axes. `XLabels` and
`YLabels` name the columns and rows in the tooltip only; they are not drawn
next to the grid. Put the meaning of rows and columns in the title, as the
example does, or in labels of your own around the chart.

## Candlestick

A [`CandlestickSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#CandlestickSeries)
draws the open, high, low and close of each period. It is the standard view of
a price over time.

{{< screenshot name="charts-candlestick" alt="A candlestick chart of thirty days of made-up prices, green for rising days and red for falling days, with a moving-average line" >}}

{{< example file="extensions/charts/series/candlestick/main.go" region="candlestick" >}}

Each
[`CandlestickData`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#CandlestickData)
has `Open`, `High`, `Low` and `Close`, one per category. The body runs from
open to close and the wick from low to high. A candle whose close is at or
above its open uses `UpColor` (green by default); any other uses `DownColor`
(red by default). `BarWidth` is the body width as a fraction of the space per
category (default 0.6).

A candlestick series shares the axes with other cartesian series, so an
indicator such as a moving average is just a `LineSeries` in the same option.
To add candles as they arrive, call `AppendCandles` on the series or
`AppendCandlestick` on the chart (see [Live data](../#live-data)). To show
volume in a pane of its own under the prices, use a
[chart group](../#interaction).
