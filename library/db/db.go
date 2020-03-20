package db

import (
	"errors"
	"fmt"
	"github.com/ucanme/fastgo/conf"
	"time"

	"github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var gDB *gorm.DB

func Init() error {
	if conf.Config.Database.Enable {
		_db, err := ConnectToDB(&conf.Config.Database)
		if err != nil {
			//prometheus.ModuleApiError("db", "init")
			return err
		}
		gDB = _db
		return nil
	}
	return errors.New("db disabled")
}

func ConnectToDB(cfg *conf.DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s@%s/%s?charset=utf8mb4&parseTime=True&loc=Local&timeout=10s&readTimeout=10s&writeTimeout=10s",
		cfg.UserPassword, cfg.Write.HostPort, cfg.DB)

	c, err := mysql.ParseDSN(dsn)
	if err != nil {
		//prometheus.ModuleApiError("db", "parse_dsn")
		return nil, err
	}

	db, err := gorm.Open("mysql", c.FormatDSN())
	if err != nil {
		//prometheus.ModuleApiError("db", "open_mysql")
		return nil, err
	}

	if cfg.Conn.MaxLifeTime > 0 {
		db.DB().SetConnMaxLifetime(time.Duration(conf.Config.Database.Conn.MaxLifeTime) * time.Second)
	}

	if cfg.Conn.MaxIdle > 0 {
		db.DB().SetMaxIdleConns(conf.Config.Database.Conn.MaxIdle)
	}

	if cfg.Conn.MaxOpen > 0 {
		db.DB().SetMaxOpenConns(conf.Config.Database.Conn.MaxOpen)
	}

	db.LogMode(true)
	return db, nil
}

func DB() *gorm.DB {
	return gDB
}
