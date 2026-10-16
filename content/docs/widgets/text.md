---
title: Displaying text
description: Label, Hyperlink, RichText with Markdown and HTML, TextGrid, icons, separators and animated GIFs.
weight: 8
---

These widgets show text and small images that the user reads but does not
edit. For editable text, see [Text input](../text-input/).

{{< screenshot name="typography" alt="The gallery's Typography page: Display, Headline and Title samples from the MD3 type scale" >}}

## Which one to use

| You want to show | Use |
| --- | --- |
| A line or paragraph of text in one style | `Label` |
| A link that opens a web page or runs a function | `Hyperlink` |
| Text with mixed styles, links, lists or code | `RichText` |
| A Markdown or HTML document | `NewRichTextFromMarkdown`, `NewRichTextFromHTML` |
| Monospaced output such as logs or a terminal | `TextGrid` |
| A themed icon | `Icon` |
| A file's type as an icon | `FileIcon` |
| A thin dividing line | `Separator` |
| An animated GIF | `AnimatedGif` |

## Label

{{< example file="widgets/text/main.go" region="label" >}}

[`Label`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Label) shows
text in a single style. `SetText` changes it and is safe to call from any
goroutine.

- `TextStyle` sets bold, italic, monospace and so on, and `Alignment` aligns
  the text. `NewLabelWithStyle` sets both at once.
- `SizeName` picks a text size from the theme, such as
  `theme.SizeNameHeadingText`.
- `Wrapping` wraps long text (`bellui.TextWrapWord` breaks between words).
  `Truncation` cuts it off instead; `bellui.TextTruncateEllipsis` adds "…".
  With truncation on, `ShowTooltipWhenTruncated` shows the full text on hover.
- `Importance` colours the text: `LowImportance` uses the disabled colour,
  `HighImportance` the primary colour, and `DangerImportance`,
  `WarningImportance` and `SuccessImportance` the matching theme colours.
- `Selectable` lets the user select and copy the text.

A wrapping or truncating label asks for almost no width, so give it room
with a layout. `IntrinsicSize()` tells you the size the text would need on one
line. `NewLabelWithData` connects the text to a `binding.String`.

## Hyperlink

{{< example file="widgets/text/main.go" region="link" >}}

[`Hyperlink`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Hyperlink)
draws text in the primary colour and opens `URL` in the default browser when
tapped. Set `OnTapped` to run your own function instead; then the URL may be
`nil`. `ColorName` picks another theme colour, and `SetURLFromString` sets the
URL from a string. Space activates a focused link.

## RichText

[`RichText`](https://pkg.go.dev/github.com/belluikit/bellui/widget#RichText)
shows a list of segments, each with its own style. Text segments are
[`TextSegment`](https://pkg.go.dev/github.com/belluikit/bellui/widget#TextSegment)
values with a
[`RichTextStyle`](https://pkg.go.dev/github.com/belluikit/bellui/widget#RichTextStyle).
The package has ready-made styles: `RichTextStyleInline`,
`RichTextStyleParagraph`, `RichTextStyleStrong`, `RichTextStyleEmphasis`,
`RichTextStyleHeading`, `RichTextStyleSubHeading`, `RichTextStyleBlockquote`,
`RichTextStyleCodeInline` and `RichTextStyleCodeBlock`.

{{< example file="widgets/richtext/main.go" region="segments" >}}

Inline segments flow on together in one paragraph. A segment whose style is
not inline starts its own block. Other segment types add links
(`HyperlinkSegment`), images (`ImageSegment`), lists (`ListSegment`), tables
(`TableSegment`), separators (`SeparatorSegment`) and syntax-highlighted code
(`CodeSegment`).

After you change `Segments`, call `Refresh()`. Set `Wrapping` for text that
should wrap, and `Selectable` to let the user copy it.

### Markdown

{{< example file="widgets/richtext/main.go" region="markdown" >}}

`NewRichTextFromMarkdown` parses Markdown into segments. Headings, emphasis,
lists, links, code, block quotes and images work, and so do tables,
strikethrough and task lists. A fenced code block that names its language is
syntax-highlighted for Go, JSON, YAML, Starlark and Protocol Buffers.
`ParseMarkdown` replaces the content of an existing `RichText` and
`AppendMarkdown` adds to it.

### HTML

{{< example file="widgets/richtext/main.go" region="html" >}}

`NewRichTextFromHTML` reads a subset of HTML: paragraphs, headings, block
quotes, `pre`, lists, tables, `b`/`strong`, `i`/`em`, `u`, `s`/`del`,
`code`, links, images, `br` and `hr`. It is for formatted text, not web
pages: there is no CSS layout and no scripting. `ParseHTML` and `AppendHTML`
work like their Markdown twins.

{{< callout type="warning" >}}
By default, images in Markdown or HTML that point to a web address or a local
file are not loaded; a placeholder is shown instead. This stops untrusted
content from making network requests or reading files. Set
`AllowRemoteContent = true` only for content you trust.
{{< /callout >}}

## TextGrid

[`TextGrid`](https://pkg.go.dev/github.com/belluikit/bellui/widget#TextGrid)
is a grid of monospaced cells, one character each, for log viewers, code
previews and terminals.

{{< example file="widgets/textgrid/main.go" region="grid" >}}

- `SetText` replaces everything and `Append` adds lines at the end. `Text()`
  returns the content without styles.
- `SetRune`, `SetCell` and `SetRow` change single cells or rows.
- `SetStyle`, `SetRowStyle` and `SetStyleRange` colour cells. A style is any
  [`TextGridStyle`](https://pkg.go.dev/github.com/belluikit/bellui/widget#TextGridStyle);
  `CustomTextGridStyle` takes a text colour, a background colour and a
  `TextStyle`.
- `ShowLineNumbers` and `ShowWhitespace` add a line-number column and
  visible whitespace. `ScrollToBottom()` follows new output.

## Icon, FileIcon and Separator

{{< example file="widgets/text/main.go" region="icon" >}}

[`Icon`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Icon) shows a
resource, usually one of the theme icons such as `theme.InfoIcon()`, and
recolours it to match the theme. `SetResource` swaps it.
[`FileIcon`](https://pkg.go.dev/github.com/belluikit/bellui/widget#FileIcon)
takes a file URI and shows an icon labelled with the file's type.
[`Separator`](https://pkg.go.dev/github.com/belluikit/bellui/widget#Separator)
is a thin dividing line in a theme colour.

[`AnimatedGif`](https://pkg.go.dev/github.com/belluikit/bellui/widget#AnimatedGif)
plays a GIF from a URI (`NewAnimatedGif`) or a resource
(`NewAnimatedGifFromResource`). Call `Start()` to play and `Stop()` to pause
on the current frame. A GIF has no size of its own on screen, so call
`SetMinSize` or let the layout size it.
