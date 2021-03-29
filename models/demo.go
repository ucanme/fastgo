package models

import "github.com/jinzhu/gorm"

type Demo struct {
	gorm.Model
}


type PreOrder struct {
	gorm.Model
	Name string `json:"name"`
	Phone string `json:"phone"`
	Date string `json:"date"`
	PlaceId int `json:"place_id"`
}

