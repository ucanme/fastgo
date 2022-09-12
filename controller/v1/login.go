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
	"github.com/ucanme/fastgo/internal/manager"
	"github.com/ucanme/fastgo/internal/session"
	"github.com/ucanme/fastgo/library/db"
	"github.com/ucanme/fastgo/library/httputil"
	"github.com/ucanme/fastgo/library/log"
	"github.com/ucanme/fastgo/models"
	"github.com/ucanme/fastgo/util"
	"io/ioutil"
	"net/url"
	"strconv"
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
	ErrorCode int `json:"error_no"`
	ErrorMsg string `json:"error_msg"`
}

func Cmd(c *gin.Context) {
	input := CmdReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}

	if input.CmdType == "start_arm" || input.CmdType == "stop_arm" {
		log.LogNotice(map[string]interface{}{"input":input})
	} else {
		data, _ := json.Marshal(input)

		resp, err := util.Post(conf.Config.PlatformApi.Host+"/v1/cmd", data, map[string]string{}, map[string]string{})
		if err != nil {
			response.Fail(c, consts.REQUEST_FAIL_CODE, consts.REQUEST_FAIL.Error())
			return
		}
		apiResp := CmdResp{}
		err = json.Unmarshal(resp, &apiResp)
		if err != nil {
			response.Fail(c, consts.REQUEST_FAIL_CODE, consts.REQUEST_FAIL.Error())
			return
		}
		if apiResp.ErrorCode != 0 {
			response.Fail(c, consts.REQUEST_FAIL_CODE, apiResp.ErrorMsg)
			return
		}
	}
	response.Success(c, nil)
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
	Timestamp int64 `json:"timestamp"`
}


type ArmReportPayload []ArmReportInfo
type ArmReportInfo struct {
	ArmSn          string    `json:"arm_sn"`
	Status         int       `json:"status"`
	JointPostion   []float64 `json:"joint_postion"`
	ActuralPostion []float64 `json:"actural_postion"`
}

func Report(c *gin.Context)  {
	input := ReportReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		fmt.Println("err1",err)
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}

	if input.EventType == "arm_status_report"{
		armReportPayload := ArmReportPayload{}
		err := json.Unmarshal([]byte(input.Payload),&armReportPayload)
		if err != nil{
			fmt.Println("err",err)
			response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
			return
		}
		for _,armInfo := range armReportPayload{
			err := db.DB().Where("arm_sn = ?",armInfo.ArmSn).Updates(map[string]interface{
			}{
				"status" :armInfo.Status,
				"joint_position" : armInfo.JointPostion,
				"actual_position" : armInfo.ActuralPostion,
			}).Error
			if err != nil{
				response.Fail(c, consts.DB_EXEC_ERR_CODE, consts.DB_EXEC_ERR.Error())
			}
		}

	}else{
	moveUnitParamList := MoveUnit{}
	err := json.Unmarshal([]byte(input.Payload),&moveUnitParamList)
	if err != nil{
		fmt.Println("err",err)
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}


	if len(moveUnitParamList) == 0{
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}

	for _,moveUnitParam := range moveUnitParamList {
		moveUnit := &models.MoveUnit{
			MoveUnitSn:         moveUnitParam.MoveUnitSn,
			Soc:                moveUnitParam.Soc,
			Status:             moveUnitParam.Status,
			Speed:              moveUnitParam.Speed,
			CurrentStationCode: moveUnitParam.CurrentStationCode,
			IsInStation:        moveUnitParam.IsInStation,
			RingAngle:          moveUnitParam.RingAngle,
			RingStatus:         moveUnitParam.RingStatus,
			WorkDuration:       moveUnitParam.WorkDuration,
			Timestamp:          moveUnitParam.Timestamp,
		}

		updates := map[string]interface{}{
			"soc":                  moveUnit.Soc,
			"status":               moveUnit.Status,
			"speed":                moveUnit.Speed,
			"current_station_code": moveUnit.CurrentStationCode,
			"is_in_station":        moveUnit.IsInStation,
			"ring_angle":           moveUnit.RingAngle,
			"ring_status":          moveUnit.RingStatus,
			"work_duration":        moveUnit.WorkDuration,
			"timestamp":            moveUnit.Timestamp,
		}

		ret := db.DB().Table("move_unit").Debug().Where("move_unit_sn = ?", moveUnitParam.MoveUnitSn).Update(updates)
		if ret.Error != nil {
			response.Fail(c, consts.DB_EXEC_ERR_CODE, consts.DB_EXEC_ERR.Error())
			return
		}
		if ret.RowsAffected == 0 {
			response.Fail(c, consts.DB_ROWS_AFFECTED_ZERO_CODE, consts.DB_ROWS_AFFECTED_ZERO_ERR.Error())
			return
		}
	}	}
	response.Success(c,nil)
}




