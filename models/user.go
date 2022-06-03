package models

import "github.com/jinzhu/gorm"

type Account struct {
	AccountId string `json:"account_id" gorm:"column:account_id"`
	Password string `json:"password" gorm:"column:password"`
	gorm.Model
}

func (Account)TableName()string  {
	return "account"
}