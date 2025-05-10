package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	srv := http.Server{
		Addr:    ":3000",
		Handler: Routes(),
	}

	log.Println("Listening on port :3000")
	err := srv.ListenAndServe()
	if err != nil {
		log.Panic(err)
	}
}

func Routes() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/inlägg", articlesPage)
	mux.HandleFunc("/inlägg/{id}", articlePage)
	mux.HandleFunc("/exempel", exampleEditPage)

	fs := http.FileServer(http.Dir("./src/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	return mux
}

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/html/base.html", "src/html/index.html"))
	tmpl.Execute(w, nil)
}

func articlesPage(w http.ResponseWriter, r *http.Request) {
	a := getArticles()
	tmpl := template.Must(template.ParseFiles("src/html/base.html", "src/html/articles.html"))
	tmpl.Execute(w, a)
}

func articlePage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	a := getArticle(id)
	tmpl := template.Must(template.ParseFiles("src/html/base.html", "src/html/article.html"))
	tmpl.Execute(w, a)
}

func exampleEditPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/html/base.html", "src/html/example/example-edit.html"))
	tmpl.Execute(w, nil)
}
