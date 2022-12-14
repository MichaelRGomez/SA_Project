//Filename: qapi/cmd/api/routes.go

package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Restful http router
func (app *application) routes() *httprouter.Router {
	//creating instance of a http router
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/v1/randomstring/:seed", app.makeRandomStringHandler)
	return router
}