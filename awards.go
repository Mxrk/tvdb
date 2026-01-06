package tvdb

import (
	"context"
	"fmt"

	"github.com/mxrk/tvdb/models"
)

// GetAllAwards returns a list of award base records.
func (c *Client) GetAllAwards(ctx context.Context) ([]models.AwardBaseRecord, error) {
	var response struct {
		Status string                   `json:"status"`
		Data   []models.AwardBaseRecord `json:"data"`
	}

	if err := c.get(ctx, "/awards", nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetAward returns a single award base record.
func (c *Client) GetAward(ctx context.Context, id int64) (*models.AwardBaseRecord, error) {
	var response struct {
		Status string                 `json:"status"`
		Data   models.AwardBaseRecord `json:"data"`
	}

	path := fmt.Sprintf("/awards/%d", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetAwardExtended returns a single award extended record.
func (c *Client) GetAwardExtended(ctx context.Context, id int64) (*models.AwardExtendedRecord, error) {
	var response struct {
		Status string                     `json:"status"`
		Data   models.AwardExtendedRecord `json:"data"`
	}

	path := fmt.Sprintf("/awards/%d/extended", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetAwardCategory returns a single award category base record.
func (c *Client) GetAwardCategory(ctx context.Context, id int64) (*models.AwardCategoryBaseRecord, error) {
	var response struct {
		Status string                         `json:"status"`
		Data   models.AwardCategoryBaseRecord `json:"data"`
	}

	path := fmt.Sprintf("/awards/categories/%d", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetAwardCategoryExtended returns a single award category extended record.
func (c *Client) GetAwardCategoryExtended(ctx context.Context, id int64) (*models.AwardCategoryExtendedRecord, error) {
	var response struct {
		Status string                             `json:"status"`
		Data   models.AwardCategoryExtendedRecord `json:"data"`
	}

	path := fmt.Sprintf("/awards/categories/%d/extended", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}
