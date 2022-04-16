/*
Author: Connor Sanders
Copyright: Connor Sanders 2020
Version: 0.0.1
Released: 12/10/2020

-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
		Golang Frontend Boilerplate V0.0.1
-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
*/

package web

import (
	"fmt"
	rice "github.com/GeertJohan/go.rice"
	"github.com/julienschmidt/httprouter"
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"spa_app/pkg/models"
	"spa_app/pkg/session"
	"strings"
)

// splitPathURL
func splitPathURL(r *http.Request) []string {
	reqUrl := r.URL.Path
	var splitURL []string
	splitURL = strings.Split(reqUrl, "/")
	splitURL = splitURL[1:]
	return splitURL
}

// Pages
type Pages struct {
	SessionManager *session.Manager
	TemplateMap    template.FuncMap
	Templates      *template.Template
	TemplateBox    *rice.Box
	Statics        *rice.HTTPBox
}

// InitializePages
func InitializePages(sm *session.Manager) Pages {
	templateMap := template.FuncMap{
		"Upper": func(s string) string {
			return strings.ToUpper(s)
		},
	}
	templates := template.New("").Funcs(templateMap)
	var templateBox *rice.Box
	// Load and parse templates (from binary or disk)
	return Pages{
		SessionManager: sm,
		TemplateMap:    templateMap,
		Templates:      templates,
		TemplateBox:    templateBox,
	}
}

// InitializeTemplates
func (p *Pages) InitializeTemplates() {
	// newTemplate
	var templateBox *rice.Box
	newTemplate := func(path string, fileInfo os.FileInfo, _ error) error {
		if path == "" {
			return nil
		}
		/*
		 * takeRelativeTo function will take the absolute path 'path' which is by default passed to
		 * our 'newTemplate' by Walk function, and will eliminate the intial part of the path up to the end of the
		 * specified directory 'afterDir' ('templates' in this case). Then it will return the rest starting from
		 * the very end of afterDir. If the specified afterDir has more than 1 occurances in the path,
		 * only the first occurance will be considered and the other occurances will be ignored.
		 * eg, If path = "/home/Projects/go/website/templates/html/index.html", then
		 * relativPath := takeRelativeTo(path, "templates") returns "/html/index.html" ;
		 * If path = "/home/Projects/go/website/templates/testing.html", then ;
		 * relativPath := takeRelativeTo(path, "templates") returns "/testing.html" ;
		 * If path = "/home/Projects/go/website/templates/html/templates/components/footer.html", then
		 * relativPath := takeRelativeTo(path, "templates") returns "/html/templates/components/footer.html" .
		 */
		takeRelativeTo := func(givenpath string, afterDir string) string {
			if strings.Contains(givenpath, afterDir+string(filepath.Separator)) {
				wantedpart := strings.SplitAfter(givenpath, afterDir)[1:]
				return filepath.Join(wantedpart...)
			}
			return givenpath
		}
		//if path is a directory, skip Parsing template. Trying to Parse a template from a directory caused an error, now fixed.
		if !fileInfo.IsDir() {
			//get relative path starting from the end of 'templates' .
			relativPath := takeRelativeTo(path, "templates")
			templateString, err := templateBox.String(relativPath)
			if err != nil {
				log.Panicf("Unable to extract: path=%s, err=%s", relativPath, err)
			}
			if _, err = p.Templates.New(filepath.Join("templates", relativPath)).Parse(templateString); err != nil {
				log.Panicf("Unable to parse: path=%s, err=%s", relativPath, err)
			}
		}
		return nil
	}
	templateBox = rice.MustFindBox("templates")
	templateBox.Walk("", newTemplate)
	static := rice.MustFindBox("static").HTTPBox()
	p.Statics = static
	p.TemplateBox = templateBox
}

// renderTemplate - Render a template given a model
func (p *Pages) renderTemplate(w http.ResponseWriter, tmpl string, pInt interface{}) {
	err := p.Templates.ExecuteTemplate(w, tmpl, pInt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// authCheck
func (p *Pages) authCheck(r *http.Request) (models.Auth, *http.Cookie) {
	cookie, err := r.Cookie("SessionID")
	auth := models.Auth{Authenticated: false}
	authenticated := p.SessionManager.IsLoggedIn(r)
	if err != nil || !authenticated {
		return auth, cookie
	}
	auth = p.SessionManager.GetSession(cookie)
	return auth, cookie
}

// Protected
func (p *Pages) Protected(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth, _ := p.authCheck(r)
		model := models.LoginModel{Name: "login", Title: "Login", Auth: auth}
		model.BuildRoute()
		if !auth.Authenticated {
			p.renderTemplate(w, "templates/login.html", &model)
			return
		} else {
			next.ServeHTTP(w, r)
		}
	})
}

