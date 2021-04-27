package cron

import (
	"github.com/jinzhu/gorm"
	"github.com/robfig/cron"
	"github.com/ucanme/fastgo/conf"
	"github.com/ucanme/fastgo/library/db"
	"github.com/ucanme/fastgo/library/log"
	"github.com/ucanme/fastgo/models"
	"github.com/urfave/cli"
	"time"
)

func Init() {
	cronJob := cron.NewWithLocation(time.Local)
	err := cronJob.AddFunc("0 */4 * * * *", func() {
		ProducePreAppointMent()
	})
	if err != nil {
		panic(err)
	}
	cronJob.Start()
}

var Util = cli.Command{
	Name:   "util",
	Usage:  "produce",
	Action: ProducePreAppointMentUtil,
}

func ProducePreAppointMentUtil(c *cli.Context)  {
	ProducePreAppointMent()
}

func ProducePreAppointMent() {
	conf.Init("./config.toml")
	db.Init()
	appoint := models.Appointment{}
	err := db.DB().Order("id desc").Limit(1).First(&appoint).Error
	if err!= nil && err!=gorm.ErrRecordNotFound{
		log.LogError(map[string]interface{}{"cron_produce_data_err":err.Error()})
		return
	}
	maxDate := time.Now().Format("2006-01-02")
	if err != gorm.ErrRecordNotFound{
		maxDate = appoint.Date
	}

	t,err := time.Parse("2006-01-02",maxDate)
	if err !=nil{
		log.LogError(map[string]interface{}{"cron_produce_date_parse_err":err.Error()})
		return
	}
	date := t.Add(25*time.Hour)
	a := models.Appointment{
		Date:     date.Format("2006-01-02"),
		Hour:     0,
		Minute:   0,
		Status:   0,
		Name:     "",
		OpenId:   "",
		PhoneNum: "",
	}

	for Hour :=9;Hour < 11;Hour++{
		miniute := 0
		if Hour == 9{
			miniute = 30
		}
		for miniute < 60{
			tmp := a
			tmp.Hour = Hour
			tmp.Minute = miniute
			db.DB().Save(&tmp)
			miniute = miniute+10
		}
	}


	for Hour :=13;Hour < 16;Hour++{
		miniute := 0
		if Hour == 13{
			miniute = 30
		}else{
			miniute = 0
		}
		for miniute < 60{
			tmp := a
			tmp.Hour = Hour
			tmp.Minute = miniute
			db.DB().Save(&tmp)
			miniute = miniute+10
		}
	}

}
