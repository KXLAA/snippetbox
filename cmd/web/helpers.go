package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

// The serverError helper writes an error message and stack trace to the errorLog,
// then sends a generic 500 Internal Server Error response to the user.
func (app *application) serverError(response http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())

	//make sure error is logged from function it originated from
	app.errorLog.Output(2, trace)

	http.Error(response, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// The clientError helper sends a specific status code and corresponding description
// to the user. We'll use this  to send responses like 400 "Bad
// Request" when there's a problem with the request that the user sent.

func (app *application) clientError(response http.ResponseWriter, status int) {
	http.Error(response, http.StatusText(status), status)
}

// For consistency, we'll also implement a notFound helper. This is simply a
// convenience wrapper around clientError which sends a 404 Not Found response to
// the user.
func (app *application) notFound(response http.ResponseWriter) {
	app.clientError(response, http.StatusNotFound)
}

func (app *application) render(response http.ResponseWriter, request *http.Request, name string, templateData *templateData) {
	template, ok := app.templateCache[name]
	if !ok {
		app.serverError(response, fmt.Errorf("the template %s does not exist", name))
		return
	}

	err := template.Execute(response, templateData)
	if err != nil {
		app.serverError(response, err)
	}

}
