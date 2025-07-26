package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	"github.com/joho/godotenv"
)

const (
	cfgPath = "config/config.yml"
)

type Config struct {
	RestServer  RestServer `yaml:"rest_server"`
	Integration Integration
}

type RestServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

type Integration struct {
	ClientID    string `env:"CLIENT_ID"`
	SecrestKey  string `env:"SECRET_KEY"`
	RedirectURL string `env:"REDIRECT_URL"`
	AuthCode    string `env:"AUTH_CODE"`
}

func InitConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		return nil, err
	}

	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	if err := cleanenv.ReadEnv(&cfg.Integration); err != nil {
		return nil, err
	}

	return &cfg, nil
}
