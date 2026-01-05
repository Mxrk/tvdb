package tvdb

import (
	"context"
	"fmt"

	"github.com/mxrkw/tvdb/models"
)

// GetCharacterBase returns a character base record.
func (c *Client) GetCharacterBase(ctx context.Context, id int64) (*models.Character, error) {
	var response struct {
		Status string           `json:"status"`
		Data   models.Character `json:"data"`
	}

	path := fmt.Sprintf("/characters/%d", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}
