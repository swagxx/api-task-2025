package serve

import (
	"api-task-2025/pkg/middleware"
	"net/http"
)

func InitRouter(quote *Handler, router *http.ServeMux) {
	router.Handle("POST /quotes", middleware.RequireJSON(quote.Create()))
	router.HandleFunc("GET /quotes", quote.Get())
	router.HandleFunc("GET /quotes/random", quote.GetRandom())
	router.HandleFunc("DELETE /quotes/{id}", quote.Delete())
}
