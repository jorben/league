package auth

import (
	"bytes"
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
const ProviderGithub = "github"

type GithubData struct {
	Username string `json:"login"`
	Id       uint   `json:"id"`
	Avatar   string `json:"avatar_url"`
	Nickname string `json:"name"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
}

type GithubOAuth struct {
	cfg config.OAuthProvider
}

// NewGithubOAuth 新建Github oAuth实例
func NewGithubOAuth(c config.OAuthProvider) *GithubOAuth {
	return &GithubOAuth{cfg: c}
}

func (g *GithubOAuth) GetLoginUrl(ctx *gin.Context) (string, error) {
	// Create the dynamic redirect URL for login
	return fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s",
		g.cfg.ClientId,
		url.QueryEscape(g.cfg.CallbackUri),
	), nil
}

func (g *GithubOAuth) GetUserinfo(ctx *gin.Context) (*model.UserSocialInfo, error) {
	code := ctx.Query("code")
	if len(code) == 0 {
		return nil, errors.New("缺少必要参数:code")
	}
	token, err := getGithubAccessToken(ctx, g.cfg.ClientId, g.cfg.ClientSecret, code)
	if err != nil {
		return nil, errors.New("无法获取access token，请确认code是否正确")
	}

	githubData, err := getGithubData(ctx, token)
	if err != nil {
		return nil, errors.New("获取用户信息失败，请稍后重试")
	}

	log.Debugf(ctx, "Github data: %s", githubData)
	data := &GithubData{}
	err = json.Unmarshal([]byte(githubData), data)
	if err != nil {
		log.Errorf(ctx, "Unmarshal github data failed, data: %s, err: %s", githubData, err.Error())
		return nil, errors.New("解析用户数据失败")
	}

	// 缺少必要数据
	if data.Id == 0 {
		log.Errorf(ctx, "Get github userinfo failed, data: %s", githubData)
		return nil, errors.New("获取用户数据失败，请确认access token是否有效")
	}

	return &model.UserSocialInfo{
		Source:   ProviderGithub,
		OpenId:   fmt.Sprintf("%d", data.Id),
		Email:    data.Email,
		Avatar:   data.Avatar,
		Username: data.Username,
		Nickname: data.Nickname,
		Bio:      data.Bio,
	}, nil

}

func getGithubData(ctx *gin.Context, accessToken string) (string, error) {
	// Get request to a set URL
	req, err := http.NewRequest(
		"GET",
		"https://api.github.com/user",
		nil,
	)
	if err != nil {
		log.Errorf(ctx, "API Request creation failed, err: %s", err.Error())
		return "", err
	}

	// Set the Authorization header before sending the request
	// Authorization: token XXXXXXXXXXXXXXXXXXXXXXXXXXX
	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

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

func getGithubAccessToken(ctx *gin.Context, clientId string, clientSecret string, code string) (string, error) {

	// Set us the request body as JSON
	requestBodyMap := map[string]string{
		"client_id":     clientId,
		"client_secret": clientSecret,
		"code":          code,
	}
	requestJSON, _ := json.Marshal(requestBodyMap)

	// POST request to set URL
	req, err := http.NewRequest(
		"POST",
		"https://github.com/login/oauth/access_token",
		bytes.NewBuffer(requestJSON),
	)
	if err != nil {
		log.Errorf(ctx, "Request creation failed, err: %s", err.Error())
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")

	// Get the response
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf(ctx, "Request failed, err: %s", err.Error())
		return "", err
	}

	// Response body converted to stringified JSON
	respbody, _ := io.ReadAll(resp.Body)

	// Represents the response received from Github
	type githubAccessTokenResponse struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}

	// Convert stringified JSON to a struct object of type githubAccessTokenResponse
	var ghresp githubAccessTokenResponse
	err = json.Unmarshal(respbody, &ghresp)
	if err != nil {
		log.Errorf(ctx, "Unmarshal github response failed, err: %s", err.Error())
		return "", err
	}

	// Return the access token (as the rest of the
	// details are relatively unnecessary for us)
	return ghresp.AccessToken, nil
}
