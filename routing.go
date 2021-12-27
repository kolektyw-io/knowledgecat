package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func generateRoutes(r *mux.Router) {
	r.HandleFunc("/{namespace}/{title}", ArticleHandler)
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/{namespace}/", NamespaceHandler)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
}
