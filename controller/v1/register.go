package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/gookit/validate"
	"github.com/jinzhu/gorm"
	"github.com/ucanme/fastgo/consts"
	"github.com/ucanme/fastgo/controller/response"
	"github.com/ucanme/fastgo/library/db"
	"github.com/ucanme/fastgo/models"
)

func Register(c *gin.Context)  {
	type Register struct {
		UserId string `json:"user_id" validate:"required|minLen:7|maxLen:15"`
		Password string `json:"password" validate:"required|minLen:7|maxLen:20"`
	}
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

	user := &models.User{}
	err := db.DB().Where("user_id=?",input.UserId).First(user).Error

	if err==nil{
		response.Fail(c, consts.USER_EXISTS_CODE, consts.USER_EXISTS_ERR.Error())
		return
	}

	if err != gorm.ErrRecordNotFound{
		response.Fail(c, consts.SYSMSTEM_ERR_CODE, consts.SYSMSTEM_ERR.Error())
		return
	}

	err = db.DB().Create(&user).Error
	if err!=nil{
		response.Fail(c, consts.DB_EXEC_ERR_CODE, consts.DB_EXEC_ERR.Error())
		return
	}
	response.Success(c,nil)
}
