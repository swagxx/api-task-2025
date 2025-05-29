package db

import (
	"api-task-2025/config"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func ConnectDB(cfg *config.Config) (*sql.DB, error) {
	connectStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.DB.DBHost, cfg.DB.DBPort, cfg.DB.DBUser, cfg.DB.DBPassword, cfg.DB.DBName, cfg.DB.DBSSLMode,
	)
	db, err := sql.Open("postgres", connectStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("ping to database failed: %v", err)
	}
	return db, nil
}
