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

	err :=doRegister(input.OpenId,input.PhoneNum,input.AvatarUrl,input.Score)
	if err!=nil{
		response.Fail(c,400,"注册失败")
		return
	}
	response.Success(c,[]string{})
}

//用户注册
func doRegister(openId,phoneNum,avatarUrl string,score int) error {
	user := models.User{}
	err := db.DB().Where("open_id=?",openId).First(&user).Error
	if err!=nil && err!=gorm.ErrRecordNotFound{
		return err
	}

	user.OpenId = openId
	if phoneNum != ""{
		user.PhoneNum = phoneNum
	}
	if user.UserId == ""{
		y  := strconv.Itoa(time.Now().Year())
		user.UserId = y[len(y)-2:]+ time.Now().Month().String()+strconv.Itoa(time.Now().Hour())+strconv.Itoa(time.Now().Minute())+strconv.Itoa(time.Now().Second())+strconv.Itoa(int(rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000)))
	}

	if score != 0{
		user.Score =score
	}
	if avatarUrl != ""{
		user.AvatarUrl = avatarUrl
	}
	if err := db.DB().Save(&user).Error;err!=nil{
		return err
	}
	return nil;
}