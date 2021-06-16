package models

import "github.com/jinzhu/gorm"

type Article struct {
	gorm.Model
	Content string `json:"text"  gorm:"text"`
	CateId int `json:"cate_id" gorm:"int"`
	ImgUrl string `json:"img_url"`
	Status int `json:"status"`
	Additional string `json:"additional"`
	Additional01 string `json:"additional_01"`
	Additional02 int `json:"additional_02"`
	Additional03 int `json:"additional_03"`
	Additional04 string `json:"additional_04"`

}