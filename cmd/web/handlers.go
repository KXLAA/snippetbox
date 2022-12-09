package main

import (
	"fmt"
	"net/http"
	"strconv"
	"text/template"
)

func (app *application) home(response http.ResponseWriter, request *http.Request) {
	//Make sure this handler is only executed when url is = "/"
	if request.URL.Path != "/" {
		http.NotFound(response, request)
		return
	}

	//parse the html files
	files := []string{
		//template page for this route, this must come first
		"./ui/html/home.page.html",
		//layout & partial templates
		"./ui/html/base.layout.html",
		"./ui/html/footer.partial.html",
	}
	template, err := template.ParseFiles(files...)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(response, err)
		return
	}

	//Execute the parsed html template with any dynamic data or nil if none
	err = template.Execute(response, nil)
	if err != nil {
		app.errorLog.Println(err.Error())
		app.serverError(response, err)
	}
}

func (app *application) showSnippet(response http.ResponseWriter, request *http.Request) {
	queryId := request.URL.Query().Get("id")
	id, err := strconv.Atoi(queryId)

	if err != nil || id < 1 {
		app.notFound(response)
		return
	}

	fmt.Fprintf(response, "Display a specific snippet with ID %d...", id)
}

func (app *application) createSnippet(response http.ResponseWriter, request *http.Request) {

	if request.Method != http.MethodPost {
		response.Header().Set("Allow", http.MethodPost)
		app.clientError(response, http.StatusMethodNotAllowed)
		return

	}
	response.Write([]byte("Create a new snippet..."))
}
