#!/bin/sh
# Compiles every example the site shows, against the bellui version pinned in
# examples/go.mod, and checks that pages and examples refer to each other.
#
# Run from the repository root:  scripts/check-examples.sh
set -eu

cd "$(dirname "$0")/.."

echo "== build and vet examples/ (bellui $(cd examples && go list -m -f '{{.Version}}' github.com/belluikit/bellui))"
(
	cd examples
	GOWORK=off go build -o /dev/null ./...
	GOWORK=off go vet ./...
)

status=0

# Every {{< example ... >}} must name a file that exists.
refs=$(grep -rhoE '\{\{< *example +(file=)?"[^"]+"' content | sed -E 's/.*"([^"]+)"/\1/' | sort -u)
for f in $refs; do
	if [ ! -f "examples/$f" ]; then
		echo "missing example: examples/$f (referenced from content/)"
		status=1
	fi
done

# Every example file must be shown somewhere, or it is dead weight that is
# compiled but documents nothing.
for f in $(cd examples && find . -name '*.go' ! -name '*_test.go' | sed 's#^\./##' | sort); do
	if ! printf '%s\n' "$refs" | grep -qx "$f"; then
		echo "unreferenced example: examples/$f"
		status=1
	fi
done

n=$(cd examples && find . -name '*.go' ! -name '*_test.go' | wc -l | tr -d ' ')
echo "== $n example files, $(printf '%s\n' "$refs" | grep -c .) referenced"
exit $status
