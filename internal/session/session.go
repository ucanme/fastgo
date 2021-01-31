package session

import "github.com/ucanme/fastgo/library/session"
var Manager *session.Manager
var err error
func Init()  {
	//注册redis provider
	provider := session.NewProvider("redis_prodvider","redis")
	session.RegisterProvider("redis_prodvider",provider)
	//实例化manager
	Manager,err = session.NewManager("redis_prodvider","login_session",123213)
}