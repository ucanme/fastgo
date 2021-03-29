package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/ucanme/fastgo/consts"
	"github.com/ucanme/fastgo/controller/response"
	"github.com/ucanme/fastgo/library/db"
	"github.com/ucanme/fastgo/models"
)

type orderReq struct {
	Id  int `json:"id"`
	Date string `json:"date"`
	Name string `json:"name"`
	Phone string `json:"phone"`
}
func PreOrder(c *gin.Context)  {
		input := orderReq{}
		if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
			response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
			return
		}
		preOrder := models.PreOrder{
			Name:    input.Name,
			Phone:   input.Phone,
			Date:    input.Date,
			PlaceId: input.Id,
		}
		err := db.DB().Create(&preOrder).Error
		if err !=nil{
			response.Fail(c, 400, "预约失败")
			return
		}
}