package auth

import (
	"github.com/gin-gonic/gin"
	"league/config"
	"league/model"
)

// ProviderGoogle 渠道名称
const ProviderGoogle = "google"

type GoogleOAuth struct {
	cfg config.OAuthProvider
}

// NewGoogleOAuth 新建Github oAuth实例
func NewGoogleOAuth(c config.OAuthProvider) *GoogleOAuth {
	return &GoogleOAuth{cfg: c}
}

func (g *GoogleOAuth) GetLoginUrl(ctx *gin.Context) (string, error) {
	// Create the dynamic redirect URL for login
	return "", nil
}

func (g *GoogleOAuth) GetUserinfo(ctx *gin.Context) (*model.UserSocialInfo, error) {
	return nil, nil
}
