package main

import (
	"log"
	"net/http"
)

// This is a handler function which writes a byte slice
// containing "Hello from Snippetbox" as the response body.
// Byte slices are a list of bytes that represent
// UTF-8 encodings of Unicode code points
// See -> https://medium.com/@tyler_brewer2/bits-bytes-and-byte-slices-in-go-8a99012dcc8f
func home(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Hello from Snippetbox"))
}

func main() {
	//Initialize a new NewServeMux, this is a router that is responsible for
	//registering our routes or paths with their respective handler functions
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)

	log.Println("Starting server on :4000")
	//Start a new web server on port 4000
	err := http.ListenAndServe(":4000", mux)

	//If there is an error we log fatal an exit
	//Errors returned by http.ListenAndServe are always non nil
	//log.Fatal is a shortcut for log.Print(v); os.Exit(1)
	// os.Exit exits the go program immediately
	log.Fatal(err)
}
