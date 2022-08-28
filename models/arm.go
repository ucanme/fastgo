package models

type Arm struct {
	ArmSn string `json:"arm_sn" gorm:"column:arm_sn"`
	ArmName string `json:"arm_name" gorm:"column:arm_name"`
	Deleted int `json:"deleted" gorm:"column:deleted"`
	Model
}

func (Arm)TableName()string  {
	return "arm"
}

