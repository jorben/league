package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Nickname string `gorm:"size:255;not null;default:''" json:"nickname"`
	Email    string `gorm:"size:255;not null;default:''" json:"email"`             // 邮件
	Avatar   string `gorm:"size:255;not null;default:''" json:"avatar"`            // 头像
	Phone    string `gorm:"size:64;not null;default:''" json:"phone"`              // 手机
	Bio      string `gorm:"size:255;not null;default:''" json:"bio"`               // 简介
	Gender   uint8  `gorm:"default:0;not null;comment:0-未知,1-男,2-女" json:"gender"` //性别
	Status   uint8  `gorm:"default:0;not null;comment:0-正常,1-禁用" json:"status"`    // 用户状态
}

type UserList struct {
	Count int64   //总记录数
	List  []*User // 用户信息
}

type UserWithExt struct {
	User
	Group  []string `json:"group"`  // 用户权限组
	Source []string `json:"source"` // 用户绑定的登录源
}

type UserListWithExt struct {
	Count int64
	List  []*UserWithExt
}
