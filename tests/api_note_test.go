package tests

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"testing"

	"github.com/antennaio/go-notes/api/app"
	"github.com/antennaio/go-notes/api/note"
	"github.com/antennaio/go-notes/api/user"
	"github.com/antennaio/go-notes/lib/env"
	"github.com/stretchr/testify/assert"
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

	assert.Equal(t, "[]", strings.TrimSpace(response.Body.String()))
}

func TestGetNotes(t *testing.T) {
	generator.generateNotes(u, 3)
	defer generator.truncateNotes()

	request, _ := http.NewRequest("GET", "/v1/note", nil)
	request.Header.Add("Authorization", "BEARER "+token)
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusOK, response.Code)

	var notes []*note.Note
	if errJSON := json.Unmarshal([]byte(response.Body.String()), &notes); errJSON != nil {
		t.Error(errJSON)
	}

	assert.Equal(t, 3, len(notes), "Expected 3 notes, got %v", len(notes))
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
