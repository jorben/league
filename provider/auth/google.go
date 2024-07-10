package auth

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"io"
	"league/config"
	"league/log"
	"league/model"
	"net/http"
)

// ProviderGoogle 渠道名称
const ProviderGoogle = "google"

type GoogleData struct {
	Id       string `json:"id"`
	Avatar   string `json:"picture"`
	Nickname string `json:"name"`
	Email    string `json:"email"`
}

type GoogleOAuth struct {
	cfg                config.OAuthProvider
	googleOAuth2Config *oauth2.Config
}

// NewGoogleOAuth 新建Github oAuth实例
func NewGoogleOAuth(c config.OAuthProvider) *GoogleOAuth {
	return &GoogleOAuth{
		cfg: c,
		googleOAuth2Config: &oauth2.Config{
			RedirectURL:  c.CallbackUri,
			ClientID:     c.ClientId,
			ClientSecret: c.ClientSecret,
			Scopes: []string{
				"https://www.googleapis.com/auth/userinfo.email",
				"https://www.googleapis.com/auth/userinfo.profile",
			},
			Endpoint: google.Endpoint,
		},
	}
}

func (g *GoogleOAuth) GetLoginUrl(ctx *gin.Context, redirect string) (string, error) {
	// 暂未找到 google auth 携带二次回调的方法，注入callback中会无法匹配
	return g.googleOAuth2Config.AuthCodeURL(g.cfg.State), nil
}

func (g *GoogleOAuth) GetUserinfo(ctx *gin.Context) (*model.UserSocialInfo, error) {
	// 校验state
	state := ctx.Query("state")
	if len(state) == 0 || state != g.cfg.State {
		return nil, errors.New("state参数不正确")
	}
	code := ctx.Query("code")
	if len(code) == 0 {
		return nil, errors.New("缺少必要参数:code")
	}
	// 获取access token
	token, err := g.googleOAuth2Config.Exchange(ctx, code)
	if err != nil {
		log.Errorf(ctx, "Code exchange failed, err: %s", err.Error())
		return nil, err
	}

	// 获取用户信息
	googleData, err := getGoogleData(ctx, token.AccessToken)
	if err != nil {
		return nil, errors.New("获取用户信息失败，请稍后重试")
	}

	log.Debugf(ctx, "Google data: %s", googleData)
	data := &GoogleData{}
	err = json.Unmarshal([]byte(googleData), data)
	if err != nil {
		log.Errorf(ctx, "Unmarshal google data failed, data: %s, err: %s", googleData, err.Error())
		return nil, errors.New("解析用户数据失败")
	}

	// 缺少必要数据
	if len(data.Id) == 0 {
		log.Errorf(ctx, "Get google userinfo failed, data: %s", googleData)
		return nil, errors.New("获取用户数据失败，请确认access token是否有效")
	}

	return &model.UserSocialInfo{
		Source:   ProviderGoogle,
		OpenId:   data.Id,
		Email:    data.Email,
		Avatar:   data.Avatar,
		Username: data.Email,
		Nickname: data.Nickname,
		Bio:      "",
	}, nil
}

func getGoogleData(ctx *gin.Context, accessToken string) (string, error) {
	// Get request to a set URL
	req, err := http.NewRequest(
		"GET",
		"https://www.googleapis.com/oauth2/v2/userinfo?access_token="+accessToken,
		nil,
	)
	if err != nil {
		log.Errorf(ctx, "API Request creation failed, err: %s", err.Error())
		return "", err
	}

	// Make the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf(ctx, "Request failed, err: %s", err.Error())
		return "", err
	}

	// Read the response as a byte slice
	respbody, _ := io.ReadAll(resp.Body)

	// Convert byte slice to string and return
	return string(respbody), nil
}
