package service

import (
	"github.com/gin-gonic/gin"
	"league/common/errs"
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

// SaveMenu 创建/更新菜单项
func (m *MenuService) SaveMenu(menu *model.Menu) (bool, error) {
	// parent不为空 则检查parent是否存在
	if len(menu.Parent) > 0 {
		_, err := m.MenuDal.GetMenuItem(&model.Menu{Key: menu.Parent, Type: menu.Type})
		if err != nil {
			log.Errorf(m.Ctx, "Query parent failed, err: %s", err.Error())
			return false, err
		}
	}
	id, err := m.MenuDal.SaveMenu(menu)
	if err != nil {
		log.Errorf(m.Ctx, "Save menu failed, err: %s", err.Error())
		return false, err
	}
	return id > 0, nil
}

// DeleteMenu 删除菜单项
func (m *MenuService) DeleteMenu(id uint) error {
	menu := &model.Menu{
		ID: id,
	}
	hasChildren, err := m.MenuDal.HasChildren(menu)
	if err != nil {
		log.Errorf(m.Ctx, "Query children failed, err: %s", err.Error())
		return err
	}
	if hasChildren {
		return errs.ErrorHasChildren
	}
	err = m.MenuDal.DeleteMenu(menu)
	if err != nil {
		log.Errorf(m.Ctx, "Delete menu failed, err: %s", err.Error())
		return err
	}
	return nil
}
