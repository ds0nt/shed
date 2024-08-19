package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ds0nt/shed/domain/users"
	"github.com/stretchr/testify/assert"
)

func TestService_createUserHandler(t *testing.T) {
	s := NewService()
	defer s.Store.Close()

	// Prepare the request payload
	reqPayload := &users.User{
		Email:    "test@test.com",
		Password: "test_password",
	}
	reqBody, err := json.Marshal(reqPayload)
	if err != nil {
		t.Fatalf("Failed to prepare test request payload: %v", err)
	}

	// Create a new request
	req := httptest.NewRequest(http.MethodPost, "/users/create", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rec := httptest.NewRecorder()

	// Call the handler
	c := s.Echo.NewContext(req, rec)
	if assert.NoError(t, s.createUserHandler(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}
}

func TestService_loginUserHandler(t *testing.T) {
	s := NewService()
	defer s.Store.Close()

	// Create a test user
	testUser := &users.User{
		Email:    "test@test.com",
		Password: "test_password",
	}
	// Hash the test user's password before storing
	hashedPassword, _ := s.hashPassword(testUser.Password)
	testUser.Password = hashedPassword

	key := users.NewUserKey(testUser.Email)
	_ = s.Store.CreateJSON(context.Background(), "users", key.String(), &testUser)

	// Prepare the login request
	loginPayload := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    testUser.Email,
		Password: "test_password", // Use the original password here
	}
	reqBody, err := json.Marshal(loginPayload)
	if err != nil {
		t.Fatalf("Failed to prepare test request payload: %v", err)
	}

	// Create a new request
	req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rec := httptest.NewRecorder()

	// Call the handler
	c := s.Echo.NewContext(req, rec)
	if assert.NoError(t, s.loginUserHandler(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}

func TestService_rejectLoginHandler(t *testing.T) {
	s := NewService()
	defer s.Store.Close()

	// Prepare the login request
	loginPayload := struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}{
		Email:    "invalid@test.com",
		Password: "invalid_password",
	}
	reqBody, err := json.Marshal(loginPayload)
	if err != nil {
		t.Fatalf("Failed to prepare test request payload: %v", err)
	}

	// Create a new request
	req := httptest.NewRequest(http.MethodPost, "/users/login", bytes.NewBuffer(reqBody))
	req.Header.Set("Content-Type", "application/json")

	// Create a response recorder
	rec := httptest.NewRecorder()

	// Call the handler
	c := s.Echo.NewContext(req, rec)
	if assert.NoError(t, s.loginUserHandler(c)) {
		assert.Equal(t, http.StatusUnauthorized, rec.Code)
	}
}
