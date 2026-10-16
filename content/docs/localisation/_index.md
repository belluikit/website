---
title: Localisation
description: Translate your app, switch language at runtime, and format dates, numbers and money the way each locale expects.
weight: 7
---

The [`lang`](https://pkg.go.dev/github.com/belluikit/bellui/lang) package
translates text and formats values for the user's **locale**: their language
and region, such as `en-US` or `de-DE`.

Bellui's own text, such as "Cancel" in dialogs, the month names in the date
picker and the items of the text edit menu, is already translated into 15
languages: Chinese (Simplified), Czech, English, French, German, Greek,
Japanese, Polish, Portuguese, Portuguese (Brazil), Russian, Spanish, Swedish,
Tamil and Ukrainian. You add translations for your app's own text.

{{< cards >}}
  {{< card link="translations/" title="Translations" subtitle="Translation files, plurals, keys and grammatical variants." >}}
  {{< card link="switching-language/" title="Switching language" subtitle="Change the locale while the app runs, with text that follows." >}}
  {{< card link="formatting/" title="Formats" subtitle="Dates, times, numbers, money, lists, distances and sorting." >}}
  {{< card link="right-to-left/" title="Right-to-left languages" subtitle="Arabic, Hebrew, Persian and Urdu: mirroring, fonts and text direction." >}}
{{< /cards >}}

## How the locale is chosen

At startup Bellui reads the user's preferred locales from the operating
system.

- **Translations** use the closest language that has translations, among
  Bellui's own and the ones you add. When none matches, English is used.
- **Formats, sorting and text direction** use the full locale, including the
  region. Users with `en-US` and `en-GB` get the same English text, but
  dates as 03/09/2026 and 09/03/2026, and prices in dollars and pounds.

[`lang.SystemLocale`](https://pkg.go.dev/github.com/belluikit/bellui/lang#SystemLocale)
returns the operating system's locale.
[`lang.CurrentLocale`](https://pkg.go.dev/github.com/belluikit/bellui/lang#CurrentLocale)
returns the locale in use, which differs when the app has chosen one with
[`lang.SetLocale`](https://pkg.go.dev/github.com/belluikit/bellui/lang#SetLocale).

On Android, Bellui reads the locale again when the app starts, when it comes
back to the foreground, and when the user changes the app's language in the
system settings, so a language change made while the app was in the
background is picked up.

## Right-to-left languages

A locale written in a right-to-left script, such as Arabic, Hebrew, Persian
or Urdu, also mirrors the layout. This happens automatically: the app does
not need to call anything. Bellui has no translation of its own text into
these languages, and no bundled text font with their letters, so check
[Right-to-left languages](right-to-left/) before you ship one. It also
lists what does not mirror and the limits of mixed-direction text in v1.0.0.
