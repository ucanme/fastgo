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

func Demo(c *gin.Context) {
	response.Success(c, map[string]interface{}{"Hello": "World"})
	return
}


type VolunterAddReq struct {
	ArticleId int `json:"article_id"`
	Name string `json:"name"`
	Address string `json:"address"`
	PhoneNum string `json:"phone_num"`
	OpenId string `json:"open_id"`
	Length string `json:"length"`
	IdNo string `json:"id_no"`
}
func VolunteerAdd(c *gin.Context)  {
	input := VolunterAddReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	v := models.Volunteer{
		Name:      input.Name,
		Address:   input.Address,
		PhoneNum:  input.PhoneNum,
		ArticleId: input.ArticleId,
		OpenId : input.OpenId,
		Length : input.Length,
		IdNo : input.IdNo,
	}
	err := db.DB().Save(&v).Error
	if err!=nil{
		response.Fail(c, consts.DB_EXEC_ERR_CODE,"dberr")
		return
	}
	response.Success(c,v)
	return
}

type VolunterDeleteReq struct {
	Id int `json:"id"`
}
func VolunteerDelete(c *gin.Context)  {
	input := VolunterDeleteReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	if input.Id <= 0{
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}

	err := db.DB().Delete(&models.Volunteer{},input.Id).Error
	if err!=nil{
		response.Fail(c, consts.DB_EXEC_ERR_CODE,"dberr")
		return
	}
	response.Success(c,nil)
	return
}


func VolunteerList(c *gin.Context)  {
	input := VolunterDeleteReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	if input.Id <= 0{
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}

	volunteers := []models.Volunteer{}
	err := db.DB().Where("article_id = ?",input.Id).Find(&volunteers).Error
	if err!=nil && err!=gorm.ErrRecordNotFound{
		response.Fail(c, consts.DB_EXEC_ERR_CODE,"dberr")
		return
	}
	response.Success(c,volunteers)
	return
}