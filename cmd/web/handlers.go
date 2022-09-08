package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	var feedUrls = []string{
		"https://www.jeroentrimbach.com/index.xml",
		"https://azurecomcdn.azureedge.net/en-us/updates/feed/",
	}

	f := app.parseFeeds(feedUrls)
	app.render(w, r, "home.page.tmpl", &templateData{
		Feeds: f,
	})
}
