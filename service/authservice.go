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
func (a *AuthService) LoginBySource(info model.UserSocialInfo) (string, error) {
	// 检查是否已存在该用户
	user := a.UserDal.GetUserBySource(info.Source, info.OpenId)
	if user == nil {
		// 不存在则插入用户
		id, err := a.UserDal.SignUpBySource(info)
		if err != nil {
			log.Errorf(a.Ctx, "UserDal.SignUpBySource failed, err: %s", err.Error())
			return "", err
		}
		return a.SignJwtString(id)
	} else {
		// 存在则更新信息
		_, err := a.UserDal.UpdateBySource(info)
		if err != nil {
			log.Errorf(a.Ctx, "UserDal.UpdateBySource failed, err: %s", err.Error())
			return "", nil
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
func (a *AuthService) SignJwtString(id uint) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Issuer:    "league",
		ExpiresAt: jwt.NewNumericDate(now.Add(2 * time.Hour)),
		NotBefore: jwt.NewNumericDate(now.Add(-1 * time.Minute)),
		IssuedAt:  jwt.NewNumericDate(now),
		ID:        fmt.Sprintf("%d", id),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS384, claims)
	sign, err := token.SignedString(a.signKey)
	if err != nil {
		log.Errorf(a.Ctx, "Jwt SignedString failed, err: %s", err.Error())
	}
	return sign, err
}

// VerifyJwtString 校验JWT
func (a *AuthService) VerifyJwtString(s string) (string, error) {

	token, err := jwt.ParseWithClaims(s, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return a.signKey, nil
	})
	if err != nil {
		log.Errorf(a.Ctx, "Jwt parse failed, err: %s", err.Error())
		return "", err
	} else if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok {
		log.Debugf(a.Ctx, "Check login passed, UserId: %s", claims.ID)
		return claims.ID, nil
	} else {
		log.Errorf(a.Ctx, "Unknown claims type, token: %s", s)
		return "", errors.New("unknown claims type")
	}
}

// TODO: jwt失效逻辑，token续签
