package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"text/template"
)

var feedURLs = []string{
	"https://azurecomcdn.azureedge.net/en-us/updates/feed/",
}

type application struct {
	errorLog      *log.Logger
	infoLog       *log.Logger
	templateCache map[string]*template.Template
}

func main() {
	port := flag.String("port", ":4000", "port on which the server will listen")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	templateCache, err := newTemplateCache("./ui/html/")
	if err != nil {
		errorLog.Fatal(err)
	}

	app := &application{
		errorLog:      errorLog,
		infoLog:       infoLog,
		templateCache: templateCache,
	}

	srv := &http.Server{
		Addr:     *port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on http://localhost%s", *port)
	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}
