#!/bin/sh
# Captures a light and a dark screenshot of one example app, rendered by
# Bellui itself, into static/images/gallery/<name>-{light,dark}.png.
#
#   scripts/screenshots/shoot.sh <example dir under examples/> <name> [WxH] [delay-seconds] [args...]
#   scripts/screenshots/shoot.sh extensions/dag/canvas dag 900x600
#
# It copies the example next to itself under a "_shot" directory (which
# go build ./... ignores), replaces its bellui.ShowAndRun(w) call with a
# version that resizes the window, waits, captures Canvas().Capture() in
# both theme variants and quits, and runs it with BELLUI_SCALE=2.
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
	"image/png"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/theme"
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
		for _, v := range []string{"light", "dark"} {
			bellui.DoAndWait(func() {
				if v == "light" {
					a.Settings().SetThemeVariant(theme.VariantLight)
				} else {
					a.Settings().SetThemeVariant(theme.VariantDark)
				}
			})
			time.Sleep(time.Duration(delay * float64(time.Second)))
			bellui.DoAndWait(func() {
				fn := os.Getenv("SHOT_OUT") + "-" + v + ".png"
				f, err := os.Create(fn)
				if err == nil {
					err = png.Encode(f, w.Canvas().Capture())
					f.Close()
				}
				log.Printf("SHOT %s err=%v", fn, err)
			})
		}
		bellui.DoAndWait(a.Quit)
	}()
	bellui.ShowAndRun(w)
}
EOF

out="$PWD/static/images/gallery/$name"
log="$moddir/_shot/$name.log"
if ! (cd "$tmp" && BELLUI_SCALE=2 SHOT_SIZE="$size" SHOT_DELAY="$delay" SHOT_OUT="$out" GOWORK=off go run . "$@") > "$log" 2>&1; then
	tail -20 "$log" >&2
	rm -rf "$moddir/_shot"
	exit 1
fi
grep 'SHOT ' "$log"
rm -rf "$moddir/_shot"
ls -la "$out-light.png" "$out-dark.png"
