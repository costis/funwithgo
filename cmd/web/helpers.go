package main

import "net/http"
import "runtime/debug"

func (app *application) clientError(w http.ResponseWriter, statusCode int, message string) {
	app.infoLog.Println("Client error occurred...")
	w.WriteHeader(statusCode)
	w.Write([]byte(http.StatusText(statusCode)))
}

func (app *application) clientErrorNotAllowed(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
}

func (app *application) clientErrorNotFound(w http.ResponseWriter) {
	http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
}

func (app *application) serverError(w http.ResponseWriter, err error) {
	app.errorLog.Println("Server error occurred", err, string(debug.Stack()))
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(http.StatusText(http.StatusInternalServerError)))
}

func (app *application) notFound(w http.ResponseWriter, msg string) {
	app.clientError(w, http.StatusNotFound, msg)
}
