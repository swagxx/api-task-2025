package serve

import (
	"api-task-2025/database/model"
	"api-task-2025/internal/repository/repository"
	"api-task-2025/internal/service/quoteservice"
	"database/sql"
	"encoding/json"
	_ "github.com/lib/pq"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func testApp(t *testing.T) http.Handler {
	db, err := sql.Open("postgres", "postgres://postgres:Prestigio123@localhost:5432/db_quotes?sslmode=disable")
	if err != nil {
		t.Fatalf("Error connecting to db: %v", err)
	}
	repo := repository.NewRepository(db)
	service := quoteservice.NewQuoteService(repo)
	handler := NewHandler(service)

	router := http.NewServeMux()

	InitRouter(handler, router)
	return router
}

func TestHandler_Create(t *testing.T) {
	router := testApp(t)

	test := []struct {
		name        string
		req         string
		expectCode  int
		contentType string
	}{
		{
			name:        "success",
			req:         `{"author":"Papich_21213", "quote":"Unique quote"}`,
			expectCode:  http.StatusCreated,
			contentType: "application/json",
		},
		{
			name:        "author empty",
			req:         `{author: "", quote:"This empty"}`,
			expectCode:  http.StatusBadRequest,
			contentType: "application/json",
		},
		{
			name:        "quote empty",
			req:         `{"author":"Papich", "quote":""}`,
			expectCode:  http.StatusBadRequest,
			contentType: "application/json",
		},
		{
			name:        "wrong content type",
			req:         `author=Test&quote=test`,
			expectCode:  http.StatusUnsupportedMediaType,
			contentType: "application/x-www-form-urlencoded",
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("POST", "/quotes", strings.NewReader(tt.req))
			req.Header.Set("Content-Type", tt.contentType)

			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)

			if w.Code != tt.expectCode {
				t.Errorf("want %d, got %d. Response: %s", tt.expectCode, w.Code, w.Body.String())
			}

			if w.Code == http.StatusCreated {
				var response map[string]interface{}
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Errorf("could not decode response: %v", err)
				}

				if _, ok := response["id"]; !ok {
					t.Error("response should contain id field")
				}
			}
		})
	}
}

func TestGetByAuthor(t *testing.T) {
	router := testApp(t)

	test := []struct {
		name        string
		url         string
		expectCode  int
		expectItems bool
	}{
		{
			name:        "get by author",
			url:         "/quotes?author=Karrigan",
			expectCode:  http.StatusOK,
			expectItems: true,
		},
	}

	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.url, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			if w.Code != tt.expectCode {
				t.Errorf("want %d, got %d", tt.expectCode, w.Code)
			}

			if tt.expectCode == http.StatusOK {
				var response model.Quotes
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Errorf("could not decode response: %v", err)
				}
				if tt.expectItems && len(response.Quotes) == 0 {
					t.Error("expected items in response, got empty array")
				}
				if !tt.expectItems && len(response.Quotes) > 0 {
					t.Error("expected empty array, got items")
				}
			}
		})
	}
}

func TestGetRandom(t *testing.T) {
	router := testApp(t)
	test := []struct {
		name       string
		url        string
		expectCode int
	}{
		{
			name:       "get random quote",
			url:        "/quotes/random",
			expectCode: http.StatusOK,
		},
	}
	for _, tt := range test {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", tt.url, nil)
			w := httptest.NewRecorder()

			router.ServeHTTP(w, req)

			if w.Code != tt.expectCode {
				t.Errorf("want %d, got %d", tt.expectCode, w.Code)
			}
			if tt.expectCode == http.StatusOK {
				var response model.Quote
				if err := json.NewDecoder(w.Body).Decode(&response); err != nil {
					t.Errorf("could not decode response: %v", err)
				}
			}
		})
	}
}
