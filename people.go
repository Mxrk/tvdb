package tvdb

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/mxrk/tvdb/models"
)

// PeopleResponse represents a paginated people response.
type PeopleResponse struct {
	Data  []models.PeopleBaseRecord `json:"data"`
	Links *models.Links             `json:"links,omitempty"`
}

// GetAllPeople returns a paginated list of people base records.
func (c *Client) GetAllPeople(ctx context.Context, page *int) (PeopleResponse, error) {
	var response struct {
		Status string                    `json:"status"`
		Data   []models.PeopleBaseRecord `json:"data"`
		Links  *models.Links             `json:"links,omitempty"`
	}

	query := url.Values{}
	if page != nil {
		query.Set("page", strconv.Itoa(*page))
	}

	if err := c.get(ctx, "/people", query, &response); err != nil {
		return PeopleResponse{}, err
	}

	return PeopleResponse{
		Data:  response.Data,
		Links: response.Links,
	}, nil
}

// GetPeopleBase returns a people base record.
func (c *Client) GetPeopleBase(ctx context.Context, id int64) (*models.PeopleBaseRecord, error) {
	var response struct {
		Status string                  `json:"status"`
		Data   models.PeopleBaseRecord `json:"data"`
	}

	path := fmt.Sprintf("/people/%d", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// PeopleExtendedOptions are options for getting an extended people record.
type PeopleExtendedOptions struct {
	Meta string // "translations" to include translations
}

// GetPeopleExtended returns a people extended record.
func (c *Client) GetPeopleExtended(ctx context.Context, id int64, opts *PeopleExtendedOptions) (*models.PeopleExtendedRecord, error) {
	var response struct {
		Status string                      `json:"status"`
		Data   models.PeopleExtendedRecord `json:"data"`
	}

	query := url.Values{}
	if opts != nil && opts.Meta != "" {
		query.Set("meta", opts.Meta)
	}

	path := fmt.Sprintf("/people/%d/extended", id)
	if err := c.get(ctx, path, query, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetPeopleTranslation returns a people translation record.
func (c *Client) GetPeopleTranslation(ctx context.Context, id int64, language string) (*models.Translation, error) {
	var response struct {
		Status string             `json:"status"`
		Data   models.Translation `json:"data"`
	}

	path := fmt.Sprintf("/people/%d/translations/%s", id, url.PathEscape(language))
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetAllPeopleTypes returns a list of people type records.
func (c *Client) GetAllPeopleTypes(ctx context.Context) ([]models.PeopleType, error) {
	var response struct {
		Status string              `json:"status"`
		Data   []models.PeopleType `json:"data"`
	}

	if err := c.get(ctx, "/people/types", nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
