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

import "net/http"

// Form
type Form struct {
	Name   string  `json:"name,omitempty"`
	Type   string  `json:"type,omitempty"`
	Class  string  `json:"class,omitempty"`
	Id     string  `json:"id,omitempty"`
	Fields []Field `json:"fields,omitempty"`
	Button Button  `json:"button,omitempty"`
	Method string  `json:"method,omitempty"`
	Action string  `json:"action,omitempty"`
}

// NewForm
// TODO FINISH THIS NEW FORM FUNCTION
func (f *Form) NewForm(auth Auth, fields []Field, button Button) Form {
	return Form{Fields: fields, Button: button}
}

// Initialize
func InitializeForm(formMeta []string, fieldStrs [][]string, button Button) Form {
	newForm := Form{}
	var fields []Field
	// Field Vector String Array, this is order
	// Name, Class, Id, Type, Label, DefaultVal
	fields = LoadFields(fieldStrs)
	newForm.Name = formMeta[0]
	newForm.Type = formMeta[1]
	newForm.Class = formMeta[2]
	newForm.Id = formMeta[3]
	newForm.Method = formMeta[4]
	newForm.Action = formMeta[5]
	newForm.Fields = fields
	newForm.Button = button
	return newForm
}

// Initialize
func InitializeSettingsForm(user User) Form {
	// Field Vector String Array, this is order
	// NAME, TYPE, CLASS, ID, METHOD, ACTION
	formMeta := []string{"Update", "User", "form1", "form2", "PATCH", ""}
	// Name, Class, Id, Type, Label, DefaultVal
	initStr := []string{"UserName", "update", "username", "text", "Username", user.Username}
	fNameStr := []string{"first_name", "update", "name", "text", "First Name", user.FirstName}
	lNameStr := []string{"last_name", "update", "name", "text", "Last Name", user.LastName}
	EmailStr := []string{"email", "update", "email", "text", "Email", user.Email}
	fieldStrs := [][]string{initStr, fNameStr, lNameStr, EmailStr}
	form := InitializeForm(formMeta, fieldStrs, Button{Name: "update", Class: "form1", Id: "form2", Type: "submit", Label: "Submit"})
	return form
}

// RegistrationForm
type RegistrationForm struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	UserName  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	CPassword string `json:"cpassword,omitempty"`
}

// Register
func (rm *RegistrationForm) Register() Auth {
	auth := Auth{}
	auth.Register(rm.FirstName, rm.LastName, rm.Email, rm.UserName, rm.Password)
	return auth
}

// LoginForm
type LoginForm struct {
	UserName string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
}

// Register
func (rm *LoginForm) Login() Auth {
	auth := Auth{}
	auth.Authenticate(rm.UserName, rm.Password)
	return auth
}

// CreateUserForm
type CreateUserForm struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	UserName  string `json:"username,omitempty"`
	Role      string `json:"role,omitempty"`
	GroupUuid string `json:"groupuuid,omitempty"`
	Password  string `json:"password,omitempty"`
	CPassword string `json:"cpassword,omitempty"`
}

// LoadRequest
func (cuf *CreateUserForm) LoadRequest(r *http.Request) {
	cuf.FirstName = r.FormValue("first_name")
	cuf.LastName = r.FormValue("last_name")
	cuf.Email = r.FormValue("email")
	cuf.UserName = r.FormValue("username")
	cuf.Role = r.FormValue("role")
	cuf.GroupUuid = r.FormValue("group_uuid")
	cuf.Password = r.FormValue("password")
}

// UpdateSettings - Update user for info display
func (cuf *CreateUserForm) Create(auth Auth) int {
	statusCode := 201
	// TODO Add User info to user struct below
	user := User{}
	// Other stuff maybe?
	user.Create(auth)
	// TODO - Capture response status and render a success or error message
	return statusCode
}

// UpdateUserForm
type UpdateUserForm struct {
	Uuid      string `json:"uuid,omitempty"` // HIDDEN FIELD IN HTML MODEL
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	UserName  string `json:"username,omitempty"`
	Role      string `json:"role,omitempty"`
	GroupUuid string `json:"groupuuid,omitempty"`
	Password  string `json:"password,omitempty"`
	CPassword string `json:"cpassword,omitempty"`
	Form      Form   `json:"form,omitempty"`
}

// Initialize
func (uuf *UpdateUserForm) Initialize(uuid string) {
	// Field Vector String Array, this is order
	// NAME, TYPE, CLASS, ID, METHOD, ACTION
	formMeta := []string{"Update", "User", "form1", "form2", "PATCH", ""}
	// Name, Class, Id, Type, Label, DefaultVal
	initStr := []string{"uuid", "update", "uuid", "text", "Uuid", uuid}
	fNameStr := []string{"first_name", "update", "name", "text", "First Name", ""}
	lNameStr := []string{"last_name", "update", "name", "text", "Last Name", ""}
	fieldStrs := [][]string{initStr, fNameStr, lNameStr}
	form := InitializeForm(formMeta, fieldStrs, Button{Name: "update", Class: "form1", Id: "form2", Type: "submit", Label: "Submit"})
	uuf.Form = form

}

