package tvdb

import (
	"context"

	"github.com/mxrkw/tvdb/models"
)

// GetAllLanguages returns a list of language records.
func (c *Client) GetAllLanguages(ctx context.Context) ([]models.Language, error) {
	var response struct {
		Status string            `json:"status"`
		Data   []models.Language `json:"data"`
	}

	if err := c.get(ctx, "/languages", nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetAllCountries returns a list of country records.
func (c *Client) GetAllCountries(ctx context.Context) ([]models.Country, error) {
	var response struct {
		Status string           `json:"status"`
		Data   []models.Country `json:"data"`
	}

	if err := c.get(ctx, "/countries", nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetAllGenders returns a list of gender records.
func (c *Client) GetAllGenders(ctx context.Context) ([]models.Gender, error) {
	var response struct {
		Status string          `json:"status"`
		Data   []models.Gender `json:"data"`
	}

	if err := c.get(ctx, "/genders", nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetAllContentRatings returns a list of content rating records.
func (c *Client) GetAllContentRatings(ctx context.Context) ([]models.ContentRating, error) {
	var response struct {
		Status string                 `json:"status"`
		Data   []models.ContentRating `json:"data"`
	}

	if err := c.get(ctx, "/content/ratings", nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetEntityTypes returns the active entity types.
func (c *Client) GetEntityTypes(ctx context.Context) ([]models.EntityType, error) {
	var response struct {
		Status string              `json:"status"`
		Data   []models.EntityType `json:"data"`
	}

	if err := c.get(ctx, "/entities", nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetAllInspirationTypes returns a list of inspiration type records.
func (c *Client) GetAllInspirationTypes(ctx context.Context) ([]models.InspirationType, error) {
	var response struct {
		Status string                   `json:"status"`
		Data   []models.InspirationType `json:"data"`
	}

	if err := c.get(ctx, "/inspiration/types", nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetAllSourceTypes returns a list of source type records.
func (c *Client) GetAllSourceTypes(ctx context.Context) ([]models.SourceType, error) {
	var response struct {
		Status string              `json:"status"`
		Data   []models.SourceType `json:"data"`
	}

	if err := c.get(ctx, "/sources/types", nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