// Render Broken Page - for missing/error routes
func (p *Pages) BrokenPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	p.renderTemplate(w, "templates/missing.html", nil)
}

// Render Index Page
func (p *Pages) IndexPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO - Get Auth from Session Manager Here
	auth, cookie := p.authCheck(r)
	model := models.IndexModel{Name: "register", Title: "Register", Auth: auth}
	if !auth.Authenticated {
		p.renderTemplate(w, "templates/index.html", &model)
		return
	}
	auth = p.SessionManager.GetSession(cookie)
	model.Name = "Home"
	model.Title = "Home"
	model.Auth = auth
	p.renderTemplate(w, "templates/index.html", &model)
}

// RegistrationHandler
func (p *Pages) RegistrationHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO - Get Auth from Session Manager Here
	auth, _ := p.authCheck(r)
	model := models.IndexModel{Name: "register", Title: "Register", Auth: auth}
	if r.Method != http.MethodPost {
		p.renderTemplate(w, "templates/index.html", &model)
		return
	}
	register := models.RegistrationForm{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Email:     r.FormValue("email"),
		UserName:  r.FormValue("username"),
		Password:  r.FormValue("password"),
	}
	auth = register.Register()
	if auth.Status != http.StatusCreated {
		p.renderTemplate(w, "templates/index.html", &model)
		return
	}
	cookie := p.SessionManager.NewSession(auth)
	http.SetCookie(w, cookie)
	model = models.IndexModel{Name: "home", Title: "Home", Auth: auth}
	model.Auth = auth
	p.renderTemplate(w, "templates/index.html", &model)
}

// Render LoginPage Page
func (p *Pages) LoginPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO - Get Auth from Session Manager Here
	auth, _ := p.authCheck(r)
	model := models.IndexModel{Name: "home", Title: "Home", Auth: auth}
	if auth.Authenticated {
		p.renderTemplate(w, "templates/index.html", &model)
		return
	}
	lModel := models.LoginModel{Name: "login", Title: "Login", Auth: auth}
	p.renderTemplate(w, "templates/login.html", &lModel)
}

// LoginHandler
func (p *Pages) LoginHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO - Get Auth from Session Manager Here
	auth, cookie := p.authCheck(r)
	model := models.LoginModel{Name: "login", Title: "Login", Auth: auth}
	if r.Method != http.MethodPost {
		p.renderTemplate(w, "templates/login.html", &model)
		return
	}
	login := models.LoginForm{
		UserName: r.FormValue("username"),
		Password: r.FormValue("password"),
	}
	auth = login.Login()
	if auth.Status != http.StatusOK {
		p.renderTemplate(w, "templates/login.html", &model)
		return
	}
	cookie = p.SessionManager.NewSession(auth)
	http.SetCookie(w, cookie)
	model.Auth = auth
	p.renderTemplate(w, "templates/login.html", &model)
}

// LogoutHandler
func (p *Pages) LogoutHandler(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	// TODO - Get Auth from Session Manager Here, Build in Check to ensure logout worked
	auth, cookie := p.authCheck(r)
	model := models.LoginModel{Name: "login", Title: "Login", Auth: auth}
	if auth.Authenticated {
		model.Auth.Invalidate()
		_ = p.SessionManager.DeleteSession(cookie)
	}
	p.renderTemplate(w, "templates/login.html", &model)
}

// Render About Page
func (p *Pages) AboutPage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth, _ := p.authCheck(r)
	model := models.AboutModel{Title: "About", SubRoute: ps.ByName("child"), Name: "about", Auth: auth}
	model.BuildRoute()
	if !auth.Authenticated {
		p.renderTemplate(w, "templates/about.html", &model)
		return
	}
	model.Auth = auth
	p.renderTemplate(w, "templates/about.html", &model)
}

// Render Variable Page
func (p *Pages) VariablePage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	auth, _ := p.authCheck(r)
	model := models.VariableModel{Title: "Variable", SubRoute: ps.ByName("child"), Name: "variable", Auth: auth}
	model.BuildRoute()
	if !auth.Authenticated {
		p.renderTemplate(w, "templates/variable.html", &model)
		return
	}
	model.Auth = auth
	p.renderTemplate(w, "templates/variable.html", &model)
}

