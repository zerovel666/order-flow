package response

import (
	"encoding/json"
	"net/http"
)

type Message struct {
	Lang    string `json:"lang"`
	Message string `json:"message"`
}

type Response struct {
	Status   bool      `json:"status"`
	Code     string    `json:"code"`
	Messages []Message `json:"messages"`
	Data     any       `json:"data,omitempty"`
	Errors   any       `json:"errors,omitempty"`
}

func SuccessAPIResponse(code string, messages []Message, data any) Response {
	return Response{
		Status:   true,
		Code:     code,
		Messages: messages,
		Data:     data,
	}
}

func ErrorResponse(code string, messages []Message, data any, errors any) Response {
	return Response{
		Status:   false,
		Code:     code,
		Messages: messages,
		Data:     data,
		Errors:   errors,
	}
}

func WriteJSON(w http.ResponseWriter, httpStatus int, resp Response) {
	w.Header().Set("Content-type", "application/json; charset=utf-8")
	w.WriteHeader(httpStatus)
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
