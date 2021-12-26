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
var configfile *string

// init does basic initialization of the KnowledgeCat applications - establishes connection
// with databases, performs migrations if needed.
func init() {
	datadir = flag.String("datadir", "data/", "Select location of data directory.")
	bindaddr = flag.String("bindaddr", "0.0.0.0:8000", "Bind address (with port)")
	configfile = flag.String("configfile", "config.toml", "Location of configuration file")
}

// main() is main entrypoint for application.
func main() {
	fmt.Println("Data directory:", *datadir)
	fmt.Println("Application will listen on:", *bindaddr)
	r := mux.NewRouter()
	http.Handle("/", r)
	srv := &http.Server{
		Handler:      r,
		Addr:         *bindaddr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
