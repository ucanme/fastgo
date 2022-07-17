package models

import "gorm.io/gorm"



type ProductionLine struct {
	ProductionLineId int `json:"production_line_id" gorm:"column:production_line_id"`
	ProductionLineName string `json:"production_line_name" gorm:"column:production_line_name"`
	gorm.Model
}




type Station struct {
	StationCode string `json:"station_code"  gorm:"column:station_code"`
	ProductionLineId int `json:"production_line_id" gorm:"column:production_line_id"`
	StationID string `json:"station_id" gorm:"column:station_id"`
	gorm.Model
}