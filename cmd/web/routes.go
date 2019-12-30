package main

import "net/http"

func (app *application) routes() *http.ServeMux {
	fileServer := http.FileServer(http.Dir("./ui/static"))
	mux := http.NewServeMux()
	mux.HandleFunc("/", app.showHome)
	mux.HandleFunc("/snippet/create", app.createSnippet)
	mux.HandleFunc("/snippet", app.showSnippet)
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
