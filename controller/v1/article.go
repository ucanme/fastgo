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
	err := db.DB().Order("created_at desc").Where("id=?",input.Id).Find(&as).Error
	if err==gorm.ErrRecordNotFound{
		response.Fail(c,400,"no article")
	}
	response.Success(c,as)
}




type addReq struct {
	CateId int `json:"cate_id"`
	Content string `json:"content"`
	ImgUrl string `json:"img_url"`
	Additional string `json:"additional"`
	Additional01 string `json:"additional_01"`
	Additional02 int `json:"additional_02"`
	Additional03 int `json:"additional_03"`
	Additional04 string `json:"additional_04"`
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
		ImgUrl: input.ImgUrl,
		Additional: input.Additional,
		Additional01: input.Additional01,
		Additional02: input.Additional02,
		Additional03: input.Additional03,
		Additional04: input.Additional04,
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



type EditReq struct {
	Id int `json:"id" binding:"required"`
	Content string `json:"content"`
	ImgUrl string `json:"img_url"`
	Additional string `json:"additional"`
	Additional01 string `json:"additional_01"`
	Additional02 int `json:"additional_02"`
	Additional03 int `json:"additional_03"`
	Additional04 int `json:"additional_04"`
}

func EditArticle(c *gin.Context)  {
	input := EditReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	if input.Content == "" && input.ImgUrl == "" && input.Additional== ""{
		response.Fail(c, consts.PARAM_ERR_CODE,"修改字段不能同时为空")
		return
	}
	if input.Content!=""{
		err := db.DB().Table("articles").Where("id=?",input.Id).Update("content",input.Content).Error
		if err!=nil{
			response.Fail(c, consts.DB_EXEC_ERR_CODE,err.Error())
			return
		}
	}
	if input.ImgUrl != ""{
		err := db.DB().Table("articles").Where("id=?",input.Id).Update("img_url",input.Additional).Error
		if err!=nil{
			response.Fail(c, consts.DB_EXEC_ERR_CODE,err.Error())
			return
		}
	}

	if input.Additional != ""{
		err := db.DB().Table("articles").Where("id=?",input.Id).Update("additional",input.Additional).Error
		if err!=nil{
			response.Fail(c, consts.DB_EXEC_ERR_CODE,err.Error())
			return
		}
	}
	if input.Additional01 != ""{
		err := db.DB().Table("articles").Where("id=?",input.Id).Update("additional_01",input.Additional01).Error
		if err!=nil{
			response.Fail(c, consts.DB_EXEC_ERR_CODE,err.Error())
			return
		}
	}
	if input.Additional02 != 0{
		err := db.DB().Table("articles").Where("id=?",input.Id).Update("additional_02",input.Additional02).Error
		if err!=nil{
			response.Fail(c, consts.DB_EXEC_ERR_CODE,err.Error())
			return
		}
	}
	if input.Additional03 != 0{
		err := db.DB().Table("articles").Where("id=?",input.Id).Update("additional_03",input.Additional03).Error
		if err!=nil{
			response.Fail(c, consts.DB_EXEC_ERR_CODE,err.Error())
			return
		}
	}

	response.Success(c,nil)
}




