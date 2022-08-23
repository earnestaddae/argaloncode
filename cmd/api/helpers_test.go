package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReadJSON(t *testing.T) {
	sampleJSON := map[string]interface{}{
		"foo": "bar",
	}

	body, _ := json.Marshal(sampleJSON)

	var decodedJSON struct {
		Foo string `json:"foo"`
	}

	req, err := http.NewRequest("POST", "/v1/divisions", bytes.NewReader(body))
	if err != nil {
		t.Log(err)
	}

	rr := httptest.NewRecorder()
	defer req.Body.Close()

	err = testApp.readJSON(rr, req, &decodedJSON)
	if err != nil {
		t.Error("failed to decode json", err)
	}
}

func TestWriteJSON(t *testing.T) {
	rr := httptest.NewRecorder()
	payload := jsonResponse{
		Error: false,
		Data:  envelope{"Foo": "Bar"},
	}
	headers := make(http.Header)
	headers.Add("FOO", "BAR")
	err := testApp.writeJSON(rr, http.StatusOK, payload, headers)
	if err != nil {
		t.Errorf("failed to write JSON: %v", err)
	}

	// test production environment
	testApp.config.env = "production"
	err = testApp.writeJSON(rr, http.StatusOK, payload, headers)
	if err != nil {
		t.Errorf("failed to write JSON in production: %v", err)
	}
	testApp.config.env = "environment"
}
