package models

import "gorm.io/gorm"



type ProductionLine struct {
	ProductionLineId int `json:"production_line_id" gorm:"column:production_line_id"`
	ProductionLineName string `json:"production_line_name" gorm:"column:production_line_name"`
	gorm.Model
}


