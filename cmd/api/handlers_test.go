package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealthcheck(t *testing.T) {
	rr := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/v1/healthcheck", nil)
	handler := http.HandlerFunc(testApp.healthcheck)
	handler.ServeHTTP(rr, req)

	if rr.Code != http.StatusOK {
		t.Error("Divisions returned a wrong status code", rr.Code)
	}
}
