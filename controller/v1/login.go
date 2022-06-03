package v1

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"github.com/ucanme/fastgo/consts"
	"github.com/ucanme/fastgo/controller/response"
	"github.com/ucanme/fastgo/internal/session"
	"github.com/ucanme/fastgo/library/db"
	"github.com/ucanme/fastgo/models"
	"net/url"
)

type loginReq struct {
	UserID string `json:"user_id"`
	Password string `json:"password"`
}

type loginResp struct {
	UserID string `json:"user_id"`
	//SessionKey string `json:"session_key"`
	ErrNo int `json:"errcode"`
	
}

type SessionData struct {
	UserID string `json:"user_id"`
	SessionKey string `json:"session_key"`
}

func LoginOut(c *gin.Context)  {
	cookie, _ := c.Cookie("login_session")
	sid, _ := url.QueryUnescape(cookie)
	if sid == ""{
		response.Fail(c,400,"请登陆")
		return
	}
	session.Manager.SessionDestroy(c)
}


func Login(c *gin.Context)  {
	input := loginReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}

	account := models.Account{}
	err := db.DB().Where("account_id = ?",input.UserID).First(&account).Error

	if account.AccountId == "" || err == gorm.ErrRecordNotFound{
		response.Fail(c, 400, "登陆失败 用户不存存在")
		return
	}

	if err!=nil{
		response.Fail(c, 400, "登陆失败 系统错误")
		return
	}


	passStr := Md5(input.Password)
	if passStr != account.Password{
		response.Fail(c, 400, "登陆失败 密码错误")
		return
	}
	s := SessionData{
		UserID:    account.AccountId,
		//SessionKey: .SessionKey,
	}
	d,_:=json.Marshal(s)
	session.Manager.SessionStart(c,string(d))
	response.Success(c,nil)
}

func Md5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}