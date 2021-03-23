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

type listReq struct {
	CateId int `json:"cate_id"`
}
func ListArticle(c *gin.Context)  {
	input := listReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	as := []models.Article{}
	err := db.DB().Order("created_at desc").Where("cate_id=?",input.CateId).Find(&as).Error
	if err==gorm.ErrRecordNotFound{
		response.Fail(c,400,"no article")
	}
	response.Success(c,as)
}



type getReq struct {
	Id int `json:"id"`
}
func GetArticle(c *gin.Context)  {
	input := getReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	as := []models.Article{}
	err := db.DB().Order("created_at desc").Where("cate_id=?",input.Id).Find(&as).Error
	if err==gorm.ErrRecordNotFound{
		response.Fail(c,400,"no article")
	}
	response.Success(c,as)
}




type addReq struct {
	CateId int `json:"cate_id"`
	Content string `json:"content"`
}
func AddArticle(c *gin.Context)  {
	input := addReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	a := models.Article{
		Model:   gorm.Model{},
		Content: input.Content,
		CateId:  input.CateId,
	}
	err := db.DB().Save(&a).Error
	if err!=nil{
		response.Fail(c, consts.PARAM_ERR_CODE, err.Error())
		return
	}
	response.Success(c,nil)
}

type delReq struct {
	Id int `json:"id"`
}

func DeleteArticle(c *gin.Context)  {
	input := delReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}

	a := models.Article{}
	err := db.DB().Delete(a,input.Id).Error
	if err!=nil{
		response.Fail(c, consts.PARAM_ERR_CODE,err.Error())
		return
	}
	response.Success(c,nil)
}




