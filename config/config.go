package config

import (
	"errors"
	conf "github.com/ldigit/config"
	"league/log"
)

type OAuthConfig struct {
	Source       string `yaml:"source"`
	ClientId     string `yaml:"client_id"`
	ClientSecret string `yaml:"client_secret"`
	CallbackUri  string `yaml:"callback_uri"`
}

type DbConfig struct {
	Dsn             string `yaml:"dsn"`
	MaxOpenConns    int    `yaml:"max_open_conns"`
	MaxIdleConns    int    `yaml:"max_idle_conns"`
	ConnMaxLifetime int    `yaml:"conn_max_lifetime"`
	CheckInterval   int    `yaml:"check_interval"`
}

type LeagueConfig struct {
	Db   DbConfig           `yaml:"db"`
	Auth []OAuthConfig      `yaml:"auth"`
	Log  []log.OutputConfig `yaml:"log"`
	Jwt  struct {
		SignKey string `yaml:"sign_key"`
	} `yaml:"jwt"`
}

// CasbinDefaultPolicies 默认策略，当策略表为空时会初始化以下策略
var CasbinDefaultPolicies = [][]string{
	{"anyone", "/health", "GET", "allow"},
	{"anyone", "/auth/login*", "GET", "allow"},
	{"anyone", "/auth/callback*", "GET", "allow"},
	{"member", "/auth/renew*", "GET", "allow"},
}

var leagueConfig *LeagueConfig

// GetConfig 获取配置实例
func GetConfig() *LeagueConfig {
	return leagueConfig
}

// LoadConfig 解析配置
func LoadConfig(path string) (*LeagueConfig, error) {
	leagueConfig = &LeagueConfig{}
	if err := conf.LoadAndDecode(path, leagueConfig); err != nil {

		return nil, errors.New("configuration is empty, please check the config file path")
	}
	//conf.SetGlobalConfig(leagueConfig)
	return leagueConfig, nil
}
