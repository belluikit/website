# bellui.com

The website and documentation for [Bellui](https://github.com/belluikit/bellui),
built with [Hugo](https://gohugo.io) and the [Hextra](https://github.com/imfing/hextra)
theme.

## Working on it

You need Hugo 0.146 or later (the standard edition is enough) and Go.

```sh
hugo server            # http://localhost:1313, rebuilds on save
scripts/check-examples.sh   # compile every code example against the pinned bellui
hugo --minify          # production build into public/
```

The theme is vendored in `_vendor/`, so a build needs no network access. To
move to a newer Hextra, run `hugo mod get github.com/imfing/hextra@<version>`
and then `hugo mod vendor`, and commit both.

## Code examples are compiled

Every Go snippet on the site lives in `examples/` as a real file, and pages
include it with the `example` shortcode:

```
{{< example "hello/main.go" >}}                      the whole file
{{< example file="layout/main.go" region="grid" >}}  only one region
```

A region is the lines between `//docs:start grid` and `//docs:end grid`; the
marker lines are never shown. `scripts/check-examples.sh` builds and vets the
`examples/` module against the bellui version pinned in `examples/go.mod`, and
fails if a page names a missing file or an example is shown nowhere. Do not
paste Go code into a page by hand: it will stop compiling and nobody will
notice. Shell commands and other languages are fine inline.

When bellui is released, bump the pin and rerun the check:

```sh
cd examples && go get github.com/belluikit/bellui@<version> && go mod tidy
```

How pages are written is in [WRITING.md](WRITING.md).

## Deploying (Cloudflare Pages)

The site is built and served by Cloudflare Pages from this repository.

| Setting | Value |
| --- | --- |
| Framework preset | Hugo |
| Build command | `hugo --minify` |
| Build output directory | `public` |
| Environment variable | `HUGO_VERSION` = `0.168.0` |

The custom domain is `bellui.com`; `www.bellui.com` and `bellui.io` redirect to it.

## License

Everything in this repository, the pages and the code in `examples/`, is
released under the [BSD 3-Clause license](LICENSE), the same as Bellui. The
vendored Hextra theme is MIT licensed; its license is in
[`third_party/hextra/LICENSE`](third_party/hextra/LICENSE).

Bellui began as a fork of [Fyne](https://fyne.io). Nothing here is copied from
Fyne's documentation, which is not published under an open license.
