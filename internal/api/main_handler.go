package api

import (
	"net/http"
	"order-flow/internal/api/user_v1"
)

func HandleRequest() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /health-check", HealthCheck)

	user_v1.HandleRequest(mux)
	http.ListenAndServe(":8000", mux)
}
