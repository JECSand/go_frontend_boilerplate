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
	"net/http"
)

// Promise
type Promise struct {
	Channel     chan *http.Response
	httpRequest *http.Request
}

// worker
func (p *Promise) worker(done chan *http.Response) {
	client := &http.Client{}
	resp, _ := client.Do(p.httpRequest)
	done <- resp
	<-done
}

// execute
func (p *Promise) execute() {
	done := make(chan *http.Response)
	go p.worker(done)
	p.Channel = done
}

// Fetch
func Fetch(r *http.Request) *Promise {
	promise := Promise{httpRequest: r}
	promise.execute()
	return &(promise)
}