// Render Account Page
func (p *Pages) AccountPage(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	fmt.Println("VarsTEST:", params.ByName("child"))
	model := models.AccountModel{Title: "Account", SubRoute: params.ByName("child"), Name: "account", Auth: auth}
	model.Initialize()
	if !auth.Authenticated {
		lModel := models.LoginModel{Title: "Login", Name: "login", Auth: auth}
		p.renderTemplate(w, "templates/login.html", &lModel)
		return
	}
	if model.SubRoute == "settings" {
		model.Title = "Account Settings"
		model.Name = "Account Settings"

	}
	p.renderTemplate(w, "templates/account.html", &model)
}

// AccountSettingsHandler
func (p *Pages) AccountSettingsHandler(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.authCheck(r)
	lModel := models.LoginModel{Name: "Login", Title: "Login", Auth: auth}
	if r.Method != http.MethodPatch {
		p.renderTemplate(w, "templates/login.html", &lModel)
		return
	}
	user := models.User{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Email:     r.FormValue("email"),
		Username:  r.FormValue("username"),
	}
	settings := models.SettingsForm{}
	settings.Load(user)
	form := models.InitializeSettingsForm(user)
	status := settings.UpdateSettings(auth)
	model := models.AccountModel{Name: "account", Title: "Account Settings", SubRoute: "settings", Auth: auth}
	model.BuildRoute()
	if status != http.StatusOK {
		p.renderTemplate(w, "templates/index.html", &model)
		return
	}
	model.Form = form
	p.renderTemplate(w, "templates/account.html", &model)
}

// Render Admin Page
func (p *Pages) AdminPage(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	//fmt.Println("VarsTEST:", params.ByName("id"))
	//vars returns and empty map array
	subRoute := params.ByName("child")
	updateId := params.ByName("id")
	fmt.Println("ID Admin:", updateId)
	//the subroute gives an index error and crashes the app if sURL[1]. If it is zero however it will work but it will never
	//work for create since that needs to be vars 1 for the if to catch it.
	fmt.Println("subRoute1:", subRoute)
	model := models.AdminModel{
		Name:     "admin",
		Title:    "Admin Settings",
		Route:    "admin",
		SubRoute: subRoute,
		Auth:     auth,
		Id:       updateId,
		Method:   "GET",
	}
	model.Initialize()
	// 3: TODO Render EDIT FORM based on subRoute (either groups or users in this scenario)
	if !auth.Authenticated {
		lModel := models.LoginModel{Title: "Login", Name: "login", Auth: auth}
		p.renderTemplate(w, "templates/login.html", &lModel)
		return
	}
	if model.SubRoute == "usermenu" {
		model.Title = "Admin Settings"
		model.Name = "Admin Settings"
	}

	p.renderTemplate(w, "templates/admin.html", &model)
}

// AdminUsermenuHandler
func (p *Pages) AdminUsermenuHandler(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.authCheck(r)
	lModel := models.LoginModel{Name: "Login", Title: "Login", Auth: auth}
	if r.Method != http.MethodPatch {
		p.renderTemplate(w, "templates/login.html", &lModel)
		return
	}
	user := models.User{
		FirstName: r.FormValue("first_name"),
		LastName:  r.FormValue("last_name"),
		Email:     r.FormValue("email"),
		Username:  r.FormValue("username"),
	}
	settings := models.SettingsForm{}
	settings.Load(user)
	//form := models.InitializeSettingsForm(user)
	status := settings.UpdateSettings(auth)
	model := models.AdminModel{Name: "admin", Title: "Admin Usermenu", SubRoute: "usermenu", Auth: auth}
	model.BuildRoute()
	if status != http.StatusOK {
		p.renderTemplate(w, "templates/index.html", &model)
		return
	}
	//model.Form = form
	p.renderTemplate(w, "templates/admin.html", &model)
}

