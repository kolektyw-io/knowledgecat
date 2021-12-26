package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"path"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	namespace := vars["namespace"]
	file := vars["title"]

	data, err := os.ReadFile(path.Join(*datadir, namespace, file+".md"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Println(err)
		_, _ = fmt.Fprintf(w, "Article not found")
	}

	_, _ = fmt.Fprintf(w, string(data))
}

func HomeHandler(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
	data, err := os.ReadFile(path.Join(*datadir, "index.md"))
	if err != nil {
		fmt.Println(err)
	}

	_, _ = fmt.Fprintf(w, string(data))
}
