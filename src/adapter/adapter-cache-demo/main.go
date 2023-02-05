package main

import (
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis"
)

// Cache is the adapter contract between the application and the cache library.
type Cache interface {
	Put(key string, value interface{})
	PutAll(map[string]interface{})
	Get(key string) interface{}
	GetAll(keys []string) map[string]interface{}
	Clean(key string)
	CleanAll()
}

// RedisCache holds a Redis connection pool.
type RedisCache struct {
	conn *redis.Pool
}

// Put adds an entry in the cache.
func (rc *RedisCache) Put(key string, value interface{}) {
	if _, err := rc.conn.Get().Do("SET", key, value); err != nil {
		fmt.Println(err)
	}
}

// PutAll adds the entries of a map in the cache.
func (rc *RedisCache) PutAll(entries map[string]interface{}) {
	c := rc.conn.Get()
	for k, v := range entries {
		if err := c.Send("SET", k, v); err != nil {
			fmt.Println(err)
		}
	}

	if err := c.Flush(); err != nil {
		fmt.Println(err)
	}
}

// Get gets an entry from the cache.
func (rc *RedisCache) Get(key string) interface{} {
	value, err := redis.String(rc.conn.Get().Do("GET", key))
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return value
}

// GetAll gets all the entries of a map from the cache.
func (rc *RedisCache) GetAll(keys []string) map[string]interface{} {
	// Converts []string to []interface{} since Go doesn't do it explicitly
	// because it doesn't want the syntax to hide a O(n) operation.
	intKeys := make([]interface{}, len(keys))
	for i, _ := range keys {
		intKeys[i] = keys[i]
	}

	c := rc.conn.Get()
	entries := make(map[string]interface{})
	values, err := redis.Strings(c.Do("MGET", intKeys...))
	if err != nil {
		fmt.Println(err)
		return entries
	}

	for i, k := range keys {
		entries[k] = values[i]
	}

	return entries
}

// Clean cleans a entry from the cache.
func (rc *RedisCache) Clean(key string) {
	if _, err := rc.conn.Get().Do("DEL", key); err != nil {
		fmt.Println(err)
	}
}

// CleanAll cleans the entire cache.
func (rc *RedisCache) CleanAll() {
	if _, err := rc.conn.Get().Do("FLUSHDB"); err != nil {
		fmt.Println(err)
	}
}

// GetCachingMechanism initializes and returns a caching mechanism.
func NewRedisCache() Cache {
	cache := &RedisCache{
		conn: &redis.Pool{
			MaxIdle:     7,
			MaxActive:   30,
			IdleTimeout: 60 * time.Second,
			Dial: func() (redis.Conn, error) {
				conn, err := redis.Dial("tcp", "localhost:6379")
				if err != nil {
					fmt.Println(err)
					return nil, err
				}

				if _, err := conn.Do("SELECT", 0); err != nil {
					conn.Close()
					fmt.Println(err)
					return nil, err
				}

				return conn, nil
			},
		},
	}
	return cache
}

func main() {
	rc := NewRedisCache()
	rc.Put("网管叨逼叨", "rub fish")
}
