package api

import (
	"fmt"
	"net/http"
	"order-flow/internal/app"
	"order-flow/internal/config"
)

func HandleRequest(a *app.App) error {
	mux := http.NewServeMux()

	userHandler := a.BuildUserCreate()
	userHandler.Register(mux)

	mux.HandleFunc("GET /health-check", HealthCheck)

	return http.ListenAndServe(
		fmt.Sprintf(":%d", config.Env.Port),
		mux,
	)
}
