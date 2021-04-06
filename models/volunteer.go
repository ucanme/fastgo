package models

import "github.com/jinzhu/gorm"

type Volunteer struct {
	gorm.Model
	Name string `json:"name" gorm:"varchar(255)"`
	Address string `json:"address" gorm:"varchar(255)"`
	PhoneNum string `json:"phone_num" gorm:"varchar(255)"`
	ArticleId int `json:"article_id" gorm:"int"`
	OpenId string `json:"open_id" gorm:"varchar(255)"`
	Length string `json:"length" gorm:"varchar(255)"`
	IdNo string `json:"id_no" gorm:"varchar(255)"`
}
