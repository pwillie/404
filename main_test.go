package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestStatus(t *testing.T) {

	uri := "http://acme.com/status"

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	healthcheck_handler(w, req)
	resp := w.Result()

	if resp.StatusCode != 200 {
		t.Errorf("Received non-200 response: %d\n", resp.StatusCode)
	}
}

func TestNotFound(t *testing.T) {

	uri := "http://acme.com/some/url"

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		t.Fatal(err)
	}

	w := httptest.NewRecorder()
	default_handler(w, req)
	resp := w.Result()

	if resp.StatusCode != 404 {
		t.Errorf("Received non-404 response: %d\n", resp.StatusCode)
	}
}
