package main

import (
	"fmt"
	"net/http"
)

func (app *application) home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		fmt.Fprintf(w, "Not found")
		return
	}

	fmt.Fprintf(w, "Nothing to see here..")
}
