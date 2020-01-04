package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"html/template"
	"log"
	"mrmambo.dev/snippetbox/pkg/models/mysql"
	"net/http"
	"os"
)

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippets *mysql.SnippetModel
	articles *mysql.Articles
	tplCache map[string]*template.Template
}

func main() {
	port := flag.String("port", ":4000", "Listening port")
	dsn := flag.String("dsn", "root:password@/snippetbox?parseTime=true", "MySQL DSN")
	flag.Parse()

	errorLog := log.New(os.Stdout, "INFO ", log.Ldate|log.Ltime)
	infoLog := log.New(os.Stderr, "INFO ", log.Ldate|log.Lshortfile)

	db, err := openDB(*dsn)
	if err != nil {
		errorLog.Fatal(err)
	}
	defer db.Close()

	cache, err := NewTemplateCache("./ui/html")
	if err != nil {
		panic(err)
	}

	app := &application{
		errorLog: errorLog,
		infoLog:  infoLog,
		snippets: &mysql.SnippetModel{DB: db},
		articles: &mysql.Articles{DB: db},
		tplCache: cache,
	}

	infoLog.Printf("Starting web server on port %s", *port)

	srv := &http.Server{
		Addr:     *port,
		Handler:  app.routes(),
		ErrorLog: errorLog,
	}

	err = srv.ListenAndServe()
	errorLog.Fatal(err)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
