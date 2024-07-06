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

// GetUserinfo 获取用户基本信息
func (u *UserDal) GetUserinfo(where *model.User, offset int, limit int) (*model.UserList, error) {
	var count int64
	var list []*model.User
	// 计算总条数
	if err := u.db.Model(where).Where(where).Count(&count).Error; err != nil {
		return nil, err
	}
	// 获取数据
	if err := u.db.Model(where).Where(where).Offset(offset).Limit(limit).Find(&list).Error; err != nil {
		return nil, err
	}

	return &model.UserList{
		Count: count,
		List:  list,
	}, nil
}

// UpdateBySource 更新第三方渠道来源用户信息
func (u *UserDal) UpdateBySource(info *model.UserSocialInfo) (bool, error) {
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
func (u *UserDal) SignUpBySource(info *model.UserSocialInfo) (uint, error) {
	var id uint
	err := u.db.Transaction(func(tx *gorm.DB) error {
		user := &model.User{
			Nickname: info.Nickname,
			Email:    info.Email,
			Avatar:   info.Avatar,
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
	result := u.db.Where(data).First(data)
	if result.RowsAffected == 0 {
		return nil
	}
	return data
}
