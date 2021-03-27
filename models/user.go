package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserId string `json:"user_id" gorm:"type:varchar(64) not null"`
	OpenId string `json:"open_id"  gorm:"type:varchar(64) not null"`
	PhoneNum string `json:"phone_num" gorm:"type:varchar(64) not null"`
	LastLoginTime int64 `json:"last_login_time" gorm:"type:varchar(64) not null"`
	AvatarUrl string `json:"avatar_url" gorm:"type:varchar(128) not null"`
	Status int8 `json:"status" gorm:"type:tinyint not null"`
	Score int64  `json:"avatar_url" gorm:"type:varchar(128) not null"`
}