type MoveUnitAddRequest struct {
	MoveUnitSn string `json:"move_unit_sn" binding:"required"`
}
func MoveUnitAdd(c *gin.Context)  {
	input := MoveUnitAddRequest{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	moveUnit := models.MoveUnit{}
	err := db.DB().Where("move_unit_sn=?",input.MoveUnitSn).First(&moveUnit).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		response.Fail(c,consts.DB_QUERY_ERR_CODE,consts.DB_QUERY_FAIL.Error())
		return
	}

	if err == nil && moveUnit.ID > 0{
		err = db.DB().Table("move_unit").Where("move_unit_sn=?",input.MoveUnitSn).Update(map[string]interface{}{
			"deleted":0,
		}).Error
	}else{
		moveUnit = models.MoveUnit{MoveUnitSn: input.MoveUnitSn}

		err = db.DB().Create(&moveUnit).Error
		if err != nil{
			response.Fail(c,consts.DB_EXEC_ERR_CODE,consts.DB_EXEC_ERR.Error())
			return
		}
	}

	response.Success(c,nil)
}

func MoveUnitDelete(c *gin.Context)  {
	input := MoveUnitAddRequest{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	moveUnit := models.MoveUnit{}
	err := db.DB().Where("move_unit_sn=?",input.MoveUnitSn).First(&moveUnit).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		response.Fail(c,consts.DB_QUERY_ERR_CODE,consts.DB_QUERY_FAIL.Error())
		return
	}

	if err == gorm.ErrRecordNotFound{
		response.Fail(c,consts.DB_QUERY_NOT_EXIST_CODE,consts.DB_QUERY_NOT_EXIST_ERR.Error())
		return
	}

	err = db.DB().Table("move_unit").Where("move_unit_sn=?",input.MoveUnitSn).Update("deleted",1).Error
	if err != nil{
		response.Fail(c,consts.DB_EXEC_ERR_CODE,consts.DB_EXEC_ERR.Error())
		return
	}
	response.Success(c,nil)
}

type MoveUnitBindRequest struct {
	MoveUnitSn string `json:"move_unit_sn"`
	ProductionLineID int `json:"production_line_id"`

}

func MoveUnitBind(c *gin.Context)  {
	input := MoveUnitBindRequest{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}

	moveUnit := models.MoveUnit{}
	err := db.DB().Where("move_unit_sn = ?",input.MoveUnitSn).First(&moveUnit).Error
	if (err != nil && err != gorm.ErrRecordNotFound) ||  moveUnit.MoveUnitSn=="" || moveUnit.Deleted == 1 {
		response.Fail(c, consts.DB_QUERY_ERR_CODE,"小车不存在")
		return
	}
	if moveUnit.ProductionLineId != 0{
		response.Fail(c,consts.DB_QUERY_ERR_CODE,"小车已经绑定产线")
		return
	}

	err = db.DB().Where("move_unit_sn = ?").Update(map[string]interface{}{"production_line_id":input.ProductionLineID}).Error
	if err != nil{
		response.Fail(c,consts.DB_EXEC_ERR_CODE,consts.DB_EXEC_ERR.Error())
		return
	}
	response.Success(c,nil)
}


type MoveUnitUpdateRequest struct {
	MoveUnitSn string `json:"move_unit_sn"`
	WorkStatus int `json:"work_status"` // 0停用1启用
}

