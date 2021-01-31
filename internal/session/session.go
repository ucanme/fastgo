package session

import (
	"fmt"
	"github.com/ucanme/fastgo/library/session"
	"github.com/ucanme/fastgo/library/session/storage/redis"
)
var Manager *session.Manager
var err error
func Init()  {
	//注册redis provider
	redis.Init()
	provider := session.NewProvider("redis")
	session.RegisterProvider("redis_prodvider",provider)
	//实例化manager
	Manager,err = session.NewManager("redis_prodvider","login_session",123213)
	fmt.Println(Manager)
	if err!=nil{
		panic(err)
	}
}