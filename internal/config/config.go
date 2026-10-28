package config

import (
	"sync"

	"github.com/caarlos0/env/v11"
)

type Config struct {
	AppEnv     string `env:"APP_ENV" envDefault:"dev"`
	Port       int    `env:"PORT" envDefault:"8000"`
	DbHost     string `env:"DB_HOST,required"`
	DbPort     int    `env:"DB_PORT" envDefault:"5432"`
	DbUser     string `env:"DB_USER,required"`
	DbPassword string `env:"DB_PASSWORD,required"`
	DbName     string `env:"DB_NAME,required"`
}

var (
	Env  *Config
	once sync.Once
	err  error
)

// Once.Do Гаранитирует что выполнится единожды
func Load() (*Config, error) {
	once.Do(func() {
		var c Config
		if err = env.Parse(&c); err == nil {
			Env = &c
		}
	})

	return Env, err
}
