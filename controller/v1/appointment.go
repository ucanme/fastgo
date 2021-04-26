package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"github.com/ucanme/fastgo/consts"
	"github.com/ucanme/fastgo/controller/response"
	"github.com/ucanme/fastgo/library/db"
	"github.com/ucanme/fastgo/models"
	"sort"
	"time"
)

type AvailableDaysListResp struct {
	Days []string `json:"days"`
}
func AvailableDaysList(c *gin.Context)  {
	yesterDay := time.Now().Add(-time.Hour*24).Format("2006-01-02")
	all := []models.Appointment{}
	err := db.DB().Where("date > ?",yesterDay).Find(&all).Error
	if err != nil{
		response.Success(c,nil)
		return
	}

	dateMap := map[string]int{}
	avaliDates := []string{}
	for _,v := range all{
		if _,ok := dateMap[v.Date];!ok{
			dateMap[v.Date] = 1
			avaliDates = append(avaliDates,v.Date)
		}
	}

	sort.Slice(avaliDates, func(i, j int) bool {
		return avaliDates[i] < avaliDates[j]
	})

	if len(avaliDates) > 10 {
		avaliDates = avaliDates[:10]
	}
	response.Success(c,AvailableDaysListResp{Days: avaliDates})
}





type AvailableHourListReq struct {
	Day string `json:"day"`
}
type AvailableHourListResp []*AvailableHourInfo
type AvailableHourInfo struct {
	Hour int `json:"hour"`
	Status int `json:"status"`
}
func AvaliableHoursList(c *gin.Context)  {
	input := AvailableHourListReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	appoints := []models.Appointment{}
	err := db.DB().Where("date=?",input.Day).Find(&appoints).Error
	if err!=nil && err != gorm.ErrRecordNotFound{
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}



	hourMap := map[int]*AvailableHourInfo{}
	avaliHours := []int{}
	for _,v := range appoints{
		if _,ok := hourMap[v.Hour];!ok{
			hourMap[v.Hour] = &AvailableHourInfo{
				Hour:   v.Hour,
				Status: 1,
			}
			avaliHours = append(avaliHours,v.Hour)
		}
		if v.Status == 0 {
			hourMap[v.Hour].Status = 0
		}
	}

	sort.Slice(avaliHours, func(i, j int) bool {
		return avaliHours[i] < avaliHours[j]
	})
	availableHourListResp := AvailableHourListResp{}
	for _,v := range avaliHours{
		availableHourListResp =append(availableHourListResp,hourMap[v])
	}
	response.Success(c,availableHourListResp)
}

type  AvaliableMinutesListReq struct {
	Day string
	Hour int
}


type AvaliableMinutesListResp struct {
	Minutes []AvaliableMinuteInfo `json:"minutes"`
}

type AvaliableMinuteInfo struct {
	Minute int `json:"minute"`
	Status int `json:"status"`
}

func AvaliableMinutesList(c *gin.Context)  {
	input := AvaliableMinutesListReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	appoints := []models.Appointment{}
	err := db.DB().Where("date=? && hour=?",input.Day,input.Hour).Find(&appoints).Error
	if err!=nil && err != gorm.ErrRecordNotFound{
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}



	sort.Slice(appoints, func(i, j int) bool {
		return appoints[i].Minute < appoints[j].Minute
	})

	avaliableMinutesListResp := AvaliableMinutesListResp{}
	for _,v := range appoints{
		avaliableMinuteInfo := AvaliableMinuteInfo{
			Minute: v.Minute,
			Status: v.Status,
		}
		avaliableMinutesListResp.Minutes = append(avaliableMinutesListResp.Minutes,avaliableMinuteInfo)
	}
	response.Success(c,avaliableMinutesListResp)
}

type MakeApointmentReq struct {
	Day      string `json:"day" binding:"required"`
	Hour     int    `json:"hour" binding:"required"`
	Minute   int    `json:"minute" binding:"required"`
	OpenID   string `json:"open_id" binding:"required"`
	Name     string `json:"name" binding:"required"`
	PhoneNum string `json:"phone_num" binding:"required"`
}

func MakeApointment(c *gin.Context)  {
	input := MakeApointmentReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	appint := models.Appointment{}
	err := db.DB().Where("day=? && hour = ? && minute = ?").First(&appint).Error
	if err!=nil{
		response.Fail(c, consts.MAKE_APPOINT_FAIL_CODE, consts.MAKE_APPOINT_FAIL.Error())
		return
	}

	if appint.Status== 1 {
		response.Fail(c, consts.TIME_ALREADY_OCCUPY_CODE, "当前时间已经被占用")
		return
	}
	appint.Status = 1
	appint.OpenId = input.OpenID
	appint.Name = input.Name
	appint.PhoneNum = input.PhoneNum
	err = db.DB().Save(&appint).Error
	if err!=nil{
		response.Fail(c, consts.MAKE_APPOINT_FAIL_CODE, consts.MAKE_APPOINT_FAIL.Error())
		return
	}
	response.Success(c,nil)
}

type AppointmentListReq struct {
	OpenId string `json:"open_id" binding:"required"`
}
func AppointmentList(c *gin.Context)  {
	input := AppointmentListReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	appoints := models.Appointment{}
	err := db.DB().Where("open_id=?",input.OpenId).Find(&appoints).Error
	if err!=nil{
		response.Fail(c,consts.GET_APPOINT_LIST_FAIL_CODE,err.Error())
		return
	}
	response.Success(c,appoints)
}



type CancelAppointmentReq struct {
	OpenID       string `json:"open_id" binding:"required"`
	AppintmentID int    `json:"appintment_id" binding:"required"`
}
func CancelAppointment(c *gin.Context)  {
	input := CancelAppointmentReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	appoint := models.Appointment{}
	err := db.DB().Where("open_id=? and id = ?",input.OpenID,input.AppintmentID).First(&appoint).Error
	if err!=nil{
		response.Fail(c,consts.GET_APPOINT_LIST_FAIL_CODE,err.Error())
		return
	}
	appoint.Status = 0
	err = db.DB().Save(&appoint).Error
	if err!=nil{
		response.Fail(c,consts.GET_APPOINT_LIST_FAIL_CODE,err.Error())
		return
	}
	response.Success(c,nil)
}


type SignInReq struct {
	OpenID       string `json:"open_id" binding:"required"`
	AppintmentID int    `json:"appintment_id" binding:"required"`
}
func SignIn(c *gin.Context)  {
	input := SignInReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	appoint := models.Appointment{}
	err := db.DB().Where("open_id=? and id = ?",input.OpenID,input.AppintmentID).First(&appoint).Error
	if err!=nil{
		response.Fail(c,consts.SIGN_IN_FAIL_CODE,"未查到预约记录")
		return
	}
	appoint.Status = 2
	err = db.DB().Save(&appoint).Error
	if err!=nil{
		response.Fail(c,consts.SIGN_IN_FAIL_CODE,err.Error())
		return
	}
	response.Success(c,nil)
}
