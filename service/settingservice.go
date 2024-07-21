package service

import (
	"github.com/gin-gonic/gin"
	"league/dal"
	"league/log"
	"league/model"
)

type SettingService struct {
	Ctx    *gin.Context
	ApiDal *dal.ApiDal
}

// NewSettingService 新建SettingService实例
func NewSettingService(ctx *gin.Context) *SettingService {
	return &SettingService{
		Ctx:    ctx,
		ApiDal: dal.NewApiDal(ctx),
	}
}

// GetApiList 获取API列表
func (s *SettingService) GetApiList(offset int, limit int) (*model.ApiList, error) {

	list, err := s.ApiDal.GetApiList(&model.Api{}, offset, limit)
	if err != nil {
		log.Errorf(s.Ctx, "Get api list failed, err: %s", err.Error())
		return nil, err
	}
	return list, nil
}

// SaveApi 创建/更新接口信息
func (s *SettingService) SaveApi(api *model.Api) (bool, error) {
	id, err := s.ApiDal.SaveApi(api)
	if err != nil {
		log.Errorf(s.Ctx, "Save api failed, err: %s", err.Error())
		return false, err
	}
	return id > 0, nil
}

// DeleteApi 删除接口
func (s *SettingService) DeleteApi(id uint) error {
	api := &model.Api{
		ID: id,
	}
	err := s.ApiDal.DeleteApi(api)
	if err != nil {
		log.Errorf(s.Ctx, "Delete api failed, err: %s", err.Error())
		return err
	}
	return nil
}
