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

	//As / is a subtree path which matches every URL path we
	//need to restrict it to only matching "/" by doing this check
	if request.URL.Path != "/" {
		http.NotFound(response, request)
		return
	}

	response.Write([]byte("Hello from Snippetbox"))
}

// Add a showSnippet handler function.
func showSnippet(response http.ResponseWriter, request *http.Request) {
	response.Write([]byte("Display a specific snippet..."))
}

// Add a createSnippet handler function.
func createSnippet(response http.ResponseWriter, request *http.Request) {

	//Need to restrict calls to this handler to POST requests only
	if request.Method != http.MethodPost {

		// Let the user know what method is allowed by setting the headers of the return
		response.Header().Set("Allow", http.MethodPost)

		//if the request is not POST we send a 405 status code "Not Allowed" & a message body
		http.Error(response, "Method Not Allowed", http.StatusMethodNotAllowed)

		//return so that subsequent code in the body is not executed
		return

	}
	response.Write([]byte("Create a new snippet..."))
}

func main() {
	//Initialize a new NewServeMux, this is a router that is responsible for
	//registering our routes or paths with their respective handler functions
	mux := http.NewServeMux()

	//This path is a subtree path (because it ends in a trailing slash)
	//Another example of a subtree path would be "/static/"
	//These kinds of paths are matched when the start or a request URL
	//path matches the subtree path, these paths acts like they
	//have a wild card at the end eg "/**" or "/static/**"
	mux.HandleFunc("/", home)

	//Our paths here are fixed paths, Go's servemux will only
	//match these paths and call the handlers when the URL path
	//exactly matches the fixed path
	mux.HandleFunc("/snippet", showSnippet)
	mux.HandleFunc("/snippet/create", createSnippet)

	log.Println("Starting server on :4000")
	//Start a new web server on port 4000
	err := http.ListenAndServe(":4000", mux)

	//If there is an error we log fatal an exit
	//Errors returned by http.ListenAndServe are always non nil
	//log.Fatal is a shortcut for log.Print(v); os.Exit(1)
	// os.Exit exits the go program immediately
	log.Fatal(err)
}
