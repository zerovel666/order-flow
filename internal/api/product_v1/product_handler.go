package product_v1

import (
	"context"
	"net/http"
	"order-flow/internal/config/response"
	"order-flow/internal/modules/products/infrastructure/dto"
)

type ProductService interface {
	Create(ctx context.Context, dto dto.CreateProductRequestDTO) response.Response
	Update(ctx context.Context, dto dto.UpdateProductRequestDTO) response.Response
	UnActive(ctx context.Context) response.Response
	GetAll(ctx context.Context) response.Response
	GetMyShop(ctx context.Context) response.Response
}

type ProductHandler struct {
	service ProductService
}

func NewProductHandler(service ProductService) *ProductHandler {
	return &ProductHandler{service: service}
}

func (h *ProductHandler) Register(mux *http.ServeMux, AuthMiddleware, PolicyMiddleware func(http.Handler) http.Handler) {
	mux.Handle("POST /product", AuthMiddleware(PolicyMiddleware(http.HandlerFunc(h.Create))))
	mux.Handle("PUT /product", AuthMiddleware(PolicyMiddleware(http.HandlerFunc(h.Update))))
	mux.Handle("DELETE /product/{id}", AuthMiddleware(PolicyMiddleware(http.HandlerFunc(h.UnActive))))
	mux.Handle("GET /product", AuthMiddleware(http.HandlerFunc(h.GetAll)))
	mux.Handle("GET /product/shop", AuthMiddleware(PolicyMiddleware(http.HandlerFunc(h.GetMyShop))))
}

func (h *ProductHandler) Create(w http.ResponseWriter, r *http.Request) {

}

func (h *ProductHandler) Update(w http.ResponseWriter, r *http.Request) {

}

func (h *ProductHandler) UnActive(w http.ResponseWriter, r *http.Request) {
	// id := r.PathValue("id")
}

func (h *ProductHandler) GetAll(w http.ResponseWriter, r *http.Request) {

}

func (h *ProductHandler) GetMyShop(w http.ResponseWriter, r *http.Request) {

}
