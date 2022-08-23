package main

import "net/http"

// logError is generic helper for logging error message
func (app *application) logError(r *http.Request, err error) {
	app.errorLog.Println(err, map[string]string{
		"request_method": r.Method,
		"request_url":    r.URL.String(),
	})
}

// errorResponse is used for creating specific errors
func (app *application) errorResponse(rw http.ResponseWriter, r *http.Request, status int, message any) {
	errorMessage := envelope{"message": message}
	payload := jsonResponse{
		Error: true,
		Data:  errorMessage,
	}
	err := app.writeJSON(rw, status, payload, nil)
	if err != nil {
		app.logError(r, err)
		rw.WriteHeader(http.StatusInternalServerError)
	}
}

// serverResponse fires 500 Internal Server Error status code
func (app *application) serverErrorResponse(rw http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "the server encountered a problem and could not process the request"
	app.errorResponse(rw, r, http.StatusInternalServerError, message)
}

// badRequestResponse fires 400 Bad Request status code
func (app *application) badRequestResponse(rw http.ResponseWriter, r *http.Request, err error) {
	app.errorResponse(rw, r, http.StatusBadRequest, err.Error())
}
