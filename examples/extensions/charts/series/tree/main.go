// A tree layout: an organisation chart drawn left to right.
package main

import (
	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/geo"

	"github.com/belluikit/bellui/x/charts/chart"
	"github.com/belluikit/bellui/x/charts/series"
)

func main() {
	a := app.New()
	w := a.NewWindow("Tree series")
	w.Resize(geo.NewSize(640, 420))

	//docs:start tree
	c := chart.New(chart.Option{
		Title: &chart.Title{Text: "Engineering"},
		Series: []series.Series{
			&series.TreeSeries{
				SeriesName: "Org",
				Orient:     "LR", // left to right; the default "TB" is top to bottom
				NodeSize:   7,
				Data: series.TreeNode{Name: "CTO", Children: []series.TreeNode{
					{Name: "Platform", Children: []series.TreeNode{
						{Name: "Build"}, {Name: "Runtime"}, {Name: "Storage"},
					}},
					{Name: "Apps", Children: []series.TreeNode{
						{Name: "Desktop"}, {Name: "Mobile"},
					}},
					{Name: "Quality", Children: []series.TreeNode{
						{Name: "Testing"}, {Name: "Release"},
					}},
				}},
			},
		},
	})
	//docs:end tree

	w.SetContent(c)
	bellui.ShowAndRun(w)
}
