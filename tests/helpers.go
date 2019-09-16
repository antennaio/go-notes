package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
)

func recordResponse(router *chi.Mux, r *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, r)
	return rr
}

func verifyResponseCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d, got %d\n", expected, actual)
	}
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
