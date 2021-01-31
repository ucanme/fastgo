package session

import (
	"fmt"
	"github.com/ucanme/fastgo/library/session/storage/redis"
	"testing"
)

func TestNewProvider(t *testing.T) {
	redis.Init()
	p := NewProvider("test","redis")
	if p==nil{
		t.Logf("%s","provider is nil")
	}
	RegisterProvider("test",p)

	s,err:= p.SessionInit("key123")
	if err!=nil{
		t.Logf("%s",err.Error())
	}
	fmt.Println(s)
}

func TestNewManager(t *testing.T) {
	m,err := NewManager("test","login_cookie",1000)
	if err!=nil{
		t.Logf("%s",err.Error())
	}
	fmt.Println(m)
}