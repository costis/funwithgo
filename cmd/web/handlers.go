package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
)

func (a *application)showHome(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allowed", "GET")
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Not allowed"))
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/home.page.gohtml",
		"./ui/html/base.layout.gohtml",
		"./ui/html/footer.gohtml",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		a.infoLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		a.infoLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (a *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allowed", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Not allowed"))
		return
	}

	w.Write([]byte("Creating a snippet"))
}

func (a *application) showSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allowed", http.MethodGet)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Not allowed"))
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if id < 1 || err != nil {
		http.NotFound(w, r)
	}

	w.Write([]byte(fmt.Sprintf("Showing snippet with id: %d", id)))
}
