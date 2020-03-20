package cron

import (
	"fmt"
	"github.com/robfig/cron"
	"github.com/urfave/cli"
	"time"
)

func Init() {
	cronJob := cron.NewWithLocation(time.Local)
	err := cronJob.AddFunc("0 */4 * * * *", func() {
		Demo()
	})
	if err != nil {
		panic(err)
	}
	cronJob.Start()
}

var Util = cli.Command{
	Name:   "util",
	Usage:  "park_info",
	Action: Demo,
}

func Demo() {
	fmt.Println("demo")
}
