package tvdb

import (
	"context"
	"fmt"

	"github.com/mxrk/tvdb/models"
)

// GetAllGenres returns a list of genre records.
func (c *Client) GetAllGenres(ctx context.Context) ([]models.GenreBaseRecord, error) {
	var response struct {
		Status string                   `json:"status"`
		Data   []models.GenreBaseRecord `json:"data"`
	}

	if err := c.get(ctx, "/genres", nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetGenreBase returns a genre record.
func (c *Client) GetGenreBase(ctx context.Context, id int64) (*models.GenreBaseRecord, error) {
	var response struct {
		Status string                 `json:"status"`
		Data   models.GenreBaseRecord `json:"data"`
	}

	path := fmt.Sprintf("/genres/%d", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}
