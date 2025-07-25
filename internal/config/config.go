package config

import "github.com/ilyakaznacheev/cleanenv"

const (
	cfgPath = "config/config.yml"
)

type Config struct {
	RestServer RestServer `yaml:"rest_server"`
}

type RestServer struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
}

func InitConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(cfgPath, &cfg); err != nil {
		return nil, err
	}

	// TODO: init env file

	return &cfg, nil
}
