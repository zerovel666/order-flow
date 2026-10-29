package response

import "order-flow/internal/config/type_error"

func PgErrResponse(typeErr type_error.PGConstError, err error) Response {
	switch typeErr {
	case type_error.DUPLICATE:
		return duplicatePgErr(typeErr, err)
	default:
		return undefiendPgErr(typeErr, err)
	}
}

func duplicatePgErr(typeErr type_error.PGConstError, err error) Response {
	return ErrorResponse(
		string(typeErr),
		[]Message{
			{Lang: "RU", Message: "Запись уже существует"},
			{Lang: "EN", Message: "Record already exists"},
		},
		nil,
		err.Error(),
	)
}

func undefiendPgErr(typeErr type_error.PGConstError, err error) Response {
	return ErrorResponse(
		string(typeErr),
		[]Message{
			{Lang: "RU", Message: "Неизвестная ошибка"},
			{Lang: "EN", Message: "Undefiend error"},
		},
		nil,
		err.Error(),
	)
}
