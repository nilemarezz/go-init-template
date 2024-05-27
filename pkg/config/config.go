package config

import (
	"fmt"

	"github.com/spf13/viper"
)

type Config struct {
	Database DBConfig
	Log      LogConfig
	App      AppConfig
}

type DBConfig struct {
	Host     string
	Port     int
	User     string
	DBName   string
	Password string
	SSLMode  string
}

type LogConfig struct {
	Path string
}

type AppConfig struct {
	Port string
}

func LoadConfig(env string) (Config, error) {
	var cfg Config

	viper.SetConfigName("config." + env)
	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	fmt.Println(cfg)

	return cfg, nil
}
