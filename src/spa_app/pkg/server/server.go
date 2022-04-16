/*
Author: Connor Sanders
Copyright: Connor Sanders 2020
Version: 0.0.1
Released: 12/10/2020

-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
		Golang Frontend Boilerplate V0.0.1
-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
*/

package server

import (
	"flag"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"log"
	"net/http"
	"spa_app/pkg/configuration"
)

// Server is a struct that stores the API Apps high level attributes such as the router, config, and services
type Server struct {
	Config configuration.Configuration
	Router *httprouter.Router
}

// NewServer is a function used to initialize a new Server struct
func NewServer(config configuration.Configuration, router *httprouter.Router) *Server {
	s := Server{Config: config, Router: router}
	return &s
}

// Start starts the initialized server
func (s *Server) Start() {
	listen := flag.String("-listen", ":"+s.Config.Port, "Interface and port to listen on")
	flag.Parse()
	fmt.Println("Listening on", *listen)
	log.Fatal(http.ListenAndServe(*listen, s.Router))
}
