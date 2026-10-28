package app

import (
	"context"
	"log"
	"order-flow/internal/api/user_v1"
	"order-flow/internal/modules/users/repositories"
	"order-flow/internal/modules/users/services"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Pool   *pgxpool.Pool
	Ctx    context.Context
	Logger *log.Logger
}

func (a *App) BuildUserCreate() *user_v1.UserHandler {
	userRepo := repositories.NewUserRepository(a.Pool)
	userService := services.NewUserService(userRepo)
	userHandler := user_v1.NewUserHandler(userService)
	return userHandler
}
