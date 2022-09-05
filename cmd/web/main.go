package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

var feedURLs = []string{
	"https://azurecomcdn.azureedge.net/en-us/updates/feed/",
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	port := flag.String("port", ":4000", "port on which the server will listen")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	srv := &http.Server{
		Addr:     *port,
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Printf("Starting server on %s", *port)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
