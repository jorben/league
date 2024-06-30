package dal

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"league/database"
	"league/model"
)

type UserDal struct {
	db *gorm.DB
}

// NewUserDal 创建User数据访问层实例
func NewUserDal(ctx *gin.Context) *UserDal {
	return &UserDal{
		db: database.GetInstance().WithContext(ctx),
	}
}

// UpdateBySource 更新第三方渠道来源用户信息
func (u *UserDal) UpdateBySource(info model.UserSocialInfo) (bool, error) {
	// 查询id
	oldInfo := u.GetUserBySource(info.Source, info.OpenId)
	if oldInfo == nil {
		return false, errors.New(fmt.Sprintf("User not found, source: %s, openId: %s", info.Source, info.OpenId))
	}
	result := u.db.Updates(model.UserSocialInfo{
		Model:    gorm.Model{ID: oldInfo.ID},
		Email:    info.Email,
		Avatar:   info.Avatar,
		Username: info.Username,
		Nickname: info.Nickname,
		Bio:      info.Bio,
		Phone:    info.Phone,
		Gender:   info.Gender,
	})
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected == 1, nil
}

// SignUpBySource 第三方渠道首次注册
func (u *UserDal) SignUpBySource(info model.UserSocialInfo) (uint, error) {
	var id uint
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user := &model.User{
			Nickname: info.Nickname,
			Email:    info.Email,
			Phone:    info.Phone,
			Gender:   info.Gender,
			Bio:      info.Bio,
		}
		// 插入用户主表
		if err := tx.Create(user).Error; err != nil {
			return err
		}

		// 获取用户ID
		if err := tx.Raw("SELECT LAST_INSERT_ID()").Scan(&id).Error; err != nil {
			return err
		}

		info.BindUserId = id
		// 插入第三方用户信息表
		if err := tx.Create(&info).Error; err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return 0, err
	}
	return id, nil
}

// GetUserBySource 获取第三方渠道用户信息
func (u *UserDal) GetUserBySource(source string, openid string) *model.UserSocialInfo {
	data := &model.UserSocialInfo{
		Source: source,
		OpenId: openid,
	}
	result := u.db.First(data)
	if result.RowsAffected == 0 {
		return nil
	}
	return data
}
