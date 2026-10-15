// A web view showing a local page that talks to the app over the bridge.
package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"strings"

	"github.com/belluikit/bellui"
	"github.com/belluikit/bellui/app"
	"github.com/belluikit/bellui/container"
	"github.com/belluikit/bellui/geo"
	"github.com/belluikit/bellui/webview"
	"github.com/belluikit/bellui/widget"
)

//docs:start page
const page = `<!doctype html>
<html><body style="font-family: sans-serif">
<input id="name" value="Ada"> <button id="hi">Greet</button>
<p id="out"></p>
<script>
  document.getElementById("hi").onclick = async function () {
    var name = document.getElementById("name").value;
    document.getElementById("out").textContent =
      await window.bellui.call("greet", {name: name});
  };
  window.bellui.handle("title", function () { return document.title || "untitled"; });
</script>
</body></html>`

//docs:end page

func main() {
	a := app.New()
	w := a.NewWindow("Web view")

	status := widget.NewLabel("")

	//docs:start host
	host, view := webview.NewHost("data:text/html," + url.PathEscape(page))
	host.SetMinSize(geo.NewSize(400, 300))
	//docs:end host

	//docs:start bridge
	// A data: URL has no origin of its own, so trust exactly that.
	view.EnableBridge(webview.TrustOrigins(webview.OriginOpaque))

	view.Handle("greet", func(m webview.Message) (any, error) {
		var arg struct{ Name string }
		if err := m.Decode(&arg); err != nil {
			return nil, err
		}
		if arg.Name == "" {
			return nil, errors.New("name is empty")
		}
		return "Hello, " + arg.Name + "!", nil
	})
	//docs:end bridge

	//docs:start navigate
	// Keep the view on local content: refuse links to anywhere else.
	view.OnNavigate = func(r webview.NavigationRequest) bool {
		return strings.HasPrefix(r.URL, "data:")
	}
	view.OnLoadChanged = func(e webview.LoadEvent) {
		status.SetText("Page: " + e.State.String())
	}
	//docs:end navigate

	//docs:start call
	ask := widget.NewButton("Ask the page for its title", func() {
		view.Call("title", nil, func(result json.RawMessage, err error) {
			if err != nil {
				log.Println("call failed:", err)
				return
			}
			status.SetText("Page title: " + string(result))
		})
	})
	//docs:end call

	title := widget.NewEntry()
	title.SetPlaceHolder("New page title")

	//docs:start evalfunc
	setTitle := widget.NewButton("Set page title", func() {
		// The entry's text arrives in the page as data, never as code.
		view.EvalFunc("document.title = title;", webview.Args{"title": title.Text})
	})
	//docs:end evalfunc

	if !webview.Supported {
		status.SetText("No web engine on this platform")
	}

	w.SetContent(container.NewBorder(container.NewVBox(ask, container.NewBorder(nil, nil, nil, setTitle, title), status), nil, nil, nil, host))
	w.Resize(geo.NewSize(640, 480))
	bellui.ShowAndRun(w)
}
