package main

import (
	"net/http"
)

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
		return
	}
}

// divide handles the division of two operands
func (app *application) divide(rw http.ResponseWriter, r *http.Request) {
	var input struct {
		A float64 `json:"a"`
		B float64 `json:"b"`
	}

	err := app.readJSON(rw, r, &input)
	if err != nil {
		app.badRequestResponse(rw, r, err)
		return
	}

	x, y, err := app.Division(input.A, input.B)
	if err != nil {
		app.badRequestResponse(rw, r, err)
		return
	}

	payload := jsonResponse{
		Error: false,
		Data: envelope{
			"a divided by b": x,
			"b divided by a": y,
		},
	}

	err = app.writeJSON(rw, http.StatusCreated, payload, nil)
	if err != nil {
		app.serverErrorResponse(rw, r, err)
		return
	}
}
