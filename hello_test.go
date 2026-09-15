package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/hello?name=", nil)

	rr := httptest.NewRecorder()

	helloHandler(rr, req)

	got := rr.Body.String()
	want := "hello Guest\n"
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}

}
