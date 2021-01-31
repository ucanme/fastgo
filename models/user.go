package models

type User struct {
	UserId string `json:"user_id" gorm:"varchar(32) not null unique"`
	Password string `json:"password" gorm:"varchar(64) not null unique"`
}
