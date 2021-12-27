package main

import (
	"fmt"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/html"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
	"os"
	"path"
)

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	namespace := vars["namespace"]
	file := vars["title"]
	context := make(map[string]interface{})
	prepareDefaultContext(r, context)

	// FIXME: This is realy ugly but works for now
	articles := getArticles(namespace)
	arts := make([]Article, 0)
	for _, n := range articles {
		arts = append(arts, Article{
			File:  n.File[0 : len(n.File)-3],
			Title: n.Title,
		})
	}
	context["articles"] = arts

	articleText, err := os.ReadFile(path.Join(*datadir, namespace, file+".md"))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		// FIXME: Add 404 handling
		fmt.Println(err)
	}

	htmlFlags := html.SkipHTML | html.CommonFlags
	opts := html.RendererOptions{Flags: htmlFlags}
	renderer := html.NewRenderer(opts)
	context["content"] = template.HTML(markdown.ToHTML(articleText, nil, renderer))

	t, _ := template.ParseFiles("templates/article.html")
	err = t.Execute(w, context)

	if err != nil {
		fmt.Println(err)
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	context := make(map[string]interface{})
	prepareDefaultContext(r, context)
	context["namespaces"] = retrieveAvailableNamespaces()
	w.WriteHeader(http.StatusOK)
	t, _ := template.ParseFiles("templates/index.html")
	err := t.Execute(w, context)

	if err != nil {
		fmt.Println(err)
	}
}

func NamespaceHandler(w http.ResponseWriter, r *http.Request) {
	context := make(map[string]interface{})
	prepareDefaultContext(r, context)
	vars := mux.Vars(r)
	namespace := vars["namespace"]
	context["namespaces"] = retrieveAvailableNamespaces()
	articles := getArticles(namespace)
	arts := make([]Article, 0)
	for _, n := range articles {
		arts = append(arts, Article{
			File:  n.File[0 : len(n.File)-3],
			Title: n.Title,
		})
	}
	context["articles"] = arts

	w.WriteHeader(http.StatusOK)
	t, _ := template.ParseFiles("templates/category.html")
	err := t.Execute(w, context)

	if err != nil {
		fmt.Println(err)
	}
}
