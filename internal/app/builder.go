package app

import (
	"context"
	"log"
	"order-flow/internal/api/product_v1"
	"order-flow/internal/api/user_v1"
	productRepo "order-flow/internal/modules/product/repository"
	productService "order-flow/internal/modules/product/service"
	userRepo "order-flow/internal/modules/user/repository"
	userService "order-flow/internal/modules/user/service"

	"github.com/jackc/pgx/v5/pgxpool"
)

type App struct {
	Pool   *pgxpool.Pool
	Ctx    context.Context
	Logger *log.Logger

	UserRepo    *userRepo.UserRepository
	UserService *userService.UserService
	UserHandler *user_v1.UserHandler

	ProductRepo    *productRepo.ProductRepository
	ProductService *productService.ProductService
	ProductHandler *product_v1.ProductHandler
}

func (a *App) CompositeRoot() {
	a.buildUserHandle()
}

func (a *App) buildUserHandle() *App {
	userRepo := userRepo.NewUserRepository(a.Pool)
	userService := userService.NewUserService(userRepo)
	userHandler := user_v1.NewUserHandler(userService)

	a.UserRepo = userRepo
	a.UserService = userService
	a.UserHandler = userHandler

	return a
}

func (a *App) buildProductHandle() *App {
	productRepo := productRepo.NewProductRepository(a.Pool)
	productService := productService.NewProductService(productRepo)
	productHandler := product_v1.NewProductHandler(productService)

	a.ProductRepo = productRepo
	a.ProductService = productService
	a.ProductHandler = productHandler

	return a
}
