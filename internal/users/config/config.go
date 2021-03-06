package config

import "github.com/caarlos0/env/v6"

type Config struct {
	AppHost      string `env:"APP_HOST" envDefault:"127.0.0.1"`
	AppPort      int    `env:"APP_PORT" envDefault:"8801"`
	ConsulHost   string `env:"CONSUL_HOST" envDefault:"127.0.0.1"`
	ConsulPort   int    `env:"CONSUL_PORT" envDefault:"8500"`
	DatabaseHost string `env:"DATABASE_HOST" envDefault:"127.0.0.1"`
	DatabasePort int    `env:"DATABASE_PORT" envDefault:"27017"`
	DatabaseUser string `env:"DATABASE_USER" envDefault:"root"`
	DatabasePass string `env:"DATABASE_PASS" envDefault:"password"`
	DatabaseName string `env:"DATABASE_NAME" envDefault:"users_db"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
