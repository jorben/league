package dal

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"league/database"
	"league/model"
)

type MenuDal struct {
	db *gorm.DB
}

// NewMenuDal 创建User数据访问层实例
func NewMenuDal(ctx *gin.Context) *MenuDal {
	return &MenuDal{
		db: database.GetInstance().WithContext(ctx),
	}
}

// GetMenus 根据菜单类型获取完整菜单列表
func (m *MenuDal) GetMenus(t string) ([]*model.Menu, error) {
	var roots []*model.Menu
	result := m.db.Where(&model.Menu{Type: t}).Where("parent = ?", "").Order("`order` asc").Find(&roots)
	if result.RowsAffected == 0 {
		return nil, nil
	}
	// 遍历所有一级节点，获取子节点
	for _, root := range roots {
		if err := m.findChildren(root); err != nil {
			return nil, err
		}
	}
	return roots, nil
}

func (m *MenuDal) findChildren(parent *model.Menu) error {
	if err := m.db.Where(&model.Menu{Parent: parent.Key, Type: parent.Type}).Order("`order` asc").Find(&parent.Children).Error; err != nil {
		return err
	}
	for _, child := range parent.Children {
		if err := m.findChildren(child); err != nil {
			return err
		}
	}
	return nil
}

// GetMenuTypes 获取菜单类别
func (m *MenuDal) GetMenuTypes() ([]string, error) {
	var result []string
	if err := m.db.Model(&model.Menu{}).Distinct("type").Pluck("type", &result).Error; err != nil {
		return nil, err
	}
	return result, nil
}
