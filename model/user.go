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

type UserSocialInfo struct {
	gorm.Model
	Source     string `gorm:"uniqueIndex:idx_s_id;not null;size:255" json:"source"`
	OpenId     string `gorm:"uniqueIndex:idx_s_id;not null;size:255" json:"open_id"` // 第三方系统传递过来的用户ID
	BindUserId uint   `gorm:"not null" json:"bind_user_id"`                          // 关联的用户ID
	Email      string `gorm:"size:255;not null;default:''" json:"email"`
	Avatar     string `gorm:"size:255;not null;default:''" json:"avatar"`
	Username   string `gorm:"size:255;not null;default:''" json:"username"`
	Nickname   string `gorm:"size:255;not null;default:''" json:"nickname"`
	Bio        string `gorm:"size:255;not null;default:''" json:"bio"`
	Phone      string `gorm:"size:64;not null;default:''" json:"phone"`              // 手机
	Gender     uint8  `gorm:"default:0;not null;comment:0-未知,1-男,2-女" json:"gender"` //性别
}

type UserList struct {
	Count int64   //总记录数
	List  []*User // 用户信息
}

type UserWithExt struct {
	User
	Group  []string          `json:"group"`  // 用户权限组
	Source []*UserSocialInfo `json:"source"` // 用户绑定的登录源
}

type UserListWithExt struct {
	Count int64
	List  []*UserWithExt
}
