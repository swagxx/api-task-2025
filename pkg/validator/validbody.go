package validator

import (
	"api-task-2025/database/model"
	"errors"
)

func ValidBody(data interface{}) error {
	switch v := data.(type) {
	case model.Quote:
		if v.Author == "" {
			return errors.New("author is required")
		}

		if v.Quote == "" {
			return errors.New("quote is required")
		}
	}
	return nil
}