/*
// UserMenuHandler
func (p *Pages) UsermenuPage(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	fmt.Println("VarsTEST:", params.ByName("child"))
	model := models.AdminModel{Title: "Admin", SubRoute: params.ByName("child"), Name: "admin", Auth: auth}
	model.Initialize()
	if !auth.Authenticated {
		lModel := models.LoginModel{Title: "Login", Name: "login", Auth: auth}
		p.renderTemplate(w, "templates/login.html", &lModel)
		return
	}
	if model.SubRoute == "usermenu" {
		model.Title = "Admin Settings"
		model.Name = "Admin Settings"

	}
	p.renderTemplate(w, "templates/account.html", &model)
}

func (p *Pages) GroupmenuPage(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	fmt.Println("VarsTEST:", params.ByName("child"))
	model := models.AdminModel{Title: "Admin", SubRoute: params.ByName("child"), Name: "admin", Auth: auth}
	model.Initialize()
	if !auth.Authenticated {
		lModel := models.LoginModel{Title: "Login", Name: "login", Auth: auth}
		p.renderTemplate(w, "templates/login.html", &lModel)
		return
	}
	if model.SubRoute == "groupmenu" {
		model.Title = "Admin Settings"
		model.Name = "Admin Settings"

	}
	p.renderTemplate(w, "templates/account.html", &model)
}
*/

/*
======================== ADMIN HANDLER FUNCTIONS ====================================
*/

/*
// CreateUserForm
type CreateUserForm struct {
	FirstName    string                        `json:"firstname,omitempty"`
	LastName     string                        `json:"lastname,omitempty"`
	Email        string                        `json:"email,omitempty"`
	UserName     string                        `json:"username,omitempty"`
	Role         string                        `json:"role,omitempty"`
	GroupUuid    string                        `json:"groupuuid,omitempty"`
	Password     string                        `json:"password,omitempty"`
	CPassword    string                        `json:"cpassword,omitempty"`
}
*/

// CreateHandler
func (p *Pages) CreateHandler(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	subRoute := params.ByName("child")
	id := params.ByName("id")
	model := models.AdminModel{
		Name:     "admin",
		Route:    "admin",
		SubRoute: subRoute,
		Auth:     auth,
		Id:       id,
		Method:   "POST",
	}
	// Create New User
	if subRoute == "users" {
		createUserForm := models.CreateUserForm{}
		createUserForm.LoadRequest(r)
		model.Title = "Create New User"
		model.UserTable.CreateUserForm = createUserForm
		model.UserTable.CreateUserForm.Create(auth)
		fmt.Println("CREATE USER")
		// Create New Group
	} else if subRoute == "groups" {
		createGroupForm := models.CreateGroupForm{}
		createGroupForm.LoadRequest(r)
		model.Title = "Create New Group"
		model.GroupTable.CreateGroupForm = createGroupForm
		model.GroupTable.CreateGroupForm.Create(auth)
		fmt.Println("CREATE GROUP")
	}
	model.Id = ""
	model.Method = "GET"
	model.Initialize()
	p.renderTemplate(w, "templates/admin.html", &model)
}

// UpdateHandler
func (p *Pages) UpdateHandler(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	subRoute := params.ByName("child")
	updateId := params.ByName("id")
	model := models.AdminModel{
		Name:     "admin",
		Route:    "admin",
		SubRoute: subRoute,
		Auth:     auth,
		Id:       updateId,
		Method:   "PATCH",
	}
	// Update User
	if subRoute == "users" && updateId != "" {
		updateUserForm := models.UpdateUserForm{}
		updateUserForm.LoadRequest(r)
		model.Title = "Update User: " + updateUserForm.UserName
		model.UserTable.UpdateUserForm = updateUserForm
		model.UserTable.UpdateUserForm.Update(auth, model.Id)
		fmt.Println("UPDATE USER")
		// Update Group
	} else if subRoute == "groups" && updateId != "" {
		updateGroupForm := models.UpdateGroupForm{}
		updateGroupForm.LoadRequest(r)
		model.Title = "Update Group: " + updateGroupForm.Name
		model.GroupTable.UpdateGroupForm = updateGroupForm
		model.GroupTable.UpdateGroupForm.Update(auth, model.Id)
		fmt.Println("UPDATE GROUP")
	}
	model.Id = ""
	model.Method = "GET"
	model.Initialize()
	p.renderTemplate(w, "templates/admin.html", &model)
}

// DeleteHandler
func (p *Pages) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	auth, _ := p.authCheck(r)
	params := httprouter.ParamsFromContext(r.Context())
	subRoute := params.ByName("child")
	delId := params.ByName("id")
	model := models.AdminModel{
		Name:     "admin",
		Route:    "admin",
		SubRoute: subRoute,
		Auth:     auth,
		Id:       delId,
		Method:   "DELETE",
	}
	// Delete User
	if subRoute == "users" && delId != "" {
		fmt.Println("DELETE USER")
		// Delete Group
	} else if subRoute == "groups" && delId != "" {
		fmt.Println("DELETE GROUP")
	}
	p.renderTemplate(w, "templates/admin.html", &model)
}
