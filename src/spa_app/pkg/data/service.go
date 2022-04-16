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

// DatabaseRequest
type DatabaseRequest struct {
	AuthToken string
}

// GetGroups
func (dr *DatabaseRequest) GetGroups(uuid string) *Dispatcher {
	d := InitializeDispatcher()
	endpoint := "/groups"
	if uuid != "" {
		endpoint = "/groups/" + uuid
	}
	d.Req = Request{Endpoint: endpoint, Type: "GET", Body: nil}
	d.Req.DefaultHeaders()
	headerEntry := []string{"Auth-Token", dr.AuthToken}
	d.Req.Headers = append(d.Req.Headers, headerEntry)
	d.Execute("")
	return d
}

// PostGroup
func (dr *DatabaseRequest) PostGroup(bodyStr string) *Dispatcher {
	d := InitializeDispatcher()
	var body = []byte(bodyStr)
	endpoint := "/groups"
	d.Req = Request{Endpoint: endpoint, Type: "POST", Body: body}
	d.Req.DefaultHeaders()
	headerEntry := []string{"Auth-Token", dr.AuthToken}
	d.Req.Headers = append(d.Req.Headers, headerEntry)
	d.Execute("")
	return d
}

// PatchGroup
func (dr *DatabaseRequest) PatchGroup(uuid string, bodyStr string) *Dispatcher {
	d := InitializeDispatcher()
	var body = []byte(bodyStr)
	endpoint := "/groups/" + uuid
	d.Req = Request{Endpoint: endpoint, Type: "PATCH", Body: body}
	d.Req.DefaultHeaders()
	headerEntry := []string{"Auth-Token", dr.AuthToken}
	d.Req.Headers = append(d.Req.Headers, headerEntry)
	d.Execute("")
	return d
}

// DeleteGroup
func (dr *DatabaseRequest) DeleteGroup(uuid string) *Dispatcher {
	d := InitializeDispatcher()
	endpoint := "/groups/" + uuid
	d.Req = Request{Endpoint: endpoint, Type: "DELETE", Body: nil}
	d.Req.DefaultHeaders()
	headerEntry := []string{"Auth-Token", dr.AuthToken}
	d.Req.Headers = append(d.Req.Headers, headerEntry)
	d.Execute("")
	return d
}

// GetUsers
func (dr *DatabaseRequest) GetUsers(uuid string) *Dispatcher {
	d := InitializeDispatcher()
	endpoint := "/users"
	if uuid != "" {
		endpoint = "/users/" + uuid
	}
	d.Req = Request{Endpoint: endpoint, Type: "GET", Body: nil}
	d.Req.DefaultHeaders()
	headerEntry := []string{"Auth-Token", dr.AuthToken}
	d.Req.Headers = append(d.Req.Headers, headerEntry)
	d.Execute("")
	return d
}

// PostUser
func (dr *DatabaseRequest) PostUser(bodyStr string) *Dispatcher {
	d := InitializeDispatcher()
	var body = []byte(bodyStr)
	endpoint := "/users"
	d.Req = Request{Endpoint: endpoint, Type: "POST", Body: body}
	d.Req.DefaultHeaders()
	headerEntry := []string{"Auth-Token", dr.AuthToken}
	d.Req.Headers = append(d.Req.Headers, headerEntry)
	d.Execute("")
	return d
}

// PatchUser
func (dr *DatabaseRequest) PatchUser(uuid string, bodyStr string) *Dispatcher {
	d := InitializeDispatcher()
	var body = []byte(bodyStr)
	endpoint := "/users/" + uuid
	d.Req = Request{Endpoint: endpoint, Type: "PATCH", Body: body}
	d.Req.DefaultHeaders()
	headerEntry := []string{"Auth-Token", dr.AuthToken}
	d.Req.Headers = append(d.Req.Headers, headerEntry)
	d.Execute("")
	return d
}

// DeleteUser
func (dr *DatabaseRequest) DeleteUser(uuid string) *Dispatcher {
	d := InitializeDispatcher()
	endpoint := "/users/" + uuid
	d.Req = Request{Endpoint: endpoint, Type: "DELETE", Body: nil}
	d.Req.DefaultHeaders()
	headerEntry := []string{"Auth-Token", dr.AuthToken}
	d.Req.Headers = append(d.Req.Headers, headerEntry)
	d.Execute("")
	return d
}

// GetTodos
func (dr *DatabaseRequest) GetTodos(uuid string) *Dispatcher {
	d := InitializeDispatcher()
	endpoint := "/todos"
	if uuid != "" {
		endpoint = "/todos/" + uuid
	}
	d.Req = Request{Endpoint: endpoint, Type: "GET", Body: nil}
	d.Req.DefaultHeaders()
	headerEntry := []string{"Auth-Token", dr.AuthToken}
	d.Req.Headers = append(d.Req.Headers, headerEntry)
	d.Execute("")
	return d
}

// PostTodo
func (dr *DatabaseRequest) PostTodo(bodyStr string) *Dispatcher {
	d := InitializeDispatcher()
	var body = []byte(bodyStr)
	endpoint := "/todos"
	d.Req = Request{Endpoint: endpoint, Type: "POST", Body: body}
	d.Req.DefaultHeaders()
	headerEntry := []string{"Auth-Token", dr.AuthToken}
	d.Req.Headers = append(d.Req.Headers, headerEntry)
	d.Execute("")
	return d
}

// PatchTodo
func (dr *DatabaseRequest) PatchTodo(uuid string, bodyStr string) *Dispatcher {
	d := InitializeDispatcher()
	var body = []byte(bodyStr)
	endpoint := "/todos/" + uuid
	d.Req = Request{Endpoint: endpoint, Type: "PATCH", Body: body}
	d.Req.DefaultHeaders()
	headerEntry := []string{"Auth-Token", dr.AuthToken}
	d.Req.Headers = append(d.Req.Headers, headerEntry)
	d.Execute("")
	return d
}

// DeleteTodo
func (dr *DatabaseRequest) DeleteTodo(uuid string) *Dispatcher {
	d := InitializeDispatcher()
	endpoint := "/todos/" + uuid
	d.Req = Request{Endpoint: endpoint, Type: "DELETE", Body: nil}
	d.Req.DefaultHeaders()
	headerEntry := []string{"Auth-Token", dr.AuthToken}
	d.Req.Headers = append(d.Req.Headers, headerEntry)
	d.Execute("")
	return d
}
