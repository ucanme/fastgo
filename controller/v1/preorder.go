package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
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

		preOrder := models.PreOrder{}
		err := db.DB().Where("phone=? && date=? && place_id=?",input.Phone,input.Date,input.Id).Order("created_at desc").First(&preOrder).Error
		if err!=nil && err!=gorm.ErrRecordNotFound{
			response.Fail(c, 400, "预约失败")
			return
		}

		if err == gorm.ErrRecordNotFound{
			preOrder.Phone = input.Phone
			preOrder.Date = input.Date
			preOrder.PlaceId = input.Id
		}
		preOrder.Name = input.Name
		err = db.DB().Save(&preOrder).Error
		if err !=nil{
			response.Fail(c, 400, "预约失败")
			return
		}
		response.Fail(c, 400, "预约失败")
		return
}