// LoadRequest
func (uuf *UpdateUserForm) LoadRequest(r *http.Request) {
	uuf.Uuid = r.FormValue("uuid")
	uuf.FirstName = r.FormValue("first_name")
	uuf.LastName = r.FormValue("last_name")
	uuf.Email = r.FormValue("email")
	uuf.UserName = r.FormValue("username")
	uuf.Role = r.FormValue("role")
	uuf.GroupUuid = r.FormValue("group_uuid")
	uuf.Password = r.FormValue("password")
}

// Load UpdateUserForm when AccountModel is Initialized
func (uuf *UpdateUserForm) Load(user User) {
	uuf.Uuid = user.Uuid
	uuf.UserName = user.Username
	uuf.FirstName = user.FirstName
	uuf.LastName = user.LastName
	uuf.Email = user.Email
	uuf.GroupUuid = user.GroupUuid
	uuf.Role = user.Role
}

// UpdateSettings - Update user for info display
func (uuf *UpdateUserForm) Update(auth Auth, uuid string) int {
	statusCode := 200
	if uuid == "" {
		uuid = auth.UserUuid
	}
	user := User{
		Uuid:      uuid,
		Username:  uuf.UserName,
		FirstName: uuf.FirstName,
		LastName:  uuf.LastName,
		Email:     uuf.Email,
		GroupUuid: uuf.GroupUuid,
		Role:      uuf.Role,
		Password:  uuf.Password,
	}
	user.Update(auth, "Admin")
	// TODO - Capture response status and render a success or error message
	return statusCode
}

// CreateGroupForm
type CreateGroupForm struct {
	Name string `json:"name,omitempty"`
}

// LoadRequest
func (cgf *CreateGroupForm) LoadRequest(r *http.Request) {
	cgf.Name = r.FormValue("name")
}

// UpdateSettings - Update user for info display
func (cgf *CreateGroupForm) Create(auth Auth) int {
	statusCode := 201
	group := Group{
		Name: cgf.Name,
	}
	group.Create(auth)
	// TODO - Capture response status and render a success or error message
	return statusCode
}

// UpdateGroupForm
type UpdateGroupForm struct {
	Uuid string `json:"uuid,omitempty"` // HIDDEN FIELD IN HTML MODEL
	Name string `json:"name,omitempty"`
}

// LoadRequest
func (ugf *UpdateGroupForm) LoadRequest(r *http.Request) {
	ugf.Uuid = r.FormValue("uuid")
	ugf.Name = r.FormValue("name")
}

// Load CreateGroupForm when AccountModel is Initialized
func (ugf *UpdateGroupForm) Load(group Group) {
	ugf.Name = group.Name
}

// UpdateSettings - Update user for info display
func (ugf *UpdateGroupForm) Update(auth Auth, uuid string) int {
	statusCode := 200
	if uuid == "" {
		uuid = auth.GroupUuid
	}
	group := Group{
		Uuid: uuid,
		Name: ugf.Name,
	}
	group.Update(auth)
	// TODO - Capture response status and render a success or error message
	return statusCode
}

// AccountSettingsForm
type AccountSettingsForm struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	UserName  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	CPassword string `json:"cpassword,omitempty"`
}

// Change Account Settings
func (rm *AccountSettingsForm) Update() Auth {
	auth := Auth{}
	auth.Register(rm.FirstName, rm.LastName, rm.Email, rm.UserName, rm.Password)
	return auth
}

// AdminSettingsForm
type AdminUsermenuForm struct {
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
	UserName  string `json:"username,omitempty"`
	Password  string `json:"password,omitempty"`
	CPassword string `json:"cpassword,omitempty"`
}

// Change admin Settings
func (aum *AdminUsermenuForm) Update() Auth {
	auth := Auth{}
	auth.Register(aum.FirstName, aum.LastName, aum.Email, aum.UserName, aum.Password)
	return auth
}

// UpdatePasswordForm
type UpdatePasswordForm struct {
	Password    string `json:"password,omitempty"`
	NewPassword string `json:"newpassword,omitempty"`
	CPassword   string `json:"cpassword,omitempty"`
}

/* Change Account Settings
func (rm *UpdatePasswordForm) Update() Auth {
	auth := Auth{}
	auth.Register(rm.Password, rm.NewPassword)
	return auth
}
*/

// OLD FORM
// SettingsForm
type SettingsForm struct {
	Username  string `json:"username,omitempty"`
	FirstName string `json:"firstname,omitempty"`
	LastName  string `json:"lastname,omitempty"`
	Email     string `json:"email,omitempty"`
}

// Load SettingsForm when AccountModel is Initialized
func (sm *SettingsForm) Load(user User) {
	sm.Username = user.Username
	sm.FirstName = user.FirstName
	sm.LastName = user.LastName
	sm.Email = user.Email
}

// UpdateSettings - Update user for info display
func (sm *SettingsForm) UpdateSettings(auth Auth) int {
	statusCode := 200
	user := User{
		Uuid:      auth.UserUuid,
		Username:  sm.Username,
		FirstName: sm.FirstName,
		LastName:  sm.LastName,
		Email:     sm.Email,
	}
	user.Update(auth, "Settings")
	// TODO - Capture response status and render a success or error message
	return statusCode
}
