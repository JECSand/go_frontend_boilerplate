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
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"spa_app/pkg/data"
)

// loadGroup
func loadGroup(resp *http.Response) Group {
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := resp.Body.Close(); err != nil {
		panic(err)
	}
	var group Group
	err = json.Unmarshal(body, &group)
	if err != nil {
		panic(err)
	}
	return group
}

// loadGroups
func loadGroups(resp *http.Response) []Group {
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := resp.Body.Close(); err != nil {
		panic(err)
	}
	var groups []Group
	err = json.Unmarshal(body, &groups)
	if err != nil {
		panic(err)
	}
	return groups
}

// loadUser
func loadUser(resp *http.Response) User {
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
	return user
}

// loadUsers
func loadUsers(resp *http.Response) []User {
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
	fmt.Println("body:", body)
	if err != nil {
		panic(err)
	}
	if err := resp.Body.Close(); err != nil {
		panic(err)
	}
	var users []User
	err = json.Unmarshal(body, &users)
	if err != nil {
		panic(err)
	}
	return users
}

// loadTodo
func loadTodo(resp *http.Response) Todo {
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := resp.Body.Close(); err != nil {
		panic(err)
	}
	var todo Todo
	err = json.Unmarshal(body, &todo)
	if err != nil {
		panic(err)
	}
	return todo
}

