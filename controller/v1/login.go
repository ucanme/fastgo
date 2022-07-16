package v1

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"github.com/ucanme/fastgo/conf"
	"github.com/ucanme/fastgo/consts"
	"github.com/ucanme/fastgo/controller/response"
	"github.com/ucanme/fastgo/internal/session"
	"github.com/ucanme/fastgo/library/db"
	"github.com/ucanme/fastgo/library/log"
	"github.com/ucanme/fastgo/models"
	"github.com/ucanme/fastgo/util"
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
	cookie, err := c.Cookie("login_session")
	log.LogNotice(map[string]interface{}{"cookie":cookie,"err":err})
	sid, err := url.QueryUnescape(cookie)
	log.LogNotice(map[string]interface{}{"sid":sid,"err":err})
	if sid == ""{
		log.LogNotice(map[string]interface{}{"sid":sid,"err":err})
		response.Fail(c,consts.ACCOUTN_NOT_LOGIN,"请登陆")
		return
	}
	err = session.Manager.SessionDestroy(c)
	log.LogNotice(map[string]interface{}{"delete session err":err})
	response.Success(c,map[string]interface{}{})
	return
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


type CmdReq struct {
	CmdType string `json:"cmd_type"`
	Payload string `json:"payload"`
}

type CmdResp struct {
	ErrorCode int `json:"error_code"`
	ErrorMsg string `json:"error_msg"`
}
func Cmd(c *gin.Context)  {
	input := CmdReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}

	data,_ := json.Marshal(input)

	resp,err := util.Post(conf.Config.PlatformApi.Host+"/v1/cmd",data, map[string]string{}, map[string]string{})
	fmt.Println("err----",err)
	if err != nil{
		response.Fail(c, consts.REQUEST_FAIL_CODE, consts.REQUEST_FAIL.Error())
		return
	}
	apiResp := CmdResp{}
	err = json.Unmarshal(resp,&apiResp)
	if err == nil{
		response.Fail(c, consts.REQUEST_FAIL_CODE, consts.REQUEST_FAIL.Error())
		return
	}
	if apiResp.ErrorCode != 0 {
		response.Fail(c, consts.REQUEST_FAIL_CODE, apiResp.ErrorMsg)
		return
	}
	response.Success(c,nil)
}


type ReportReq struct {
	EventType string `json:"event_type"`
	Payload string `json:"payload"`
}


type MoveUnit []struct {
	MoveUnitSn string `json:"move_unit_sn"`
	CurrentStationCode string `json:"current_station_code"`
	IsInStation int `json:"is_in_station"`
	Status int `json:"status"`
	Soc int `json:"soc"`
	Speed float64 `json:"speed"`
	RingAngle float32 `json:"ring_angle"`
	RingStatus int `json:"ring_status"`
	WorkDuration int `json:"work_duration"`
	ProductionLineId int `json:"production_line_id" gorm:"column:production_line_id"`
}

func Report(c *gin.Context)  {
	input := ReportReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	moveUnitParamList := MoveUnit{}
	err := json.Unmarshal([]byte(input.Payload),&moveUnitParamList)
	if err != nil{
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}

	if len(moveUnitParamList) == 0{
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}

	moveUnitList := []models.MoveUnit{}
	for _,moveUnitParam := range moveUnitParamList{
		moveUnit := models.MoveUnit{
			MoveUnitSn:         moveUnitParam.MoveUnitSn,
			Soc:                moveUnitParam.Soc,
			Status:             moveUnitParam.Status,
			Speed:              moveUnitParam.Speed,
			CurrentStationCode: moveUnitParam.CurrentStationCode,
			IsInStation:        moveUnitParam.IsInStation,
			RingAngle:          moveUnitParam.RingAngle,
			RingStatus:         moveUnitParam.RingStatus,
			WorkDuration:       moveUnitParam.WorkDuration,
			ProductionLineId:   moveUnitParam.ProductionLineId,
		}
		moveUnitList = append(moveUnitList,moveUnit)
		err = db.DB().Updates(moveUnitList).Error
		if err != nil{
			response.Fail(c, consts.DB_EXEC_ERR_CODE, consts.DB_EXEC_ERR.Error())
			return
		}
	}
	response.Success(c,nil)
}