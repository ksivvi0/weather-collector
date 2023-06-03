package config

import "github.com/ilyakaznacheev/cleanenv"

type Config struct {
}

func InitConfig() (*Config, error) {
	cfg := &Config{}
	err := cleanenv.ReadEnv(cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}
