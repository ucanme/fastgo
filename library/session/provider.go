package session

import (
	"github.com/ucanme/fastgo/library/session/common"
	"github.com/ucanme/fastgo/library/session/storage/redis"
)

type Provider interface {
	SessionInit(sid string) (common.Session, error)
	SessionRead(sid string) (common.Session, error)
	SessionDestroy(sid string) error
	SessionGC(maxLifeTime int64)
}


var providers = make(map[string]Provider)

//注册一个能通过名称来获取的 session provider 管理器
func RegisterProvider(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provider is nil")
	}

	if _, p := providers[name]; p {
		panic("session: Register provider is existed")
	}

	providers[name] = provider
}


//生成新的provider
func NewProvider(name string,provider_type string) Provider {
	var provider Provider
	switch provider_type {
	case "redis":
		var sess SessionInterface
		store := redis.RedisStore{}
		sess = store
		provider = &sessProvider{
			session: &sess,
		}
		return provider
	default:
		return nil
	}
}
