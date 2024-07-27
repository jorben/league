package dal

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"league/common/errs"
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

// GetMenuItem 查询单个菜单项
func (m *MenuDal) GetMenuItem(menu *model.Menu) (*model.Menu, error) {
	result := &model.Menu{}
	if err := m.db.Where(menu).First(result).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errs.ErrorRecordNotFound
		}
		return nil, err
	}
	return result, nil
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

// SaveMenu 更新菜单项，会更新零值，无主键则创建记录
func (m *MenuDal) SaveMenu(menu *model.Menu) (uint, error) {
	result := m.db.Model(menu).Save(menu)
	if result.Error != nil {
		return 0, result.Error
	}
	return menu.ID, nil
}

// DeleteMenu 删除菜单项
func (m *MenuDal) DeleteMenu(menu *model.Menu) error {
	return m.db.Where(menu).Delete(menu).Error
}

// HasChildren 检查一个菜单项是否含有子菜单
func (m *MenuDal) HasChildren(menu *model.Menu) (bool, error) {
	var num int64
	if len(menu.Key) > 0 && len(menu.Type) > 0 {
		if err := m.db.Model(menu).Where(&model.Menu{Parent: menu.Key, Type: menu.Type}).Count(&num).Error; err != nil {
			return false, err
		}
		return num > 0, nil
	} else if menu.ID > 0 {
		if err := m.db.Model(menu).Where(&model.Menu{ID: menu.ID}).First(menu).Error; err != nil {
			return false, err
		}
		return m.HasChildren(menu)
	} else {
		return false, errors.New("missing query conditions")
	}
}
