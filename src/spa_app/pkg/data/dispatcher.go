/*
Author: Connor Sanders
Copyright: Connor Sanders 2020
Version: 0.0.1
Released: 12/10/2020

-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
		Golang Frontend Boilerplate V0.0.1
-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
*/

package data

import (
	"bytes"
	"net/http"
	"net/url"
	"os"
)

// Request
type Request struct {
	Endpoint string
	Headers  [][]string
	Body     []byte
	Type     string
}

// DefaultHeaders
func (re *Request) DefaultHeaders() {
	var headers [][]string
	headerEntry := []string{"Content-Type", "application/json"}
	headers = append(headers, headerEntry)
	re.Headers = headers
}

// Dispatcher
type Dispatcher struct {
	APIHost string
	Req     Request
	Res     *http.Response
	Promise *Promise
}

// InitializeDispatcher
func InitializeDispatcher() *Dispatcher {
	d := Dispatcher{APIHost: os.Getenv("API_HOST")}
	return &d
}

// headers
func (d *Dispatcher) headers(r *http.Request) *http.Request {
	for _, headerStr := range d.Req.Headers {
		r.Header.Set(headerStr[0], headerStr[1])
	}
	return r
}

// Execute Request
func (d *Dispatcher) Execute(resType string) {
	apiUrl := d.APIHost + d.Req.Endpoint
	u, _ := url.ParseRequestURI(apiUrl)
	urlStr := u.String()
	var r *http.Request
	if d.Req.Body == nil {
		r, _ = http.NewRequest(d.Req.Type, urlStr, nil) // No body
	} else {
		r, _ = http.NewRequest(d.Req.Type, urlStr, bytes.NewBuffer(d.Req.Body)) // body
	}
	r = d.headers(r)
	reqPromise := Fetch(r)
	if resType == "discard" {
		<-reqPromise.Channel
	} else {
		d.Promise = reqPromise
	}
}

// Resolve Request
func (d *Dispatcher) Resolve() {
	resp := <-d.Promise.Channel
	d.Res = resp
}

// Authenticate a user session
func (d *Dispatcher) Authenticate(bodyStr string, authType string) {
	var body = []byte(bodyStr)
	d.Req = Request{Endpoint: "/auth", Type: "POST", Body: body}
	d.Req.DefaultHeaders()
	if authType == "registration" {
		d.Req.Endpoint = "/auth/register"
	}
	d.Execute("default")
	d.Resolve()
}

// Invalidate a user session
func (d *Dispatcher) Invalidate(authToken string) {
	d.Req = Request{Endpoint: "/auth", Type: "DELETE", Body: nil}
	d.Req.DefaultHeaders()
	headerEntry := []string{"Auth-Token", authToken}
	d.Req.Headers = append(d.Req.Headers, headerEntry)
	d.Execute("discard")
}
