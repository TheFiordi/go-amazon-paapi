package operations

import (
	"context"
	"net/http"

	"go-amazon-paapi/models"
)

// GetVariations invokes the GetVariations API.
func GetVariations(ctx context.Context, client CreatorsClient, req *models.GetVariationsRequest) (*models.GetVariationsResponse, error) {
	var response models.GetVariationsResponse

	err := client.Do(ctx, http.MethodPost, "/catalog/v1/getVariations", req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
