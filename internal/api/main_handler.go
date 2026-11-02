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
	mux := http.NewServeMux()
	mux.HandleFunc("GET /swagger/", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	a.CompositeRoot()
	AuthMiddleware := middleware.AuthMiddleware(a.UserRepo)
	PolicyMiddleware := middleware.PolicyMiddleware()
	a.UserHandler.Register(mux, AuthMiddleware)
	a.ProductHandler.Register(mux, AuthMiddleware, PolicyMiddleware)

	mux.HandleFunc("GET /health-check", HealthCheck)

	return http.ListenAndServe(
		fmt.Sprintf(":%d", config.Env.Port),
		mux,
	)
}
