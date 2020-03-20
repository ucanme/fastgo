package main

import (
	"fmt"
	"github.com/ucanme/fastgo/cmd/server"
	"github.com/ucanme/fastgo/cmd/tool"
	"github.com/urfave/cli"
	"os"
	"time"
)

func main() {
	var (
		BuildGitHash  = "0000000"
		BuildGitTag   = "v1.0.1"
		BuildGitCount = "0"
		BuildTime     = ""
		buildTime     time.Time
	)

	buildTime, err := time.Parse("2006-01-02T15:04:05", BuildTime)
	app := cli.NewApp()
	app.Name = "demo"
	app.Version = fmt.Sprintf("%s.%s.%s", BuildGitTag, BuildGitHash, BuildGitCount)
	app.Usage = "http server"
	app.Compiled = buildTime
	app.Commands = []cli.Command{
		server.Server,
		tool.InitDB,
	}
	app.Before = func(c *cli.Context) error {
		fmt.Printf("Hello fastgo, version %s\n", app.Version)
		return nil
	}

	err = app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
