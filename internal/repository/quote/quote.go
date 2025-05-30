package quote

import (
	"api-task-2025/database/model"
	"crypto/rand"
	"database/sql"
	"fmt"
	"math/big"
)

const (
	QuoteTable = "quotes"
	QuoteCheck = "quote"
)

type Quote struct {
	db *sql.DB
}

func NewQuote(db *sql.DB) *Quote {
	return &Quote{db}
}

func (q *Quote) Create(quote model.Quote) (int, error) {
	tx, err := q.db.Begin()
	if err != nil {
		return 0, fmt.Errorf("create quote transaction error: %v", err)
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	checkQuery := fmt.Sprintf("SELECT EXISTS(SELECT 1 FROM %s WHERE %s = $1)", QuoteTable, QuoteCheck)
	var exists bool
	_ = tx.QueryRow(checkQuery, quote.Quote).Scan(&exists)
	if exists {
		return 0, fmt.Errorf("quote already exists")
	}

	var id int
	query := fmt.Sprintf(
		"INSERT INTO %s (author, quote) VALUES ($1,$2) RETURNING id", QuoteTable,
	)
	err = tx.QueryRow(query, quote.Author, quote.Quote).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("create quote row error: %v", err)
	}
	return id, nil
}

func (q *Quote) GetAll() ([]model.Quotes, error) {
	quoteMap := make(map[string][]string)

	query := fmt.Sprintf("SELECT author, quote FROM %s", QuoteTable)
	rows, err := q.db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("get all quotes error: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var quote model.Quote
		if err := rows.Scan(&quote.Author, &quote.Quote); err != nil {
			return nil, fmt.Errorf("scan quote error: %v", err)
		}
		quoteMap[quote.Author] = append(quoteMap[quote.Author], quote.Quote)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("get all quotes error: %v", err)
	}

	var result []model.Quotes
	for k, v := range quoteMap {
		result = append(result, model.Quotes{Author: k, Quotes: v})
	}
	return result, nil
}

func (q *Quote) GetRandomQuote() (*model.Quote, error) {
	var count int
	var quote model.Quote

	queryCount := fmt.Sprintf("SELECT count(*) FROM %s", QuoteTable)
	err := q.db.QueryRow(queryCount).Scan(&count)

	if count == 0 {
		return nil, fmt.Errorf("no quotes available")
	}

	if err != nil {
		return nil, fmt.Errorf("count quotes error: %v", err)
	}

	offset, err := rand.Int(rand.Reader, big.NewInt(int64(count)))

	if err != nil {
		return nil, fmt.Errorf("random offset error: %v", err)
	}

	querySelect := fmt.Sprintf("SELECT author, quote FROM %s LIMIT 1 OFFSET $1", QuoteTable)
	err = q.db.QueryRow(querySelect, offset.Int64()).Scan(&quote.Author, &quote.Quote)

	if err != nil {
		return nil, fmt.Errorf("cannot get random quotes: %v", err)
	}

	return &quote, nil

}

func (q *Quote) GetByAuthor(author string) (*model.Quotes, error) {
	query := fmt.Sprintf("SELECT author, quote FROM %s WHERE author=$1", QuoteTable)
	rows, err := q.db.Query(query, author)
	if err != nil {
		return nil, fmt.Errorf("query error: %v", err)
	}
	defer rows.Close()

	quotes := &model.Quotes{
		Author: author,
		Quotes: []string{},
	}

	for rows.Next() {
		var quote model.Quote
		if err := rows.Scan(&quote.Author, &quote.Quote); err != nil {
			return nil, fmt.Errorf("scan error : %v", err)
		}
		quotes.Quotes = append(quotes.Quotes, quote.Quote)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %v", err)
	}
	return quotes, nil
}

func (q *Quote) DeleteByID(id int) error {
	query := fmt.Sprintf("DELETE FROM %s WHERE id=$1", QuoteTable)
	res, err := q.db.Exec(query, id)
	if err != nil {
		return fmt.Errorf("failed to delete quote: %w", err)
	}
	rows, _ := res.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("quote with id %d not found", id)
	}
	return nil
}
