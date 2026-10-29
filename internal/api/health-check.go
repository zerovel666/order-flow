package api

import (
	"net/http"
	"order-flow/internal/config/response"
)

// HealthCheck godoc
// @Summary Проверка состояния сервиса
// @Tags health
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /health-check [get]
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	messages := []response.Message{
		{Lang: "ru", Message: "Сервер работает"},
		{Lang: "en", Message: "Server is working"},
	}

	response.WriteJSON(w, http.StatusOK,
		response.SuccessAPIResponse("SERVER_LIFE", messages, nil),
	)
}
