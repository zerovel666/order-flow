package user_v1

import (
	"context"
	"net/http"
	"order-flow/internal/config/request"
	"order-flow/internal/config/response"
	"order-flow/internal/modules/user/infrasturcture/dto"
)

type UserService interface {
	Create(ctx context.Context, dto dto.CreateUserRequestDTO) response.Response
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(mux *http.ServeMux) {
	mux.HandleFunc("POST /user", h.Create)
}

func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	dto, err := request.Parse[dto.CreateUserRequestDTO](r)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	resp := h.service.Create(r.Context(), dto)
	response.WriteJSON(w, 200, resp)
}
