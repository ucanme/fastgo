package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/ucanme/fastgo/controller/response"
	"github.com/ucanme/fastgo/library/db"
	"github.com/ucanme/fastgo/models"
	"sort"
	"time"
)


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
	response.Success(c,avaliDates)
}