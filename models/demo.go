package models

import "github.com/jinzhu/gorm"

type Demo struct {
	gorm.Model
}


type PreOrder struct {
	gorm.Model
	Name string `json:"name"`
	PhoneNum string `json:"phone_num"`
	Date string `json:"date"`
	PlaceId int `json:"place_id"`
	PersonCnt int `json:"person_cnt"`
	Addition string `json:"addition"`
	Addtion01 string `json:"addtion_01"`
}

