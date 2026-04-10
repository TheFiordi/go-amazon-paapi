package auth

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"
)

const (
	DefaultLwAEndpoint = "https://api.amazon.com/auth/o2/token"
	defaultScope       = "creatorsapi::default"
	grantType          = "client_credentials"

	tokenExpiryBuffer = 60 * time.Second
)

// Config holds the credentials required to initialize the Authenticator.
type Config struct {
	ClientID     string
	ClientSecret string
	LwAEndpoint  string // Optional: overrides the default endpoint
	HTTPClient   *http.Client
}

// Authenticator manages generation and caching of the LwA token (v3.x).
type Authenticator struct {
	config Config

	mu          sync.RWMutex
	accessToken string
	expiresAt   time.Time
}

// tokenRequest represents the JSON body sent to Amazon.
type tokenRequest struct {
	GrantType    string `json:"grant_type"`
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Scope        string `json:"scope"`
}

// tokenResponse represents a successful response from Amazon.
type tokenResponse struct {
	AccessToken string `json:"access_token"`
	Scope       string `json:"scope"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"` // In seconds
}

// tokenErrorResponse represents an error formatted by Amazon LwA.
type tokenErrorResponse struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

// NewAuthenticator creates a new instance of Authenticator.
func NewAuthenticator(cfg Config) *Authenticator {
	if cfg.LwAEndpoint == "" {
		cfg.LwAEndpoint = DefaultLwAEndpoint
	}
	if cfg.HTTPClient == nil {
		cfg.HTTPClient = &http.Client{Timeout: 10 * time.Second}
	}
	return &Authenticator{
		config: cfg,
	}
}

// GetToken returns a valid token from the cache or generates a new one.
// It is completely thread-safe.
func (a *Authenticator) GetToken(ctx context.Context) (string, error) {
	// Fast read with RLock to check for a valid token
	a.mu.RLock()
	token := a.accessToken
	expiresAt := a.expiresAt
	a.mu.RUnlock()

	// If the token exists and is not expiring soon, use the cache
	if token != "" && time.Now().Add(tokenExpiryBuffer).Before(expiresAt) {
		return token, nil
	}

	// Full lock to refresh the token
	a.mu.Lock()
	defer a.mu.Unlock()

	// Double-check locking: another goroutine might have refreshed the token
	// while waiting to acquire the write lock.
	if a.accessToken != "" && time.Now().Add(tokenExpiryBuffer).Before(a.expiresAt) {
		return a.accessToken, nil
	}

	// Make the HTTP call to refresh the token
	return a.refreshToken(ctx)
}

// refreshToken performs the HTTP request.
// WARNING: do not call this method directly, use GetToken.
func (a *Authenticator) refreshToken(ctx context.Context) (string, error) {
	reqBody := tokenRequest{
		GrantType:    grantType,
		ClientID:     a.config.ClientID,
		ClientSecret: a.config.ClientSecret,
		Scope:        defaultScope,
	}

	jsonBody, err := json.Marshal(reqBody)
	if err != nil {
		return "", fmt.Errorf("failed to marshal token request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, a.config.LwAEndpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return "", fmt.Errorf("failed to create token request: %w", err)
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := a.config.HTTPClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to execute token request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		var apiErr tokenErrorResponse
		if json.Unmarshal(body, &apiErr) == nil && apiErr.Error != "" {
			return "", fmt.Errorf("auth error [%s]: %s", apiErr.Error, apiErr.ErrorDescription)
		}
		return "", fmt.Errorf("token endpoint returned status %d: %s", resp.StatusCode, string(body))
	}

	var tResp tokenResponse
	if err := json.NewDecoder(resp.Body).Decode(&tResp); err != nil {
		return "", fmt.Errorf("failed to decode token response: %w", err)
	}

	// Update the in-memory cache
	a.accessToken = tResp.AccessToken
	a.expiresAt = time.Now().Add(time.Duration(tResp.ExpiresIn) * time.Second)

	return a.accessToken, nil
}
