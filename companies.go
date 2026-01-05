package tvdb

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/mxrkw/tvdb/models"
)

// CompaniesResponse represents a paginated companies response.
type CompaniesResponse struct {
	Data  []models.Company `json:"data"`
	Links *models.Links    `json:"links,omitempty"`
}

// GetAllCompanies returns a paginated list of company records.
func (c *Client) GetAllCompanies(ctx context.Context, page *int) (CompaniesResponse, error) {
	var response struct {
		Status string           `json:"status"`
		Data   []models.Company `json:"data"`
		Links  *models.Links    `json:"links,omitempty"`
	}

	query := url.Values{}
	if page != nil {
		query.Set("page", strconv.Itoa(*page))
	}

	if err := c.get(ctx, "/companies", query, &response); err != nil {
		return CompaniesResponse{}, err
	}

	return CompaniesResponse{
		Data:  response.Data,
		Links: response.Links,
	}, nil
}

// GetCompany returns a company record.
func (c *Client) GetCompany(ctx context.Context, id int64) (*models.Company, error) {
	var response struct {
		Status string         `json:"status"`
		Data   models.Company `json:"data"`
	}

	path := fmt.Sprintf("/companies/%d", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetCompanyTypes returns all company type records.
func (c *Client) GetCompanyTypes(ctx context.Context) ([]models.CompanyType, error) {
	var response struct {
		Status string               `json:"status"`
		Data   []models.CompanyType `json:"data"`
	}

	if err := c.get(ctx, "/companies/types", nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
