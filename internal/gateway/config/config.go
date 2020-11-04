package config

import "github.com/caarlos0/env/v6"

const DevMode = "dev"
const TestMode = "test"

type Config struct {
	AppHost    string `env:"APP_HOST" envDefault:"127.0.0.1"`
	AppPort    string `env:"APP_PORT" envDefault:"8080"`
	AppMode    string `env:"APP_MODE" envDefault:"prod"`
	ConsulHost string `env:"CONSUL_HOST" envDefault:"127.0.0.1"`
	ConsulPort string `env:"CONSUL_PORT" envDefault:"8500"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}

func (cfg *Config) IsProdMode() bool {
	return cfg.AppMode != DevMode && cfg.AppMode != TestMode
}

func (cfg *Config) IsDevMode() bool {
	return cfg.AppMode == DevMode
}

func (cfg *Config) IsTestMode() bool {
	return cfg.AppMode == TestMode
}
