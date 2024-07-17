package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"league/dal"
	"league/log"
	"league/model"
	"strconv"
)

type UserService struct {
	Ctx       *gin.Context
	UserDal   *dal.UserDal
	CasbinDal *dal.CasbinDal
}

// NewUserService 新建AuthService实例
func NewUserService(ctx *gin.Context) *UserService {
	return &UserService{
		Ctx:       ctx,
		UserDal:   dal.NewUserDal(ctx),
		CasbinDal: dal.NewCasbinDal(ctx),
	}
}

// GetGroups 获取用户组及所包含的用户ID
func (u *UserService) GetGroups() (map[string][]string, error) {
	groups, err := u.CasbinDal.GetAllGroups()
	if err != nil {
		log.Errorf(u.Ctx, "Casbin get all groups failed, err: %s", err.Error())
		return nil, err
	}
	if len(groups) == 0 {
		return nil, nil
	}
	result := make(map[string][]string, len(groups))
	// 遍历用户组 获取用户
	for _, group := range groups {
		users, err := u.CasbinDal.GetGroupUsers(group)
		if err != nil {
			result[group] = nil
			log.Errorf(u.Ctx, "Casbin get group users failed, group: %s, err: %s", group, err.Error())
			continue
		}
		result[group] = users
	}
	return result, nil
}

// JoinGroup 加入用户组
func (u *UserService) JoinGroup(id uint, group string, newFlag bool) (bool, error) {
	strId := strconv.Itoa(int(id))
	if !newFlag {
		// 非新增用户组常见 检查用户组是否存在
		users, err := u.CasbinDal.GetGroupUsers(group)
		if err != nil {
			log.Errorf(u.Ctx, "Casbin get group users failed, group: %s, err: %s", group, err.Error())
			return false, err
		}
		if len(users) == 0 {
			return false, errors.New(fmt.Sprintf("%s group is not exist", group))
		}
	}
	result, err := u.CasbinDal.JoinGroups(strId, []string{group})
	if err != nil {
		log.Errorf(u.Ctx, "Casbin join groups failed, id: %d, group: %s, err: %s", id, group, err.Error())
		return false, err
	}
	return result, nil
}

// ExitGroup 退出用户组
func (u *UserService) ExitGroup(id uint, group string) (bool, error) {
	strId := strconv.Itoa(int(id))
	result, err := u.CasbinDal.ExitGroup(strId, group)
	if err != nil {
		log.Errorf(u.Ctx, "Casbin exit group failed, id: %d, group: %s, err: %s", id, group, err.Error())
		return false, err
	}
	return result, nil
}

func (u *UserService) DeleteUser(id uint) (bool, error) {
	// 解绑所有该用户的登录渠道
	if _, err := u.UnbindUserSource(id, ""); err != nil {
		return false, err
	}
	// 删除用户主表记录
	result, err := u.UserDal.DeleteUser(&model.User{
		Model: gorm.Model{ID: id},
	})
	if err != nil {
		log.Errorf(u.Ctx, "Delete user failed, err: %s", err.Error())
		return false, err
	}
	log.Debugf(u.Ctx, "Delete user id: %d, result: %v", id, result)
	return result, nil
}

// UnbindUserSource 解绑用户登录渠道
func (u *UserService) UnbindUserSource(id uint, source string) (bool, error) {
	info := &model.UserSocialInfo{
		BindUserId: id,
		Source:     source,
	}
	result, err := u.UserDal.DeleteUserSource(info)
	if err != nil {
		log.Errorf(u.Ctx, "Delete user source failed, err: %s", err.Error())
		return false, err
	}
	return result, nil
}

// UpdateUserStatus 更新用户状态
func (u *UserService) UpdateUserStatus(id uint, status uint8) (bool, error) {
	user, err := u.GetUserInfo(id)
	if err != nil {
		return false, err
	}
	if user == nil {
		return false, errors.New("user not found")
	}
	if user.Status == status {
		log.Debug(u.Ctx, "Current status is equal to target status")
		return true, nil
	}
	user.Status = status
	_, err = u.UserDal.UpdateUser(user)
	if err != nil {
		log.Errorf(u.Ctx, "Update user failed, err: %s", err.Error())
		return false, err
	}
	return true, nil
}

// GetUserInfo 根据用户ID获取用户基本信息
func (u *UserService) GetUserInfo(id uint) (*model.User, error) {
	user := &model.User{
		Model: gorm.Model{ID: id},
	}
	result, err := u.UserDal.GetUserList(user, 0, 1)
	if err != nil {
		log.Errorf(u.Ctx, "Get userinfo failed, err: %s", err.Error())
		return nil, err
	}
	if result == nil || result.Count == 0 || len(result.List) == 0 {
		return nil, nil
	}
	return result.List[0], nil
}

// GetUserList 查询用户列表
func (u *UserService) GetUserList(search string, offset int, limit int, order []string) (*model.UserListWithExt, error) {

	log.Debugf(u.Ctx, "search: %s, offset: %d, limit: %d, order: %v", search, offset, limit, order)
	result, err := u.UserDal.GetUserListbySearch(search, offset, limit, order)
	if err != nil {
		log.Errorf(u.Ctx, "Get userlist by search failed, err: %s", err.Error())
		return nil, err
	}

	var userListExt []*model.UserWithExt

	if result != nil && result.Count > 0 {
		var strUserIdList []string
		var intUserIdList []uint
		for _, user := range result.List {
			strUserIdList = append(strUserIdList, strconv.Itoa(int(user.ID)))
			intUserIdList = append(intUserIdList, user.ID)
		}
		// 批量查询角色信息
		group, err := u.CasbinDal.GetUserGroup(strUserIdList, "g")
		if err != nil {
			log.Errorf(u.Ctx, "Get user group list failed, err: %s", err.Error())
			return nil, err
		}
		log.Debugf(u.Ctx, "user group: %v", group)

		// 批量查询登录来源信息
		source, err := u.UserDal.GetUserSource(intUserIdList)
		if err != nil {
			log.Errorf(u.Ctx, "Get user source list failed, err: %s", err.Error())
			return nil, err
		}
		//log.Debugf(u.Ctx, "user source: %v", source)

		for _, user := range result.List {
			userExt := &model.UserWithExt{
				User: *user,
			}
			// 补充角色信息
			if g, exists := group[strconv.Itoa(int(user.ID))]; exists {
				userExt.Group = g
			}
			// 补充登录来源信息
			if s, exists := source[user.ID]; exists {
				userExt.Source = s
			}
			userListExt = append(userListExt, userExt)
		}
	}
	return &model.UserListWithExt{
		Count: result.Count,
		List:  userListExt,
	}, nil
}
