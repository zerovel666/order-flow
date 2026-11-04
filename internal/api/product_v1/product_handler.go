package product_v1

import (
	"context"
	"net/http"
	"order-flow/internal/config/request"
	"order-flow/internal/config/response"
	productDTO "order-flow/internal/modules/product/infrastructure/dto"
	userDTO "order-flow/internal/modules/user/infrastructure/dto"
	"strconv"
)

type ProductService interface {
	Create(ctx context.Context, dto productDTO.CreateProductRequestDTO) response.Response
	UnActive(ctx context.Context, product_id int) response.Response
	GetAll(ctx context.Context) response.Response
}

type ProductHandler struct {
	service ProductService
}

func NewProductHandler(service ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) Register(mux *http.ServeMux, AuthMiddleware func(http.Handler) http.Handler, PolicyMiddleware func(http.Handler, ...userDTO.Role) http.Handler) {
	mux.Handle("POST /product", AuthMiddleware(PolicyMiddleware(http.HandlerFunc(h.Create), userDTO.SellerRole)))
	mux.Handle("DELETE /product/{id}", AuthMiddleware(PolicyMiddleware(http.HandlerFunc(h.UnActive), userDTO.SellerRole)))
	mux.Handle("GET /product", AuthMiddleware(http.HandlerFunc(h.GetAll)))
}

// Create godoc
// @Summary Создание товара
// @Description Создаёт товар и его остаток. Доступно продавцу (SELLER) и администратору (ADMIN). Ошибки сервиса возвращаются с HTTP 200 и status=false.
// @Tags Product
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param request body dto.CreateProductRequestDTO true "Данные товара и начальный остаток"
// @Success 200 {object} response.Response{data=dto.CreateProductResponseDTO} "Товар с вложенным inventory; при ошибке status=false"
// @Failure 400 {object} response.Response "Некорректное тело запроса или отсутствуют обязательные поля"
// @Failure 401 {object} response.Response "Необходима авторизация"
// @Failure 422 {object} response.Response "Недостаточно прав"
// @Router /product [post]
func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	dto, err := request.Parse[productDTO.CreateProductRequestDTO](r)
	if err != nil {
		response.BadRequest(w, err.Error())
		return
	}

	resp := h.service.Create(r.Context(), dto)
	response.WriteJSON(w, 200, resp)
}

// UnActive godoc
// @Summary Деактивация товара
// @Description Устанавливает is_active=false без удаления товара. Доступно SELLER и ADMIN. Ошибки сервиса, включая отсутствие товара, возвращаются с HTTP 200 и status=false.
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Param id path int true "ID товара" minimum(1)
// @Success 200 {object} response.Response "Результат деактивации; при ошибке status=false"
// @Failure 400 {object} response.Response "ID должен быть положительным целым числом"
// @Failure 401 {object} response.Response "Необходима авторизация"
// @Failure 422 {object} response.Response "Недостаточно прав"
// @Router /product/{id} [delete]
func (h *ProductHandler) UnActive(w http.ResponseWriter, r *http.Request) {
	value := r.PathValue("id")
	product_id, err := strconv.Atoi(value)
	if err != nil || product_id <= 0 {
		response.BadRequest(w, "id должен быть положительным целым числом")
		return
	}
	resp := h.service.UnActive(r.Context(), product_id)
	response.WriteJSON(w, 200, resp)
}

// GetAll godoc
// @Summary Список товаров
// @Description Возвращает все товары, включая неактивные, в порядке ID. Доступно авторизованным пользователям. Ошибки сервиса возвращаются с HTTP 200 и status=false.
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Success 200 {object} response.Response{data=[]dto.Product} "Список товаров; при ошибке status=false"
// @Failure 401 {object} response.Response "Необходима авторизация"
// @Router /product [get]
func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	resp := h.service.GetAll(r.Context())
	response.WriteJSON(w, 200, resp)
}
