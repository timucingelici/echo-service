package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestEchoServerMirrorsTheRequest(t *testing.T) {

	// Request body
	body := "Hello, 世界"

	// Request
	req, err := http.NewRequest("POST", "/some-endpoint", bytes.NewBufferString(body))

	if err != nil {
		t.Fatal(err)
	}

	// Request Recorder
	rr 		:= httptest.NewRecorder()

	// Server
	handler := http.HandlerFunc(echoHandler)
	handler.ServeHTTP(rr, req)

	// Check if the returned status is OK
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Unexpected response code. Got %v, want %v", status, http.StatusOK)
	}

	// Check if the returned body is same as the body we sent
	if rr.Body.String() != body {
		t.Errorf("Unexpected response. Got %v, want %v", rr.Body.String(), body)
	}
}
