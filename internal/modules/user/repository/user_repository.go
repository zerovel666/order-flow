package repository

import (
	"context"
	"order-flow/internal/modules/user/infrasturcture/dto"

	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository struct {
	pool *pgxpool.Pool
}

func NewUserRepository(pool *pgxpool.Pool) *UserRepository {
	return &UserRepository{pool: pool}
}

func (r *UserRepository) Create(ctx context.Context, u dto.User) (dto.User, error) {
	err := r.pool.QueryRow(
		ctx,
		`INSERT INTO users (name, password, role, api_key_hash)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, created_at, updated_at`,
		u.Name, u.Password, u.Role, u.APIKeyHash,
	).Scan(&u.ID, &u.CreatedAt, &u.UpdatedAt)

	if err != nil {
		return dto.User{}, err
	}

	return u, nil
}
