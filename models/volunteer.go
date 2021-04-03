package models

import "github.com/jinzhu/gorm"

type Volunteer struct {
	Model gorm.Model
	Name string `json:"name" gorm:"varchar(255)"`
	Address string `json:"address" gorm:"varchar(255)"`
	PhoneNum string `json:"phone_num" gorm:"varchar(255)"`
	ArticleId string `json:"article_id" gorm:"int"`
	OpenId string `json:"open_id" gorm:"varchar(255)"`
}
