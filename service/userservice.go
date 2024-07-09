package service

import (
	"errors"
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
