package dal

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"league/database"
	"league/log"
	"league/model"
)

type CasbinDal struct {
	e *casbin.Enforcer
}

func NewCasbinDal() *CasbinDal {
	db := database.GetInstance()
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(db, &model.CasbinRule{}, "t_casbin_rule")
	if err != nil {
		log.Errorf("New casbin database adapter failed, err: %s", err.Error())
	}
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		log.Errorf("New casbin enforcer failed, err: %s", err.Error())
	}
	return &CasbinDal{
		e: enforcer,
	}
}

// IsAllow 权限校验
func (c *CasbinDal) IsAllow(req model.CasbinReq) bool {
	if ok, err := c.e.Enforce(req.UserId, req.Path, req.Method); err != nil {
		log.Errorf("Casbin enforce failed, err: %s", err.Error())
		// 系统故障 权限默认不开放
		return false
	} else if !ok {
		return false
	}
	return true
}
