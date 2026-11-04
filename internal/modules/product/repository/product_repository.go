package repository

import (
	"context"
	"order-flow/internal/modules/product/infrastructure/dto"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ProductRepository struct {
	pool *pgxpool.Pool
}

func NewProductRepository(pool *pgxpool.Pool) *ProductRepository {
	return &ProductRepository{pool: pool}
}

func (r *ProductRepository) Pool() *pgxpool.Pool {
	return r.pool
}

func (r *ProductRepository) Create(ctx context.Context, data dto.Product) (dto.Product, error) {
	err := r.pool.QueryRow(ctx, `
		INSERT INTO products (sku, name, price, currency, is_active)
		VALUES ($1, $2, $3, $4, $5)
		RETURNING id, created_at, updated_at
	`, data.Sku, data.Name, data.Price, data.Currency, data.IsActive).
		Scan(&data.ID, &data.CreatedAt, &data.UpdatedAt)
	return data, err
}

func (r *ProductRepository) Update(ctx context.Context, data dto.Product) error {
	result, err := r.pool.Exec(ctx, `
		UPDATE products SET sku = $2, name = $3, price = $4,
			currency = $5, is_active = $6, updated_at = now()
		WHERE id = $1
	`, data.ID, data.Sku, data.Name, data.Price, data.Currency, data.IsActive)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}

func (r *ProductRepository) GetAll(ctx context.Context) ([]dto.Product, error) {
	rows, err := r.pool.Query(ctx, `
		SELECT id, sku, name, price, currency, is_active, created_at, updated_at
		FROM products ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	products := make([]dto.Product, 0)
	for rows.Next() {
		product, err := scanProduct(rows)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, rows.Err()
}

func (r *ProductRepository) GetById(ctx context.Context, id int) (dto.Product, error) {
	return scanProduct(r.pool.QueryRow(ctx, `
		SELECT id, sku, name, price, currency, is_active, created_at, updated_at
		FROM products WHERE id = $1
	`, id))
}

func scanProduct(row pgx.Row) (dto.Product, error) {
	var product dto.Product
	err := row.Scan(&product.ID, &product.Sku, &product.Name, &product.Price,
		&product.Currency, &product.IsActive, &product.CreatedAt, &product.UpdatedAt)
	return product, err
}

func (r *ProductRepository) CreateInventory(ctx context.Context, data dto.Inventory) (dto.Inventory, error) {
	err := r.pool.QueryRow(ctx, `
		INSERT INTO inventory (product_id, available_quantity) VALUES ($1, $2)
		RETURNING updated_at
	`, data.ID, data.AvailableQuantity).Scan(&data.UpdatedAt)
	return data, err
}

func (r *ProductRepository) UnActive(ctx context.Context, id int) error {
	result, err := r.pool.Exec(ctx, `
		UPDATE products SET is_active = false, updated_at = now() WHERE id = $1
	`, id)
	if err != nil {
		return err
	}
	if result.RowsAffected() == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
