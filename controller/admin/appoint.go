package admin

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"github.com/ucanme/fastgo/consts"
	"github.com/ucanme/fastgo/controller/response"
	"github.com/ucanme/fastgo/library/db"
	"github.com/ucanme/fastgo/models"
	"time"
)


type AppointmentListReq struct {
	Day string
}
func AppointmentList(c *gin.Context)  {
	input := AppointmentListReq{}
	var err error
	if err = c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	appoints := []models.Appointment{}
	if input.Day != ""{
		err = db.DB().Where("day=?",input.Day).Find(&appoints).Error
	}else{
		yesterDay := time.Now().Add(-time.Hour*24).Format("2006-01-02")
		oneMonthDay := time.Now().Add(time.Hour*30).Format("2006-01-02")
		err = db.DB().Where("day > ? && day < ? ",yesterDay,oneMonthDay).Find(appoints).Error
	}
	if err!=nil{
		response.Fail(c, consts.ADMIN_GET_APPOINT_FAIL_CODE, "获取记录失败")
		return
	}
	response.Success(c,appoints)
}


type SigninAppointmentListReq struct {
	Day string `json:"day" binding:"required"`
}
func SigninAppointmentList(c *gin.Context)  {
	input := SigninAppointmentListReq{}
	var err error
	if err = c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	appoints := []models.Appointment{}
	err = db.DB().Where("day=?",input.Day).Find(&appoints).Error
	if err!=nil && err!=gorm.ErrRecordNotFound{
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
}

