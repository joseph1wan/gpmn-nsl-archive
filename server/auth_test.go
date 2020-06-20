package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestLoginReturnsStatusOK(t *testing.T) {
	app := setupTestApp()

	w := httptest.NewRecorder()

	reqUser := LoginRequest{
		Email:    "dan@admin.com",
		Password: "admin",
	}

	reqBytes, err := json.Marshal(reqUser)
	reqBody := bytes.NewReader(reqBytes)

	req, err := http.NewRequest("POST", "/login", reqBody)
	require.NoError(t, err)
	app.server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code) //t,
	resp, err := ioutil.ReadAll(w.Body)
	assert.Equal(t, "{\"authToken\":\"ZGFuQGFkbWluLmNvbTphZG1pbg==\",\"userID\":1}", string(resp))
}

func TestLoginReturnsStatusBadRequest(t *testing.T) {
	app := setupTestApp()

	w := httptest.NewRecorder()

	req, err := http.NewRequest("POST", "/login", strings.NewReader("bad data"))
	require.NoError(t, err)
	app.server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}

func TestLoginReturnsStatusUnauthorized(t *testing.T) {
	app := setupTestApp()

	w := httptest.NewRecorder()

	reqUser := LoginRequest{
		Email:    "dan@admin.com",
		Password: "badpw",
	}

	reqBytes, err := json.Marshal(reqUser)
	reqBody := bytes.NewReader(reqBytes)

	req, err := http.NewRequest("POST", "/login", reqBody)
	require.NoError(t, err)
	app.server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusUnauthorized, w.Code) //t,
	resp, err := ioutil.ReadAll(w.Body)
	assert.Equal(t, "{\"error\":\"bad user/pw combo\"}", string(resp))
}
