package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"league/config"
	"league/log"
	"league/model"
	"league/service"
	"net/http"
	"net/url"
	"sync"
)

var once sync.Once
var oAuthGithubConfig config.OAuthConfig

type GithubData struct {
	Username string `json:"login"`
	Id       uint   `json:"id"`
	Avatar   string `json:"avatar_url"`
	Nickname string `json:"name"`
	Email    string `json:"email"`
	Bio      string `json:"bio"`
	Status   string `json:"status"`
	Message  string `json:"message"`
}

func getConfig() *config.OAuthConfig {
	once.Do(func() {
		cfg := config.GetConfig()
		for _, authConfig := range cfg.Auth {
			if authConfig.Source == "github" {
				oAuthGithubConfig = authConfig
			}
		}

	})
	return &oAuthGithubConfig
}

func LoginGithub(ctx *gin.Context) {
	// Create the dynamic redirect URL for login
	redirectURL := fmt.Sprintf(
		"https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s",
		getConfig().ClientId,
		url.QueryEscape(getConfig().CallbackUri),
	)
	ctx.Redirect(http.StatusFound, redirectURL)
}

func CallbackGithub(ctx *gin.Context) {
	code := ctx.Query("code")
	if len(code) == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"ret": -1, "msg": "code is empty"})
		return
	}
	token, err := getGithubAccessToken(ctx, code)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"ret": -1, "msg": "Get github access token failed"})
		return
	}

	githubData, err := getGithubData(ctx, token)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"ret": -1, "msg": "Get github data failed"})
		return
	}

	log.Debug(ctx, "Github data: %s", githubData)
	data := &GithubData{}
	err = json.Unmarshal([]byte(githubData), data)
	if err != nil {
		log.Errorf(ctx, "Unmarshal github data failed, data: %s, err: %s", githubData, err.Error())
		ctx.JSON(http.StatusBadGateway, gin.H{"ret": -1, "msg": "Unmarshal github data failed"})
		return
	}

	// 缺少必要数据
	if data.Id == 0 {
		log.Errorf(ctx, "Get github userinfo failed, data: %s", githubData)
		ctx.JSON(http.StatusUnauthorized, gin.H{"ret": -1, "msg": fmt.Sprintf("%s:%s", data.Status, data.Message)})
		return
	}

	authService := service.NewAuthService(ctx)
	id, err := authService.LoginBySource(model.UserSocialInfo{
		Source:   "Github",
		OpenId:   fmt.Sprintf("%d", data.Id),
		Email:    data.Email,
		Avatar:   data.Avatar,
		Username: data.Username,
		Nickname: data.Nickname,
		Bio:      data.Bio,
	})
	if err != nil {
		ctx.JSON(http.StatusOK, gin.H{"ret": -1, "msg": "第三方渠道登录/注册失败"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"ret": 0, "msg": "ok", "data": id})
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
	}

	// Set the Authorization header before sending the request
	// Authorization: token XXXXXXXXXXXXXXXXXXXXXXXXXXX
	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	req.Header.Set("Authorization", authorizationHeaderValue)

	// Make the request
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Errorf(ctx, "Request failed, err: %s", err.Error())
	}

	// Read the response as a byte slice
	respbody, _ := io.ReadAll(resp.Body)

	// Convert byte slice to string and return
	return string(respbody), nil
}

func getGithubAccessToken(ctx *gin.Context, code string) (string, error) {

	clientID := getConfig().ClientId
	clientSecret := getConfig().ClientSecret

	// Set us the request body as JSON
	requestBodyMap := map[string]string{
		"client_id":     clientID,
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
