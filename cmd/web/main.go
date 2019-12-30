package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "Listening port")
	flag.Parse()

	errorLog := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	infoLog := log.New(os.Stderr, "INFO ", log.Ldate|log.Lshortfile)

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
	}

	infoLog.Printf("Starting web server on port %s", *addr)

	srv := &http.Server{
		Addr:     *addr,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
