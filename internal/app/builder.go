package app

import (
	"context"
	"log"
	"order-flow/internal/api/user_v1"
	"order-flow/internal/modules/user/repository"
	"order-flow/internal/modules/user/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Pool   *pgxpool.Pool
	Ctx    context.Context
	Logger *log.Logger
}

func (a *App) BuildUserCreate() *user_v1.UserHandler {
	userRepo := repository.NewUserRepository(a.Pool)
	userService := service.NewUserService(userRepo)
	userHandler := user_v1.NewUserHandler(userService)
	return userHandler
}
