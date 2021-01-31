package redis

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"testing"
)

func Test_Conn(t *testing.T)  {
	pool := Conn()
	conn := pool.Get()
	ret,err := redis.String(conn.Do("set","key","val12321312`		1"))
	if err!=nil{
		t.Logf("%s",err)
	}
	fmt.Println(ret)

	reply,err := redis.String(conn.Do("get","key"))
	if err!=nil{
		t.Logf("%s",err)
	}
	fmt.Println(reply)
}