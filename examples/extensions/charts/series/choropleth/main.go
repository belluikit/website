// A choropleth: regions from GeoJSON, filled by value. The outlines here are
// simplified by hand and built in code, so the example needs no network.
package main

import (
	"encoding/json"
	"image/color"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/series"
)

// outlines are rough state borders as longitude, latitude pairs.
var outlines = map[string][][2]float64{
	"Montana":      {{-116.05, 49}, {-104.05, 49}, {-104.05, 45}, {-111.05, 45}, {-111.05, 44.5}, {-113, 44.5}, {-114.5, 45.5}, {-114.5, 46.6}, {-116.05, 47.9}},
	"Wyoming":      {{-111.05, 45}, {-104.05, 45}, {-104.05, 41}, {-111.05, 41}},
	"Colorado":     {{-109.05, 41}, {-102.05, 41}, {-102.05, 37}, {-109.05, 37}},
	"Utah":         {{-114.05, 42}, {-111.05, 42}, {-111.05, 41}, {-109.05, 41}, {-109.05, 37}, {-114.05, 37}},
	"Arizona":      {{-114.05, 37}, {-109.05, 37}, {-109.05, 31.33}, {-111.07, 31.33}, {-114.8, 32.5}, {-114.6, 35}, {-114.05, 36.2}},
	"New Mexico":   {{-109.05, 37}, {-103, 37}, {-103, 32}, {-106.6, 32}, {-106.6, 31.8}, {-108.2, 31.8}, {-108.2, 31.33}, {-109.05, 31.33}},
	"North Dakota": {{-104.05, 49}, {-97.2, 49}, {-96.6, 45.94}, {-104.05, 45.94}},
	"South Dakota": {{-104.05, 45.94}, {-96.6, 45.94}, {-96.45, 42.5}, {-98.5, 43}, {-104.05, 43}},
	"Nebraska":     {{-104.05, 43}, {-98.5, 43}, {-96.45, 42.5}, {-95.3, 40}, {-102.05, 40}, {-102.05, 41}, {-104.05, 41}},
	"Kansas":       {{-102.05, 40}, {-95.3, 40}, {-94.6, 39.1}, {-94.6, 37}, {-102.05, 37}},
	"Oklahoma":     {{-103, 37}, {-94.6, 37}, {-94.43, 33.6}, {-97, 33.8}, {-100, 34.56}, {-100, 36.5}, {-103, 36.5}},
}

// featureCollection turns the outlines into a GeoJSON document. Each feature
// gets a "name" property, which is what RegionData is keyed by.
func featureCollection() []byte {
	type feature struct {
		Type       string            `json:"type"`
		Properties map[string]string `json:"properties"`
		Geometry   struct {
			Type        string         `json:"type"`
			Coordinates [][][2]float64 `json:"coordinates"`
		} `json:"geometry"`
	}
	var fs []feature
	for name, ring := range outlines {
		f := feature{Type: "Feature", Properties: map[string]string{"name": name}}
		f.Geometry.Type = "Polygon"
		f.Geometry.Coordinates = [][][2]float64{append(ring, ring[0])} // close the ring
		fs = append(fs, f)
	}
	doc, _ := json.Marshal(map[string]any{"type": "FeatureCollection", "features": fs})
	return doc
}

func main() {
	a := app.New()
	w := a.NewWindow("Choropleth series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start choropleth
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Population density", SubText: "People per square mile, approximate"},
		Series: []series.Series{
			&series.ChoroplethSeries{
				SeriesName: "Density",
				GeoJSON:    featureCollection(), // or GeoJSONURL to download it
				MapCenter:  [2]float64{-105, 40},
				MapZoom:    4.2,
				RegionData: map[string]float64{
					"Montana": 7, "Wyoming": 6, "Colorado": 56, "Utah": 40,
					"Arizona": 63, "New Mexico": 17, "North Dakota": 11,
					"South Dakota": 12, "Nebraska": 25, "Kansas": 36, "Oklahoma": 58,
				},
				ColorStops: []series.HeatmapColorStop{
					{Position: 0, Color: color.NRGBA{R: 237, G: 248, B: 233, A: 255}},
					{Position: 1, Color: color.NRGBA{R: 0, G: 109, B: 44, A: 255}},
				},
			},
		},
	})
	//docs:end choropleth

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
