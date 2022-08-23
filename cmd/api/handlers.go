package main

import "net/http"

// jsonResponse is the struct to display information to the client
type jsonResponse struct {
	Error bool `json:"error"`
	Data  any  `json:"data"`
}

// envelope is used to encapsulate data for jsonResponse
type envelope map[string]any

// healthcheck displays the status of the api
func (app *application) healthcheck(rw http.ResponseWriter, r *http.Request) {
	env := envelope{
		"status": "available",
		"application_information": map[string]string{
			"environment": app.config.env,
			"version":     app.config.version,
		},
	}
	err := app.writeJSON(rw, http.StatusOK, env, nil)
	if err != nil {
		app.serverErrorResponse(rw, r, err)
	}
}
