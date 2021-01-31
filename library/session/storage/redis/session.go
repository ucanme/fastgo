package redis

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	uuid "github.com/satori/go.uuid"
	"github.com/ucanme/fastgo/library/session/common"
)

type RedisStore struct {
}

var connPool *redis.Pool
func Init()  {
	connPool = Conn()

}

func (s RedisStore)Set(key string, value common.Session) error  {
	data,_ := json.Marshal(value)
	_,err :=connPool.Get().Do("set","key",string(data))
	return err
}

func (s RedisStore) Get(key string) (common.Session,error) {
	str,err := redis.String(connPool.Get().Do("get",key))
	sess := common.Session{}
	if err != nil{
		return sess,err
	}
	err = json.Unmarshal([]byte(str),&sess)
	return sess,err
}


func (s RedisStore) Delete(key string) error  {
	_,err:= connPool.Get().Do("delete",key)
	return err

}

func (s RedisStore) SessionID() string{
	return "redis"+uuid.NewV4().String()
}


