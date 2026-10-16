---
title: Formats
description: Format and parse dates, times, numbers, money, lists and distances for the user's locale, and sort text the way they expect.
weight: 3
---

Each function in this page uses the current locale, including its region. They
follow `lang.SetLocale` without extra work.

{{< example file="localisation/formats/main.go" region="formats" >}}

## Dates and times

| Function | Returns |
| --- | --- |
| [`FormatDate(t)`](https://pkg.go.dev/github.com/belluikit/bellui/lang#FormatDate) | the short numeric date, such as `03/09/2026` (en-US) or `09.03.2026` (de-DE) |
| [`FormatTime(hour, minute)`](https://pkg.go.dev/github.com/belluikit/bellui/lang#FormatTime) | `2:05 PM` or `14:05` |
| [`FormatDateTime(t)`](https://pkg.go.dev/github.com/belluikit/bellui/lang#FormatDateTime) | both, separated by a space |
| [`FormatMonthYear(t)`](https://pkg.go.dev/github.com/belluikit/bellui/lang#FormatMonthYear) | the translated month name and the year |
| [`DatePatternHint()`](https://pkg.go.dev/github.com/belluikit/bellui/lang#DatePatternHint) | a placeholder for input fields, such as `MM/DD/YYYY` |
| [`WeekStartDay()`](https://pkg.go.dev/github.com/belluikit/bellui/lang#WeekStartDay) | the first day of the week: Sunday in the US, Monday in Germany |

`ParseDate` and `ParseDateTime` read text in the same format back into a
`time.Time`. `ParseDate` accepts single-digit days and months, such as
`1/5/2026`. `DatePattern`, `TimePattern` and `DateTimePattern` return the Go
layout strings, if you want to call `time.Format` yourself.

[`lang.Prefers24HourClock`](https://pkg.go.dev/github.com/belluikit/bellui/lang#Prefers24HourClock)
reports whether the region normally uses a 24-hour clock. Most regions do.
The exceptions include the United States, Canada, Australia, India and
Korea. Use it to decide whether a time input needs an AM/PM switch:

{{< example file="localisation/formats/main.go" region="clock" >}}

Dates use the Gregorian calendar in every locale.

## Numbers and money

- [`FormatDecimal(v, digits)`](https://pkg.go.dev/github.com/belluikit/bellui/lang#FormatDecimal)
  formats a number with the locale's decimal and grouping separators. A
  negative `digits` shows as many decimals as needed and no trailing zeros.
- [`ParseLocaleFloat(s)`](https://pkg.go.dev/github.com/belluikit/bellui/lang#ParseLocaleFloat)
  reads a number typed in the locale's format. See the example below.
- `DecimalSeparator()` and `GroupSeparator()` return the separators.
- [`FormatCurrency(amount, code)`](https://pkg.go.dev/github.com/belluikit/bellui/lang#FormatCurrency)
  formats money with two decimals and the currency's symbol, placed before
  or after the number as the region does: `$19.99`, `19,99 €`. The code is
  an ISO 4217 code such as `"EUR"`.
  [`DefaultCurrency()`](https://pkg.go.dev/github.com/belluikit/bellui/lang#DefaultCurrency)
  returns the region's currency code, and `CurrencySymbol(code)` the symbol.

{{< example file="localisation/formats/main.go" region="parse" >}}

The locale's separators decide the result. Under `de-DE` this text reads as
1234.5. Under `en-US` the same text reads as 1.2345, because there the comma
is the grouping separator and is dropped.

## Lists and units

[`FormatList`](https://pkg.go.dev/github.com/belluikit/bellui/lang#FormatList)
joins items with the locale's separators and conjunction: "red, green, and
blue" in English, "red, green und blue" in German, "red、green、blue" in
Japanese. Use it instead of joining with `", "` and `"and"` yourself.

[`FormatDistance(meters)`](https://pkg.go.dev/github.com/belluikit/bellui/lang#FormatDistance)
and [`FormatMass(grams)`](https://pkg.go.dev/github.com/belluikit/bellui/lang#FormatMass)
take a value in metres or grams and choose the unit. The United States, the
United Kingdom, Liberia and Myanmar get feet and miles, or pounds. Elsewhere
the result is in metres or kilometres, grams or kilograms. Values are rounded
to two decimals.
[`FormatUnit(value, unit)`](https://pkg.go.dev/github.com/belluikit/bellui/lang#FormatUnit)
formats a value in a unit you choose, from metres, kilometres, miles, feet,
grams, kilograms and pounds. The unit's label takes the right plural form,
such as Polish "1 mila", "2 mile", "5 mil".

{{< callout type="warning" >}}
In bellui v1.0.0, `FormatUnit`, `FormatDistance` and `FormatMass` log a "Unit
format failure" error when the value has a fractional part, such as 1.24
miles. The text they return is still correct.
{{< /callout >}}

## Sorting text

Sorting strings with `<` puts them in code-point order, which is wrong for
most languages: "Zoë" before "apple", and accented letters at the end. Use
[`lang.Compare`](https://pkg.go.dev/github.com/belluikit/bellui/lang#Compare):

{{< example file="localisation/formats/main.go" region="sort" >}}

In German, this sorts "Ärger" with the A words and "Österreich" with the O
words. In Swedish, Ä and Ö come after Z. `Compare` is safe to call from any
goroutine and follows `SetLocale`.

- [`lang.CompareFold`](https://pkg.go.dev/github.com/belluikit/bellui/lang#CompareFold)
  ignores case and accents. Use it to match search text.
- [`lang.NewCollator`](https://pkg.go.dev/github.com/belluikit/bellui/lang#NewCollator)
  makes a reusable comparer for sorting large slices. Options include
  `lang.CollateNumeric`, which sorts "file2" before "file10". A collator
  keeps the locale it was made with and must not be shared between
  goroutines.

Sorting by locale adds about 1.75 MB of collation tables to your binary, only
if you call one of these functions.

The full program:

{{< example "localisation/formats/main.go" >}}
