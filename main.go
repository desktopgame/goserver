package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/main", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "web/views/MainPage.html")
	})
	fs := http.FileServer(http.Dir("web/dist/js/"))
	r.Handle("/js/*", http.StripPrefix("/js/", fs))
	http.ListenAndServe(":3000", r)
}
