package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"github.com/ucanme/fastgo/consts"
	"github.com/ucanme/fastgo/controller/response"
	"github.com/ucanme/fastgo/library/db"
	"github.com/ucanme/fastgo/models"
	"strings"
	"time"
)

type orderReq struct {
	ArticleId  int `json:"article_id"`
	Name string `json:"name"`
	PhoneNum string `json:"phone_num"`
	Date string `json:"date"`
	PersonCnt int `json:"person_cnt"`
	OpenId string `json:"open_id"`
	Addition string `json:"addition"`
	Addition01 string `json:"addition_01"`
}
func PreOrder(c *gin.Context)  {
	input := orderReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}

	a := models.Article{}
	err := db.DB().Where("id=?",input.ArticleId).First(&a).Error
	if err!=nil{
		response.Fail(c, 400, "活动不存在")
		return
	}

	preOrder := models.PreOrder{}
	err = db.DB().Where("article_id=? && date=? && open_id = ?",input.OpenId,input.Date,input.ArticleId).Order("created_at desc").First(&preOrder).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		response.Fail(c, 400, "系统错误")
		return
	}
	if err == nil{
		response.Fail(c, 400, "已经预约")
		return
	}
	orders := []models.PreOrder{}
	err = db.DB().Where("article_id=? && date=?").Find(&orders).Error
	if err != nil && err != gorm.ErrRecordNotFound{
		response.Fail(c, 400, "系统错误")
		return
	}

	//总报名人数
	sumCnt := 0
	for _,v:=range orders{
		sumCnt = sumCnt+ v.PersonCnt
	}

	if sumCnt + input.PersonCnt > a.Additional02{
		response.Fail(c, 400, "人数超过限制")
		return
	}

	pre := models.PreOrder{
		Name:      input.Name,
		PhoneNum:     input.PhoneNum,
		Date:      input.Date,
		PlaceId:   input.ArticleId,
		PersonCnt: input.PersonCnt,
		Addition: input.Addition,
		Addtion01: input.Addition01,
	}
	err = db.DB().Save(&pre).Error
	if err!=nil{
		response.Fail(c, 400, "系统错误，预约失败")
		return
	}
	response.Success(c,0)
}


type PlaceOrderLeft struct {
	Date string `json:"date"`
	LeftCnt int `json:"left_cnt"`
}

func PlaceOrderInfo(c *gin.Context)  {
	input := orderReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}

	a := models.Article{}
	err := db.DB().Where("id=?",input.ArticleId).First(&a).Error
	if err!=nil{
		response.Fail(c, 400, "场馆不存在")
		return
	}

	preOrders := []models.PreOrder{}


	err = db.DB().Where("place_id=? && date>= ?  && date <= ?",input.ArticleId,time.Now().Format("2006-01-02"),time.Now().Add(time.Hour*720).Format("2006-01-02")).Find(&preOrders).Error
	fmt.Println("err",err)
	if err!=nil && err != gorm.ErrRecordNotFound{
		fmt.Println(err)
		response.Fail(c, 400, "系统错误")
		return
	}

	//已预约信息
	var placeOrderedMap =  map[string]int{}
	for _,v := range preOrders{
		placeOrderedMap[v.Date] = a.Additional02 - v.PersonCnt
	}


	var ret = []PlaceOrderLeft{}
	//可预约的信息
	fmt.Println(a)
	dates := strings.Split(a.Additional01,",")
	for _,v := range dates{
		if v  >= time.Now().Format("2006-01-02") && v < time.Now().Add(30*time.Hour*240).Format("2006-01-02"){
			ret = append(ret,PlaceOrderLeft{
				Date:    v,
				LeftCnt: a.Additional02-placeOrderedMap[v],
			})
		}
	}

	response.Success(c,ret)
}



func UserPlaceOrderList(c *gin.Context)  {
	input := orderReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}


	preOrders := []models.PreOrder{}

	err := db.DB().Where("open_id = ? && date>= ?  && date <= ?",input.OpenId,time.Now().Format("2006-01-02"),time.Now().Add(time.Hour*720).Format("2006-01-02")).Find(&preOrders).Error
	fmt.Println("err",err)
	if err!=nil && err != gorm.ErrRecordNotFound{
		fmt.Println(err)
		response.Fail(c, 400, "系统错误")
		return
	}

	articleId := []int{}
	for _,v := range preOrders{
		articleId = append(articleId,v.PlaceId)
	}


	articles := []models.Article{}
	err = db.DB().Where("id in (?)",articleId).Find(&articles).Error
	if err!=nil && err!= gorm.ErrRecordNotFound{
		response.Fail(c, 400, "系统错误")
		return
	}
	articleMap := map[uint]models.Article{}
	for _,v := range articles{
		articleMap[v.ID] = v
	}

	type OrderInfo struct {
		PreOder models.PreOrder
		Article models.Article
	}
	type Resp []OrderInfo
	var resp = Resp{}
	a := OrderInfo{}
	for _,v := range preOrders{
		a.PreOder = v
		if _,ok:= articleMap[uint(v.PlaceId)];ok{
			a.Article = articleMap[uint(v.PlaceId)]
		}
		resp = append(resp,a)
	}
	response.Success(c,resp)
}


func PlaceOrderList(c *gin.Context)  {
	input := orderReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	if input.ArticleId == 0{
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
	}


	preOrders := []models.PreOrder{}

	err := db.DB().Where("place_id = ? && date>= ?  && date <= ?",input.ArticleId,time.Now().Add(-time.Hour*720).Format("2006-01-02"),time.Now().Add(time.Hour*720).Format("2006-01-02")).Find(&preOrders).Error
	fmt.Println("err",err)
	if err!=nil && err != gorm.ErrRecordNotFound{
		fmt.Println(err)
		response.Fail(c, 400, "系统错误")
		return
	}


	article := models.Article{}
	err = db.DB().Where("id in (?)",input.ArticleId).First(&article).Error
	if err!=nil{
		response.Fail(c, 400, "不存在预约场所")
		return
	}

	type OrderInfo struct {
		PreOrders []models.PreOrder `json:"pre_orders"`
		Article models.Article `json:"article"`
	}
	type Resp OrderInfo
	var resp = Resp{}
	for _,v := range preOrders{
		resp.PreOrders = append(resp.PreOrders,v)
	}
	resp.Article = article
	response.Success(c,resp)


}

