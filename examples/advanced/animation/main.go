// Animations: a looping pulse, a colour change, a retargetable slide and a spring.
package main

import (
	"image/color"
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/canvas"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/widget"
)

var (
	blue = color.NRGBA{R: 0x1E, G: 0x88, B: 0xE5, A: 0xFF}
	pink = color.NRGBA{R: 0xD8, G: 0x1B, B: 0x60, A: 0xFF}
)

func main() {
	a := app.New()
	w := a.NewWindow("Animation")

	box := canvas.NewRectangle(blue)
	box.CornerRadius = 8
	box.Resize(geo.NewSize(60, 60))
	box.Move(geo.NewPos(20, 20))
	stage := container.NewWithoutLayout(box)

	//docs:start basic
	// Tick gets the progress, from 0 to 1, once per frame.
	pulse := bellui.NewAnimation(800*time.Millisecond, func(p float32) {
		size := 60 + 20*p
		box.Resize(geo.NewSize(size, size))
	})
	pulse.AutoReverse = true
	pulse.RepeatCount = bellui.AnimationRepeatForever
	pulse.Curve = bellui.AnimationEaseInOut
	//docs:end basic

	//docs:start colour
	toPink := canvas.NewColorRGBAAnimation(blue, pink, canvas.DurationStandard, func(c color.Color) {
		box.FillColor = c
		box.Refresh()
	})
	//docs:end colour

	//docs:start motion
	makePink := func() {
		if !a.Settings().ShowAnimations() {
			box.FillColor = pink // Jump straight to the end state.
			box.Refresh()
			return
		}
		toPink.Start()
	}
	//docs:end motion

	//docs:start retarget
	// AnimateTo starts from wherever the box is now, even mid-animation.
	x := canvas.NewAnimatedFloat32(20)
	slideTo := func(target float32) {
		x.AnimateTo(target, canvas.DurationMedium4, bellui.AnimationEaseStandard, func(v float32) {
			box.Move(geo.NewPos(v, box.Position().Y))
		})
	}
	//docs:end retarget

	//docs:start spring
	drop := canvas.NewSpringFloat32Animation(20, 200, canvas.SpringBouncy, func(v float32) {
		box.Move(geo.NewPos(box.Position().X, v))
	})
	//docs:end spring

	controls := container.NewGridWithColumns(3,
		widget.NewButton("Pulse", pulse.Start),
		widget.NewButton("Stop pulse", pulse.Stop),
		widget.NewButton("Pink", makePink),
		widget.NewButton("Left", func() { slideTo(20) }),
		widget.NewButton("Right", func() { slideTo(300) }),
		widget.NewButton("Drop", drop.Start),
	)

	w.SetContent(container.NewBorder(controls, nil, nil, nil, stage))
	w.Resize(geo.NewSize(480, 400))
	bellui.ShowAndRun(w)
}
