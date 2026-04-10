package go_amazon_paapi_test

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"testing"

	paapi "go-amazon-paapi"
	"go-amazon-paapi/locale"
	"go-amazon-paapi/models"
)

// roundTripperFunc allows us to use a function as an http.RoundTripper
type roundTripperFunc func(*http.Request) (*http.Response, error)

func (f roundTripperFunc) RoundTrip(req *http.Request) (*http.Response, error) {
	return f(req)
}

func TestClientBuilder(t *testing.T) {
	// Custom HTTP client mocking auth response
	customClient := &http.Client{
		Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewBufferString(`{"access_token": "fake_token", "expires_in": 3600}`)),
				Header:     make(http.Header),
			}, nil
		}),
	}

	builder := paapi.New(
		paapi.WithMarketplace(locale.Italy),
		paapi.WithHttpClient(customClient),
	)

	if builder == nil {
		t.Fatal("expected builder to not be nil")
	}

	client, err := builder.CreateClient("partnerTag", "clientID", "clientSecret")
	if err != nil {
		t.Fatalf("expected no error, got: %v", err)
	}

	if client == nil {
		t.Fatal("expected client to not be nil")
	}
}

func TestClientDoInjectsData(t *testing.T) {
	partnerTag := "tag"
	clientId := "id"
	clientSecret := "secret"

	// Custom HTTP client mocking auth response and then capturing the main request
	var lastReq *http.Request
	customClient := &http.Client{
		Transport: roundTripperFunc(func(req *http.Request) (*http.Response, error) {
			if req.URL.Path == "/auth/o2/token" {
				return &http.Response{
					StatusCode: http.StatusOK,
					Body:       io.NopCloser(bytes.NewBufferString(`{"access_token": "fake_token", "expires_in": 3600}`)),
					Header:     make(http.Header),
				}, nil
			}
			lastReq = req
			return &http.Response{
				StatusCode: http.StatusNoContent,
				Body:       io.NopCloser(bytes.NewBufferString("")),
				Header:     make(http.Header),
			}, nil
		}),
	}

	client, err := paapi.New(
		paapi.WithMarketplace(locale.Italy),
		paapi.WithHttpClient(customClient),
	).CreateClient(partnerTag, clientId, clientSecret)

	if err != nil {
		t.Fatalf("expected no error creating client, got: %v", err)
	}

	reqBody := &models.GetItemsRequest{}

	// Because of our mock, this Do should succeed.
	err = client.Do(context.Background(), http.MethodPost, "/foo", reqBody, nil)
	if err != nil {
		t.Fatalf("expected no error from Do, got: %v", err)
	}

	// Check if partner tag and marketplace were injected into the payload
	if reqBody.PartnerTag != partnerTag {
		t.Errorf("expected partner tag '%s', got '%s'", partnerTag, reqBody.PartnerTag)
	}

	if reqBody.Marketplace != "www.amazon.it" {
		t.Errorf("expected marketplace 'www.amazon.it', got '%s'", reqBody.Marketplace)
	}

	// Check if headers were set correctly
	if lastReq != nil {
		if lastReq.Header.Get("x-marketplace") != "www.amazon.it" {
			t.Errorf("expected x-marketplace header to be 'www.amazon.it', got '%s'", lastReq.Header.Get("x-marketplace"))
		}
	} else {
		t.Fatal("expected a request to be made but lastReq is nil")
	}
}
