package api

import (
	"net/http"
	"order-flow/internal/config/response"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	messages := []response.Message{
		{Lang: "ru", Message: "Сервер работает"},
		{Lang: "en", Message: "Server is working"},
	}

	response.WriteJSON(w, http.StatusOK,
		response.SuccessAPIResponse("SERVER_LIFE", messages, nil),
	)
}
