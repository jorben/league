package config

import (
	"errors"
	conf "github.com/ldigit/config"
	"league/log"
)

type LeagueConfig struct {
	Db struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"db"`
	Log []log.OutputConfig `yaml:"log"`
}

// GetConfig 获取解析好的配置
func GetConfig(path string) (*LeagueConfig, error) {
	raw := conf.GetGlobalConfig()
	if raw != nil {
		return raw.(*LeagueConfig), nil
	}
	cfg := loadConfig(path)
	if cfg == nil {
		return nil, errors.New("configuration is empty, please check the config file path")
	}
	return cfg, nil
}

func loadConfig(path string) *LeagueConfig {
	cfg := &LeagueConfig{}
	if err := conf.LoadAndDecode(path, cfg); err != nil {
		return nil
	}
	conf.SetGlobalConfig(cfg)
	return cfg
}
