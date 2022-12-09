package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	//Command line flags
	addr := flag.String("addr", ":4000", "HTTP network address") //Local Host address
	flag.Parse()

	//Router
	mux := http.NewServeMux()

	//Handlers
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//Static File server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	log.Printf("Starting server on %s", *addr)
	err := http.ListenAndServe(*addr, mux)
	log.Fatal(err)
}
