package main

import (
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
	mux.HandleFunc("/works", articlesPage)
	mux.HandleFunc("/about", aboutPage)
	mux.HandleFunc("/works/{id}", articlePage)
	mux.HandleFunc("/exempel", exampleEditPage)

	fs := http.FileServer(http.Dir("./src/static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	return mux
}
