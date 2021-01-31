package session

import "github.com/ucanme/fastgo/library/session/common"

type SessionInterface interface {
	Set(key string, value common.Session) error //设置Session
	Get(key string) (common.Session,error) //获取Session
	Delete(key string) error     //删除Session
	SessionID() string                //当前SessionID
}

