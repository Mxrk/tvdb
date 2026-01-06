package tvdb

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/mxrk/tvdb/models"
)

// EpisodesResponse represents a paginated episodes response.
type EpisodesResponse struct {
	Data  []models.EpisodeBaseRecord `json:"data"`
	Links *models.Links              `json:"links,omitempty"`
}

// GetAllEpisodes returns a paginated list of episode base records.
func (c *Client) GetAllEpisodes(ctx context.Context, page *int) (EpisodesResponse, error) {
	var response struct {
		Status string                     `json:"status"`
		Data   []models.EpisodeBaseRecord `json:"data"`
		Links  *models.Links              `json:"links,omitempty"`
	}

	query := url.Values{}
	if page != nil {
		query.Set("page", strconv.Itoa(*page))
	}

	if err := c.get(ctx, "/episodes", query, &response); err != nil {
		return EpisodesResponse{}, err
	}

	return EpisodesResponse{
		Data:  response.Data,
		Links: response.Links,
	}, nil
}

// GetEpisodeBase returns an episode base record.
func (c *Client) GetEpisodeBase(ctx context.Context, id int64) (*models.EpisodeBaseRecord, error) {
	var response struct {
		Status string                   `json:"status"`
		Data   models.EpisodeBaseRecord `json:"data"`
	}

	path := fmt.Sprintf("/episodes/%d", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// EpisodeExtendedOptions are options for getting an extended episode record.
type EpisodeExtendedOptions struct {
	Meta string // "translations" to include translations
}

// GetEpisodeExtended returns an episode extended record.
func (c *Client) GetEpisodeExtended(ctx context.Context, id int64, opts *EpisodeExtendedOptions) (*models.EpisodeExtendedRecord, error) {
	var response struct {
		Status string                       `json:"status"`
		Data   models.EpisodeExtendedRecord `json:"data"`
	}

	query := url.Values{}
	if opts != nil && opts.Meta != "" {
		query.Set("meta", opts.Meta)
	}

	path := fmt.Sprintf("/episodes/%d/extended", id)
	if err := c.get(ctx, path, query, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetEpisodeTranslation returns an episode translation record.
func (c *Client) GetEpisodeTranslation(ctx context.Context, id int64, language string) (*models.Translation, error) {
	var response struct {
		Status string             `json:"status"`
		Data   models.Translation `json:"data"`
	}

	path := fmt.Sprintf("/episodes/%d/translations/%s", id, url.PathEscape(language))
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}
