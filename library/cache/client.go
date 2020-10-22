package cache

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

func Connect(address, pass string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 240 * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", address, redis.DialPassword(pass))
			if err != nil {
				return nil, err
			}

			return c, nil
		},
	}
}

// string
func Set(pool *redis.Pool, key, val string) (bool, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("SET", key, val))
}

func Get(pool *redis.Pool, key string) ([]byte, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Bytes(conn.Do("GET", key))
}

func Del(pool *redis.Pool, key string) (int, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("DEL", key))
}

func Incr(pool *redis.Pool, key string) (int, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("INCR", key))
}

func IncrBy(pool *redis.Pool, key string, increment int) (int, error) {
	conn := pool.Get()
	defer conn.Close()
	args := redis.Args{}.Add(key).Add(increment)
	return redis.Int(conn.Do("INCRBY", args...))
}

func Expire(pool *redis.Pool, key string, ttl int64) (int, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Int(conn.Do("EXPIRE", key, ttl))
}

func Setnx(pool *redis.Pool, key, value string) ([]byte, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Bytes(conn.Do("SETNX", key, value))
}

func Setex(pool *redis.Pool, key, value string, ttl int64) ([]byte, error) {
	conn := pool.Get()
	defer conn.Close()
	return redis.Bytes(conn.Do("SETEX", key, value, ttl))
}
