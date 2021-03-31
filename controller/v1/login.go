package v1

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ucanme/fastgo/conf"
	"github.com/ucanme/fastgo/consts"
	"github.com/ucanme/fastgo/controller/response"
	"github.com/ucanme/fastgo/internal/session"
	"github.com/ucanme/fastgo/util"
)

type loginReq struct {
	Code string `json:"code" binding:"required"`
}

type loginResp struct {
	OpenId string `json:"openid"`
	SessionKey string `json:"session_key"`
	ErrNo int `json:"errcode"`
	
}

type SessionData struct {
	OpenId string `json:"open_id"`
	SessionKey string `json:"session_key"`
}
func Login(c *gin.Context)  {
	input := loginReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}

	data,err := util.Get("https://api.weixin.qq.com/sns/jscode2session", map[string]string{"js_code":input.Code,"grant_type":"authorization_code","appid":conf.Config.Wechat.ApiKey,"secret":conf.Config.Wechat.ApiSecret})
	fmt.Println("data",string(data),"appid",conf.Config.Wechat.ApiKey)
	if err!=nil{
		fmt.Println(string(data))
		response.Fail(c, 400, "登陆失败")
	}
 	l:=loginResp{}
	err = json.Unmarshal(data,&l)
	if err!=nil || l.ErrNo !=0{
		response.Fail(c, 400, "登陆失败")
		return
	}
	s := SessionData{
		OpenId:    l.OpenId,
		SessionKey: l.SessionKey,
	}
	d,_:=json.Marshal(s)
	session.Manager.SessionStart(c,string(d))
	response.Success(c,nil)
}
