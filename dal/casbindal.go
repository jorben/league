package dal

import (
	"errors"
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"league/config"
	"league/database"
	"league/log"
	"league/model"
)

type CasbinDal struct {
	e   *casbin.Enforcer
	ctx *gin.Context
	db  *gorm.DB
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
		db:  db,
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

// GetRules 获取规则集
func (c *CasbinDal) GetRules(rule *model.CasbinRule) ([]*model.CasbinRule, error) {
	var result []*model.CasbinRule
	if err := c.db.Where(rule).Find(&result).Error; err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserGroup 批量获取用户的角色组
func (c *CasbinDal) GetUserGroup(id []string, ptype string) (map[string][]string, error) {
	var result []*model.CasbinRule
	if err := c.db.Where("ptype = ? AND v0 IN ?", ptype, id).Find(&result).Error; err != nil {
		return nil, err
	}
	group := make(map[string][]string, len(result))
	for _, rule := range result {
		if _, exists := group[rule.V0]; exists {
			group[rule.V0] = append(group[rule.V0], rule.V1)
		} else {
			group[rule.V0] = []string{rule.V1}
		}
	}
	return group, nil
}

// JoinGroups 为用户加入用户组
func (c *CasbinDal) JoinGroups(id string, groups []string) (bool, error) {
	return c.e.AddRolesForUser(id, groups)
}

// ExitGroup 为用户退出用户组
func (c *CasbinDal) ExitGroup(id string, groups string) (bool, error) {
	return c.e.DeleteRoleForUser(id, groups)
}

// GetGroupUsers 获取指定用户组的所有用户
func (c *CasbinDal) GetGroupUsers(group string) ([]string, error) {
	return c.e.GetUsersForRole(group)
}

// GetAllGroups 获取所有用户组
func (c *CasbinDal) GetAllGroups() ([]string, error) {
	return c.e.GetAllRoles()
}
