package dal

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"league/database"
	"league/model"
)

type ApiDal struct {
	db *gorm.DB
}

// NewApiDal 创建User数据访问层实例
func NewApiDal(ctx *gin.Context) *ApiDal {
	return &ApiDal{
		db: database.GetInstance().WithContext(ctx),
	}
}

// GetApiList 获取Api列表数据
func (a *ApiDal) GetApiList(where *model.Api, offset int, limit int) (*model.ApiList, error) {
	var count int64
	var list []*model.Api
	// 计算总条数
	if err := a.db.Model(where).Where(where).Count(&count).Error; err != nil {
		return nil, err
	}
	if count > 0 {
		// 获取数据
		if err := a.db.Model(where).Where(where).Offset(offset).Limit(limit).Order("name ASC, path ASC").Find(&list).Error; err != nil {
			return nil, err
		}
	}

	return &model.ApiList{
		Count: count,
		List:  list,
	}, nil
}

// SaveApi 更新API，会更新零值，无主键则创建记录
func (a *ApiDal) SaveApi(api *model.Api) (uint, error) {
	result := a.db.Model(api).Save(api)
	if result.Error != nil {
		return 0, result.Error
	}
	return api.ID, nil
}

// DeleteApi 删除API接口
func (a *ApiDal) DeleteApi(api *model.Api) error {
	return a.db.Where(api).Delete(api).Error
}
