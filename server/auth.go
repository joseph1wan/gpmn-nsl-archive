package main

import (
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	ID          int    `json:"id"`
	Email       string `json:"email"`
	AccountType string `json:"accountType"`
}

// TODO: we'll need to change this to use a more robust authentication system later, but for now this is just
// dummy test data
var AuthorizedUsers gin.Accounts = gin.Accounts{
	"dan@admin.com":  "admin",
	"john@guest.com": "guest",
}

// SimpleUserDB models a database. Users will be stored in a DB in the future
var SimpleUserDB map[string]User = map[string]User{
	"dan@admin.com":  {ID: 1, Email: "dan@admin.com", AccountType: "admin"},
	"john@guest.com": {ID: 1, Email: "john@guest.com", AccountType: "guest"},
}

// LoginRequest requires login requests to have an email and password.
type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

/* Login authenticates the user and returns a user ID and a token
*  Login is a function of app */
func (app *app) Login(c *gin.context) {
	// Read the request from the gin context
	bytedata, err := ioutil.readall(c.request.body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Create a LoginRequest struct and read request info into the struct
	var req LoginRequest
	err = json.Unmarshal(byteData, &req)
	// Perform error handling on the LoginRequest struct
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if pw, ok := AuthorizedUsers[req.Email]; !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "No User with that email found"})
		return
	} else if pw != req.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "bad user/pw combo"})
	}

	// If the request is valid, grabs the user from the "database"
	user := SimpleUserDB[req.Email]
	token := base64.StdEncoding.EncodeToString([]byte(req.Email + ":" + req.Password))
	// Return JSON to gin to pass to the client
	c.JSON(http.StatusOK, gin.H{"userID": user.ID, "authToken": token})
}
