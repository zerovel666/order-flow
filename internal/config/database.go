package config

import (
	"context"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func DSN() (string, error) {
	if Env == nil {
		return "", errors.New("config: Load() was not called or failed")
	}

	u := url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(Env.DbUser, Env.DbPassword),
		Host:     fmt.Sprintf("%s:%d", Env.DbHost, Env.DbPort),
		Path:     Env.DbName,
		RawQuery: "sslmode=disable",
	}

	return u.String(), nil
}

func OpenPull(ctx context.Context) (*pgxpool.Pool, error) {
	dsn, err := DSN()
	if err != nil {
		return nil, err
	}

	cfg, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, err
	}

	cfg.MaxConns = 100
	cfg.MinConns = 2
	cfg.MaxConnLifetime = time.Hour
	cfg.MaxConnIdleTime = 30 * time.Minute
	cfg.HealthCheckPeriod = time.Minute

	pool, err := pgxpool.NewWithConfig(ctx, cfg)
	if err != nil {
		return nil, err
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, err
	}

	return pool, nil
}
