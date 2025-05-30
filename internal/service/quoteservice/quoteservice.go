package quoteservice

import (
	"api-task-2025/database/model"
	"api-task-2025/internal/repository/repository"
	"fmt"
)

type QuoteService struct {
	repo repository.QuoteStore
}

func NewQuoteService(repo repository.QuoteStore) *QuoteService {
	return &QuoteService{
		repo: repo,
	}
}

func (q *QuoteService) Create(quote model.Quote) (int, error) {
	res, err := q.repo.Create(quote)
	if err != nil {
		return 0, fmt.Errorf("create quote: %w", err)
	}
	return res, nil
}

func (q *QuoteService) GetAll() ([]model.Quotes, error) {
	quotes, err := q.repo.GetAll()
	if err != nil {
		return nil, fmt.Errorf("get all quotes: %w", err)
	}
	return quotes, nil
}

func (q *QuoteService) GetRandomQuote() (*model.Quote, error) {
	quote, err := q.repo.GetRandomQuote()
	if err != nil {
		return nil, fmt.Errorf("get random quote: %w", err)
	}
	return quote, nil
}

func (q *QuoteService) GetByAuthor(author string) (*model.Quotes, error) {
	quotes, err := q.repo.GetByAuthor(author)
	if err != nil {
		return nil, fmt.Errorf("get quotes by author %s: %w", author, err)
	}
	return quotes, nil
}

func (q *QuoteService) DeleteByID(id int) error {
	if err := q.repo.DeleteByID(id); err != nil {
		return fmt.Errorf("delete quote by id %d: %w", id, err)
	}
	return nil
}
