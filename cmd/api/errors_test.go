package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestErrorResponse(t *testing.T) {
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Log(err)
	}

	err = testApp.errorResponse(rr, req, http.StatusInternalServerError, errors.New("an error occcured"))
	if err != nil {
		t.Error(err)
	}
	testJSONPayload(t, rr)

	err = testApp.badRequestResponse(rr, req, errors.New("bad request"))
	if err != nil {
		t.Error(err)
	}
	testJSONPayload(t, rr)

	err = testApp.serverErrorResponse(rr, req, errors.New("server error"))
	if err != nil {
		t.Error(err)
	}
	testJSONPayload(t, rr)
}

func testJSONPayload(t *testing.T, rr *httptest.ResponseRecorder) {
	var requestPayload jsonResponse
	decoder := json.NewDecoder(rr.Body)
	err := decoder.Decode(&requestPayload)
	if err != nil {
		t.Error("received error when decoding", err)
	}

	if !requestPayload.Error {
		t.Error("error set to false")
	}
}
