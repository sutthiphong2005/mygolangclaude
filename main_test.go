package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGreeting(t *testing.T) {
	got := greeting()
	want := "Hello, World!"
	if got != want {
		t.Errorf("greeting() = %q, want %q", got, want)
	}
}

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	helloHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rr.Code, http.StatusOK)
	}
	got := strings.TrimSpace(rr.Body.String())
	want := "Hello, World!"
	if got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}

func TestDemoHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/demo", nil)
	rr := httptest.NewRecorder()

	demoHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rr.Code, http.StatusOK)
	}
	got := strings.TrimSpace(rr.Body.String())
	want := "demo"
	if got != want {
		t.Errorf("body = %q, want %q", got, want)
	}
}
