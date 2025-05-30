package validator

import (
	"api-task-2025/database/model"
	"errors"
)

func ValidBody(data model.Quote) error {
	if data.Author == "" || data.Quote == "" {
		return errors.New("author or Quote is empty")
	}
	return nil
}
