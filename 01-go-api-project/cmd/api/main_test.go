package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMain(t *testing.T) {
	router := SetupRouter()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/ping", nil)

	if err != nil {
		t.Fatalf("Failed to create a request. Error: %e", err)
	}

	rw := httptest.NewRecorder()

	router.ServeHTTP(rw, req)

	if status := rw.Result().StatusCode; status != http.StatusOK {
		t.Fatalf("Received status: %d, when %d expected", status, http.StatusOK)
	}

	expected := "pong"
	if response := rw.Body.String(); response != expected {
		t.Fatalf("Received response %s when %s expected\n", response, expected)
	}
}

func TestOk01(t *testing.T) {}
func TestOk02(t *testing.T) {}
