---
title: Translations
description: Write translation files, load them, and translate plain text, templated text, plurals and grammatical variants.
weight: 1
---

## Translation files

Translations are JSON files, one per language. Each maps a message ID to its
translation. A plain message is a string. A message with plural forms is an
object with one entry per form:

```json
{
  "Welcome, {{.Name}}!": "Willkommen, {{.Name}}!",
  "files_selected": {
    "one": "{{.Count}} Datei ausgewählt",
    "other": "{{.Count}} Dateien ausgewählt"
  },
  "button.save": "Speichern"
}
```

Name each file after its language: `de.json`, `pt_BR.json`, `zh_Hans.json`.
A prefix before the language is allowed, as in `app.de.json`.

`{{.Name}}` is a placeholder filled from data you pass when you translate.
The syntax is Go's [`text/template`](https://pkg.go.dev/text/template).

## Load the files

Keep the files in a directory next to your code and embed it:

{{< example file="localisation/translations/main.go" region="embed" >}}

Then load them once at startup, before you build the UI:

{{< example file="localisation/translations/main.go" region="load" >}}

[`lang.AddTranslationsFS`](https://pkg.go.dev/github.com/belluikit/bellui/lang#AddTranslationsFS)
loads every file in the directory. If one file fails, it loads the others and
returns the first error. To load a single file, use
[`AddTranslations`](https://pkg.go.dev/github.com/belluikit/bellui/lang#AddTranslations)
with a resource whose name ends in the language, such as `fr.json`, or
[`AddTranslationsForLocale`](https://pkg.go.dev/github.com/belluikit/bellui/lang#AddTranslationsForLocale)
with the bytes and the locale.

## Translate text

{{< example file="localisation/translations/main.go" region="use" >}}

| Function | Short name | Use it for |
| --- | --- | --- |
| [`lang.Localize`](https://pkg.go.dev/github.com/belluikit/bellui/lang#Localize) | `lang.L` | Text whose ID is the English text itself. |
| [`lang.LocalizePlural`](https://pkg.go.dev/github.com/belluikit/bellui/lang#LocalizePlural) | `lang.N` | Text with plural forms, chosen by a count. |
| [`lang.LocalizeKey`](https://pkg.go.dev/github.com/belluikit/bellui/lang#LocalizeKey) | `lang.X` | Text looked up by a separate ID, with a fallback. |
| [`lang.LocalizePluralKey`](https://pkg.go.dev/github.com/belluikit/bellui/lang#LocalizePluralKey) | `lang.XN` | Plural text looked up by ID, with a fallback. |

Template data is the last argument: a map, or a struct with exported fields.

**Plural forms.** The count picks the form. Which forms a language has comes
from its CLDR plural rules (the Unicode project's data on how each language
counts): English has `one` and `other`, Polish and Russian also have `few`
and `many`. Write the forms your language needs, and always include `other`.
The count only chooses the form. To show the number, put it in the template
data as well, as the example does with `Count`.

**IDs or English text.** With `L` and `N` the ID is the English text, so
the source reads naturally. `L` needs no English file. `N` without an English
file uses the ID for every count, so a count of 1 reads "1 files selected".
Add an English file with the `one` form for it. With `X` and `XN` the ID
is a short key such as `button.save`, and you pass the English text as the
fallback. Keys stay stable when you reword the English.

## Grammatical variants

Some sentences change with something other than a count, such as the
grammatical gender of a person. Write one message per variant, with the
variant name after a dot, plus an `.other` variant:

```json
{
  "profile_updated.female": "{{.Name}} updated her profile",
  "profile_updated.male": "{{.Name}} updated his profile",
  "profile_updated.other": "{{.Name}} updated their profile"
}
```

[`lang.LSelect`](https://pkg.go.dev/github.com/belluikit/bellui/lang#LSelect)
picks the variant. An unknown selector uses `.other`. If that is missing too,
it returns the prefix itself.
[`lang.LSelectPlural`](https://pkg.go.dev/github.com/belluikit/bellui/lang#LSelectPlural)
picks a variant and then a plural form inside it, so each variant is a
plural object.

## When a translation is missing

Bellui chooses one language for the whole app. If your app has no file for
that language, your messages are not found:

- `L` returns the text you wrote, which is English.
- `X` and `XN` return the fallback you passed.
- `N` and `LSelect` return the message ID itself, such as `files_selected`.

Bellui also logs a "Translation failure" error for each missing message.

This happens whenever the user's language is one of the 15 languages Bellui
itself is translated into but your app is not. A French user of an app with
only English and German files sees Bellui's own text in French and your text
in English, or as IDs where you used `N` or `LSelect`. To avoid that, ship
a file for every language you support, and use `XN` with a fallback rather
than `N` with a key.

The full program:

{{< example "localisation/translations/main.go" >}}
