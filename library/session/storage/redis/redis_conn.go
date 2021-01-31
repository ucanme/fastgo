package redis

import (
	"github.com/garyburd/redigo/redis"
)

func Conn()  *redis.Pool{
	return &redis.Pool{
		MaxIdle: 80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", ":6379")
			if err != nil {
				panic(err.Error())
			}

			if _, err := c.Do("AUTH", "123456"); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
	}
}
