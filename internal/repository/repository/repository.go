package repository

import (
	"api-task-2025/database/model"
	"api-task-2025/internal/repository/quote"
	"database/sql"
)

type QuoteStore interface {
	Create(quote model.Quote) (int, error)
	GetAll() ([]model.Quotes, error)
	GetRandomQuote() (*model.Quote, error)
	GetByAuthor(author string) (*model.Quotes, error)
	DeleteByID(id int) error
}

type Repository struct {
	QuoteStore
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		QuoteStore: quote.NewQuote(db),
	}
}
