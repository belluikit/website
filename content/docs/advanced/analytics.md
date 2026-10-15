---
title: Analytics
description: Opt-in product analytics that records nothing until the user consents, batches events and survives being offline.
weight: 2
---

The [`analytics`](https://pkg.go.dev/github.com/belluikit/bellui/analytics)
package records events such as "export clicked" and sends them in batches to a
backend you choose. It is built around consent: a tracker records nothing
until your app turns it on, and it never invents an identifier for the device
or the user.

The package is optional. Bellui itself does not import it, so an app that
does not use it carries none of its code.

## Set up a tracker

{{< example file="advanced/analytics/main.go" region="setup" >}}

- [`analytics.New`](https://pkg.go.dev/github.com/belluikit/bellui/analytics#New)
  creates a tracker that delivers to a sink (the backend adapter, see below).
  The tracker starts **disabled**.
- `SetDefault` makes it the tracker behind the package-level functions
  `analytics.Track`, `Identify`, `Flush` and `SetEnabled`. Before you install
  one, those functions do nothing, so they are always safe to call.
- [`Attach`](https://pkg.go.dev/github.com/belluikit/bellui/analytics#Attach)
  connects the tracker to the app's lifecycle. It records `session_start`
  when the app starts, records `session_end` and sends the queue when the app
  goes to the background, and sends the queue again when the app stops. On
  the desktop, "goes to the background" means the last of the app's windows
  lost focus, so switching to another app also records `session_end`.
- `EnableAutoCapture` is optional. It records taps without any code per
  widget; see [Automatic tap capture](#automatic-tap-capture).

## Ask for consent

Nothing is recorded until `SetEnabled(true)`. Tie it to a setting the user
controls:

{{< example file="advanced/analytics/main.go" region="consent" >}}

While the tracker is disabled:

- `Track` and `Identify` drop their events. Nothing is queued in memory,
  written to disk or sent.
- No identifier exists. The package never creates an advertising ID, a
  fingerprint or an install ID.

Turning it off again with `SetEnabled(false)` drops the queue and deletes the
copy on disk. A real app should save the user's choice with its other
preferences and call `SetEnabled` with it at startup.

## Record events

{{< example file="advanced/analytics/main.go" region="track" >}}

`Track` takes an event name and a map of properties. To tie later events to a
user, call `Identify` with an ID your app owns, for example its own account
ID, and optional traits. Call it only once the user has consented. The tracker
adds that ID to every following event.

Each [`Event`](https://pkg.go.dev/github.com/belluikit/bellui/analytics#Event)
also carries a `SessionID`. It is random, made fresh every time the app
starts, and never saved as an identity.

## When events are sent

`Track` never touches the network. It adds the event to a queue and returns.
The queue is sent in one request when any of these happens:

| Trigger | Default | Change it with |
| --- | --- | --- |
| The queue reaches a size | 30 events | `WithBatchSize` |
| A timer fires while the app is in the foreground | every 60 seconds | `WithInterval` |
| The app goes to the background | always, with `Attach` | |
| You call `Flush` | | |

The timer stops while the app is in the background. On a phone the radio is
expensive to wake, so sending rarely and in batches saves battery.

When sending fails, for example offline, the events go back into the queue.
The tracker waits before trying again, doubling the wait each time from
5 seconds up to 10 minutes, with some randomness (`WithBackoff`). Being offline
is normal and is not reported as an error.

The queue is saved to disk after every change, so events survive the app
being killed. They are sent, oldest first, the next time the tracker is
enabled. The saved queue holds at most 1000 events (`WithMaxSpool`); beyond
that the oldest are dropped.

{{< callout type="info" >}}
Create the tracker after `app.New()`. The queue is saved in the app's
storage only if an app exists when `New` runs. Otherwise it lives in memory
and is lost when the app exits.
{{< /callout >}}

## Choose a backend

A [`Sink`](https://pkg.go.dev/github.com/belluikit/bellui/analytics#Sink)
delivers one batch per call. Two are included:

- [`analytics/posthog`](https://pkg.go.dev/github.com/belluikit/bellui/analytics/posthog):
  `posthog.New(apiKey)` for PostHog's US cloud, or
  `posthog.NewWithHost(apiKey, host)` for the EU cloud or a self-hosted
  instance. `Identify` becomes a `$identify` event.
- [`analytics/ga4`](https://pkg.go.dev/github.com/belluikit/bellui/analytics/ga4):
  `ga4.New(measurementID, apiSecret)` for the Google Analytics 4 Measurement
  Protocol. It sends at most 25 events per request and splits larger batches.

Both need an ID for the client. They use the ID you passed to `Identify` and
fall back to the session ID, so they never create one either.

Pick one at startup, for example from configuration:

{{< example file="advanced/analytics/main.go" region="choose" >}}

To use another backend, write your own sink. `Send` gets the batch and returns
an error to have it retried later. This one just prints:

{{< example file="advanced/analytics/main.go" region="sink" >}}

For an HTTP backend,
[`analytics.PostJSONGzip`](https://pkg.go.dev/github.com/belluikit/bellui/analytics#PostJSONGzip)
encodes a value as JSON, gzips it and posts it in one request. It returns an
error for any status outside 2xx and gives up after 30 seconds.

## Automatic tap capture

`EnableAutoCapture(t)` records a `ui.activate` event each time the user
activates a control: a tap or click, a screen-reader action, or Space or Enter
on a focused button-like widget. The event's properties come from the
accessibility information the widget already provides:

| Property | Value |
| --- | --- |
| `role` | the widget's role, such as `button` or `checkbox` |
| `verb` | `press` for buttons, links and menu items; `toggle` for check boxes and switches; `select` for list items, tree items, tabs and table cells |
| `label` | the text a screen reader would read for it, if any |
| `state` | flags such as `checked` or `disabled`, if any are set |
| `via` | `keyboard` when activated from the keyboard |

Other roles, such as labels and text fields, are not recorded. Auto-capture
never records what a user typed or a widget's value. A password field is
skipped completely. Every event goes through `Track`, so the consent rule
applies to it too.

To exclude a region, such as an account or payment form, wrap it:

{{< example file="advanced/analytics/main.go" region="nocapture" >}}

[`NoCapture`](https://pkg.go.dev/github.com/belluikit/bellui/analytics#NoCapture)
changes nothing about layout, look or input. `WithSampling(n)` records one in
`n` of the eligible events, and `DisableAutoCapture()` turns it off.

## Gotchas

- **`Attach` takes over the lifecycle hooks.** It sets all four callbacks on
  `a.Lifecycle()`, replacing any you set before. If you set your own after
  `Attach`, you replace the tracker's.
- **`Close` does not send.** Call `Flush` first if you want the queue
  delivered before the tracker stops.
- **Labels are recorded as-is.** If a button's text contains something
  personal, such as a person's name, put it inside `NoCapture`.