func MoveUnitUpdate(c *gin.Context)  {
	input := MoveUnitUpdateRequest{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	err := db.DB().Table("move_unit").Where("move_unit_sn=?",input.MoveUnitSn).Update(map[string]interface{}{"work_status":input.WorkStatus})
	if err != nil{
		response.Fail(c, consts.DB_EXEC_ERR_CODE, consts.DB_EXEC_ERR.Error())
		return
	}
	response.Success(c,nil)
}


type ProductionLineRequest struct {
	ProductionLineId string
}


type MoveUnitListRequest struct {
	ProductionLineID int `json:"production_line_id"`
	MoveUnitSN string `json:"move_unit_sn"`
}
func MoveUnitList(c *gin.Context)  {
	input := MoveUnitListRequest{}
	c.ShouldBindWith(&input, binding.JSON);
	db := db.DB().Table("move_unit")
	if input.ProductionLineID > 0 {
		db = db.Where("production_line_id = ?",input.ProductionLineID)
	}
	if input.MoveUnitSN != ""{
		db = db.Where("move_unit_sn = ?",input.MoveUnitSN)
	}
	moveUnitList := []models.MoveUnit{}
	err := db.Find(&moveUnitList).Error
	if err != nil{
		response.Fail(c,consts.DB_QUERY_ERR_CODE,consts.DB_EXEC_ERR.Error())
		return
	}
	response.Success(c,moveUnitList)
}



type MoveUnitUnbindRequest struct {
	MoveUnitSn string `json:"move_unit_sn"`
}
func MoveUnitUnbind(c *gin.Context)  {
	input := MoveUnitUpdateRequest{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	err := db.DB().Table("move_unit").Where("move_unit_sn=?",input.MoveUnitSn).Update(map[string]interface{}{"production_line_id":0,"work_status": 0})
	if err != nil{
		response.Fail(c, consts.DB_EXEC_ERR_CODE, consts.DB_EXEC_ERR.Error())
		return
	}
	response.Success(c,nil)
}


type ProductionLineMoveReq struct {
	ProductionLineID int `json:"production_line_id" binding:"required"`
}





type MoveCmdInfo struct {
	MoveUnitSn        string `json:"move_unit_sn"`
	SourceStationCode string `json:"source_station_code"`
	TargetStationCode string `json:"target_station_code"`
}

func Forward(c *gin.Context)  {
	var input ProductionLineMoveReq
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}

	moveUnits := []models.MoveUnit{}
	err := db.DB().Table("move_unit").Where("production_line_id = ? && work_status =1",input.ProductionLineID).Find(&moveUnits).Error
	if err == gorm.ErrRecordNotFound{
		response.Fail(c, consts.NO_MOVEUNIT_IN_WORK_CODE, consts.NO_MOVEUNIT_IN_WORK_ERR.Error())
		return
	}

	cmdList := []MoveCmdInfo{}
	for _,moveUnit := range moveUnits{

		stationCode := moveUnit.CurrentStationCode
		if stationCode == ""{
			log.LogError(map[string]interface{}{"info":"move_unit_not_station station_code_empty","moveUnit" : moveUnit})
			continue
		}

		if _, ok := manager.Manager.ProductionLineStationMap[moveUnit.ProductionLineId]; !ok {
			log.LogError(map[string]interface{}{"info":" system error ","ProductionLineStationMap" :  manager.Manager.ProductionLineStationMap,"production_line ": moveUnit.ProductionLineId})
			continue
		}

		index,ok := manager.Manager.ProductionLineStationIndex[moveUnit.ProductionLineId][moveUnit.CurrentStationCode];
		if !ok{
			log.LogError(map[string]interface{}{"info":" system error ","moveUnit.CurrentStationCode not in  ProductionLineStationMap" :  manager.Manager.ProductionLineStationMap[moveUnit.ProductionLineId],"current_station_code ": moveUnit.CurrentStationCode})
			continue
		}

		stationCnt := len(manager.Manager.ProductionLineStationSort[moveUnit.ProductionLineId])

		//已经到达终点
		if index +1 >= stationCnt-1 {
			continue
		}
		nextStataion := manager.Manager.ProductionLineStationSort[moveUnit.ProductionLineId][index +1]
		cmdInfo := MoveCmdInfo{
			MoveUnitSn:        moveUnit.MoveUnitSn,
			SourceStationCode: moveUnit.CurrentStationCode,
			TargetStationCode: nextStataion.StationCode,
		}

		cmdList = append(cmdList,cmdInfo)
	}

	log.LogNotice(map[string]interface{}{"move cmdList" : cmdList})

	if len(cmdList) == 0{
		response.Fail(c, consts.NO_MOVEUNIT_CAN_MOVE_FORWARD_CODE, consts.NO_MOVEUNIT_CAN_MOVE_FORWARD.Error())
		return
	}

	payload,_ :=json.Marshal(cmdList)
	
	cmd := CmdReq{
		CmdType: "move",
		Payload: string(payload),
	}

	data,_ := json.Marshal(cmd)

	resp,err := util.Post(conf.Config.PlatformApi.Host+"/v1/cmd",data, map[string]string{}, map[string]string{})
	if err != nil{
		response.Fail(c, consts.REQUEST_FAIL_CODE, consts.REQUEST_FAIL.Error())
		return
	}
	apiResp := CmdResp{}
	err = json.Unmarshal(resp,&apiResp)
	if err != nil{
		response.Fail(c, consts.REQUEST_FAIL_CODE, consts.REQUEST_FAIL.Error())
		return
	}
	if apiResp.ErrorCode != 0 {
		response.Fail(c, consts.REQUEST_FAIL_CODE, apiResp.ErrorMsg)
		return
	}
	response.Success(c,nil)
}



func Stop(c *gin.Context)  {
	var input ProductionLineMoveReq
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}


	err := db.DB().Table("move_unit").Where("production_line_id = ?",input.ProductionLineID).Update("work_status",0).Error
	if err != nil{
		response.Fail(c, consts.DB_EXEC_ERR_CODE, consts.DB_EXEC_ERR.Error())
		return
	}

	response.Success(c,nil)
}

