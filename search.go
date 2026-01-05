package tvdb

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/mxrkw/tvdb/models"
)

// SearchOptions are options for searching.
type SearchOptions struct {
	Query       string // primary search string
	Q           string // alias for Query (deprecated)
	Type        string // "movie", "series", "person", or "company"
	Year        *int   // restrict to specific year
	Company     string // restrict to specific company
	Country     string // 3-character country code
	Director    string // full director name
	Language    string // 3-character language code
	PrimaryType string // type of company (e.g., "Production Company")
	Network     string // network name
	RemoteID    string // IMDB or EIDR ID
	Offset      *int   // offset results
	Limit       *int   // limit results
}

// SearchResponse represents a search response.
type SearchResponse struct {
	Data  []models.SearchResult `json:"data"`
	Links *models.Links         `json:"links,omitempty"`
}

// Search performs a search across series, movies, people, and companies.
func (c *Client) Search(ctx context.Context, opts SearchOptions) (SearchResponse, error) {
	var response struct {
		Status string                `json:"status"`
		Data   []models.SearchResult `json:"data"`
		Links  *models.Links         `json:"links,omitempty"`
	}

	query := url.Values{}
	if opts.Query != "" {
		query.Set("query", opts.Query)
	}
	if opts.Q != "" {
		query.Set("q", opts.Q)
	}
	if opts.Type != "" {
		query.Set("type", opts.Type)
	}
	if opts.Year != nil {
		query.Set("year", strconv.Itoa(*opts.Year))
	}
	if opts.Company != "" {
		query.Set("company", opts.Company)
	}
	if opts.Country != "" {
		query.Set("country", opts.Country)
	}
	if opts.Director != "" {
		query.Set("director", opts.Director)
	}
	if opts.Language != "" {
		query.Set("language", opts.Language)
	}
	if opts.PrimaryType != "" {
		query.Set("primaryType", opts.PrimaryType)
	}
	if opts.Network != "" {
		query.Set("network", opts.Network)
	}
	if opts.RemoteID != "" {
		query.Set("remote_id", opts.RemoteID)
	}
	if opts.Offset != nil {
		query.Set("offset", strconv.Itoa(*opts.Offset))
	}
	if opts.Limit != nil {
		query.Set("limit", strconv.Itoa(*opts.Limit))
	}

	if err := c.get(ctx, "/search", query, &response); err != nil {
		return SearchResponse{}, err
	}

	return SearchResponse{
		Data:  response.Data,
		Links: response.Links,
	}, nil
}

// SearchByRemoteId searches by a specific remote ID (e.g., IMDB, EIDR).
func (c *Client) SearchByRemoteId(ctx context.Context, remoteId string) ([]models.SearchByRemoteIdResult, error) {
	var response struct {
		Status string                          `json:"status"`
		Data   []models.SearchByRemoteIdResult `json:"data"`
	}

	path := fmt.Sprintf("/search/remoteid/%s", url.PathEscape(remoteId))
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
