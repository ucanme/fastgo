package redis

import (
	"github.com/garyburd/redigo/redis"
	"github.com/ucanme/fastgo/conf"
)

func Conn()  *redis.Pool{
	return &redis.Pool{
		MaxIdle: 80,
		MaxActive: 12000, // max number of connections
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", conf.Config.Redis.Host)
			if err != nil {
				panic(err.Error())
			}

			if _, err := c.Do("AUTH", conf.Config.Redis.Password); err != nil {
				c.Close()
				return nil, err
			}
			return c, err
		},
	}
}
