/*
Author: Connor Sanders
Copyright: Connor Sanders 2020
Version: 0.0.1
Released: 12/10/2020

-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
		Golang Frontend Boilerplate V0.0.1
-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
*/

package models

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"spa_app/pkg/data"
	"strings"
	"time"
)

// Auth
type Auth struct {
	Username      string
	UserUuid      string
	GroupUuid     string
	Role          string
	AuthToken     string
	APIKey        string
	Authenticated bool
	LastLogin     string
	LoginIP       string
	Status        int
}

// GetAuthString
func (a *Auth) GetAuthString() string {
	authStr := "false"
	if a.Authenticated {
		authStr = "true"
	}
	var reString string
	reString = a.Username + "||" + a.UserUuid + "||" + a.GroupUuid + "||" + a.Role + "||" + authStr + "||" + a.AuthToken + "||" + a.LastLogin + "||" + a.LoginIP
	return reString
}

// LoadAuthString
func (a *Auth) LoadAuthString(authString string) {
	authBool := false
	sString := strings.Split(authString, "||")
	authStr := sString[4]
	if authStr == "true" {
		authBool = true
	}
	a.Username = sString[0]
	a.UserUuid = sString[1]
	a.GroupUuid = sString[2]
	a.Role = sString[3]
	a.Authenticated = authBool
	a.AuthToken = sString[5]
	a.LastLogin = sString[6]
	a.LoginIP = sString[7]
}

// Delete all data in Auth struct
func (a *Auth) delete() {
	a.Username = ""
	a.UserUuid = ""
	a.GroupUuid = ""
	a.Role = ""
	a.Authenticated = false
	a.AuthToken = ""
	a.LastLogin = ""
	a.LoginIP = ""
}

// Load Auth with data from a http.Response
func (a *Auth) load(resp *http.Response) {
	currentTime := time.Now().UTC()
	authToken := resp.Header["Auth-Token"]
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := resp.Body.Close(); err != nil {
		panic(err)
	}
	var user User
	err = json.Unmarshal(body, &user)
	if err != nil {
		panic(err)
	}
	a.Status = resp.StatusCode
	if a.Status == http.StatusOK {
		a.Username = user.Username
		a.UserUuid = user.Uuid
		a.GroupUuid = user.GroupUuid
		a.Role = user.Role
		a.AuthToken = authToken[0]
		a.LastLogin = currentTime.String()
		a.Authenticated = true
		// TODO - Nick - Log LOGIN IP
		a.LoginIP = "Unknown"
	}
}

// AuthHeaders
func (a *Auth) AuthHeaders(req data.Request) data.Request {
	if a.Authenticated {
		headerEntry := []string{"Auth-Token", a.AuthToken}
		req.Headers = append(req.Headers, headerEntry)
	}
	return req
}

// Register
func (a *Auth) Register(firstName string, lastName string, email string, userName string, password string) {
	bodyStr := `{"firstname": "` + firstName + `","lastname":"` + lastName + `","email":"` + email + `","username":"` + userName + `","password":"` + password + `"}`
	dispatch := data.InitializeDispatcher()
	dispatch.Authenticate(bodyStr, "registration")
	a.load(dispatch.Res)
}

// Authenticate
func (a *Auth) Authenticate(userName string, password string) {
	bodyStr := `{"username":"` + userName + `","password":"` + password + `"}`
	dispatch := data.InitializeDispatcher()
	dispatch.Authenticate(bodyStr, "login")
	a.load(dispatch.Res)
}

// Invalidate
func (a *Auth) Invalidate() {
	dispatch := data.InitializeDispatcher()
	dispatch.Invalidate(a.AuthToken)
	a.delete()
}
