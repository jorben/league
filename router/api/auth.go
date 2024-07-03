package api

import (
	"github.com/gin-gonic/gin"
	"league/common/context"
	"league/common/errs"
	"league/config"
	"league/log"
	"league/model"
	"league/provider/auth"
	"league/service"
	"net/http"
)

// AuthProvider 第三方登录渠道接口
type AuthProvider interface {
	GetLoginUrl(ctx *gin.Context) (string, error)
	GetUserinfo(ctx *gin.Context) (*model.UserSocialInfo, error)
}

// AuthLogin 登录接口
func AuthLogin(ctx *gin.Context) {
	c := &context.CustomContext{Context: ctx}
	var oAuthClient AuthProvider
	// 获取登录方式
	provider := ctx.Query("type")
	providerConfig := config.OAuthProvider{}

	if len(provider) > 0 {
		providerConfig = config.GetAuthProviderConfig(provider)
	}

	// 创建provider实例
	switch provider {
	case auth.ProviderGithub:
		oAuthClient = auth.NewGithubOAuth(providerConfig)
	case auth.ProviderGoogle:
		oAuthClient = auth.NewGoogleOAuth(providerConfig)
	default:
		c.CJSON(errs.ErrAuthUnknownProvider)
		return
	}

	// 执行provider登录构造
	url, err := oAuthClient.GetLoginUrl(ctx)
	if err != nil {
		log.Errorf(ctx, "Get %s login url failed, err: %s", provider, err.Error())
		c.CJSON(errs.ErrAuthLoginUrl)
		return
	}
	log.Debugf(ctx, "Login url: %s", url)
	ctx.Redirect(http.StatusTemporaryRedirect, url)
}

// AuthCallback 登录回调接口
func AuthCallback(ctx *gin.Context) {
	c := context.CustomContext{Context: ctx}
	var oAuthClient AuthProvider
	// 获取回调来源
	provider := ctx.Query("type")
	providerConfig := config.OAuthProvider{}

	if len(provider) > 0 {
		providerConfig = config.GetAuthProviderConfig(provider)
	}

	// 创建provider实例
	switch provider {
	case auth.ProviderGithub:
		oAuthClient = auth.NewGithubOAuth(providerConfig)
	case auth.ProviderGoogle:
		oAuthClient = auth.NewGoogleOAuth(providerConfig)
	default:
		c.CJSON(errs.ErrAuthUnknownProvider)
		return
	}

	// 从provider获取用户信息
	usi, err := oAuthClient.GetUserinfo(ctx)
	if err != nil {
		c.CJSON(errs.ErrAuthUserinfo, err.Error())
		return
	}

	authService := service.NewAuthService(ctx)
	// 注册或更新用户信息并获取jwt token
	token, err := authService.LoginBySource(usi)
	if err != nil {
		c.CJSON(errs.ErrAuthLoginFailed)
		return
	}

	c.CJSON(errs.Success, gin.H{"token": token})
}

// AuthRenew 续期JWT
func AuthRenew(ctx *gin.Context) {

}
