package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"os"
	"strconv"
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
	request, errRequest := http.NewRequest("GET", "/v1/note", nil)
	if errRequest != nil {
		t.Error(errRequest)
	}
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusUnauthorized, response.Code)
}

func TestNoResults(t *testing.T) {
	request, errRequest := http.NewRequest("GET", "/v1/note", nil)
	if errRequest != nil {
		t.Error(errRequest)
	}
	request.Header.Add("Authorization", "BEARER "+token)
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusOK, response.Code)

	assert.Equal(t, "[]", strings.TrimSpace(response.Body.String()))
}

func TestGetNotes(t *testing.T) {
	generator.generateNotes(u, 3)
	defer generator.truncateNotes()

	request, errRequest := http.NewRequest("GET", "/v1/note", nil)
	if errRequest != nil {
		t.Error(errRequest)
	}
	request.Header.Add("Authorization", "BEARER "+token)
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusOK, response.Code)

	var notes []*note.Note
	if errJSON := json.Unmarshal([]byte(response.Body.String()), &notes); errJSON != nil {
		t.Error(errJSON)
	}

	assert.Equal(t, 3, len(notes), "Expected 3 notes, got %v", len(notes))
}

func TestCreateNote(t *testing.T) {
	defer generator.truncateNotes()

	body, errBody := json.Marshal(map[string]string{
		"title":   "New note",
		"content": "Content",
	})
	if errBody != nil {
		t.Error(errBody)
	}

	request, errRequest := http.NewRequest("POST", "/v1/note", bytes.NewBuffer(body))
	if errRequest != nil {
		t.Error(errRequest)
	}
	request.Header.Add("Authorization", "BEARER "+token)
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusOK, response.Code)

	count, err := a.Pg.Model((*note.Note)(nil)).Count()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 1, count, "Expected 1 note, got %v", count)
}

func TestDeleteNote(t *testing.T) {
	notes := generator.generateNotes(u, 1)
	defer generator.truncateNotes()

	id := strconv.Itoa(notes[0].Id)
	request, errRequest := http.NewRequest("DELETE", "/v1/note/"+id, nil)
	if errRequest != nil {
		t.Error(errRequest)
	}
	request.Header.Add("Authorization", "BEARER "+token)
	response := recordResponse(a.Router, request)

	verifyResponseCode(t, http.StatusOK, response.Code)

	count, err := a.Pg.Model((*note.Note)(nil)).Count()
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, 0, count, "Expected 0 notes, got %v", count)
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
