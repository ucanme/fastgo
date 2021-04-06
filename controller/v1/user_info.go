package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"github.com/ucanme/fastgo/controller/response"
	"github.com/ucanme/fastgo/library/db"
	"github.com/ucanme/fastgo/models"
	"math/rand"
	"strconv"
	"time"
)



type RegisterReq struct {
	PhoneNum string `json:"phone_num"`
	AvatarUrl string `json:"avatar_url"`
	OpenId string `json:"open_id"`
	Score int `json:"score"`
	Name string `json:"name"`
	Length string `json:"length"`
	IdNo string `json:"id_no"`
	Address string `json:"address"`
	Addition string `json:"addition"`
	Addition01 string `json:"addition01"`

}

func Register(c *gin.Context)  {
	input := RegisterReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	//v,ok := c.Get("session")
	//if !ok && input.OpenId == "" {
	//	response.Fail(c, 400, "未登陆")
	//	return
	//}
	//if ok {
	//	data,ok1 :=  v.(SessionData)
	//	if ok1{
	//		input.OpenId = data.OpenId
	//	}
	//
	//}

	u,err :=doRegister(input)
	if err!=nil{
		response.Fail(c,400,"注册失败")
		return
	}
	response.Success(c,u)
}

//用户注册
func doRegister(req RegisterReq) (models.User,error) {
	user := models.User{}
	err := db.DB().Where("open_id=?",req.OpenId).First(&user).Error
	if err!=nil && err!=gorm.ErrRecordNotFound{
		return user,err
	}

	user.OpenId = req.OpenId
	if req.OpenId != "" && req.PhoneNum != ""{
		user.PhoneNum = req.PhoneNum
	}
	if user.UserId == ""{
		y  := strconv.Itoa(time.Now().Year())
		user.UserId = y[len(y)-2:]+ time.Now().Month().String()+strconv.Itoa(time.Now().Hour())+strconv.Itoa(time.Now().Minute())+strconv.Itoa(time.Now().Second())+strconv.Itoa(int(rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000)))
	}

	if req.Score != 0{
		user.Score =req.Score
	}
	if req.Score == -1{
		user.Score = 0
	}
	if req.AvatarUrl != ""{
		user.AvatarUrl = req.AvatarUrl
	}

	if req.Address != ""{
		user.Address = req.Address
	}

	if req.Length != ""{
		user.Length = req.Length
	}

	if req.Name != ""{
		user.Name = req.Name
	}
	if req.Addition != ""{
		user.Addition = req.Addition
	}

	if req.Addition01 != ""{
		user.Addition01 = req.Addition01
	}

	if err := db.DB().Save(&user).Error;err!=nil{
		return user,err
	}
	return user,nil;
}

func UserList(c *gin.Context)  {
	us := []models.User{}
	err := db.DB().Find(&us).Error
	if err!=nil{
		if err!=nil{
			response.Fail(c,400,"失败")
			return
		}
	}
	response.Success(c,us)
}



type GetUserInfoReq struct {
	OpenId string `json:"open_id" binding:"required"`
}
func GetUser(c *gin.Context)  {
	input := GetUserInfoReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, 400, "参数错误")
		return
	}

	openId := input.OpenId
	u := models.User{}
	err := db.DB().Where("open_id=?",openId).First(&u).Error
	if err!=nil{
		response.Fail(c, 400, "登陆失败")
		return
	}
	response.Success(c,u)

}