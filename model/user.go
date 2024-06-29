package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nickname string `gorm:"size:255;not null;default:''" json:"nickname"`
	Email    string `gorm:"size:255;not null;default:''" json:"email"`             // 邮件
	Phone    string `gorm:"size:64;not null;default:''" json:"phone"`              // 手机
	Bio      string `gorm:"size:255;not null;default:''" json:"bio"`               // 简介
	Gender   uint8  `gorm:"default:0;not null;comment:0-未知,1-男,2-女" json:"gender"` //性别
}
