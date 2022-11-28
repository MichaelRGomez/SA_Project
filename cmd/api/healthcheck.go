//Filename: qapi/cmd/api/healthcheck.go

package main

import "net/http"

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	//creatng map that'll hold the health of the server
	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	//converting map -> JSON Object
	err := app.writeJSON(w, http.StatusOK, data, nil)

	//print error if any
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}