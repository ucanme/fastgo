package models

import "gorm.io/gorm"

type MoveUnit struct {
	MoveUnitSn string `json:"move_unit_sn" gorm:"column:move_unit_sn"`
	Soc int `json:"soc" gorm:"column:soc"`
	Status int `json:"status" gorm:"column:status"`
	Speed int `json:"speed" gorm:"column:speed"`
	CurrentStationCode string `json:"current_station_code" gorm:"column:current_station_code"`
	IsInStatus int `json:"is_in_station" gorm:"column:is_in_station"`
	RingAngle float32 `json:"ring_angle" gorm:"column:ring_angle"`
	RingStatus int `json:"ring_status" gorm:"column:ring_status"`
	WorkDuration int `json:"work_duration" gorm:"column:work_duration"`
	ProductionLineId int `json:"production_line_id" gorm:"column:production_line_id"`
	gorm.Model
}

func (MoveUnit)TableName()string  {
	return "move_unit"
}

type ProductionLine struct {
	ProductionLineId int `json:"production_line_id" gorm:"column:production_line_id"`
	ProductionLineName string `json:"production_line_name" gorm:"column:production_line_name"`
	gorm.Model
}


