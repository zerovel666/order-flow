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
	Login(ctx context.Context, dto dto.LoginRequestDTO) response.Response
	GetInfoByToken(ctx context.Context, r *http.Request) response.Response
}

type UserHandler struct {
	service UserService
}

func NewUserHandler(service UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) Register(mux *http.ServeMux, AuthMiddleware func(http.Handler) http.Handler) {
	mux.Handle("POST /user", AuthMiddleware(http.HandlerFunc(h.Create)))
	mux.Handle("POST /user/login", http.HandlerFunc(h.SingUp))
	mux.Handle("GET /user/info", AuthMiddleware(http.HandlerFunc(h.GetInfoByToken)))
}

// UserRegister godoc
// @Summary Регистрация пользователя
// @Tags User
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /user [post]
func (h *UserHandler) Create(w http.ResponseWriter, r *http.Request) {
	dto, err := request.Parse[dto.CreateUserRequestDTO](r)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	resp := h.service.Create(r.Context(), dto)
	response.WriteJSON(w, 200, resp)
}

// UserLogin godoc
// @Summary Авторизация пользователя
// @Tags User
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router /user/login [post]
func (h *UserHandler) SingUp(w http.ResponseWriter, r *http.Request) {
	dto, err := request.Parse[dto.LoginRequestDTO](r)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	resp := h.service.Login(r.Context(), dto)
	response.WriteJSON(w, 200, resp)
}

// UserInfoByToken godoc
// @Summary Информация пользователя по токену
// @Tags User
// @Produce json
// @Security BearerAuth
// @Success 200 {object} map[string]interface{}
// @Failure 401 {object} map[string]interface{}
// @Router /user/info [get]
func (h *UserHandler) GetInfoByToken(w http.ResponseWriter, r *http.Request) {
	resp := h.service.GetInfoByToken(r.Context(), r)
	response.WriteJSON(w, 200, resp)
}
