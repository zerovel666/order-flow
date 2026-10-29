package type_error

import (
	"errors"

	"github.com/jackc/pgx/v5/pgconn"
)

type PGConstError string

const (
	DUPLICATE PGConstError = "DUPLICATION"
	UNDEFIEND PGConstError = "UNDEFINED"
)

func PgErr(err error) PGConstError {
	var pgErr *pgconn.PgError

	if errors.As(err, &pgErr) {
		switch pgErr.Code {
		case "23505":
			return DUPLICATE
		}
	}

	return UNDEFIEND
}
