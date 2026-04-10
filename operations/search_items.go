package operations

import (
	"context"
	"net/http"

	"go-amazon-paapi/models"
)

// SearchItems invokes the SearchItems API.
func SearchItems(ctx context.Context, client CreatorsClient, req *models.SearchItemsRequest) (*models.SearchItemsResponse, error) {
	var response models.SearchItemsResponse

	err := client.Do(ctx, http.MethodPost, "/catalog/v1/searchItems", req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
