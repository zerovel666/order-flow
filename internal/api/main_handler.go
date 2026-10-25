package api

import "net/http"

func HandleRequest() {
	mux := http.NewServeMux()

	mux.HandleFunc("/health-check", HealthCheck)
	http.ListenAndServe(":8000", mux)
}