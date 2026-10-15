// GridWrap: a scrolling grid of equally sized tiles that wraps to the width.
package main

import (
	"fmt"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/core"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/theme"
	"github.com/belluikit/bellui/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("GridWrap")
	w.Resize(geo.NewSize(480, 360))

	//docs:start grid
	photos := widget.NewGridWrap(
		func() int { return 200 },
		func() core.CanvasObject {
			return container.NewVBox(widget.NewIcon(theme.MediaPhotoIcon()), widget.NewLabel("IMG_0000"))
		},
		func(id widget.GridWrapItemID, tile core.CanvasObject) {
			tile.(*bellui.Container).Objects[1].(*widget.Label).SetText(fmt.Sprintf("IMG_%04d", id))
		},
	)
	photos.OnSelected = func(id widget.GridWrapItemID) {
		w.SetTitle(fmt.Sprintf("Photo %d of %d columns", id, photos.ColumnCount()))
	}
	//docs:end grid

	w.SetContent(photos)
	bellui.ShowAndRun(w)
}
