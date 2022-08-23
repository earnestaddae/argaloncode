package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// writeJSON takes a response status code and arbitrary data and writes a JSON response to the client
// it sends responses in JSON
func (app *application) writeJSON(rw http.ResponseWriter, status int, data any, headers ...http.Header) error {
	var output []byte

	if app.config.env == "development" {
		out, err := json.MarshalIndent(data, "", "\t")
		if err != nil {
			return err
		}
		output = out
	} else {
		out, err := json.Marshal(data)
		if err != nil {
			return err
		}
		output = out
	}

	if len(headers) > 0 {
		for key, value := range headers[0] {
			rw.Header()[key] = value
		}
	}
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(status)
	_, err := rw.Write(output)
	if err != nil {
		return err
	}
	return nil
}

// readJSON tries to read the body of a request and converts it into JSON
func (app *application) readJSON(rw http.ResponseWriter, r *http.Request, data interface{}) error {
	maxBytes := 1_048_576 // 1MB
	r.Body = http.MaxBytesReader(rw, r.Body, int64(maxBytes))

	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()
	err := dec.Decode(data)
	if err != nil {
		return err
	}

	err = dec.Decode(&struct{}{})
	if err != io.EOF {
		return errors.New("body must have only a single json value")
	}
	return nil
}
