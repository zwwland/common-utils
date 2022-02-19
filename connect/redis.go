package connect

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type RedisConfig struct {
	Host     string
	Port     string
	Password string
	DB       string
}

func NewRedisPool(rc *RedisConfig) (*redis.Pool, error) {
	p := &redis.Pool{
		Dial: func() (redis.Conn, error) {
			redisUrl := fmt.Sprintf("%s:%d", rc.Host, rc.Port)
			var c redis.Conn
			var err error
			if c, err = redis.Dial("tcp", redisUrl); err != nil {
				panic(err)
			}
			if rc.Password != "" {
				if _, err := c.Do("AUTH", rc.Password); err != nil {
					panic(err)
				}
			}
			return c, nil
		},
		// MaxIdle:    10,
	}

	return p, nil
}
