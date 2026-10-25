package api

import (
	"net/http"

	"order-flow/internal/utils/responses"
)

func HealthCheck(w http.ResponseWriter, r *http.Request) {
	messages := []responses.Message{
		{Lang: "ru", Message: "Сервер работает"},
		{Lang: "en", Message: "Server is working"},
	}

	responses.WriteJSON(w, http.StatusOK,
		responses.SuccessAPIResponse("SERVER_LIFE", messages, nil),
	)
}
