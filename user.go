package tvdb

import (
	"context"
	"fmt"

	"github.com/mxrkw/tvdb/models"
)

// GetUserInfo returns the current user's info.
func (c *Client) GetUserInfo(ctx context.Context) (*models.UserInfo, error) {
	var response struct {
		Status string          `json:"status"`
		Data   models.UserInfo `json:"data"`
	}

	if err := c.get(ctx, "/user", nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetUserInfoById returns user info by user ID.
func (c *Client) GetUserInfoById(ctx context.Context, id int64) (*models.UserInfo, error) {
	var response struct {
		Status string          `json:"status"`
		Data   models.UserInfo `json:"data"`
	}

	path := fmt.Sprintf("/user/%d", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetUserFavorites returns the current user's favorites.
func (c *Client) GetUserFavorites(ctx context.Context) (*models.Favorites, error) {
	var response struct {
		Status string           `json:"status"`
		Data   models.Favorites `json:"data"`
	}

	if err := c.get(ctx, "/user/favorites", nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// CreateUserFavorite creates a new user favorite.
func (c *Client) CreateUserFavorite(ctx context.Context, favorite models.FavoriteRecord) error {
	return c.post(ctx, "/user/favorites", favorite, nil)
}
