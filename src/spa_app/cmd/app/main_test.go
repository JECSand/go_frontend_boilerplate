/*
Author: Connor Sanders
Copyright: Connor Sanders 2020
Version: 0.0.1
Released: 12/10/2020

-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
		Golang Frontend Boilerplate V0.0.1
-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
*/

package main

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

var ta App

// Setup Tests
func setup() {
	ta = App{}
	ta.Initialize("test")
}

// TestDefaultHandler
func TestDefaultHandler(t *testing.T) {
	setup()
	server := httptest.NewServer(ta.server.Router)
	defer server.Close()
	resp, err := http.Get(server.URL + "/")
	assert.Nil(t, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(body), "Buddy")
	assert.Contains(t, string(body), "Boy")
}

// TestAboutHandler
func TestAboutHandler(t *testing.T) {
	setup()
	server := httptest.NewServer(ta.server.Router)
	defer server.Close()
	resp, err := http.Get(server.URL + "/about")
	assert.Nil(t, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(body), "Dawg")
}

// TestVariableHandler
func TestVariableHandler(t *testing.T) {
	setup()
	server := httptest.NewServer(ta.server.Router)
	defer server.Close()
	resp, err := http.Get(server.URL + "/variable/friend")
	assert.Nil(t, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(body), "Hi FRIEND")
}

// TestHandlerWithError
func TestHandlerWithError(t *testing.T) {
	setup()
	server := httptest.NewServer(ta.server.Router)
	defer server.Close()
	resp, err := http.Get(server.URL + "/broken/handler")
	assert.Nil(t, err)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	assert.Nil(t, err)
	assert.Contains(t, string(body), `html/template: "templates/missing.html" is undefined`)
	assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
}
