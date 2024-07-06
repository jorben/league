package service

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"league/dal"
	"league/log"
	"league/model"
)

type UserService struct {
	Ctx     *gin.Context
	UserDal *dal.UserDal
}

// NewUserService 新建AuthService实例
func NewUserService(ctx *gin.Context) *UserService {
	return &UserService{
		Ctx:     ctx,
		UserDal: dal.NewUserDal(ctx),
	}
}

// GetUserInfo 根据用户ID获取用户基本信息
func (u *UserService) GetUserInfo(id uint) *model.User {
	user := &model.User{
		Model: gorm.Model{ID: id},
	}
	result, err := u.UserDal.GetUserinfo(user, 0, 1)
	if err != nil {
		log.Errorf(u.Ctx, "Get userinfo failed, err: %s", err.Error())
		return nil
	}
	if result == nil || result.Count == 0 || len(result.List) == 0 {
		return nil
	}
	return result.List[0]
}
