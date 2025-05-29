package repository

import (
	"api-task-2025/internal/model"
	"api-task-2025/internal/repository/quote"
	"database/sql"
)

type QuoteStore interface {
	Create(quote model.Quote) (int, error)
	GetAll() ([]*model.Quote, error)
	GetRandomQuote() (*model.Quote, error)
	GetByAuthor(author string) ([]*model.Quote, error)
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
