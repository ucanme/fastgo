package tool

import (
	"github.com/ucanme/fastgo/conf"
	"github.com/ucanme/fastgo/library/db"
	"github.com/ucanme/fastgo/models"
	"github.com/urfave/cli"
)

// InitDB ...
var InitDB = cli.Command{
	Name:  "init-db",
	Usage: "add_db init db",
	Flags: []cli.Flag{
		cli.StringFlag{
			Name:  "conf, c",
			Value: "config.toml",
			Usage: "toml配置文件",
		},
		cli.StringFlag{
			Name:  "args",
			Value: "",
			Usage: "multiconfig cmd line args",
		},
	},
	Action: runInitDB,
}

func runInitDB(c *cli.Context) {
	conf.Init(c.String("conf"), c.String("args"))
	db.Init()

	// create table
	create := db.DB().Set("gorm:table_options", "ENGINE=InnoDB CHARSET=utf8mb4")
	//create.CreateTable()

	//db.DB().DropTableIfExists(
	//	&models.User{},
	//)

	create.AutoMigrate(
		&models.Account{},
	)
}
