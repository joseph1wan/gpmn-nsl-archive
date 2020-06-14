package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
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
