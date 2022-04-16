/*
Author: Connor Sanders
Copyright: Connor Sanders 2020
Version: 0.0.1
Released: 12/10/2020

-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
		Golang Frontend Boilerplate V0.0.1
-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
*/

package router

import (
	"github.com/julienschmidt/httprouter"
	"spa_app/pkg/web"
)

// GetRouter
func GetRouter(p web.Pages) *httprouter.Router {

	// mux handler
	router := httprouter.New()

	// Index route
	router.GET("/", p.IndexPage)
	// TODO Change "/" to "/register" and Make index (/) a slash page
	// Register new Account/User at Index Page Form
	router.POST("/", p.RegistrationHandler)

	// Login Route
	router.GET("/login", p.LoginPage)
	router.POST("/login", p.LoginHandler)

	// Logout Route
	router.GET("/logout", p.LogoutHandler)

	// About Route
	router.GET("/about", p.AboutPage)
	router.GET("/about/:child", p.AboutPage)

	// Admin Page Routes
	router.Handler("GET", "/admin", p.Protected(p.AdminPage)) // 1) GET GENERIC ADMIN VIEW

	router.Handler("GET", "/admin/:child", p.Protected(p.AdminPage)) // 2) GET ADMIN USER OR GROUP DATA TABLE
	// ID CAN BE EITHER A UUID OR CREATE - CREATE LOADS CREATE FORM
	router.Handler("GET", "/admin/:child/:id", p.Protected(p.AdminPage)) // 3) GET ADMIN USER OR GROUP UPDATE FORM
	// Admin Group Handler Routes
	router.Handler("POST", "/admin/:child", p.Protected(p.CreateHandler))
	// TO LOAD AN UPDATE FORM SPECIFICALLY WHEN APP USER CLICKS UPDATE FOR A GROUP LISTED IN GROUP DATATABLE
	// HANDLERS TO SUBMIT UPDATE FORM OR DELETE A GROUP
	router.Handler("PATCH", "/admin/:child/:id", p.Protected(p.UpdateHandler))
	router.Handler("DELETE", "/admin/:child/:id", p.Protected(p.DeleteHandler))

	//router.Handler("PATCH", "/admin/usermenu", p.Protected(p.AdminPage)) // 2) GET ADMIN USER OR GROUP DATA TABLE

	// Admin User Handler Routes
	//router.Handler("POST", "/admin/users", p.Protected(p.CreateUserHandler))
	// TO LOAD AN UPDATE FORM SPECIFICALLY WHEN APP USER CLICKS UPDATE FOR A USER LISTED IN USER DATATABLE
	//router.Handler("GET", "/admin/users/:id", p.Protected(p.AdminPage))
	// HANDLERS TO SUBMIT UPDATE FORM OR DELETE A USER
	//router.Handler("PATCH", "/admin/users/:id", p.Protected(p.UpdateUserHandler))
	//router.Handler("DELETE", "/admin/users/:id", p.Protected(p.DeleteUserHandler))

	// Account Route
	router.Handler("GET", "/account", p.Protected(p.AccountPage))
	router.Handler("GET", "/account/:child", p.Protected(p.AccountPage))

	// Account Settings Route
	router.Handler("PATCH", "/account/settings", p.Protected(p.AccountSettingsHandler))

	// Variable Route
	router.GET("/variable", p.VariablePage)
	router.GET("/variable/:child", p.VariablePage)

	// Example route that encounters an error
	router.GET("/broken/handler", p.BrokenPage)

	// Serve static assets via the "static" directory
	router.ServeFiles("/static/*filepath", p.Statics)

	return router
}
