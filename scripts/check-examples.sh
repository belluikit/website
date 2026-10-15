#!/bin/sh
# Compiles every example the site shows, against the bellui version pinned in
# examples/go.mod, and checks that pages and examples refer to each other.
#
# Run from the repository root:  scripts/check-examples.sh
set -eu

cd "$(dirname "$0")/.."

# examples/ holds one Go module per directory that has a go.mod: the toolkit
# examples in examples/, and the x/ extension examples in examples/extensions/,
# which need the extension modules as well.
for mod in $(find examples -name go.mod | sort); do
	dir=$(dirname "$mod")
	echo "== build and vet $dir (bellui $(cd "$dir" && go list -m -f '{{.Version}}' github.com/belluikit/bellui))"
	(
		cd "$dir"
		GOWORK=off go build -o /dev/null ./...
		GOWORK=off go vet ./...
	)
done

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
# compiled but documents nothing. The exception is a shots/ directory: those
# programs exist to be captured by scripts/screenshots/shoot.sh, and are
# compiled here so they keep working, but no page shows their code.
for f in $(cd examples && find . -name '*.go' ! -name '*_test.go' ! -path '*/shots/*' | sed 's#^\./##' | sort); do
	if ! printf '%s\n' "$refs" | grep -qx "$f"; then
		echo "unreferenced example: examples/$f"
		status=1
	fi
done

n=$(cd examples && find . -name '*.go' ! -name '*_test.go' | wc -l | tr -d ' ')
echo "== $n example files, $(printf '%s\n' "$refs" | grep -c .) referenced"
exit $status
