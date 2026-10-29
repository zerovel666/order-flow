package config

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"sync"
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
	Env    *Config
	once   sync.Once
	err_do error
)

// Once.Do Гаранитирует что выполнится единожды
func Load() (*Config, error) {
	once.Do(func() {
		var c Config
		scan(&c)
		Env = &c
	})

	return Env, err_do
}

func scan(c *Config) {
	file, err := os.Open(".env")
	if err != nil {
		err_do = err
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		str := strings.TrimSpace(scanner.Text())
		if str == "" || strings.HasPrefix(str, "#") {
			continue
		}

		key, value, found := strings.Cut(str, "=")
		if !found {
			continue
		}

		if err := bind(strings.TrimSpace(key), strings.TrimSpace(value), c); err != nil {
			err_do = err
			break
		}
	}

	if err := scanner.Err(); err != nil {
		err_do = err
	}
}

func bind(key, value string, c *Config) error {
	switch key {
	case "APP_ENV":
		c.AppEnv = value

	case "PORT":
		port, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		c.Port = port

	case "DB_HOST":
		c.DbHost = value

	case "DB_PORT":
		dbPort, err := strconv.Atoi(value)
		if err != nil {
			return err
		}
		c.DbPort = dbPort

	case "DB_USER":
		c.DbUser = value

	case "DB_PASSWORD":
		c.DbPassword = value

	case "DB_NAME":
		c.DbName = value
	}

	return nil
}
