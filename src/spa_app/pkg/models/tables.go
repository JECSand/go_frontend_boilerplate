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

// UserTable
type UserTable struct {
	CreateUserForm CreateUserForm
	UpdateUserForm UpdateUserForm
	Users          []User
}

// Initialize
func (ut *UserTable) Initialize(auth Auth, method string, id string) {
	if method == "GET" {
		if id == "" {
			user := User{}
			users := user.GetAll(auth)
			ut.Users = users
		} else if id == "create" {
			ut.CreateUserForm = CreateUserForm{}
		} else {
			user := User{Uuid: id}
			user.Get(auth)
			upUserForm := UpdateUserForm{}
			upUserForm.Load(user)
			ut.UpdateUserForm = upUserForm
		}
	}
}

/*
// findUserByUuid
func (ut *UserTable) findUserByUuid(userUuid string) User {
	var user User
	for _, u := range ut.Users {
		if u.Uuid == userUuid {
			user = u
			break
		}
	}
	return user
}

// Load
func (ut *UserTable) Load(auth Auth) {
	var users []User
	var user User
	users = user.GetAll(auth)
	ut.Users = users
}

// LoadForm
func (ut *UserTable) LoadForm(userUuid string) {
	user := ut.findUserByUuid(userUuid)
	if userUuid != "" {
		// Create User
		ut.UpdateUserForm.Load(user)
	}
}
*/

/*
////////////////////////////////////////////
*/

// GroupTable
type GroupTable struct {
	CreateGroupForm CreateGroupForm
	UpdateGroupForm UpdateGroupForm
	Groups          []Group
}

// Initialize
func (gt *GroupTable) Initialize(auth Auth, method string, id string) {
	if method == "GET" {
		if id == "" {
			group := Group{}
			groups := group.GetAll(auth)
			gt.Groups = groups
		} else if id == "create" {
			gt.CreateGroupForm = CreateGroupForm{}
		} else {
			group := Group{Uuid: id}
			group.Get(auth)
			upGroupForm := UpdateGroupForm{}
			upGroupForm.Load(group)
			gt.UpdateGroupForm = upGroupForm
		}
	}
}

/*
// Load
func (gt *GroupTable) findGroupByUuid(groupUuid string) Group {
	var group Group
	for _, gt := range gt.Groups {
		if gt.Uuid == groupUuid {
			group = gt
			break
		}
	}
	return group
}

// Load
func (gt *GroupTable) Load(auth Auth) {
	var groups []Group
	var group Group
	groups = group.GetAll(auth)
	gt.Groups = groups
}

// LoadForm
func (gt *GroupTable) LoadForm(groupUuid string) {
	group := gt.findGroupByUuid(groupUuid)
	if groupUuid != "" {
		// Create User
		gt.UpdateGroupForm.Load(group)
	}
}
*/
