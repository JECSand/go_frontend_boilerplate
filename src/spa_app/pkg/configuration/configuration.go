/*
Author: Connor Sanders
Copyright: Connor Sanders 2020
Version: 0.0.1
Released: 12/10/2020

-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
		Golang Frontend Boilerplate V0.0.1
-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-*-
*/

package configuration

import (
	"encoding/json"
	"os"
)

// Configuration is a struct designed to hold the applications variable configuration settings
type Configuration struct {
	Port           string
	APIHost        string
	SessionManager string
	RedisURL       string
	RedisPassword  string
}

// ConfigurationSettings is a function that reads a json configuration file and outputs a Configuration struct
func ConfigurationSettings(env string) Configuration {
	confFile := "conf.json"
	if env == "test" {
		confFile = "test_conf.json"
	}
	file, _ := os.Open(confFile)
	decoder := json.NewDecoder(file)
	configurationSettings := Configuration{}
	err := decoder.Decode(&configurationSettings)
	if err != nil {
		panic(err)
	}
	return configurationSettings
}

// InitializeEnvironmentals
func (c *Configuration) InitializeEnvironmentals() {
	os.Setenv("PORT", c.Port)
	os.Setenv("API_HOST", c.APIHost)
	os.Setenv("SESSION_MANAGER", c.SessionManager)
	os.Setenv("REDIS_URL", c.RedisURL)
	os.Setenv("REDIS_PASSWORD", c.RedisPassword)
}
