package user_v1

import (
	"context"
	"net/http"
	"order-flow/internal/config/requests"
	"order-flow/internal/config/responses"
	"order-flow/internal/modules/users/infrasturcture/structs"
)

type UserService interface {
	Create(ctx context.Context, dto structs.CreateUserRequestDTO) responses.Response
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
	dto, err := requests.Parse[structs.CreateUserRequestDTO](r)
	if err != nil {
		responses.BadRequest(w, err.Error())
		return
	}

	resp := h.service.Create(r.Context(), dto)
	responses.WriteJSON(w, 200, resp)
}
