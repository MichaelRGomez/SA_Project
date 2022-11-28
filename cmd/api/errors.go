//Filename: qapi/cmd/api/errors.go

package main

import "net/http"

// This function adds the error + message to the logger
func (app *application) logError(r *http.Request, err error) {
	app.logger.Println(err)
}

// JSON error message
func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message interface{}) {
	//creating the JSON response
	env := envelope{"error": message}
	err := app.writeJSON(w, status, env, nil)

	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// If the server was unable to process the request for some reason
func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	//logging the error
	app.logError(r, err)

	//Preparing a message with the error
	message := "the server encountered a problem and could not process the request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

// If the int was too large for the server
func (app *application) intTooLargeErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	//logging the error
	app.logError(r, err)

	//Preparing a message with the error
	message := err.Error()
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}