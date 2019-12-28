package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "Listening port")
	flag.Parse()

	infoLog := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "INFO ", log.Ldate|log.Lshortfile)
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux := http.NewServeMux()
	mux.HandleFunc("/", showHome)
	mux.HandleFunc("/snippet/create", createSnippet)
	mux.HandleFunc("/snippet", showSnippet)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	infoLog.Printf("Starting web server on port %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	errorLog.Fatal(err)}
