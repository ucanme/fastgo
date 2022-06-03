package cron

import (
	"github.com/robfig/cron"
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

}
