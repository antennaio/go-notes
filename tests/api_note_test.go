package tests

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/antennaio/go-notes/api/app"
	"github.com/antennaio/go-notes/api/note"
	"github.com/antennaio/go-notes/api/user"
	"github.com/antennaio/go-notes/lib/env"
)

var a app.App
var u *user.User
var token string
var generator *Generator

func TestUnauthorized(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/note", nil)
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusUnauthorized, response.Code)
}

func TestNoResults(t *testing.T) {
	request, _ := http.NewRequest("GET", "/v1/note", nil)
	request.Header.Add("Authorization", "BEARER "+token)
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusOK, response.Code)

	if body := strings.TrimSpace(response.Body.String()); body != "[]" {
		t.Errorf("Expected an empty array, got %s", body)
	}
}

func TestGetNotes(t *testing.T) {
	generator.generateNotes(u, 3)
	defer generator.truncateNotes()

	request, _ := http.NewRequest("GET", "/v1/note", nil)
	request.Header.Add("Authorization", "BEARER "+token)
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusOK, response.Code)

	body, errRead := ioutil.ReadAll(response.Body)
	handleError(errRead)

	var notes []*note.Note
	errJSON := json.Unmarshal(body, &notes)
	handleError(errJSON)

	if len(notes) != 3 {
		t.Errorf("Expected 3 notes, got %v", len(notes))
	}
}

func TestMain(m *testing.M) {
	env.LoadEnv("../.env")
	a = makeApp()
	migrateUp(a.Pg)
	generator = &Generator{Pg: a.Pg}
	u, token = generator.generateUser()
	code := m.Run()
	migrateDown(a.Pg)
	os.Exit(code)
}
