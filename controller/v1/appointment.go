package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/jinzhu/gorm"
	"github.com/ucanme/fastgo/consts"
	"github.com/ucanme/fastgo/controller/response"
	"github.com/ucanme/fastgo/library/db"
	"github.com/ucanme/fastgo/models"
	"sort"
	"time"
)



type AvailableDaysListResp struct {
	Days []string `json:"days"`
}
func AvailableDaysList(c *gin.Context)  {
	yesterDay := time.Now().Add(-time.Hour*24).Format("2006-01-02")
	all := []models.Appointment{}
	err := db.DB().Where("date > ?",yesterDay).Find(&all).Error
	if err != nil{
		response.Success(c,nil)
		return
	}

	dateMap := map[string]int{}
	avaliDates := []string{}
	for _,v := range all{
		if _,ok := dateMap[v.Date];!ok{
			dateMap[v.Date] = 1
			avaliDates = append(avaliDates,v.Date)
		}
	}

	sort.Slice(avaliDates, func(i, j int) bool {
		return avaliDates[i] < avaliDates[j]
	})

	if len(avaliDates) > 10 {
		avaliDates = avaliDates[:10]
	}
	response.Success(c,AvailableDaysListResp{Days: avaliDates})
}


type AvailableHourListReq struct {
	Day string `json:"day"`
}

type AvailableHourListResp []*AvailableHourInfo

type AvailableHourInfo struct {
	Hour int `json:"hour"`
	Status int `json:"status"`
}

func AvaliableHoursList(c *gin.Context)  {
	input := AvailableHourListReq{}
	if err := c.ShouldBindWith(&input, binding.JSON); err != nil {
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}
	appoints := []models.Appointment{}
	err := db.DB().Where("date=?",input.Day).Find(&appoints).Error
	if err!=nil && err != gorm.ErrRecordNotFound{
		response.Fail(c, consts.PARAM_ERR_CODE, consts.PARAM_ERR.Error())
		return
	}



	hourMap := map[int]*AvailableHourInfo{}
	avaliHours := []int{}
	for _,v := range appoints{
		if _,ok := hourMap[v.Hour];!ok{
			hourMap[v.Hour] = &AvailableHourInfo{
				Hour:   v.Hour,
				Status: 1,
			}
			avaliHours = append(avaliHours,v.Hour)
		}
		if v.Status == 0 {
			hourMap[v.Hour].Status = 0
		}
	}

	sort.Slice(avaliHours, func(i, j int) bool {
		return avaliHours[i] < avaliHours[j]
	})
	availableHourListResp := AvailableHourListResp{}
	for _,v := range avaliHours{
		availableHourListResp =append(availableHourListResp,hourMap[v])
	}
	response.Success(c,availableHourListResp)
}