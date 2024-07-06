package api

import (
	"github.com/gin-gonic/gin"
	"league/common/context"
	"league/common/errs"
	"league/model"
	"league/service"
)

// GetIndexMenus 获取前台权限范围内的菜单项
func GetIndexMenus(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	if menus, err := getMenus(ctx, model.MenuTypeIndex); err == nil {
		c.CJSON(errs.Success, menus)
	} else {
		c.CJSON(errs.ErrMenu)
	}
}

// GetAdminMenus 获取管理后台权限范围内的菜单项
func GetAdminMenus(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	if menus, err := getMenus(ctx, model.MenuTypeAdmin); err == nil {
		c.CJSON(errs.Success, menus)
	} else {
		c.CJSON(errs.ErrMenu)
	}
}

func getMenus(ctx *gin.Context, t string) ([]*model.Menu, error) {
	menuService := service.NewMenuService(ctx)
	return menuService.GetUserMenus(t)
}
