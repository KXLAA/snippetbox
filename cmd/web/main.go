package main

import (
	"flag"
	"log"
	"net/http"
	"os"
)

func main() {
	//Command line flags
	addr := flag.String("addr", ":4000", "HTTP network address") //Local Host address
	flag.Parse()

	//Custom loggers
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	//Router
	mux := http.NewServeMux()

	//Handlers
	mux.HandleFunc("/", home)
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	//Static File server
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))

	//Necessary to use our error log when Go’s HTTP server
	//encounters an error instead of the standard logger
	server := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  mux,
	}

	infoLog.Printf("Starting server on %s", *addr)
	err := server.ListenAndServe()
	errorLog.Fatal(err)
}
