package api

import (
	"fmt"
	"net/http"
	"order-flow/internal/app"
	"order-flow/internal/app/middleware"
	"order-flow/internal/config"
)

func HandleRequest(a *app.App) error {
	AuthMiddleware := middleware.AuthMiddleware(a.Pool)

	mux := http.NewServeMux()

	userHandler := a.BuildUserCreate()
	userHandler.Register(mux, AuthMiddleware)

	mux.HandleFunc("GET /health-check", HealthCheck)

	return http.ListenAndServe(
		fmt.Sprintf(":%d", config.Env.Port),
		mux,
	)
}
