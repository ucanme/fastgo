package session

import (
	"github.com/ucanme/fastgo/library/session/common"
	"time"
)

type sessProvider struct {
	session *SessionInterface
}

func (s *sessProvider) SessionInit(sid string,value string) (common.Session, error) {
	sess := common.Session{
		Sid:            sid,
		Value:          value,
		LastAccessTime: time.Now().UnixNano()/1e6,
	}
	err := (*s.session).Set(sid,sess)
	return sess ,err
}

func (s *sessProvider)SessionRead(sid string) (common.Session, error){
	v,err := (*s.session).Get(sid)
	return v,err
}


func (s *sessProvider) SessionDestroy(sid string) error{
	return (*s.session).Delete(sid)
}

func (s *sessProvider)	SessionGC(maxLifeTime int64) {

}


