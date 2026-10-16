// The toolbox: line/bar switching, and saving the chart as a PNG.
package main

import (
	"image"
	"image/png"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/dialog"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"

	"github.com/belluikit/bellui/x/charts/axis"
	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/component"
	"github.com/belluikit/bellui/x/charts/series"
)

//docs:start capture
// chartImage returns the pixels of c as they are on screen, cut out of a
// capture of the whole window.
func chartImage(w bellui.Window, c *chart.Chart) image.Image {
	cv := w.Canvas()
	pos := bellui.CurrentApp().Driver().AbsolutePositionForObject(c)
	x0, y0 := cv.PixelCoordinateForPosition(pos)
	x1, y1 := cv.PixelCoordinateForPosition(pos.Add(c.Size()))
	img := cv.Capture()
	if sub, ok := img.(interface {
		SubImage(image.Rectangle) image.Image
	}); ok {
		return sub.SubImage(image.Rect(x0, y0, x1, y1))
	}
	return img
}

//docs:end capture

func main() {
	a := app.New()
	w := a.NewWindow("Toolbox")
	w.Resize(geo.NewSize(900, 560))

	//docs:start toolbox
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Signups", SubText: "Switch between line and bar"},
		Toolbox: &component.ToolboxConfig{
			MagicType: []string{"line", "bar"}, // one button per type
		},
		Grid:  &chart.Grid{Left: 56, Right: 20, Top: 12, Bottom: 30},
		XAxis: []axis.Config{{Type: axis.TypeCategory, Data: []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}}},
		YAxis: []axis.Config{{Type: axis.TypeValue}},
		Series: []series.Series{
			&series.BarSeries{SeriesName: "Signups", Data: []float64{120, 200, 150, 80, 70, 110, 130}},
		},
	})
	//docs:end toolbox

	//docs:start save
	save := widget.NewButton("Save PNG", func() {
		img := chartImage(w, c)
		dialog.ShowFileSave(func(out bellui.URIWriteCloser, err error) {
			if err != nil || out == nil {
				return // cancelled
			}
			defer out.Close()
			if err := png.Encode(out, img); err != nil {
				dialog.ShowError(err, w)
			}
		}, w)
	})
	//docs:end save

	w.SetContent(container.NewBorder(container.NewHBox(save), nil, nil, nil, c))
	bellui.ShowAndRun(w)
}
