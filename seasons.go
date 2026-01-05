package tvdb

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/mxrkw/tvdb/models"
)

// SeasonsResponse represents a paginated seasons response.
type SeasonsResponse struct {
	Data  []models.SeasonBaseRecord `json:"data"`
	Links *models.Links             `json:"links,omitempty"`
}

// GetAllSeasons returns a paginated list of season base records.
func (c *Client) GetAllSeasons(ctx context.Context, page *int) (SeasonsResponse, error) {
	var response struct {
		Status string                    `json:"status"`
		Data   []models.SeasonBaseRecord `json:"data"`
		Links  *models.Links             `json:"links,omitempty"`
	}

	query := url.Values{}
	if page != nil {
		query.Set("page", strconv.Itoa(*page))
	}

	if err := c.get(ctx, "/seasons", query, &response); err != nil {
		return SeasonsResponse{}, err
	}

	return SeasonsResponse{
		Data:  response.Data,
		Links: response.Links,
	}, nil
}

// GetSeasonBase returns a season base record.
func (c *Client) GetSeasonBase(ctx context.Context, id int64) (*models.SeasonBaseRecord, error) {
	var response struct {
		Status string                  `json:"status"`
		Data   models.SeasonBaseRecord `json:"data"`
	}

	path := fmt.Sprintf("/seasons/%d", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetSeasonExtended returns a season extended record.
func (c *Client) GetSeasonExtended(ctx context.Context, id int64) (*models.SeasonExtendedRecord, error) {
	var response struct {
		Status string                      `json:"status"`
		Data   models.SeasonExtendedRecord `json:"data"`
	}

	path := fmt.Sprintf("/seasons/%d/extended", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetSeasonTypes returns season type records.
func (c *Client) GetSeasonTypes(ctx context.Context) ([]models.SeasonType, error) {
	var response struct {
		Status string              `json:"status"`
		Data   []models.SeasonType `json:"data"`
	}

	if err := c.get(ctx, "/seasons/types", nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetSeasonTranslation returns a season translation record.
func (c *Client) GetSeasonTranslation(ctx context.Context, id int64, language string) (*models.Translation, error) {
	var response struct {
		Status string             `json:"status"`
		Data   models.Translation `json:"data"`
	}

	path := fmt.Sprintf("/seasons/%d/translations/%s", id, url.PathEscape(language))
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}
