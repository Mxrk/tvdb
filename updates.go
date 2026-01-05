package tvdb

import (
	"context"
	"net/url"
	"strconv"

	"github.com/mxrkw/tvdb/models"
)

// UpdatesOptions are options for getting updates.
type UpdatesOptions struct {
	Since  int64  // required - Unix timestamp to get updates since
	Type   string // entity type filter (e.g., "movies", "series", "episodes")
	Action string // "delete" or "update"
	Page   *int   // page number
}

// UpdatesResponse represents an updates response.
type UpdatesResponse struct {
	Data  []models.EntityUpdate `json:"data"`
	Links *models.Links         `json:"links,omitempty"`
}

// GetUpdates returns updated entities since the specified timestamp.
// methodInt in results: 1 = created, 2 = updated, 3 = deleted
func (c *Client) GetUpdates(ctx context.Context, opts UpdatesOptions) (UpdatesResponse, error) {
	var response struct {
		Status string                `json:"status"`
		Data   []models.EntityUpdate `json:"data"`
		Links  *models.Links         `json:"links,omitempty"`
	}

	query := url.Values{}
	query.Set("since", strconv.FormatInt(opts.Since, 10))

	if opts.Type != "" {
		query.Set("type", opts.Type)
	}
	if opts.Action != "" {
		query.Set("action", opts.Action)
	}
	if opts.Page != nil {
		query.Set("page", strconv.Itoa(*opts.Page))
	}

	if err := c.get(ctx, "/updates", query, &response); err != nil {
		return UpdatesResponse{}, err
	}

	return UpdatesResponse{
		Data:  response.Data,
		Links: response.Links,
	}, nil
}
