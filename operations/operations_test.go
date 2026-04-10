package operations_test

import (
	"context"
	"net/http"
	"testing"

	"go-amazon-paapi/models"
	"go-amazon-paapi/operations"
)

type MockClient struct {
	MethodCalled string
	PathCalled   string
	ReqCalled    interface{}
	ResToReturn  interface{}
	ErrToReturn  error
}

func (m *MockClient) Do(ctx context.Context, method, path string, in interface{}, out interface{}) error {
	m.MethodCalled = method
	m.PathCalled = path
	m.ReqCalled = in

	if m.ErrToReturn != nil {
		return m.ErrToReturn
	}

	// For semplicity, we don't copy the mock response to `out` but just check if it was called.
	return nil
}

func TestGetItems(t *testing.T) {
	mockClient := &MockClient{}

	req := &models.GetItemsRequest{
		ItemIds: []string{"B08N5WRWNW"},
	}

	_, err := operations.GetItems(context.Background(), mockClient, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if mockClient.MethodCalled != http.MethodPost {
		t.Errorf("expected POST method, got %s", mockClient.MethodCalled)
	}

	if mockClient.PathCalled != "/catalog/v1/getItems" {
		t.Errorf("expected path /catalog/v1/getItems, got %s", mockClient.PathCalled)
	}

	if req.ItemIdType != models.ItemIdTypeASIN {
		t.Errorf("expected default ItemIdType to be ASIN, got %s", req.ItemIdType)
	}
}

func TestSearchItems(t *testing.T) {
	mockClient := &MockClient{}

	req := &models.SearchItemsRequest{
		Keywords: "Harry Potter",
	}

	_, err := operations.SearchItems(context.Background(), mockClient, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if mockClient.PathCalled != "/catalog/v1/searchItems" {
		t.Errorf("expected path /catalog/v1/searchItems, got %s", mockClient.PathCalled)
	}
}

func TestGetVariations(t *testing.T) {
	mockClient := &MockClient{}

	req := &models.GetVariationsRequest{
		ASIN: "B08N5WRWNW",
	}

	_, err := operations.GetVariations(context.Background(), mockClient, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if mockClient.PathCalled != "/catalog/v1/getVariations" {
		t.Errorf("expected path /catalog/v1/getVariations, got %s", mockClient.PathCalled)
	}
}

func TestGetBrowseNodes(t *testing.T) {
	mockClient := &MockClient{}

	req := &models.GetBrowseNodesRequest{
		BrowseNodeIds: []string{"283155"},
	}

	_, err := operations.GetBrowseNodes(context.Background(), mockClient, req)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if mockClient.PathCalled != "/catalog/v1/getBrowseNodes" {
		t.Errorf("expected path /catalog/v1/getBrowseNodes, got %s", mockClient.PathCalled)
	}
}
