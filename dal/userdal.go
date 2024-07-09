package dal

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"league/common/util"
	"league/database"
	"league/model"
	"strconv"
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

// UpdateUser 会更新零值，无主键则创建记录
func (u *UserDal) UpdateUser(user *model.User) (int64, error) {
	result := u.db.Save(user)
	if result.Error != nil {
		return 0, result.Error
	}
	return result.RowsAffected, nil
}

// GetUserList 获取用户基本信息
func (u *UserDal) GetUserList(where *model.User, offset int, limit int) (*model.UserList, error) {
	var count int64
	var list []*model.User
	// 计算总条数
	if err := u.db.Model(where).Where(where).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		// 获取数据
		if err := u.db.Model(where).Where(where).Offset(offset).Limit(limit).Order("ID desc").Find(&list).Error; err != nil {
			return nil, err
		}
	}

	return &model.UserList{
		Count: count,
		List:  list,
	}, nil
}

// GetUserListbySearch 用户列表查询接口
func (u *UserDal) GetUserListbySearch(searchKey string, offset int, limit int, order []string) (*model.UserList, error) {
	var count int64
	var list []*model.User
	db := u.db.Model(&model.User{})
	if len(searchKey) > 0 {
		// 判定searchKey
		intKey, err := strconv.ParseUint(searchKey, 10, 64)
		if err == nil && intKey > 0 {
			// 通过ID查询
			db = db.Where(&model.User{
				Model: gorm.Model{ID: uint(intKey)},
			})
		} else if util.IsValidEmail(searchKey) {
			// 通过邮箱查询
			db = db.Where("email LIKE ?", fmt.Sprintf("%%%s%%", searchKey))
		} else {
			// 通过昵称查询
			db = db.Where("nickname LIKE ?", fmt.Sprintf("%%%s%%", searchKey))
		}
	}

	// 计算总条数
	if err := db.Count(&count).Error; err != nil {
		return nil, err
	}

	if count > 0 {
		if len(order) > 0 {
			for _, o := range order {
				db = db.Order(o)
			}
		} else {
			db.Order("ID desc")
		}
		if err := db.Offset(offset).Limit(limit).Find(&list).Error; err != nil {
			return nil, err
		}
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

// GetUserSource 批量获取用户绑定的渠道
func (u *UserDal) GetUserSource(id []uint) (map[uint][]*model.UserSocialInfo, error) {
	var result []*model.UserSocialInfo
	if err := u.db.Model(&model.UserSocialInfo{}).Where("bind_user_id IN ?", id).Find(&result).Error; err != nil {
		return nil, err
	}
	source := make(map[uint][]*model.UserSocialInfo, len(result))
	for _, info := range result {
		if _, exists := source[info.BindUserId]; exists {
			source[info.BindUserId] = append(source[info.BindUserId], info)
		} else {
			source[info.BindUserId] = []*model.UserSocialInfo{info}
		}
	}
	return source, nil
}
