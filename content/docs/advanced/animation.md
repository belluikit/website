---
title: Animation
description: Animate any property with a per-frame callback, easing curves, springs, keyframes and composed sequences.
weight: 6
---

An animation in Bellui is a function that runs once per frame with a progress
value. You decide what that value changes: a position, a size, a color, or
anything else. The [`canvas`](https://pkg.go.dev/github.com/belluikit/bellui/canvas)
package adds ready-made helpers on top.

## The basic animation

{{< example file="advanced/animation/main.go" region="basic" >}}

[`bellui.NewAnimation(duration, tick)`](https://pkg.go.dev/github.com/belluikit/bellui#NewAnimation)
returns an [`Animation`](https://pkg.go.dev/github.com/belluikit/bellui#Animation).
Nothing moves until you call `Start()`. `Stop()` removes it again.

While it runs, `Tick` is called once per frame, just before the window is
drawn, on the UI thread, so it can change widgets directly. The value it gets
climbs from 0 towards 1, and is exactly 1 on the last frame (0 when
`AutoReverse` is set, because the last run goes backwards). Fields you can set
before starting:

| Field | Meaning |
| --- | --- |
| `Curve` | how time maps to progress (see below); `AnimationEaseInOut` when nil |
| `RepeatCount` | how many extra times to run; `AnimationRepeatForever` (-1) loops until `Stop` |
| `AutoReverse` | run backwards from 1 to 0 after each forward run |

Canvas objects such as rectangles, text and images redraw themselves when you
move or resize them. When you change another field, such as `FillColor`, call
the object's `Refresh()`.

## Easing curves

A curve is a `func(float32) float32` that maps time (0 to 1) to progress. You
can write your own. The built-in ones are:

| Curve | Feel |
| --- | --- |
| `AnimationLinear` | constant speed |
| `AnimationEaseIn` | starts slowly, speeds up |
| `AnimationEaseOut` | starts fast, slows down |
| `AnimationEaseInOut` | slow, fast, slow (the default) |
| `AnimationEaseStandard` | Material 3 standard easing |
| `AnimationEaseStandardAccelerate` | Material 3, for things leaving the screen |
| `AnimationEaseStandardDecelerate` | Material 3, for things entering the screen |
| `AnimationEaseEmphasized` | Material 3 emphasized easing |

The `canvas` package also has duration constants: `DurationShort` (150 ms),
`DurationStandard` (300 ms), and the Material 3 steps from `DurationShort1`
(50 ms) to `DurationLong2` (500 ms).

## Animating values

These `canvas` helpers work out the in-between values and hand them to your
function:

{{< example file="advanced/animation/main.go" region="colour" >}}

| Helper | Animates |
| --- | --- |
| `NewFloat32Animation(start, stop, d, fn)` | a number, such as opacity, angle or scale |
| `NewPositionAnimation(start, stop, d, fn)` | a `geo.Position` |
| `NewSizeAnimation(start, stop, d, fn)` | a `geo.Size` |
| `NewColorRGBAAnimation(start, stop, d, fn)` | a color, in a straight line through RGB |
| `NewPathAnimation(points, d, fn)` | a position through a list of points, equal time per segment |

Each returns a `*bellui.Animation`, so you can set `Curve`, `RepeatCount` and
`AutoReverse` on it before `Start()`. `NewPathAnimation` panics with fewer than
two points.

## Changing direction midway

A button that slides a panel open and closed needs to reverse from wherever the
panel is, not jump to an end first.
[`AnimatedFloat32`](https://pkg.go.dev/github.com/belluikit/bellui/canvas#AnimatedFloat32)
and `AnimatedPosition` remember their current value for exactly that:

{{< example file="advanced/animation/main.go" region="retarget" >}}

`AnimateTo` stops any animation still running and starts a new one from the
current value. `Value()` reads the current value, and `Stop()` freezes it where
it is. Use these only from the UI thread.

## Springs

A spring animation has no fixed curve. It follows a physical spring, so it can
overshoot and settle:

{{< example file="advanced/animation/main.go" region="spring" >}}

[`SpringConfig`](https://pkg.go.dev/github.com/belluikit/bellui/canvas#SpringConfig)
has `Stiffness`, `Damping`, `Mass` and `MaxDuration`. Zero fields get the
defaults: 170, 26, 1 and 2 seconds. The presets are `SpringGentle`,
`SpringBouncy` and `SpringStiff`. The animation always lasts `MaxDuration`;
once the spring has settled it simply holds the end value.
`NewSpringAnimation(config, fn)` gives you the raw progress, which can go above 1
while the spring overshoots.

## Keyframes and combinations

- `NewKeyframeAnimation(keyframes, d, fn)` moves a number through a list of
  [`Keyframe`](https://pkg.go.dev/github.com/belluikit/bellui/canvas#Keyframe)
  values, each with its own progress point and an optional curve for the
  stretch leading to it. `NewPositionKeyframeAnimation` does the same for
  positions. The first keyframe must be at progress 0, the last at 1, and they
  must be in order, and there must be at least two; otherwise the function
  panics.
- `NewParallelAnimation(a, b, ...)` runs animations together. It lasts as long
  as the longest.
- `NewSequenceAnimation(a, b, ...)` runs them one after another.
- `NewStaggeredAnimation(entries)` starts each animation after its own delay,
  for example to bring list items in one by one.

The combined animation uses each child's curve and ignores the children's
`RepeatCount` and `AutoReverse`. Set those on the combined animation instead.

## Respect the user's preference

Built-in widgets skip their animations when the app's settings say not to
animate: the user turned off "Animate widgets" in the `bellui_settings` app
(stored as `disable_animations` in the global `settings.pxf`), or the app was
built with the `no_animations` tag. Your own animations run regardless, so check the setting
and jump to the end state instead:

{{< example file="advanced/animation/main.go" region="motion" >}}

## The whole example

{{< example "advanced/animation/main.go" >}}
