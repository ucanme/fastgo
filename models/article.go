package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Content string `json:"text"  gorm:"text"`
	CateId int `json:"cate_id" gorm:"int"`
}

