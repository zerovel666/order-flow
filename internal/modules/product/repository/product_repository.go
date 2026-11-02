package repository

import (
	"context"
	"order-flow/internal/modules/product/infrastructure/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	pool *pgxpool.Pool
}

func NewProductRepository(pool *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{pool: pool}
}

func (r *ProductRepository) Create(ctx context.Context, data dto.Product) {

}

func (r *ProductRepository) Update(ctx context.Context, data dto.Product) {

}

func (r *ProductRepository) GetAll(ctx context.Context) {

}

func (r *ProductRepository) GetMyShop(ctx context.Context, userSellerId int) {

}
