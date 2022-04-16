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

import "fmt"

// IndexModel of dynamic data used for index view
type IndexModel struct {
	Title    string
	Name     string
	SubRoute string
	Route    string
	Auth     Auth
}

// LoginModel of dynamic data used for variable view
type LoginModel struct {
	Title    string
	Variable string
	Name     string
	SubRoute string
	Route    string
	Auth     Auth
}

// BuildRoute
func (lm *LoginModel) BuildRoute() {
	route := lm.Name + "/" + lm.SubRoute
	lm.Route = route
}

// AboutModel of dynamic data used for about view
type AboutModel struct {
	Title    string
	Variable string
	Name     string
	SubRoute string
	Route    string
	Auth     Auth
}

// BuildRoute
func (am *AboutModel) BuildRoute() {
	route := am.Name + "/" + am.SubRoute
	am.Route = route
}

// VariableModel of dynamic data used for variable view
type VariableModel struct {
	Title    string
	Variable string
	Name     string
	SubRoute string
	Route    string
	Auth     Auth
}

// BuildRoute
func (vm *VariableModel) BuildRoute() {
	route := vm.Name + "/" + vm.SubRoute
	vm.Route = route
}

// AdminModel
type AdminModel struct {
	Title      string
	Variable   string
	Name       string
	SubRoute   string
	Route      string
	Id         string
	Method     string
	UserTable  UserTable
	GroupTable GroupTable
	Auth       Auth
}

// BuildRoute
func (adm *AdminModel) BuildRoute() {
	route := adm.Name
	if adm.SubRoute != "" {
		route = route + "/" + adm.SubRoute
		if adm.Id != "" {
			route = "/" + adm.Id
		}
	}
	adm.Route = route
}

// Initialize a new Admin Page Data Model
func (adm *AdminModel) Initialize() {
	adm.BuildRoute()
	if adm.SubRoute == "users" {
		adm.UserTable.Initialize(adm.Auth, adm.Method, adm.Id)
	} else if adm.SubRoute == "groups" {
		adm.GroupTable.Initialize(adm.Auth, adm.Method, adm.Id)
	}
}

// AccountModel
type AccountModel struct {
	Title    string
	Variable string
	Name     string
	SubRoute string
	Route    string
	Auth     Auth
	User     User
	Form     Form
}

// BuildRoute
func (acm *AccountModel) BuildRoute() {
	route := acm.Name + "/" + acm.SubRoute
	fmt.Println("subroutetest:", acm.SubRoute)
	acm.Route = route
	fmt.Println("routetest:", acm.Route)
}

// LoadUser for info display
func (acm *AccountModel) LoadUser() {
	user := User{Uuid: acm.Auth.UserUuid}
	user.Get(acm.Auth)
	acm.User = user
	if acm.SubRoute == "settings" {
		form := InitializeSettingsForm(acm.User)
		acm.Form = form
	}
}

// Initialize a new Account Page Data Model
func (acm *AccountModel) Initialize() {
	acm.BuildRoute()
	acm.LoadUser()

}
