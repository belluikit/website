---
title: Maps
description: Tile maps, points and routes at longitude and latitude, and choropleths from GeoJSON in x/charts.
weight: 6
---

Four series work with geographic coordinates. A `MapSeries` draws a map from
image tiles. `GeoScatterSeries` and `GeoLinesSeries` draw points and routes on
top of it. A `ChoroplethSeries` colours regions read from GeoJSON, with or
without a tile map underneath.

On a chart with a map, drag to pan and scroll to zoom. Coordinates are always
longitude first, then latitude, in degrees.

## Map

A [`MapSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#MapSeries)
fills the plot area with a tiled web map.

{{< screenshot name="charts-map" alt="A tile map of western Europe from OpenStreetMap" >}}

<small>Map tiles in the screenshots on this page © [OpenStreetMap](https://www.openstreetmap.org/copyright) contributors.</small>

{{< example file="extensions/charts/series/map/main.go" region="map" >}}

`Center` is the longitude and latitude in the middle of the view, and `Zoom`
the usual web-map zoom level, from 1 (the whole world) to 18 (a street).
Fractional zoom levels work. `TileURL` is a template with `{z}`, `{x}` and
`{y}` placeholders; when it is empty, tiles come from OpenStreetMap.

Tiles are downloaded in the background over the network and appear as they
arrive. Until then the map shows a plain sea colour. The map draws an
OpenStreetMap attribution in its corner, in a light grey that is hard to see
on light tiles. If you use OpenStreetMap's tile servers, follow their
[tile usage policy](https://operations.osmfoundation.org/policies/tiles/); an
app with many users should use a tile provider of its own through `TileURL`.

## Geo scatter

A [`GeoScatterSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#GeoScatterSeries)
draws a dot at each place.

{{< screenshot name="charts-geoscatter" alt="Six red dots on a tile map of western Europe, at London, Paris, Amsterdam, Berlin, Madrid and Rome" >}}

{{< example file="extensions/charts/series/geoscatter/main.go" region="geoscatter" >}}

`Data` stores longitude, latitude pairs flat: `[lon0, lat0, lon1, lat1, ...]`.
`SymbolSize` sets the dot size (default 6) and `Color` its colour. Dots are
always circles; `Symbol` has no effect in v1.0.0.


When the chart has a `MapSeries`, every geographic series uses the map's
projection, so the points follow the map as the user pans and zooms. Set
`MapCenter` and `MapZoom` to the map's values anyway: a series uses them on
its own when there is no map.

## Geo lines

A [`GeoLinesSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#GeoLinesSeries)
draws a line from one place to another, with an arrowhead at the destination.
Use it for routes or flows between places.

{{< screenshot name="charts-geolines" alt="Four curved red arrows from Paris to London, Berlin, Madrid and Rome on a tile map" >}}

{{< example file="extensions/charts/series/geolines/main.go" region="geolines" >}}

Each [`GeoLine`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#GeoLine)
has `From` and `To` as longitude, latitude pairs. `Curved` bends the lines into
arcs, `Width` sets their thickness (default 2) and `Color` their colour.
`GeoLine.Value` is not used for drawing in v1.0.0.

## Choropleth

A [`ChoroplethSeries`](https://pkg.go.dev/github.com/belluikit/bellui/x/charts/series#ChoroplethSeries)
fills regions, such as countries or states, with a colour for their value.

{{< screenshot name="charts-choropleth" alt="Eleven US states with simplified outlines, shaded from pale to dark green by population density" >}}

{{< example file="extensions/charts/series/choropleth/main.go" region="choropleth" >}}

The regions come from a GeoJSON feature collection, either downloaded from
`GeoJSONURL` in the background or given inline in `GeoJSON`. The example builds
a small one in code from hand-simplified outlines, so it runs offline.

`RegionData` maps a region to its value. A feature is matched by its name,
read from the first of the properties `name`, `NAME`, `admin`, `ADMIN`,
`name_en`, `NAME_EN`, `iso_a3` or `ISO_A3` that it has, and otherwise by an
identifier such as `iso_a3`, `iso_a2` or `id`. Regions with no value are drawn
in translucent grey. `Min`, `Max` and `ColorStops` work as in a
[heatmap](../statistical/#heatmap); the default scale runs from pale to dark green.

Set `MapCenter` and `MapZoom` to frame the regions, or add a `MapSeries`
first to draw them over a tile map. Regions are filled without outlines. In
v1.0.0 the choropleth shows tooltips only for a short built-in list of large
countries, looked up by name.
