package middleware

import (
	"net/http"
	"order-flow/internal/config/request"
	"order-flow/internal/config/response"
	"order-flow/internal/modules/user/infrastructure/dto"
	"slices"
)

func PolicyMiddleware() func(http.Handler, ...dto.Role) http.Handler {
	return func(next http.Handler, roles ...dto.Role) http.Handler {
		allowed := append([]dto.Role{dto.AdminRole}, roles...)

		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			u, err := request.User(r)
			if err != nil {
				writeErr(w)
				return
			}

			if !slices.Contains(allowed, u.Role) {
				response.WriteJSON(w, 422, response.ErrorResponse(
					"PERMISSION_DENIED",
					[]response.Message{
						{Lang: "RU", Message: "Доступ запрещен"},
						{Lang: "EN", Message: "Permission denied"},
					},
					nil,
					nil,
				))
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
