package model

import "gorm.io/gorm"

// 增加信用卡结构体
type CreditCard struct {
	gorm.Model
	Number string
	UserID uint
}
