package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ucanme/fastgo/consts"
	"github.com/ucanme/fastgo/controller/response"
)
import 	"github.com/gookit/validate"

type loginReq struct {
	UserId string `json:"user_id" validate:"required|minLen:7|maxLen:15"`
	PassWord string `json:"pass_word" validate:"required|minLen:7|maxLen:20"`
}
func Login(c *gin.Context)  {
	input := loginReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	err := validate.Struct(input)
	if err!=nil{
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
}
