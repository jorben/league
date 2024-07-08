package auth

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"league/config"
	"league/log"
	"league/model"
	"net/http"
	"net/url"
)

// ProviderGithub 渠道名称
const ProviderWechat = "wechat"

type WechatOAuth struct {
	cfg config.OAuthProvider
}

type WechatUserInfo struct {
	OpenId    string   `json:"openid"`
	Nickname  string   `json:"nickname"`
	Sex       int      `json:"sex"`
	Province  string   `json:"province"`
	City      string   `json:"city"`
	Country   string   `json:"country"`
	Avatar    string   `json:"headimgurl"`
	Privilege []string `json:"privilege"`
	UnionId   string   `json:"unionid"`
}

// NewWechatOAuth 新建Github oAuth实例
func NewWechatOAuth(c config.OAuthProvider) *WechatOAuth {
	return &WechatOAuth{cfg: c}
}

// GetLoginUrl 获取登录URL
func (g *WechatOAuth) GetLoginUrl(ctx *gin.Context) (string, error) {

	return fmt.Sprintf(
		"https://open.weixin.qq.com/connect/qrconnect?appid=%s&scope=snsapi_login&redirect_uri=%s&state=%s",
		g.cfg.ClientId,
		url.QueryEscape(g.cfg.CallbackUri),
		g.cfg.State,
	), nil
}

// GetUserinfo 获取用户信息
func (g *WechatOAuth) GetUserinfo(ctx *gin.Context) (*model.UserSocialInfo, error) {
	state := ctx.Query("state")
	if len(state) == 0 || state != g.cfg.State {
		return nil, errors.New("state参数不正确")
	}
	code := ctx.Query("code")
	if len(code) == 0 {
		return nil, errors.New("缺少必要参数:code")
	}
	token, openId, err := getWechatAccessToken(ctx, g.cfg.ClientId, g.cfg.ClientSecret, code)
	if err != nil {
		return nil, errors.New("无法获取access token，请确认code是否正确")
	}

	wechatData, err := getWechatData(ctx, token, openId)
	if err != nil {
		return nil, errors.New("获取用户信息失败，请稍后重试")
	}

	log.Debugf(ctx, "Wechat data: %s", wechatData)
	data := &WechatUserInfo{}
	err = json.Unmarshal([]byte(wechatData), data)
	if err != nil {
		log.Errorf(ctx, "Unmarshal wechat data failed, data: %s, err: %s", wechatData, err.Error())
		return nil, errors.New("解析用户数据失败")
	}

	return &model.UserSocialInfo{
		Source:   ProviderWechat,
		OpenId:   data.OpenId,
		Avatar:   data.Avatar,
		Username: data.OpenId,
		Nickname: data.Nickname,
		Gender:   uint8(data.Sex),
	}, nil
}

func getWechatData(ctx *gin.Context, accessToken string, openId string) (string, error) {
	// Get request to a set URL
	resp, err := http.Get(
		fmt.Sprintf("https://api.weixin.qq.com/sns/userinfo?access_token=%s&openid=%s",
			accessToken, openId),
	)
	if err != nil {
		log.Errorf(ctx, "Request failed, err: %s", err.Error())
		return "", err
	}

	// Read the response as a byte slice
	respbody, _ := io.ReadAll(resp.Body)

	// Convert byte slice to string and return
	return string(respbody), nil
}

func getWechatAccessToken(ctx *gin.Context, clientId string, clientSecret string, code string) (string, string, error) {

	resp, err := http.Get(
		fmt.Sprintf("https://api.weixin.qq.com/sns/oauth2/access_token?appid=%s&secret=%s&code=%s&grant_type=authorization_code",
			clientId, clientSecret, code),
	)
	if err != nil {
		log.Errorf(ctx, "Request failed, err: %s", err.Error())
		return "", "", err
	}

	// Response body converted to stringified JSON
	respbody, _ := io.ReadAll(resp.Body)

	// Represents the response received from Github
	//{
	//	"access_token":"ACCESS_TOKEN",
	//	"expires_in":7200,
	//	"refresh_token":"REFRESH_TOKEN",
	//	"openid":"OPENID",
	//	"scope":"SCOPE",
	//	"unionid": "o6_bmasdasdsad6_2sgVt7hMZOPfL"
	//}
	type wechatAccessTokenResponse struct {
		AccessToken  string `json:"access_token"`
		ExpiresIn    int    `json:"expires_in"`
		RefreshToken string `json:"refresh_token"`
		OpenId       string `json:"openid"`
		Scope        string `json:"scope"`
		UnionId      string `json:"unionid"`
		ErrCode      int    `json:"errcode"`
		ErrMsg       string `json:"errmsg"`
	}

	// Convert stringified JSON to a struct object of type githubAccessTokenResponse
	var wxresp wechatAccessTokenResponse
	err = json.Unmarshal(respbody, &wxresp)
	if err != nil {
		log.Errorf(ctx, "Unmarshal wechat response failed, err: %s", err.Error())
		return "", "", err
	}

	// 检查必要数据
	if len(wxresp.AccessToken) == 0 || len(wxresp.OpenId) == 0 {
		log.Errorf(ctx, "Get access token failed, errcode: %d, errmsg: %s", wxresp.ErrCode, wxresp.ErrMsg)
		return "", "", errors.New("get wechat access token failed")
	}

	// Return the access token (as the rest of the
	// details are relatively unnecessary for us)
	return wxresp.AccessToken, wxresp.OpenId, nil
}
