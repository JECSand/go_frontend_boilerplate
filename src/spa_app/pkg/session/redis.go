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
	"context"
	"github.com/go-redis/redis"
	"os"
)

// RedisManager
type RedisManager struct {
	Client *redis.Client
	Ctx    context.Context
}

// InitRedisClient
func InitRedisClient() *RedisManager {
	var rm RedisManager
	var ctx = context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_HOST"),
		Password: os.Getenv("REDIS_PASSWORD"), // no password set
		DB:       0,                           // use default DB
	})
	rm.Client = rdb
	rm.Ctx = ctx
	return &rm
}

// Set
func (rm *RedisManager) Set(key string, value string) error {
	err := rm.Client.Set(rm.Ctx, key, value, 0).Err()
	if err != nil {
		return err
	}
	return nil
}

// Get
func (rm *RedisManager) Get(key string) (string, error) {
	val, err := rm.Client.Get(rm.Ctx, key).Result()
	if err != nil {
		return val, err
	}
	return val, nil
}

// Del
func (rm *RedisManager) Del(key string) error {
	_, err := rm.Client.Del(rm.Ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}
