package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/mmcdole/gofeed"
)

type feedItem struct {
	SourceTitle   string
	SourceUrl     string
	PublishedDate string
	ArticleTitle  string
	ArticleUrl    string
	Author        string
}

func (app *application) parseFeed(url string) *gofeed.Feed {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(url)
	return feed
}

func (app *application) parseFeeds(urls []string) []feedItem {
	feedItems := []feedItem{}
	for i := 0; i < len(urls); i++ {
		f := app.parseFeed(urls[i])
		for j := 0; j < len(f.Items); j++ {
			feedItems = append(feedItems, feedItem{
				SourceTitle:   f.Title,
				SourceUrl:     f.FeedLink,
				PublishedDate: f.Items[j].Published,
				ArticleTitle:  f.Items[j].Title,
				ArticleUrl:    f.Items[j].Link,
			})
		}
	}

	return feedItems
}

func (app *application) addDefaultData(td *templateData, r *http.Request) *templateData {
	if td == nil {
		td = &templateData{}
	}
	td.CurrentYear = time.Now().Year()
	return td

}

func (app *application) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		app.serverError(w, fmt.Errorf("The template %s does not exist", name))
		return
	}

	buf := new(bytes.Buffer)

	err := ts.Execute(buf, app.addDefaultData(td, r))
	if err != nil {
		app.serverError(w, err)
		return
	}

	buf.WriteTo(w)
}
