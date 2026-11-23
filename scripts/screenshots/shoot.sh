#!/bin/sh
# Captures a light and a dark screenshot of one example app, rendered by
# Bellui itself, into static/images/gallery/<name>-{light,dark}.png.
#
#   scripts/screenshots/shoot.sh <example dir under examples/> <name> [WxH] [delay-seconds] [args...]
#   scripts/screenshots/shoot.sh extensions/dag/canvas dag 900x600
#
# It copies the example next to itself under a "_shot" directory (which
# go build ./... ignores), replaces its bellui.ShowAndRun(w) call with a
# version that resizes the window, waits, captures Canvas().Capture() and
# quits, and runs it with BELLUI_SCALE=2: once with BELLUI_THEME=light and
# once with BELLUI_THEME=dark. Each variant gets its own process, because
# x/charts charts keep the colours they were first drawn in when the theme
# variant changes while the app runs (belluikit/bellui#1381).
set -eu
cd "$(dirname "$0")/../.."
src="examples/$1"; name="$2"; size="${3:-900x650}"; delay="${4:-3}"
[ $# -ge 4 ] && shift 4 || shift $#
[ -f "$src/main.go" ] || { echo "no $src/main.go" >&2; exit 1; }

mod=$(cd "$src" && go env GOMOD); moddir=$(dirname "$mod")
tmp="$moddir/_shot/$name"
rm -rf "$tmp"; mkdir -p "$tmp"
cp -R "$src"/. "$tmp"/
rm -f "$tmp"/*_test.go
grep -q 'bellui.ShowAndRun(' "$tmp/main.go" || { echo "$src/main.go has no bellui.ShowAndRun call" >&2; exit 1; }
sed -i.bak -E 's/bellui\.ShowAndRun\(([A-Za-z0-9_]+)\)/shotShowAndRun(\1)/' "$tmp/main.go" && rm -f "$tmp/main.go.bak"
# The rewrite can leave main.go's bellui import otherwise unused.
printf '\nvar _ = bellui.ShowAndRun\n' >> "$tmp/main.go"

cat > "$tmp/zz_shot.go" <<'EOF'
package main

import (
	"errors"
	"image"
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/geo"
)

func shotShowAndRun(w bellui.Window) {
	if s := strings.Split(os.Getenv("SHOT_SIZE"), "x"); len(s) == 2 {
		width, _ := strconv.Atoi(s[0])
		height, _ := strconv.Atoi(s[1])
		w.Resize(geo.NewSize(float32(width), float32(height)))
	}
	delay, _ := strconv.ParseFloat(os.Getenv("SHOT_DELAY"), 64)
	go func() {
		a := bellui.CurrentApp()
		time.Sleep(time.Duration(delay * float64(time.Second)))
		// Capture returns nil when macOS reports the window as covered by
		// another one, so try again a few times before giving up.
		var img image.Image
		for try := 0; img == nil && try < 10; try++ {
			if try > 0 {
				time.Sleep(500 * time.Millisecond)
			}
			bellui.DoAndWait(func() { img = w.Canvas().Capture() })
		}
		fn := os.Getenv("SHOT_OUT") + "-" + os.Getenv("BELLUI_THEME") + ".png"
		err := errors.New("Capture returned nil: is the window covered?")
		if img != nil {
			var f *os.File
			if f, err = os.Create(fn); err == nil {
				err = png.Encode(f, img)
				f.Close()
			}
		}
		log.Printf("SHOT %s err=%v", fn, err)
		bellui.DoAndWait(a.Quit)
	}()
	bellui.ShowAndRun(w)
}
EOF

out="$PWD/static/images/gallery/$name"
log="$moddir/_shot/$name.log"
(cd "$tmp" && GOWORK=off go build -o shot .) > "$log" 2>&1 || { tail -20 "$log" >&2; rm -rf "$moddir/_shot"; exit 1; }
for v in light dark; do
	if ! (cd "$tmp" && BELLUI_THEME=$v BELLUI_SCALE=2 SHOT_SIZE="$size" SHOT_DELAY="$delay" SHOT_OUT="$out" ./shot "$@") > "$log" 2>&1; then
		tail -20 "$log" >&2
		rm -rf "$moddir/_shot"
		exit 1
	fi
	grep 'SHOT ' "$log"
	grep -q 'SHOT .* err=<nil>' "$log" || { rm -rf "$moddir/_shot"; exit 1; }
done
rm -rf "$moddir/_shot"
ls -la "$out-light.png" "$out-dark.png"
