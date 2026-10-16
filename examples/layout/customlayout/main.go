// A custom layout that keeps its children at a fixed aspect ratio, as large
// as fits, centred in the space.
package main

import (
	"image/color"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/canvas"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"
)

//docs:start layout
// aspectLayout gives every child the largest size with the given
// width/height ratio that fits, centred in the container.
type aspectLayout struct {
	ratio float32
}

// Layout positions and sizes the children. It runs whenever the
// container is resized or refreshed.
func (l aspectLayout) Layout(objects []core.CanvasObject, size geo.Size) {
	w, h := size.Width, size.Width/l.ratio
	if h > size.Height {
		w, h = size.Height*l.ratio, size.Height
	}
	pos := geo.NewPos((size.Width-w)/2, (size.Height-h)/2)
	for _, o := range objects {
		if !o.Visible() {
			continue
		}
		o.Resize(geo.NewSize(w, h))
		o.Move(pos)
	}
}

// MinSize is the smallest size the container can have: big enough for
// every visible child's own minimum, at the right ratio.
func (l aspectLayout) MinSize(objects []core.CanvasObject) geo.Size {
	var min geo.Size
	for _, o := range objects {
		if o.Visible() {
			min = min.Max(o.MinSize())
		}
	}
	if min.Width/l.ratio < min.Height {
		min.Width = min.Height * l.ratio
	} else {
		min.Height = min.Width / l.ratio
	}
	return min
}

//docs:end layout

func main() {
	a := app.NewWithID("com.example.customlayout")
	w := a.NewWindow("Custom layout")

	//docs:start use
	screen := canvas.NewRectangle(color.NRGBA{R: 0x20, G: 0x20, B: 0x28, A: 0xff})
	caption := widget.NewLabel("16:9")
	video := container.New(aspectLayout{ratio: 16.0 / 9.0}, screen, container.NewCenter(caption))
	//docs:end use

	w.SetContent(video)
	w.Resize(geo.NewSize(600, 500))
	bellui.ShowAndRun(w)
}
