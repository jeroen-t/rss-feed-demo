package main

import (
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}

	f := app.parseFeed()
	app.render(w, r, "home.page.tmpl", &templateData{
		Feeds: f,
	})
}
