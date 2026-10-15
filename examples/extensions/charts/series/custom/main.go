// A custom series: your function draws the series with Bellui canvas objects.
package main

import (
	"fmt"
	"image/color"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/canvas"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/series"
)

type goal struct {
	name        string
	done, total float64
}

func main() {
	a := app.New()
	w := a.NewWindow("Custom series")
	w.Resize(geo.NewSize(640, 420))

	goals := []goal{
		{"Docs pages", 34, 40},
		{"Examples", 52, 60},
		{"Screenshots", 18, 30},
		{"Translations", 3, 12},
	}

	//docs:start custom
	progress := func(ctx series.RenderContext) []core.CanvasObject {
		vp := ctx.Viewport // the plot area, in pixels
		rowH := vp.H / float32(len(goals))
		track := color.NRGBA{R: 128, G: 128, B: 128, A: 60}

		var objs []core.CanvasObject
		for i, g := range goals {
			y := vp.Y + float32(i)*rowH + rowH/2

			label := canvas.NewText(fmt.Sprintf("%s  %.0f / %.0f", g.name, g.done, g.total), ctx.Color)
			label.TextSize = 13
			label.Move(geo.NewPos(vp.X, y-24))
			objs = append(objs, label)

			bg := canvas.NewRectangle(track)
			bg.CornerRadius = 5
			bg.Move(geo.NewPos(vp.X, y))
			bg.Resize(geo.NewSize(vp.W, 10))

			bar := canvas.NewRectangle(ctx.Color) // the series colour from the palette
			bar.CornerRadius = 5
			bar.Move(geo.NewPos(vp.X, y))
			bar.Resize(geo.NewSize(vp.W*float32(g.done/g.total), 10))
			objs = append(objs, bg, bar)
		}
		return objs
	}

	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Release checklist"},
		Series: []series.Series{
			&series.CustomSeries{SeriesName: "Progress", RenderFunc: progress},
		},
	})
	//docs:end custom

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
