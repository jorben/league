package dal

import (
	"errors"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"league/config"
	"league/database"
	"league/log"
	"league/model"
)

type CasbinDal struct {
	e   *casbin.Enforcer
	ctx *gin.Context
}

// initialRules 初始化规则，仅当规则表为空时注入默认规则
func initialRules(e *casbin.Enforcer) error {

	policy, err := e.GetPolicy()
	if err != nil {
		return err
	}
	if policy == nil || len(policy) == 0 {
		// 策略数据为空，初始化策略数据
		ok, err := e.AddPolicies(config.CasbinDefaultPolicies)
		if err != nil {
			return err
		} else if !ok {
			return errors.New("add default policies failed")
		}
	}
	return nil
}

func NewCasbinDal(ctx *gin.Context) *CasbinDal {
	db := database.GetInstance().WithContext(ctx)
	adapter, err := gormadapter.NewAdapterByDBWithCustomTable(db, &model.CasbinRule{}, "t_casbin_rule")
	if err != nil {
		log.Errorf(ctx, "New casbin database adapter failed, err: %s", err.Error())
		panic(err)
	}
	enforcer, err := casbin.NewEnforcer("config/rbac_model.conf", adapter)
	if err != nil {
		log.Errorf(ctx, "New casbin enforcer failed, err: %s", err.Error())
		panic(err)
	}
	if err = initialRules(enforcer); err != nil {
		log.Error(ctx, err.Error())
		panic(err)
	}
	return &CasbinDal{
		e:   enforcer,
		ctx: ctx,
	}
}

// IsAllow 权限校验
func (c *CasbinDal) IsAllow(req model.CasbinReq) bool {
	if ok, err := c.e.Enforce(req.UserId, req.Path, req.Method); err != nil {
		log.Errorf(c.ctx, "Casbin enforce failed, err: %s", err.Error())
		// 系统故障 权限默认不开放
		return false
	} else if !ok {
		return false
	}
	return true
}
