package models

type Arm struct {
	ArmSn string `json:"arm_sn" gorm:"column:arm_sn"`
	Deleted int `json:"deleted" gorm:"column:deleted"`
	Status int `json:"status" gorm:"column:status"`
	JointPosition string `json:"joint_position" gorm:"column:joint_position"`
	ActualPosition string `json:"actual_position" gorm:"column:actual_position"`
	Model
}

func (Arm)TableName()string  {
	return "arm"
}