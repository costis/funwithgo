package main

import (
	"errors"
	"fmt"
	"html/template"
	"mrmambo.dev/snippetbox/pkg/models"
	"net/http"
	"strconv"
)

func (app *application) showHome(w http.ResponseWriter, r *http.Request) {
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
		app.infoLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.Execute(w, nil)
	if err != nil {
		app.infoLog.Println(err.Error())
		http.Error(w, "Internal Server Error", 500)
	}
}

func (app *application) createSnippet(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.Header().Set("Allowed", http.MethodPost)
		w.WriteHeader(http.StatusMethodNotAllowed)
		w.Write([]byte("Not allowed"))
		return
	}

	w.Write([]byte("Creating a snippet"))
}

func (app *application) showSnippet(w http.ResponseWriter, r *http.Request) {
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

	snippet, err := app.snippets.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrorNoRecord) {
			app.notFound(w, fmt.Sprintf("No snippet with id %d found", id))
			return
		} else {
			app.serverError(w, err)
			return
		}
	}

	_, err = w.Write([]byte(snippet.String()))
	if err != nil {
		app.serverError(w, err)
	}
}

func (app *application) showArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.clientErrorNotAllowed(w)
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		app.clientErrorNotFound(w)
		return
	}

	article, err := app.articles.Get(id)
	if err != nil {
		if errors.Is(err, models.ErrorNoRecord) {
			app.clientErrorNotFound(w)
			return
		} else {
			app.serverError(w, err)
		}
	}

	//files := []string{
	//	"./ui/html/base.layout.gohtml",
	//	"./ui/html/footer.gohtml",
	//	"./ui/html/home.page.gohtml",
	//}

	//tpl, err := template.ParseFiles(files...)
	//if err != nil {
	//	app.serverError(w, err)
	//	return
	//}

	//tpl.ExecuteTemplate()

	w.Write([]byte(article.String()))
}

func (app *application) showArticles(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		app.clientErrorNotAllowed(w)
		return
	}

	articles, err := app.articles.All()
	if err != nil {
		app.serverError(w, err)
		return
	}

	for _, a := range articles {
		_, err := w.Write([]byte(a.String() + "\n"))
		app.errorLog.Print(err)
	}
}

func (app *application) createArticle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		app.clientErrorNotAllowed(w)
		return
	}

	title := r.FormValue("title")
	slug := r.FormValue("slug")

	err := app.articles.Create(title, slug)
	if err != nil {
		app.serverError(w, err)
	}

	w.Write([]byte("Article inserted\n"))
}
