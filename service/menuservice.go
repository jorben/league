package service

import (
	"github.com/gin-gonic/gin"
	"league/dal"
	"league/log"
	"league/model"
)

type MenuService struct {
	Ctx     *gin.Context
	MenuDal *dal.MenuDal
}

// NewMenuService 新建MenuService实例
func NewMenuService(ctx *gin.Context) *MenuService {

	return &MenuService{
		Ctx:     ctx,
		MenuDal: dal.NewMenuDal(ctx),
	}
}

// GetUserMenus 获取用户权限范围内的菜单
func (m *MenuService) GetUserMenus(t string) ([]*model.Menu, error) {
	menus, err := m.MenuDal.GetMenus(t)
	if err != nil {
		log.Errorf(m.Ctx, "Get menus failed, err: %s", err.Error())
		return nil, err
	}
	// TODO：菜单权限校验
	return menus, nil
}

// GetAllMenus 获取全部的菜单
func (m *MenuService) GetAllMenus() (map[string][]*model.Menu, error) {
	menuTypes, err := m.MenuDal.GetMenuTypes()
	if err != nil {
		log.Errorf(m.Ctx, "Get menu types failed, err: %s", err.Error())
		return nil, err
	}

	result := make(map[string][]*model.Menu)
	for _, menuType := range menuTypes {
		menus, err := m.MenuDal.GetMenus(menuType)
		if err != nil {
			log.Errorf(m.Ctx, "Get menus failed, err: %s", err.Error())
			continue
		}
		result[menuType] = menus
	}
	return result, nil
}
