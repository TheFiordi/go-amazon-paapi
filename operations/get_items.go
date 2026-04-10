package operations

import (
	"context"
	"net/http"

	"go-amazon-paapi/models"
)

// CreatorsClient defines the interface that our main client must satisfy.
type CreatorsClient interface {
	Do(ctx context.Context, method, path string, in interface{}, out interface{}) error
}

// GetItems invokes the Amazon Creators GetItems API.
func GetItems(ctx context.Context, client CreatorsClient, req *models.GetItemsRequest) (*models.GetItemsResponse, error) {
	// Set reasonable defaults if not provided by the user
	if req.ItemIdType == "" {
		req.ItemIdType = models.ItemIdTypeASIN
	}

	var response models.GetItemsResponse

	// Execute the call using the abstract client.
	// Note: The x-marketplace header required by the documentation should ideally be
	// injected in the client's Do() method by reading from the context or config,
	// or we could add it by extending the Do() signature to accept extra headers.
	err := client.Do(ctx, http.MethodPost, "/catalog/v1/getItems", req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
