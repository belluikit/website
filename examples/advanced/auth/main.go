// Sign in with a branded provider button, remember the user, and sign out.
package main

import (
	"log"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/auth"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/widget"
)

// Replace with the client ID your identity provider issued for this app.
const googleClientID = "1234567890-example.apps.googleusercontent.com"

func main() {
	a := app.New()
	w := a.NewWindow("Sign in")

	status := widget.NewLabel("Not signed in")

	//docs:start stored
	// A refresh token saved by an earlier sign-in, if any.
	if refresh, email, err := auth.LoadToken("google"); err != nil {
		log.Println("could not read stored token:", err)
	} else if refresh != "" {
		status.SetText("Previously signed in as " + email)
	}
	//docs:end stored

	//docs:start button
	signIn := widget.NewAuthButton("google", googleClientID, func(r auth.OIDCResult) {
		// Runs on the UI thread, so it can update widgets directly.
		status.SetText("Signed in as " + r.Name + " (" + r.Email + ")")
	})
	signIn.OnError = func(err error) {
		status.SetText("Sign-in failed: " + err.Error())
	}
	//docs:end button

	//docs:start signout
	signOut := widget.NewButton("Sign out", func() {
		signIn.CancelAuth()
		if err := auth.ClearToken("google"); err != nil {
			log.Println("could not clear token:", err)
		}
		status.SetText("Not signed in")
	})
	//docs:end signout

	//docs:start direct
	sso := widget.NewButton("Sign in with company SSO", func() {
		go func() { // Authenticate blocks until the browser flow ends.
			r, err := auth.Authenticate(auth.OIDCConfig{
				Provider: "okta",
				Issuer:   "https://example.okta.com",
				ClientID: "0oa-example-client-id",
			})
			bellui.Do(func() { // Back on the UI thread.
				if err != nil {
					status.SetText("Sign-in failed: " + err.Error())
					return
				}
				status.SetText("Signed in as " + r.Email)
			})
		}()
	})
	//docs:end direct

	w.SetContent(container.NewPadded(container.NewVBox(signIn, sso, signOut, status)))
	bellui.ShowAndRun(w)
}
