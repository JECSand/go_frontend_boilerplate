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
	"spa_app/pkg/configuration"
	"spa_app/pkg/router"
	"spa_app/pkg/server"
	"spa_app/pkg/session"
	"spa_app/pkg/web"
)

// App is a the highest level struct of the rest_api application. Stores the server, client, and config settings.
type App struct {
	server *server.Server
	config configuration.Configuration
}

// Initialize is a function used to initialize a new instantiation of the API Application
func (a *App) Initialize(env string) {
	a.config = configuration.ConfigurationSettings(env)
	a.config.InitializeEnvironmentals()
	var globalSessions *session.Manager
	globalSessions = session.NewManager()
	p := web.InitializePages(globalSessions)
	p.InitializeTemplates()
	r := router.GetRouter(p)
	a.server = server.NewServer(a.config, r)
}

// Run is a function used to run a previously initialized Application
func (a *App) Run() {
	a.server.Start()
}
