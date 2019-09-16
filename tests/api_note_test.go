package tests

import (
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/antennaio/go-notes/api/app"
	"github.com/antennaio/go-notes/api/user"
	"github.com/antennaio/go-notes/lib/env"
)

var a app.App
var u *user.User
var token string

func TestUnauthorized(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/note", nil)
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusUnauthorized, response.Code)
}

func TestGetNotes(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/note", nil)
	request.Header.Add("Authorization", "BEARER "+token)
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusOK, response.Code)

	if body := strings.TrimSpace(response.Body.String()); body != "[]" {
		t.Errorf("Expected an empty array, got %s", body)
	}
}

func TestMain(m *testing.M) {
	env.LoadEnv("../.env")
	a = makeApp()
	migrateUp(a.Pg)
	u, token = generateUser(a.Pg)
	code := m.Run()
	migrateDown(a.Pg)
	os.Exit(code)
}
