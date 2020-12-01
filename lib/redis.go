package lib

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
)

var redisPool *redis.Pool

func getRedisPool() *redis.Pool {
	if redisPool == nil {
		redisPool = &redis.Pool{
			Dial: func() (conn redis.Conn, e error) {
				conn, e = redis.Dial("tcp", "127.0.0.1:6379")
				if e != nil {
					_, _ = fmt.Fprintln(gin.DefaultWriter, "redis error : "+e.Error())
					panic(e.Error())
				}
				return
			},
		}
	}
	return redisPool
}

func GetRedis() redis.Conn {
	return getRedisPool().Get()
}
