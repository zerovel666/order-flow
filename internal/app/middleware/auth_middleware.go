package middleware

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
	"net/http"
	"order-flow/internal/config/response"
	"order-flow/internal/modules/user/repository"
	"strings"

	"github.com/jackc/pgx/v5/pgxpool"
)

func AuthMiddleware(pool *pgxpool.Pool) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")
			token, ok := strings.CutPrefix(header, "Bearer ")
			if !ok || token == "" {
				writeErr(w)
				return
			}
			sum := sha256.Sum256([]byte(token))
			token = hex.EncodeToString(sum[:])
			exists, err := checkToken(pool, r.Context(), token)
			if err != nil || !exists {
				writeErr(w)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func checkToken(pool *pgxpool.Pool, ctx context.Context, token string) (bool, error) {
	userRepo := repository.NewUserRepository(pool)
	return userRepo.CheckToken(ctx, token)
}

func writeErr(w http.ResponseWriter) {
	response.WriteJSON(
		w,
		401,
		response.ErrorResponse(
			"UNAUTHORIZATION",
			[]response.Message{
				{Lang: "RU", Message: "Неавторизованный пользователь"},
				{Lang: "EN", Message: "Unauthorization user"},
			},
			nil,
			nil,
		),
	)
}
