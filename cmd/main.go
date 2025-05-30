package main

import (
	"api-task-2025/config"
	db2 "api-task-2025/database"
	"api-task-2025/internal/repository/repository"
	"api-task-2025/internal/serve"
	"api-task-2025/internal/service/quoteservice"
	"database/sql"
	"log"
	"log/slog"
	"net/http"
)

func App(dbInstance *sql.DB) (*http.ServeMux, *serve.Server) {
	if err := db2.CreateTable(dbInstance); err != nil {
		slog.Error("error creating table: ", err)
	}

	//repository
	repo := repository.NewRepository(dbInstance)
	ser := quoteservice.NewQuoteService(repo)
	handler := serve.NewHandler(ser)

	router := http.NewServeMux()
	serve.InitRouter(handler, router)
	srv := serve.NewServer()
	return router, srv
}

func main() {
	cfg := config.MustLoad()

	dbApp, err := db2.ConnectDB(cfg)
	defer dbApp.Close()
	if err != nil {
		slog.Error("error connecting to database: ", err)
		return
	}

	router, srv := App(dbApp)

	slog.Info("Starting server on port 8080")
	if err := srv.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
