---
title: Translations
description: Write translation files, load them, and translate plain text, templated text, plurals and grammatical variants.
weight: 1
---

## Translation files

A translation file, or catalog, is a PXF file, one per language. It holds one
`messages` map from each message ID to its translation. A plain message is a
string. A message with plural forms is a block with one entry per form. This
is the German file of the example on this page:

{{< example file="localisation/translations/translations/de.pxf" lang="text" >}}

Put an ID in double quotes unless it is a single word, such as
`files_selected`. A quoted ID always works.

Name each file after its language: `de.pxf`, `pt_BR.pxf`, `zh_Hans.pxf`.
A prefix before the language is allowed, as in `app.de.pxf`. Bellui does not
read JSON catalogs: a `.json` file is an error. To convert one, move each key
into the `messages` map and write its placeholders as `{Name}`.

`{Name}` is a placeholder, filled from data you pass when you translate. See
[Placeholders](#placeholders).

## Write the English file with `bellui translate`

`bellui translate` reads your Go source, finds the strings you pass to the
`lang` functions, and writes them into a catalog:

```sh
bellui translate translations/en.pxf .
```

It creates the `translations` folder if it is missing. Before v1.3.0 the
folder had to exist already, so with an older `bellui` command, create it
first. Run it again after you change the code: it adds new
messages and leaves the ones already in the file alone. With `--update`
(`-u`) it replaces those too, with the text from the source.

The translation it writes uses `{Name}` placeholders, even where the source
still uses the older `{{.Name}}`. The ID stays as the source wrote it, since
it is what the code looks up. A plural message gets only the `other` form, so
add `one` and the other forms your language needs by hand.
Copy the file for each language you translate into.

## Load the files

Keep the files in a directory next to your code and embed it:

{{< example file="localisation/translations/main.go" region="embed" >}}

Then load them once at startup, before you build the UI:

{{< example file="localisation/translations/main.go" region="load" >}}

[`lang.AddTranslationsFS`](https://pkg.go.dev/github.com/belluikit/bellui/lang#AddTranslationsFS)
loads every file in the directory. If one file fails, it loads the others and
returns the first error. To load a single file, use
[`AddTranslations`](https://pkg.go.dev/github.com/belluikit/bellui/lang#AddTranslations)
with a resource whose name ends in the language, such as `fr.pxf`, or
[`AddTranslationsForLocale`](https://pkg.go.dev/github.com/belluikit/bellui/lang#AddTranslationsForLocale)
with the bytes of a PXF catalog and the locale.

## Translate text

{{< example file="localisation/translations/main.go" region="use" >}}

| Function | Short name | Use it for |
| --- | --- | --- |
| [`lang.Localize`](https://pkg.go.dev/github.com/belluikit/bellui/lang#Localize) | `lang.L` | Text whose ID is the English text itself. |
| [`lang.LocalizePlural`](https://pkg.go.dev/github.com/belluikit/bellui/lang#LocalizePlural) | `lang.N` | Text with plural forms, chosen by a count. |
| [`lang.LocalizeKey`](https://pkg.go.dev/github.com/belluikit/bellui/lang#LocalizeKey) | `lang.X` | Text looked up by a separate ID, with a fallback. |
| [`lang.LocalizePluralKey`](https://pkg.go.dev/github.com/belluikit/bellui/lang#LocalizePluralKey) | `lang.XN` | Plural text looked up by ID, with a fallback. |

The data values come last. See [Placeholders](#placeholders).

**Plural forms.** The count picks the form. Which forms a language has comes
from its CLDR plural rules (the Unicode project's data on how each language
counts): English has `one` and `other`, Polish and Russian also have `few`
and `many`. Write the forms your language needs, and always include `other`.
The count only chooses the form. To show the number, pass it as data as well,
as the example does with `Count`.

**IDs or English text.** With `L` and `N` the ID is the English text, so
the source reads naturally. `L` needs no English file. `N` without an English
file uses the ID for every count, so a count of 1 reads "1 files selected".
Add an English file with the `one` form for it. With `X` and `XN` the ID
is a short key such as `button.save`, and you pass the English text as the
fallback. Keys stay stable when you reword the English.

## Placeholders

A placeholder is a name in braces, such as `{Name}`. Each call takes any
number of data values after its other arguments. A value is a map with string
keys, or a struct whose exported fields give the names. With several values,
the first one that has a name fills it. This lets you pass a struct and a map
together:

{{< example file="localisation/translations/main.go" region="types" >}}

{{< example file="localisation/translations/main.go" region="several" >}}

A placeholder never calls a method. `{Count}` is not filled from a `Count()`
method: Bellui leaves it as written and logs a message once that names the
map to pass instead. Call the method yourself and put the result in a map, as
above.

A name that no data value has stays in the text as written, braces included.
So do braces that hold no name, such as `{}` or `{ x }`.

The older form `{{.Name}}` is deprecated. It is still filled the same way, and
Bellui logs each spelling once with its replacement. Write `{Name}` instead.
Other Go template actions, such as `{{if .X}}`, are not run: they are left as
written and logged once.

### Fill a string without a catalog

[`lang.Fill`](https://pkg.go.dev/github.com/belluikit/bellui/lang#Fill) fills
placeholders by the same rules, without looking the string up in a catalog.
Use it for text you format yourself:

{{< example file="localisation/translations/main.go" region="fill" >}}

It goes over the string once, so a placeholder inside a value it filled in is
not filled again. There is no escape for a literal brace.

## Grammatical variants

Some sentences change with something other than a count, such as the
grammatical gender of a person. Write one message per variant, with the
variant name after a dot, plus an `.other` variant:

```text
messages = {
  "profile_updated.female": "{Name} updated her profile"
  "profile_updated.male": "{Name} updated his profile"
  "profile_updated.other": "{Name} updated their profile"
}
```

[`lang.LSelect`](https://pkg.go.dev/github.com/belluikit/bellui/lang#LSelect)
picks the variant. An unknown selector uses `.other`. If that is missing too,
it returns the prefix itself.
[`lang.LSelectPlural`](https://pkg.go.dev/github.com/belluikit/bellui/lang#LSelectPlural)
picks a variant and then a plural form inside it, so each variant is a
plural block.

## When a translation is missing

Bellui chooses one language for the whole app. When your app has no file for
that language, it uses your English file. This happens whenever the user's
language is one of the 15 languages Bellui itself is translated into but your
app is not: a French user of an app with only English and German files sees
Bellui's own text in French and your text in English.

Without an English file, your messages are not found:

- `L` returns the text you wrote, which is English.
- `X` and `XN` return the fallback you passed.
- `N` and `LSelect` return the message ID itself, such as `files_selected`.

So ship an English file, or use `XN` with a fallback rather than `N` with a
key.

The full program:

{{< example "localisation/translations/main.go" >}}
