package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/html/base.html", "src/html/elements/nav.html", "src/html/index.html"))
	tmpl.Execute(w, nil)
}

func articlesPage(w http.ResponseWriter, r *http.Request) {
	a := getArticles()
	tmpl := template.Must(template.ParseFiles("src/html/base.html", "src/html/elements/nav.html", "src/html/articles.html"))
	tmpl.Execute(w, a)
}

func articlePage(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	a := getArticle(id)
	tmpl := template.Must(template.ParseFiles("src/html/base.html", "src/html/elements/nav.html", "src/html/article.html"))
	tmpl.Execute(w, a)
}

func aboutPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/html/base.html", "src/html/elements/nav.html", "src/html/about.html"))
	tmpl.Execute(w, nil)
}

func exampleEditPage(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("src/html/base.html", "src/html/elements/nav.html", "src/html/example/example-edit.html"))
	tmpl.Execute(w, nil)
}
