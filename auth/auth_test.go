package auth_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-amazon-paapi/auth"
)

func TestAuthenticator_GetToken_Success(t *testing.T) {
	// Create a mock auth server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"access_token":"mock_token", "expires_in": 3600, "token_type":"bearer"}`))
	}))
	defer ts.Close()

	cfg := auth.Config{
		ClientID:     "client_id",
		ClientSecret: "client_secret",
		LwAEndpoint:  ts.URL, // Use mock server
	}

	authenticator := auth.NewAuthenticator(cfg)

	token, err := authenticator.GetToken(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if token != "mock_token" {
		t.Errorf("expected token 'mock_token', got '%s'", token)
	}
}

func TestAuthenticator_GetToken_Error(t *testing.T) {
	// Create a mock auth server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(`{"error":"invalid_client", "error_description":"Client ID is invalid"}`))
	}))
	defer ts.Close()

	cfg := auth.Config{
		ClientID:     "invalid",
		ClientSecret: "invalid",
		LwAEndpoint:  ts.URL,
	}

	authenticator := auth.NewAuthenticator(cfg)

	_, err := authenticator.GetToken(context.Background())
	if err == nil {
		t.Fatalf("expected error, got none")
	}
}

func TestAuthenticator_GetToken_Cache(t *testing.T) {
	callCount := 0
	// Create a mock auth server
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		callCount++
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"access_token":"mock_token", "expires_in": 3600, "token_type":"bearer"}`))
	}))
	defer ts.Close()

	cfg := auth.Config{
		ClientID:     "client_id",
		ClientSecret: "client_secret",
		LwAEndpoint:  ts.URL,
		HTTPClient:   ts.Client(),
	}

	authenticator := auth.NewAuthenticator(cfg)

	// First call
	_, err := authenticator.GetToken(context.Background())
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if callCount != 1 {
		t.Errorf("expected 1 call to mock server, got %d", callCount)
	}

	// Second call should use cache
	_, err = authenticator.GetToken(context.Background())
	if err != nil {
		t.Fatalf("expected no error on second call, got %v", err)
	}

	if callCount != 1 {
		t.Errorf("expected 1 call to mock server after cache hit, got %d", callCount)
	}
}