func Start(c *gin.Context)  {
	var input ProductionLineMoveReq
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}


	err := db.DB().Table("move_unit").Where("production_line_id = ?",input.ProductionLineID).Update("work_status",1).Error
	if err != nil{
		response.Fail(c, consts.DB_EXEC_ERR_CODE, consts.DB_EXEC_ERR.Error())
		return
	}
	response.Success(c,nil)
}

func ArmList(c *gin.Context)  {
	arms := []models.Arm{}
	db.DB().Table("arm").AutoMigrate(models.Arm{})
	db.DB().Table("arm").Find(&arms)
	response.Success(c,arms)
}

type ArmCmdInfo struct {
	BombNo         string `json:"bomb_no"`
	Cnt            int64  `json:"cnt"`
	MajorProgramNo string `json:"major_program_no"`
	PlanEndTime    string `json:"plan_end_time"`
	PlanStartTime  string `json:"plan_start_time"`
	ProductionName string `json:"production_name"`
	ProductionNo   string `json:"production_no"`
}


type ArmCmdInput []ArmCmdInfo
func TaskIssue(c *gin.Context)  {
	var input ArmCmdInput
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	data,_ := json.Marshal(input)
	resp,err := httputil.Post(conf.Config.PlatformApi.Host+"/v1/task/issue",data, map[string]string{}, map[string]string{})
	if err!= nil{
		log.LogError(map[string]interface{}{"http post fail err ":err,"path":"/v1/task/issue"})
		response.Fail(c,consts.REQUEST_FAIL_CODE,"http post fail err :"+ err.Error())
		return
	}

	data,err = ioutil.ReadAll(resp.Body)
	if err != nil{
		response.Fail(c,consts.REQUEST_FAIL_CODE,"http post read data err :"+ err.Error())
		return
	}

	cmdResp := CmdResp{}
	err = json.Unmarshal(data,&cmdResp)
	if err != nil{
		response.Fail(c,consts.REQUEST_FAIL_CODE,"http post unmarshal data err : "+ err.Error() +  " data: "+string(data))
		return
	}
	if cmdResp.ErrorCode != 0{
			response.Fail(c,consts.REQUEST_FAIL_CODE,"http post resp data fail error : " + strconv.Itoa(cmdResp.ErrorCode)+" error_msg : " + cmdResp.ErrorMsg)
			return
	}

	log.LogNotice(map[string]interface{}{"input":input})
	response.Success(c,nil)
}