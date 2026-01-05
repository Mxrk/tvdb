package tvdb

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/mxrkw/tvdb/models"
)

// ListsResponse represents a paginated lists response.
type ListsResponse struct {
	Data  []models.ListBaseRecord `json:"data"`
	Links *models.Links           `json:"links,omitempty"`
}

// GetAllLists returns a paginated list of list base records.
func (c *Client) GetAllLists(ctx context.Context, page *int) (ListsResponse, error) {
	var response struct {
		Status string                  `json:"status"`
		Data   []models.ListBaseRecord `json:"data"`
		Links  *models.Links           `json:"links,omitempty"`
	}

	query := url.Values{}
	if page != nil {
		query.Set("page", strconv.Itoa(*page))
	}

	if err := c.get(ctx, "/lists", query, &response); err != nil {
		return ListsResponse{}, err
	}

	return ListsResponse{
		Data:  response.Data,
		Links: response.Links,
	}, nil
}

// GetList returns a list base record.
func (c *Client) GetList(ctx context.Context, id int64) (*models.ListBaseRecord, error) {
	var response struct {
		Status string                `json:"status"`
		Data   models.ListBaseRecord `json:"data"`
	}

	path := fmt.Sprintf("/lists/%d", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetListBySlug returns a list base record by slug.
func (c *Client) GetListBySlug(ctx context.Context, slug string) (*models.ListBaseRecord, error) {
	var response struct {
		Status string                `json:"status"`
		Data   models.ListBaseRecord `json:"data"`
	}

	path := fmt.Sprintf("/lists/slug/%s", url.PathEscape(slug))
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetListExtended returns a list extended record.
func (c *Client) GetListExtended(ctx context.Context, id int64) (*models.ListExtendedRecord, error) {
	var response struct {
		Status string                    `json:"status"`
		Data   models.ListExtendedRecord `json:"data"`
	}

	path := fmt.Sprintf("/lists/%d/extended", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetListTranslation returns list translation records.
func (c *Client) GetListTranslation(ctx context.Context, id int64, language string) ([]models.Translation, error) {
	var response struct {
		Status string               `json:"status"`
		Data   []models.Translation `json:"data"`
	}

	path := fmt.Sprintf("/lists/%d/translations/%s", id, url.PathEscape(language))
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
