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
	Update(ctx context.Context, dto productDTO.UpdateProductRequestDTO) response.Response
	UnActive(ctx context.Context, product_id int) response.Response
	GetAll(ctx context.Context) response.Response
	GetMyShop(ctx context.Context) response.Response
}

type ProductHandler struct {
	service ProductService
}

func NewProductHandler(service ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) Register(mux *http.ServeMux, AuthMiddleware func(http.Handler) http.Handler, PolicyMiddleware func(http.Handler, ...userDTO.Role) http.Handler) {
	mux.Handle("POST /product", AuthMiddleware(PolicyMiddleware(http.HandlerFunc(h.Create), userDTO.SellerRole)))
	mux.Handle("PUT /product", AuthMiddleware(PolicyMiddleware(http.HandlerFunc(h.Update), userDTO.SellerRole)))
	mux.Handle("DELETE /product/{id}", AuthMiddleware(PolicyMiddleware(http.HandlerFunc(h.UnActive), userDTO.SellerRole)))
	mux.Handle("GET /product", AuthMiddleware(http.HandlerFunc(h.GetAll)))
	mux.Handle("GET /product/shop", AuthMiddleware(http.HandlerFunc(h.GetMyShop)))
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {
	dto, err := request.Parse[productDTO.CreateProductRequestDTO](r)
	if err != nil {
		response.BadRequest(w, err.Error())
	}

	resp := h.service.Create(r.Context(), dto)
	response.WriteJSON(w, 200, resp)
}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {
	dto, err := request.Parse[productDTO.UpdateProductRequestDTO](r)
	if err != nil {
		response.BadRequest(w, err.Error())
	}

	resp := h.service.Update(r.Context(), dto)
	response.WriteJSON(w, 200, resp)
}

func (h *ProductHandler) UnActive(w http.ResponseWriter, r *http.Request) {
	value := r.PathValue("id")
	product_id, err := strconv.Atoi(value)
	if err != nil {
		response.BadRequest(w, err.Error())
	}
	resp := h.service.UnActive(r.Context(), product_id)
	response.WriteJSON(w, 200, resp)
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	resp := h.service.GetAll(r.Context())
	response.WriteJSON(w, 200, resp)
}

func (h *ProductHandler) GetMyShop(w http.ResponseWriter, r *http.Request) {
	resp := h.service.GetMyShop(r.Context())
	response.WriteJSON(w, 200, resp)
}
