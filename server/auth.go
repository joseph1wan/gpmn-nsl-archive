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

var SimpleUserDB map[string]User = map[string]User{
	"dan@admin.com":  {ID: 1, Email: "dan@admin.com", AccountType: "admin"},
	"john@guest.com": {ID: 1, Email: "john@guest.com", AccountType: "guest"},
}

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (app *App) Login(c *gin.Context) {
	byteData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var req LoginRequest
	err = json.Unmarshal(byteData, &req)
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

	user := SimpleUserDB[req.Email]
	token := base64.StdEncoding.EncodeToString([]byte(req.Email + ":" + req.Password))
	c.JSON(http.StatusOK, gin.H{"userID": user.ID, "authToken": token})
}
