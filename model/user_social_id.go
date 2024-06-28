package model

import (
	"gorm.io/gorm"
)

type UserSocialId struct {
	gorm.Model
	Source string `gorm:"uniqueIndex:idx_s_id;not null;size:255" json:"source"`
	OpenId string `gorm:"uniqueIndex:idx_s_id;not null;size:255" json:"open_id"` // 第三方系统传递过来的用户ID

	BindUserId uint `gorm:"not null" json:"bind_user_id"` // 关联的用户ID

}
