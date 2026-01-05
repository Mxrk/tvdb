package tvdb

import (
	"context"
	"fmt"

	"github.com/mxrkw/tvdb/models"
)

// GetArtworkBase returns a single artwork base record.
func (c *Client) GetArtworkBase(ctx context.Context, id int64) (*models.ArtworkBaseRecord, error) {
	var response struct {
		Status string                   `json:"status"`
		Data   models.ArtworkBaseRecord `json:"data"`
	}

	path := fmt.Sprintf("/artwork/%d", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetArtworkExtended returns a single artwork extended record.
func (c *Client) GetArtworkExtended(ctx context.Context, id int64) (*models.ArtworkExtendedRecord, error) {
	var response struct {
		Status string                       `json:"status"`
		Data   models.ArtworkExtendedRecord `json:"data"`
	}

	path := fmt.Sprintf("/artwork/%d/extended", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetAllArtworkStatuses returns a list of artwork status records.
func (c *Client) GetAllArtworkStatuses(ctx context.Context) ([]models.ArtworkStatus, error) {
	var response struct {
		Status string                 `json:"status"`
		Data   []models.ArtworkStatus `json:"data"`
	}

	if err := c.get(ctx, "/artwork/statuses", nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetAllArtworkTypes returns a list of artwork type records.
func (c *Client) GetAllArtworkTypes(ctx context.Context) ([]models.ArtworkType, error) {
	var response struct {
		Status string               `json:"status"`
		Data   []models.ArtworkType `json:"data"`
	}

	if err := c.get(ctx, "/artwork/types", nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
