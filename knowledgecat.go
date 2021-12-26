package main

import (
	"flag"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

// datadir and bindaddr are required runtime parameters.
var datadir *string
var bindaddr *string
var _ *string

// init does basic initialization of the KnowledgeCat applications - establishes connection
// with databases, performs migrations if needed.
func init() {
	fmt.Println("Booting KnowledgeCat...")
	fmt.Println("")
	datadir = flag.String("datadir", "data/", "Select location of data directory.")
	bindaddr = flag.String("bindaddr", "0.0.0.0:8000", "Bind address (with port)")
	_ = flag.String("configfile", "config.toml", "Location of configuration file")
}

// main() is main entrypoint for application.
func main() {
	fmt.Println("KnowledgeCat 0.0.1")
	fmt.Println("==================\n\n")
	fmt.Println("Data directory:", *datadir)
	fmt.Println("Application will listen on:", *bindaddr)
	r := mux.NewRouter()
	r.HandleFunc("/{namespace}/{title}", ArticleHandler)
	r.HandleFunc("/", HomeHandler)
	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         *bindaddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())

}
