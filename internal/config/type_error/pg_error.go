package type_error

import (
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type PGConstError string

const (
	DUPLICATE PGConstError = "DUPLICATION"
	NOT_FOUND PGConstError = "NOT_FOUND"
	UNDEFIEND PGConstError = "UNDEFINED"
)

func PgErr(err error) PGConstError {
	if errors.Is(err, pgx.ErrNoRows) {
		return NOT_FOUND
	}

	var pgErr *pgconn.PgError
	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return DUPLICATE
		}
	}

	return UNDEFIEND
}
