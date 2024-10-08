package service

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"league/config"
	"league/dal"
	"league/log"
	"league/model"
	"time"
)

type AuthService struct {
	Ctx       *gin.Context
	UserDal   *dal.UserDal
	CasbinDal *dal.CasbinDal
	signKey   []byte
}

type AuthToken struct {
	Token     string `json:"token"`
	UserId    string `json:"user_id"`
	ExpiresAt int64  `json:"expires_at"`
	NotBefore int64  `json:"not_before"`
}

// NewAuthService 新建AuthService实例
func NewAuthService(ctx *gin.Context) *AuthService {

	return &AuthService{
		Ctx:       ctx,
		UserDal:   dal.NewUserDal(ctx),
		CasbinDal: dal.NewCasbinDal(ctx),
		signKey:   []byte(config.GetConfig().Jwt.SignKey),
	}
}

// LoginBySource 从第三方登陆
func (a *AuthService) LoginBySource(info *model.UserSocialInfo) (*AuthToken, error) {
	// 检查是否已存在该用户
	user := a.UserDal.GetUserBySource(info.Source, info.OpenId)
	if user == nil {
		// 不存在则插入用户
		id, err := a.UserDal.SignUpBySource(info)
		if err != nil {
			log.Errorf(a.Ctx, "UserDal.SignUpBySource failed, err: %s", err.Error())
			return nil, err
		}
		return a.SignJwtString(id)
	} else {
		// 存在则更新信息
		_, err := a.UserDal.UpdateBySource(info)
		if err != nil {
			log.Errorf(a.Ctx, "UserDal.UpdateBySource failed, err: %s", err.Error())
			return nil, err
		}
		return a.SignJwtString(user.BindUserId)
	}

}

// IsAllow 权限校验
func (a *AuthService) IsAllow(userId string, path string, method string) bool {
	req := model.CasbinReq{
		UserId: userId,
		Path:   path,
		Method: method,
	}
	isAllow := a.CasbinDal.IsAllow(req)
	log.WithField(a.Ctx, "Path", path, "Method", method).Debugf("IsAllow: %v", isAllow)
	return isAllow
}

// SignJwtString 签发JWT
func (a *AuthService) SignJwtString(id uint) (*AuthToken, error) {
	now := time.Now()
	expiresAt := now.Add(2 * time.Hour)
	notBefore := now.Add(-5 * time.Minute)
	claims := jwt.RegisteredClaims{
		Issuer:    "league",
		ExpiresAt: jwt.NewNumericDate(expiresAt),
		NotBefore: jwt.NewNumericDate(notBefore),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        fmt.Sprintf("%d", id),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	sign, err := token.SignedString(a.signKey)
	if err != nil {
		log.Infof(a.Ctx, "Jwt SignedString failed, err: %s", err.Error())
		return nil, err
	}
	return &AuthToken{
		Token:     sign,
		UserId:    fmt.Sprintf("%d", id),
		ExpiresAt: expiresAt.Unix(),
		NotBefore: notBefore.Unix(),
	}, nil
}

// VerifyJwtString 校验JWT
func (a *AuthService) VerifyJwtString(s string) (*AuthToken, error) {

	token, err := jwt.ParseWithClaims(s, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return a.signKey, nil
	})
	if err != nil {
		log.Errorf(a.Ctx, "Jwt parse failed, err: %s", err.Error())
		return nil, err
	} else if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok {
		log.Debugf(a.Ctx, "Check login passed, UserId: %s", claims.ID)
		return &AuthToken{
			UserId:    claims.ID,
			ExpiresAt: claims.ExpiresAt.Unix(),
			NotBefore: claims.NotBefore.Unix(),
		}, nil
	} else {
		log.Errorf(a.Ctx, "Unknown claims type, token: %s", s)
		return nil, errors.New("unknown claims type")
	}
}

// TODO: jwt失效逻辑，token续签

// SavePolicy 更新或创建权限规则
func (a *AuthService) SavePolicy(policy *model.Policy) (bool, error) {
	rule := &model.CasbinRule{
		ID:      policy.ID,
		Ptype:   "p",
		V0:      policy.Subject,
		V1:      policy.Path,
		V2:      policy.Method,
		V3:      policy.Result,
		Comment: policy.Comment,
	}
	id, err := a.CasbinDal.SaveRule(rule)
	if err != nil {
		log.Errorf(a.Ctx, "Save policy failed, err: %s", err.Error())
		return false, err
	}
	return id > 0, nil
}

// DeletePolicy 删除权限规则
func (a *AuthService) DeletePolicy(id uint) error {
	rule := &model.CasbinRule{
		ID: id,
	}
	err := a.CasbinDal.DeleteRule(rule)
	if err != nil {
		log.Errorf(a.Ctx, "Delete policy failed, err: %s", err.Error())
		return err
	}
	return nil
}

// GetPolicyList 获取策略列表
func (a *AuthService) GetPolicyList() (map[string][]*model.Policy, error) {

	rules, err := a.CasbinDal.GetRules(&model.CasbinRule{Ptype: "p"})
	if err != nil {
		log.Errorf(a.Ctx, "Casbin get rules failed, err: %s", err.Error())
		return nil, err
	}
	// 批量获取接口列表
	mapApiName := make(map[string]string)
	settingService := NewSettingService(a.Ctx)
	apis, err := settingService.GetApiList(0, 9999)
	if apis != nil {
		for _, api := range apis.List {
			mapApiName[fmt.Sprintf("%s-%s", api.Method, api.Path)] = api.Name
		}
	}
	result := make(map[string][]*model.Policy)
	for _, rule := range rules {
		policy := &model.Policy{}
		policy.ID = rule.ID
		policy.Subject = rule.V0
		policy.Path = rule.V1
		policy.Method = rule.V2
		policy.Result = rule.V3
		policy.Comment = rule.Comment
		if name, ok := mapApiName[fmt.Sprintf("%s-%s", policy.Method, policy.Path)]; ok {
			policy.PathName = name
		}
		if list, exists := result[policy.Subject]; exists {
			list = append(list, policy)
			result[policy.Subject] = list
		} else {
			result[policy.Subject] = []*model.Policy{policy}
		}
	}

	return result, nil
}
