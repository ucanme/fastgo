package v1

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ucanme/fastgo/conf"
	"github.com/ucanme/fastgo/consts"
	"github.com/ucanme/fastgo/controller/response"
	"github.com/ucanme/fastgo/internal/session"
	"github.com/ucanme/fastgo/util"
)
import 	"github.com/gookit/validate"

type loginReq struct {
	Code string
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
	valid := validate.Struct(input)
	if !valid.Validate(){
		response.Fail(c, consts.PARAM_ERR_CODE, consts.VALIDATE_ERR.Error())
		return
	}

	data,err := util.Get("https://api.weixin.qq.com/sns/jscode2session?appid="+conf.Config.Wechat.ApiKey+"&secret="+conf.Config.Wechat.ApiSecret+"&js_code=JSCODE&grant_type="+input.Code+"",nil)
	if err!=nil{
		response.Fail(c, 400, "登陆失败")
	}
 	l:=loginResp{}
	err = json.Unmarshal(data,&l)
	if err!=nil || l.ErrNo !=0{
		response.Fail(c, 400, "登陆失败")
	}
	s := SessionData{
		OpenId:    l.OpenId,
		SessionKey: l.SessionKey,
	}
	d,_:=json.Marshal(s)
	session.Manager.SessionStart(c,string(d))
	response.Success(c,nil)
}
