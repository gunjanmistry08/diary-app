package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/gunjanmistry08/diary-app/internal/database"
	"github.com/gunjanmistry08/diary-app/internal/handlers"
	"github.com/gunjanmistry08/diary-app/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	handlers.RegisterRoutes(r)
	return r
}

func TestMain(m *testing.M) {
	// Use in-memory SQLite for tests
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		panic("failed to connect test database")
	}
	database.DB = db
	// Migrate tables
	db.AutoMigrate(&models.User{}, &models.DiaryEntry{})
	os.Exit(m.Run())
}

func TestRegisterHandler(t *testing.T) {
	r := setupRouter()

	tests := []struct {
		name         string
		payload      gin.H
		expectedCode int
		expectedBody string
	}{
		{
			name:         "valid registration",
			payload:      gin.H{"username": "testuser", "password": "testpass", "email": "test@example.com"},
			expectedCode: http.StatusCreated,
			expectedBody: "User registered successfully",
		},
		{
			name:         "missing fields",
			payload:      gin.H{"username": "", "password": "", "email": ""},
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid request",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.payload)
			req, _ := http.NewRequest("POST", "/auth/register", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tc.expectedCode {
				t.Errorf("expected status %d, got %d", tc.expectedCode, w.Code)
			}
			if !bytes.Contains(w.Body.Bytes(), []byte(tc.expectedBody)) {
				t.Errorf("expected body to contain %q, got %q", tc.expectedBody, w.Body.String())
			}
		})
	}
}

func TestLoginHandler(t *testing.T) {
	r := setupRouter()

	tests := []struct {
		name         string
		payload      gin.H
		expectedCode int
		expectedBody string
	}{
		{
			name:         "missing fields",
			payload:      gin.H{"username": "", "password": ""},
			expectedCode: http.StatusBadRequest,
			expectedBody: "Invalid request",
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			body, _ := json.Marshal(tc.payload)
			req, _ := http.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)

			if w.Code != tc.expectedCode {
				t.Errorf("expected status %d, got %d", tc.expectedCode, w.Code)
			}
			if !bytes.Contains(w.Body.Bytes(), []byte(tc.expectedBody)) {
				t.Errorf("expected body to contain %q, got %q", tc.expectedBody, w.Body.String())
			}
		})
	}
}
