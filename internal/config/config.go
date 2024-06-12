package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	Database DBConfig
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

type AppConfig struct {
	Name string
}

func LoadConfig(env string) (Config, error) {
	var cfg Config

	if env == "" {
		viper.SetConfigName("config" + env)
	} else {
		viper.SetConfigName("config." + env)
	}

	viper.AddConfigPath("./config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return cfg, err
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		return cfg, err
	}

	return cfg, nil
}
