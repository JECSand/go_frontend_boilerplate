/*
Author: Connor Sanders
Copyright: Connor Sanders 2020
Version: 0.0.1
Released: 12/10/2020

-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
		Golang Frontend Boilerplate V0.0.1
-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
*/

package session

import (
	"fmt"
	"github.com/gofrs/uuid"
	"math/rand"
	"net/http"
	"spa_app/pkg/models"
	"strconv"
	"strings"
	"time"
)

// generateUuid
func generateUuid() string {
	curId, err := uuid.NewV4()
	if err != nil {
		panic(err)
	}
	return curId.String()
}

// Manager
type Manager struct {
	RedisManager *RedisManager
}

// NewManager
func NewManager() *Manager {
	redisManager := InitRedisClient()
	return &Manager{RedisManager: redisManager}
}

// lookupSessionID
func (m *Manager) lookupSessionID(checkId string) (string, error) {
	val, err := m.RedisManager.Get(checkId)
	if err != nil {
		return "", err
	}
	z := strings.Split(val, ":")
	if z[0] == "" {
		return "", nil
	}
	return z[1], nil
}

// IsLoggedIn
func (m *Manager) IsLoggedIn(r *http.Request) bool {
	var expectedSessionID string
	cookie, err := r.Cookie("SessionID")
	if err != nil {
		fmt.Println(err)
		return false
	}
	// TODO - Nick - GET IP FROM r AND COMPARE WITH IP FROM  m.lookupSessionID
	expectedSessionID, err = m.lookupSessionID(cookie.Value)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if expectedSessionID != "" {
		return true
	}
	return false
}

// NewSession
func (m *Manager) NewSession(auth models.Auth) *http.Cookie {
	authStr := auth.GetAuthString()
	checkUuid := generateUuid()
	checkUuid = strings.Replace(checkUuid, "-", "", -1)
	cookieValue := checkUuid + ":" + SHA(checkUuid+strconv.Itoa(rand.Intn(100000000)))
	expire := time.Now().AddDate(0, 0, 1)
	err := m.RedisManager.Set(cookieValue, authStr)
	if err != nil {
		panic(err)
	}
	return &http.Cookie{Name: "SessionID", Value: cookieValue, Expires: expire, HttpOnly: true}
}

// GetSession
func (m *Manager) GetSession(cookie *http.Cookie) models.Auth {
	sessionID := cookie.Value
	authStr, err := m.RedisManager.Get(sessionID)
	if err != nil {
		panic(err)
	}
	auth := models.Auth{Authenticated: false}
	auth.LoadAuthString(authStr)
	return auth
}

// DeleteSession
func (m *Manager) DeleteSession(cookie *http.Cookie) error {
	sessionID := cookie.Value
	err := m.RedisManager.Del(sessionID)
	if err != nil {
		return err
	}
	return nil
}