// loadTodos
func loadTodos(resp *http.Response) []Todo {
	body, err := ioutil.ReadAll(io.LimitReader(resp.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := resp.Body.Close(); err != nil {
		panic(err)
	}
	var todos []Todo
	err = json.Unmarshal(body, &todos)
	if err != nil {
		panic(err)
	}
	return todos
}

// Group is a root struct that is used to store the json encoded data for/from a mongodb group doc.
type Group struct {
	Id               string `json:"id,omitempty"`
	GroupType        string `json:"grouptype,omitempty"`
	Uuid             string `json:"uuid,omitempty"`
	Name             string `json:"name,omitempty"`
	LastModified     string `json:"last_modified,omitempty"`
	CreationDatetime string `json:"creation_datetime,omitempty"`
}

func (g *Group) GetBodyString() string {
	bodyStr := `{"name": "` + g.Name + `"}`
	return bodyStr
}

// Get a Group
func (g *Group) Get(auth Auth) {
	dataReq := data.DatabaseRequest{AuthToken: auth.AuthToken}
	d := dataReq.GetGroups(g.Uuid)
	d.Resolve()
	getGroup := loadGroup(d.Res)
	g.Id = getGroup.Id
	g.GroupType = getGroup.GroupType
	g.Name = getGroup.Name
	g.LastModified = getGroup.LastModified
	g.CreationDatetime = getGroup.CreationDatetime
}

// Get Groups
func (g *Group) GetAll(auth Auth) []Group {
	dataReq := data.DatabaseRequest{AuthToken: auth.AuthToken}
	d := dataReq.GetGroups("")
	d.Resolve()
	return loadGroups(d.Res)
}

// Create a Group
func (g *Group) Create(auth Auth) {
	dataReq := data.DatabaseRequest{AuthToken: auth.AuthToken}
	d := dataReq.PostGroup(g.GetBodyString())
	d.Resolve()
	group := loadGroup(d.Res)
	g.Id = group.Id
	g.Uuid = group.Uuid
	g.GroupType = group.GroupType
	g.Name = group.Name
	g.LastModified = group.LastModified
	g.CreationDatetime = group.CreationDatetime
}

// Update a Group
func (g *Group) Update(auth Auth) {
	dataReq := data.DatabaseRequest{AuthToken: auth.AuthToken}
	d := dataReq.PatchGroup(g.Uuid, g.GetBodyString())
	d.Resolve()
	group := loadGroup(d.Res)
	g.Id = group.Id
	g.Uuid = group.Uuid
	g.GroupType = group.GroupType
	g.Name = group.Name
	g.LastModified = group.LastModified
	g.CreationDatetime = group.CreationDatetime
}

// Delete a Group
func (g *Group) Delete(auth Auth) {
	dataReq := data.DatabaseRequest{AuthToken: auth.AuthToken}
	_ = dataReq.DeleteGroup(g.Uuid)
}

// User is a root struct that is used to store the json encoded data for/from a mongodb user doc.
type User struct {
	Id               string `json:"id,omitempty"`
	Uuid             string `json:"uuid,omitempty"`
	Username         string `json:"username,omitempty"`
	Password         string `json:"password,omitempty"`
	FirstName        string `json:"firstname,omitempty"`
	LastName         string `json:"lastname,omitempty"`
	Email            string `json:"email,omitempty"`
	Role             string `json:"role,omitempty"`
	GroupUuid        string `json:"groupuuid,omitempty"`
	LastModified     string `json:"last_modified,omitempty"`
	CreationDatetime string `json:"creation_datetime,omitempty"`
}

// GetBodyString
func (u *User) GetBodyString(strType string) string {
	var bodyStr string
	if strType == "Settings" {
		bodyStr = `{"username": "` + u.Username + `","firstname": "` + u.FirstName + `","lastname":"` + u.LastName + `","email":"` + u.Email + `"}`
	} else if strType == "Admin" {
		bodyStr = `{"username": "` + u.Username + `","firstname": "` + u.FirstName + `","lastname":"` + u.LastName + `","email":"` + u.Email + `","password":"` + u.Password + `","groupuuid":"` + u.GroupUuid + `","role":"` + u.Role + `"}`
	}
	return bodyStr
}

// Get Users
func (u *User) GetAll(auth Auth) []User {
	dataReq := data.DatabaseRequest{AuthToken: auth.AuthToken}
	d := dataReq.GetUsers("")
	d.Resolve()
	return loadUsers(d.Res)
}

// Get a User
func (u *User) Get(auth Auth) {
	dataReq := data.DatabaseRequest{AuthToken: auth.AuthToken}
	d := dataReq.GetUsers(u.Uuid)
	d.Resolve()
	user := loadUser(d.Res)
	u.Id = user.Id
	u.Username = user.Username
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.Role = user.Role
	u.GroupUuid = user.GroupUuid
	u.LastModified = user.LastModified
	u.CreationDatetime = user.CreationDatetime
}

// Create a User
func (u *User) Create(auth Auth) {
	dataReq := data.DatabaseRequest{AuthToken: auth.AuthToken}
	d := dataReq.PostUser(u.GetBodyString("Admin"))
	d.Resolve()
	user := loadUser(d.Res)
	u.Id = user.Id
	u.Uuid = user.Uuid
	u.Username = user.Username
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.Role = user.Role
	u.GroupUuid = user.GroupUuid
	u.LastModified = user.LastModified
	u.CreationDatetime = user.CreationDatetime
}

// Update a User
func (u *User) Update(auth Auth, strType string) {
	dataReq := data.DatabaseRequest{AuthToken: auth.AuthToken}
	d := dataReq.PatchUser(u.Uuid, u.GetBodyString(strType))
	d.Resolve()
	user := loadUser(d.Res)
	u.Id = user.Id
	u.Uuid = user.Uuid
	u.Username = user.Username
	u.FirstName = user.FirstName
	u.LastName = user.LastName
	u.Email = user.Email
	u.LastModified = user.LastModified
	u.CreationDatetime = user.CreationDatetime
}

// Delete a User
func (u *User) Delete(auth Auth) {
	dataReq := data.DatabaseRequest{AuthToken: auth.AuthToken}
	_ = dataReq.DeleteUser(u.Uuid)
}

// Todo is a root struct that is used to store the json encoded data for/from a mongodb todos doc.
type Todo struct {
	Id               string `json:"id,omitempty"`
	Uuid             string `json:"uuid,omitempty"`
	Name             string `json:"name,omitempty"`
	Completed        string `json:"completed,omitempty"`
	Due              string `json:"due,omitempty"`
	Description      string `json:"description,omitempty"`
	UserUuid         string `json:"useruuid,omitempty"`
	GroupUuid        string `json:"groupuuid,omitempty"`
	LastModified     string `json:"last_modified,omitempty"`
	CreationDatetime string `json:"creation_datetime,omitempty"`
}

func (t *Todo) GetBodyString() string {
	// TODO: Nick - Finish Building Out this body String
	bodyStr := `{"name": "` + t.Name + `"}`
	return bodyStr
}

// Get a Group
func (t *Todo) Get(auth Auth) {
	dataReq := data.DatabaseRequest{AuthToken: auth.AuthToken}
	d := dataReq.GetUsers(t.Uuid)
	d.Resolve()
	user := loadUser(d.Res)
	// TODO - Nick Add Rest of Mapping Below
	t.Id = user.Id
	t.GroupUuid = user.GroupUuid
	t.LastModified = user.LastModified
	t.CreationDatetime = user.CreationDatetime
}

// Get Todos
func (t *Todo) GetAll(auth Auth) []Todo {
	dataReq := data.DatabaseRequest{AuthToken: auth.AuthToken}
	d := dataReq.GetTodos("")
	d.Resolve()
	return loadTodos(d.Res)
}

// Create a Todo
func (t *Todo) Create(auth Auth) {
	dataReq := data.DatabaseRequest{AuthToken: auth.AuthToken}
	d := dataReq.PostTodo(t.GetBodyString())
	d.Resolve()
	todo := loadTodo(d.Res)
	// TODO - Nick Add Rest of Mapping Below
	t.Id = todo.Id
	t.Uuid = todo.Uuid
	t.Name = todo.Name
	t.LastModified = todo.LastModified
	t.CreationDatetime = todo.CreationDatetime
}

// Update a Todo
func (t *Todo) Update(auth Auth) {
	dataReq := data.DatabaseRequest{AuthToken: auth.AuthToken}
	d := dataReq.PatchTodo(t.Uuid, t.GetBodyString())
	d.Resolve()
	todo := loadTodo(d.Res)
	// TODO - Nick Add Rest of Mapping Below
	t.Id = todo.Id
	t.Uuid = todo.Uuid
	t.Name = todo.Name
	t.LastModified = todo.LastModified
	t.CreationDatetime = todo.CreationDatetime
}

// Delete a Todo
func (t *Todo) Delete(auth Auth) {
	dataReq := data.DatabaseRequest{AuthToken: auth.AuthToken}
	_ = dataReq.DeleteTodo(t.Uuid)
}
