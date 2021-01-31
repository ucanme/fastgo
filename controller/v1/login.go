package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"github.com/ucanme/fastgo/consts"
	"github.com/ucanme/fastgo/controller/response"
	"github.com/ucanme/fastgo/internal/session"
	"github.com/ucanme/fastgo/library/db"
	"github.com/ucanme/fastgo/models"
)
import 	"github.com/gookit/validate"

type loginReq struct {
	UserId string `json:"user_id" validate:"required|minLen:7|maxLen:15"`
	Password string `json:"password" validate:"required|minLen:7|maxLen:20"`
}
func Login(c *gin.Context)  {
	input := loginReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	fmt.Println(input)
	valid := validate.Struct(input)
	if !valid.Validate(){
		response.Fail(c, consts.PARAM_ERR_CODE, consts.VALIDATE_ERR.Error())
		return
	}
	uid := input.UserId
	//password := input.PassWord

	user := &models.User{}
	err := db.DB().Where("user_id=?",uid).First(user).Error
	if err == gorm.ErrRecordNotFound{
		response.Fail(c, consts.USER_NOT_FOUND_CODE, consts.USER_NOT_FOUND_ERR.Error())
		return
	}

	if err!=nil{
		response.Fail(c, consts.SYSMSTEM_ERR_CODE, consts.SYSMSTEM_ERR.Error())
		return
	}

	if user.Password != input.Password{
		response.Fail(c, consts.USER_PASS_INCORRECT_CODE, consts.USER_PASS_INCORRECT_ERR.Error())
		return
	}
	session.Manager.SessionStart(c)
	response.Success(c,nil)
}
