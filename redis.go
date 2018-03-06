package main

import (
	"log"

	"github.com/garyburd/redigo/redis"
)

var rdc redis.Conn

// RedisConnect is used to create a connection to Redis and return that connection string.
func RedisConnect() {
	var err error
	rdc, err = redis.Dial("tcp", "localhost:6379")
	if err != nil {
		log.Fatal(err)
	}
}

// RedisSet is used to make a call to Redis and Set a Key and a Value
func RedisSet(k string, v string) interface{} {
	r, err := rdc.Do("SET", k, v)
	if err != nil {
		log.Fatal(err)
	}

	return r
}

// RedisGet is used to make a call to Redis and Get a value from a key
func RedisGet(k string) (string, error) {
	r, err := redis.String(rdc.Do("GET", k))
	return r, err

}
