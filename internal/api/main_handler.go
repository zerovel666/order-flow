package api

import (
	"fmt"
	"net/http"
	_ "order-flow/docs"
	"order-flow/internal/app"
	"order-flow/internal/app/middleware"
	"order-flow/internal/config"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func HandleRequest(a *app.App) error {
	AuthMiddleware := middleware.AuthMiddleware(a.Pool)

	mux := http.NewServeMux()
	mux.HandleFunc("GET /swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	userHandler := a.BuildUserCreate()
	userHandler.Register(mux, AuthMiddleware)

	mux.HandleFunc("GET /health-check", HealthCheck)

	return http.ListenAndServe(
		fmt.Sprintf(":%d", config.Env.Port),
		mux,
	)
}
