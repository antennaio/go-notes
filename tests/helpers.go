package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi"
	"github.com/stretchr/testify/assert"
)

func recordResponse(router *chi.Mux, r *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, r)
	return rr
}

func verifyResponseCode(t *testing.T, expected, actual int) {
	assert.Equal(t, expected, actual, "Expected response code %d, got %d", expected, actual)
}

func handleError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
