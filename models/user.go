package models

type Appointment struct {
	Date string `json:"date"`
	Hour int `json:"hour"`
	Minute int `json:"minute"`
	Status int `json:"status"` //0表示未预约 1表示已经预约 2 已经签到
	Name string `json:"name"`
	OpenId string `json:"open_id"`
	PhoneNum string `json:"phone_num"`
}