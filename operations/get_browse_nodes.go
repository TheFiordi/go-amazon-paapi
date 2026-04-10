package operations

import (
	"context"
	"net/http"

	"go-amazon-paapi/models"
)

// GetBrowseNodes invokes the GetBrowseNodes API.
func GetBrowseNodes(ctx context.Context, client CreatorsClient, req *models.GetBrowseNodesRequest) (*models.GetBrowseNodesResponse, error) {
	var response models.GetBrowseNodesResponse

	err := client.Do(ctx, http.MethodPost, "/catalog/v1/getBrowseNodes", req, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
