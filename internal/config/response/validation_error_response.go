package response

import (
	"fmt"
	"net/http"
)

func BadRequest(w http.ResponseWriter, msg string) {
	messages := []Message{
		{Lang: "RU", Message: fmt.Sprintf("Ошибка валидации входных данных: %s", msg)},
		{Lang: "EN", Message: fmt.Sprintf("Validation error input data: %s", msg)},
	}
	WriteJSON(w, 400, Response{
		Code:     "VALIDATION_ERROR",
		Messages: messages,
		Data:     nil,
		Errors:   msg,
	})
}
