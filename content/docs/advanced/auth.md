---
title: Sign-in with OIDC and OAuth
description: Let users sign in to a desktop app with Google, Microsoft, Apple, GitHub, Okta, Keycloak or any OpenID Connect provider.
weight: 1
---

The [`auth`](https://pkg.go.dev/github.com/belluikit/bellui/auth) package signs
a user in through their identity provider, using the browser they already use.
It returns the user's name, email and tokens. The
[`widget.AuthButton`](https://pkg.go.dev/github.com/belluikit/bellui/widget#AuthButton)
wraps the whole flow in a branded button.

The flow is OAuth 2.0 Authorization Code with PKCE (a one-time secret that
proves the app that started the sign-in is the one finishing it). It is the
flow intended for desktop apps. The package is written for desktop: it opens
the system browser and listens on a local port.

## Add a sign-in button

{{< example file="advanced/auth/main.go" region="button" >}}

`NewAuthButton` takes a provider key, the client ID your provider gave you,
and a function to call once the user is signed in. When the user taps the
button:

1. The default browser opens the provider's sign-in page.
2. The user signs in there.
3. The provider sends the browser back to the app, and the app finishes the
   sign-in in the background.
4. Your function runs on the UI thread with an
   [`auth.OIDCResult`](https://pkg.go.dev/github.com/belluikit/bellui/auth#OIDCResult).
   It can update widgets directly.

If the sign-in fails, `OnError` runs instead, also on the UI thread. A tap
while a sign-in is already in progress is ignored.

The button draws the provider's logo and brand colors. Set `Text` or `Logo`
to override them. `Scopes` is a space-separated list, for example
`"openid email profile"`, and `Issuer` sets the provider address when the key
needs one (see below).

## Providers

| Key | Provider | What you set |
| --- | --- | --- |
| `google` | Google | Client ID |
| `microsoft` | Microsoft (Entra ID), multi-tenant endpoint | Client ID |
| `apple` | Apple | Client ID |
| `github` | GitHub (OAuth 2.0, not OpenID Connect) | Client ID, and scopes (see below) |
| `okta` | Okta | Client ID and `Issuer`, your Okta address |
| `keycloak` | Keycloak | Client ID and `Issuer`, your realm address |
| any other key, such as `custom` | any OpenID Connect provider | Client ID and `Issuer` |

The logos, colors and default labels live in
[`auth.Providers`](https://pkg.go.dev/github.com/belluikit/bellui/auth#Providers).
Keycloak has no logo there. An unknown key gets a grey button labelled
"Sign in".

## What you register with the provider

Create an app registration (often called a "client") of the native, desktop
or public kind. Then check three things:

- **Redirect URI.** The app listens on `http://127.0.0.1:<port>/callback`,
  and the operating system picks a free port for each sign-in. The provider
  must accept a loopback redirect on any port.
- **No client secret.** The app sends only the client ID and the PKCE
  verifier when it exchanges the code for tokens. A registration that
  requires a client secret will reject the exchange.
- **Scopes.** OpenID Connect providers get `openid email profile` unless you
  set others.

The client ID is not a secret. It is fine to compile it into the app.

## How the sign-in is checked

For every provider, including GitHub, the package sends a random `state`
value with the request and rejects a reply that does not return it.

For every provider except GitHub, the package also:

- reads the provider's discovery document from
  `<issuer>/.well-known/openid-configuration`, and rejects it if the issuer
  inside does not match the one you asked for;
- verifies the ID token's signature against the provider's published keys,
  and checks its issuer, audience (your client ID), expiry and nonce.

The browser shows a short "Authentication Successful" page that tells the user
to return to the app. The whole flow times out after five minutes.

## GitHub

GitHub uses plain OAuth 2.0, so there is no ID token to verify. After the code
exchange the package asks the GitHub API for the user. In the result, `Sub` is
the numeric GitHub user ID, and `Name` falls back to the login when the user
has no display name. Of the three token fields, only `AccessToken` is set:
GitHub sign-ins have no refresh token and no ID token.

With no `Scopes` set, GitHub is asked for `user:email read:user`. (Before
v1.1.0 it was sent the OpenID Connect defaults instead.)

## Without the button

Call [`auth.Authenticate`](https://pkg.go.dev/github.com/belluikit/bellui/auth#Authenticate)
when you want your own UI. It blocks until the sign-in finishes, fails or
times out, so run it in a goroutine and come back to the UI thread with
`bellui.Do`:

{{< example file="advanced/auth/main.go" region="direct" >}}

[`auth.OIDCConfig`](https://pkg.go.dev/github.com/belluikit/bellui/auth#OIDCConfig)
also has `Scopes` (a slice) and `Timeout` (five minutes when zero).

## Remembering the user

The button saves the refresh token when the provider returns one. You can read
it back at the next launch:

{{< example file="advanced/auth/main.go" region="stored" >}}

[`LoadToken`](https://pkg.go.dev/github.com/belluikit/bellui/auth#LoadToken)
returns empty strings, not an error, when nothing is stored. If you call
`Authenticate` yourself, store the token with
[`SaveToken`](https://pkg.go.dev/github.com/belluikit/bellui/auth#SaveToken).
Tokens are stored per app, so the app needs a unique ID: create it with
`app.NewWithID`, or give it an ID in `manifest.pxf`. Without one,
`SaveToken`, `LoadToken` and `ClearToken` return
[`auth.ErrNoAppID`](https://pkg.go.dev/github.com/belluikit/bellui/auth#ErrNoAppID).
To sign out, remove the token:

{{< example file="advanced/auth/main.go" region="signout" >}}

How the token is stored:

- It is encrypted with AES-256-GCM in a file named `<provider>.enc`, in a
  folder named after the app's ID, inside a `bellui-auth` folder under the
  user's config directory (the one Go's `os.UserConfigDir` returns). Two apps
  that sign in with the same provider keep separate tokens, and the app ID is
  part of the encryption key, so one app cannot read another's token.
- The encryption key comes from a secret kept in the operating system's
  secret store: the Keychain on macOS, DPAPI on Windows, and the Secret
  Service (through the `secret-tool` command) on Linux.
- With no secret store available, for example on a headless Linux machine,
  the secret is a random file in the same folder, readable only by the user.
- Token files do not work on another machine. A file that cannot be decrypted
  is deleted, and the user simply signs in again.
- Versions before v1.1.0 kept one token per provider for all apps, directly in
  `bellui-auth`. Those tokens are not read any more, so users sign in once more
  after upgrading. From v1.2.0, the old file for a provider
  (`bellui-auth/<provider>.enc`) is deleted the first time any app saves or
  clears a token for that provider.

{{< callout type="info" >}}
The package does not use the refresh token for you. It stores it and gives it
back. Exchanging it for a new access token is up to your app.
{{< /callout >}}

## Gotchas

- **`CancelAuth` only stops the button.** The button stops waiting and reports
  nothing. The local listener keeps waiting for the browser until the
  five-minute timeout.
- **The user must finish in the browser.** Nothing in the app window changes
  while the browser page is open. Show a hint such as "Continue in your
  browser" if the wait could confuse people.
