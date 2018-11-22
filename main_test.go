package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestIndexHandler(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("unexpected status: got '%v' want '%v'", status, http.StatusOK)
	}

	expected := "Hello, World from Travis CI!"
	if rr.Body.String() != expected {
		t.Errorf("unexpected body: got '%v' want '%v'", rr.Body.String(), "Hello, World!")
	}
}

func TestIndexHandlerNotFound(t *testing.T) {
	req := httptest.NewRequest("GET", "/404", nil)
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(indexHandler)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNotFound {
		t.Errorf("unexpected status: got '%v' want '%v'", status, http.StatusOK)
	}
}
