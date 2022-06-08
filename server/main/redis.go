package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

var pool *redis.Pool

func initPool(maxIdle, maxActive int, idleTimeout time.Duration, address string) {
	pool = &redis.Pool{
		MaxIdle:     maxIdle,
		MaxActive:   maxActive, //最大连接数量，0表示没有限制
		IdleTimeout: idleTimeout,
		Dial: func() (redis.Conn, error) { //初始化连接代码
			return redis.Dial("tcp", address)
		},
	}
}
