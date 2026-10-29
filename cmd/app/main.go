package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"order-flow/internal/api"
	"order-flow/internal/app"
	"order-flow/internal/config"
	"order-flow/migrations"

	"github.com/jackc/pgx/v5/stdlib"
)

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}

func run() error {
	fmt.Println("Order-flow service is starting")

	if _, err := config.Load(); err != nil {
		return fmt.Errorf("load config: %w", err)
	}

	// Загрузка фонового процесса
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Открытие пула коннекшенов
	pool, err := config.OpenPull(ctx)
	if err != nil {
		return fmt.Errorf("open db pool: %w", err)
	}
	defer pool.Close()

	// Запуск миграций
	sqlDB := stdlib.OpenDBFromPool(pool)
	defer sqlDB.Close()

	if err := migrations.RunMigrations(ctx, sqlDB); err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}

	application := app.App{
		Pool:   pool,
		Ctx:    ctx,
		Logger: log.Default(),
	}

	// Подключение апишек и регистрация зависимостей
	if err := api.HandleRequest(&application); err != nil {
		return fmt.Errorf("run migrations: %w", err)
	}

	return nil
}
