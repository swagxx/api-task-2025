package serve

import (
	"api-task-2025/database/model"
	"api-task-2025/internal/service/quoteservice"
	"api-task-2025/pkg/handlerset"
	"api-task-2025/pkg/validator"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"strconv"
	"strings"
)

type Handler struct {
	service *quoteservice.QuoteService
}

func NewHandler(ser *quoteservice.QuoteService) *Handler {
	return &Handler{
		service: ser,
	}
}

func (h *Handler) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var payload model.Quote
		defer r.Body.Close()
		body, err := io.ReadAll(r.Body)
		if err != nil {
			return
		}
		if err := json.Unmarshal(body, &payload); err != nil {
			http.Error(w, "failed to read body", http.StatusBadRequest)
			return
		}
		if err := validator.ValidBody(payload); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		payload.Quote = correctQuote(payload.Quote)
		id, err := h.service.Create(payload)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		handlerset.HandlerSet(w, struct {
			ID string `json:"id"`
		}{ID: strconv.Itoa(id)}, http.StatusCreated)

	}
}

func (h *Handler) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		author := r.URL.Query().Get("author")
		switch {
		case author != "":
			h.GetAuthor(author)(w, r)
		default:
			h.GetAll()(w, r)
		}
	}
}

func (h *Handler) GetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		quotes, err := h.service.GetAll()
		if err != nil {
			slog.Error("failed to get all quotes", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		handlerset.HandlerSet(w, struct {
			AllQuotes []model.Quotes `json:"all_quotes"`
		}{
			AllQuotes: quotes,
		}, http.StatusOK)
	}
}

func (h *Handler) GetRandom() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		result, err := h.service.GetRandomQuote()
		if err != nil {
			slog.Error("failed to get random quote", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		handlerset.HandlerSet(w, result, http.StatusOK)
	}
}

func (h *Handler) GetAuthor(author string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if author == "" {
			slog.Error("author is empty")
			http.Error(w, "author parameter is required", http.StatusBadRequest)
			return
		}
		quotes, err := h.service.GetByAuthor(author)
		if err != nil {
			slog.Error("failed to get quotes")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if len(quotes.Quotes) == 0 {
			http.Error(w, "no quotes found for this author", http.StatusNotFound)
			return
		}

		handlerset.HandlerSet(w, quotes, http.StatusOK)
	}
}

func (h *Handler) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.PathValue("id")
		id, err := strconv.ParseUint(idString, 10, 32)
		if err != nil {
			slog.Error("failed to parse id", "error", err)
			handlerset.HandlerSet(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err = h.service.DeleteByID(int(id)); err != nil {
			slog.Error("failed to delete quote", "error", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		handlerset.HandlerSet(w, "successful delete", http.StatusOK)
	}
}

func correctQuote(quote string) string {
	runes := []rune(quote)
	return strings.ToUpper(string(runes[0])) + strings.ToLower(string(runes[1:]))
}
