package main

import (
	"github.com/a2fumn2022/gpmn-nsl/server/datastore"
	"testing"
	"net/http"
	"net/http/httptest"
	"encoding/json"
	"io/ioutil"
	"bytes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	//"time"
)

type DatabaseLayer struct {
	MockDB []datastore.MaintenanceRequest
}

//var mockDB DatabaseLayer

func setupTestApp() *App {
	app := App{
		conf: ServerConfig{Port: 4000},
	}
	app.setupRoutes()
	w := httptest.NewRecorder()

	//TODO: add setting up a test postgresql database
	return &app
}

func TestCreateMaintenanceRequestStatusOK(t *testing.T) {
	app := setupTestApp()
	w := httptest.NewRecorder()

	newRequest := datastore.MaintenanceRequest{
		Request:	"Fix lawnmower",
	}

	byteData, err := json.Marshal(newRequest)
	reqBody := bytes.NewReader(byteData)

	req, err := http.NewRequest("POST","/maintenance_requests", reqBody)
	require.NoError(t, err)
	app.server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	resp, err := ioutil.ReadAll(w.Body)
	assert.Equal(t, "\"request\":\"1\"", string(resp))
}