package requests

import (
	"encoding/json"
	"errors"
	"fmt"
	"mime"
	"net/http"
	"reflect"
	"strings"
)

func Parse[T any](r *http.Request) (T, error) {
	var result T

	contentType, _, err := mime.ParseMediaType(r.Header.Get("Content-type"))
	if err != nil {
		return result, err
	}

	switch contentType {
	case "application/json":
		return parseApplicationJson[T](r)
	default:
		return result, errors.New("Неизвестный тип тела запроса")
	}

}

func parseApplicationJson[T any](r *http.Request) (T, error) {
	var result T

	if reflect.TypeOf(&result).Elem().Kind() != reflect.Struct {
		return result, errors.New("T должен быть структурой")
	}

	if err := json.NewDecoder(r.Body).Decode(&result); err != nil {
		return result, err
	}

	if err := validateRequired(result); err != nil {
		return result, err
	}

	return result, nil
}

func validateRequired[T any](value T) error {
	v := reflect.ValueOf(value)
	t := v.Type()
	var missing []string

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		name := strings.Split(field.Tag.Get("json"), ",")[0]
		if !field.IsExported() || name == "-" {
			continue
		}
		for rule := range strings.SplitSeq(field.Tag.Get("validate"), ",") {
			if rule == "required" && v.Field(i).IsZero() {
				if name == "" {
					name = field.Name
				}
				missing = append(missing, name)
				break
			}
		}
	}

	if len(missing) > 0 {
		return fmt.Errorf("Обязательные поля не заполнены: %s", strings.Join(missing, ", "))
	}
	return nil
}
