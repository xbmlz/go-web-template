package config

import (
	"github.com/spf13/viper"
	"github.com/xbmlz/go-web-template/internal/database"
	"github.com/xbmlz/go-web-template/internal/logger"
	"github.com/xbmlz/go-web-template/internal/server"
	"github.com/xbmlz/go-web-template/internal/token"
)

type Config struct {
	Server   server.Config   `json:"server"`
	Database database.Config `json:"database"`
	Token    token.Config    `json:"token"`
	Log      logger.Config   `json:"log"`
}

func Init(configFile string) (*Config, error) {
	v := viper.New()
	v.SetConfigFile(configFile)
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config
	if err := v.Unmarshal(&c); err != nil {
		return nil, err
	}
	return &c, nil
}

func MustInit(configFile string) *Config {
	c, err := Init(configFile)
	if err != nil {
		panic(err)
	}
	return c
}
