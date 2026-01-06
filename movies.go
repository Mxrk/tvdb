package tvdb

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/mxrk/tvdb/models"
)

// MoviesResponse represents a paginated movies response.
type MoviesResponse struct {
	Data  []models.MovieBaseRecord `json:"data"`
	Links *models.Links            `json:"links,omitempty"`
}

// GetAllMovies returns a paginated list of movie base records.
func (c *Client) GetAllMovies(ctx context.Context, page *int) (MoviesResponse, error) {
	var response struct {
		Status string                   `json:"status"`
		Data   []models.MovieBaseRecord `json:"data"`
		Links  *models.Links            `json:"links,omitempty"`
	}

	query := url.Values{}
	if page != nil {
		query.Set("page", strconv.Itoa(*page))
	}

	if err := c.get(ctx, "/movies", query, &response); err != nil {
		return MoviesResponse{}, err
	}

	return MoviesResponse{
		Data:  response.Data,
		Links: response.Links,
	}, nil
}

// GetMovieBase returns a movie base record.
func (c *Client) GetMovieBase(ctx context.Context, id int64) (*models.MovieBaseRecord, error) {
	var response struct {
		Status string                 `json:"status"`
		Data   models.MovieBaseRecord `json:"data"`
	}

	path := fmt.Sprintf("/movies/%d", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// MovieExtendedOptions are options for getting an extended movie record.
type MovieExtendedOptions struct {
	Meta  string // "translations" to include translations
	Short bool   // reduce payload without characters, artworks, and trailers
}

// GetMovieExtended returns a movie extended record.
func (c *Client) GetMovieExtended(ctx context.Context, id int64, opts *MovieExtendedOptions) (*models.MovieExtendedRecord, error) {
	var response struct {
		Status string                     `json:"status"`
		Data   models.MovieExtendedRecord `json:"data"`
	}

	query := url.Values{}
	if opts != nil {
		if opts.Meta != "" {
			query.Set("meta", opts.Meta)
		}
		if opts.Short {
			query.Set("short", "true")
		}
	}

	path := fmt.Sprintf("/movies/%d/extended", id)
	if err := c.get(ctx, path, query, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetMovieBaseBySlug returns a movie base record by slug.
func (c *Client) GetMovieBaseBySlug(ctx context.Context, slug string) (*models.MovieBaseRecord, error) {
	var response struct {
		Status string                 `json:"status"`
		Data   models.MovieBaseRecord `json:"data"`
	}

	path := fmt.Sprintf("/movies/slug/%s", url.PathEscape(slug))
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetMovieTranslation returns a movie translation record.
func (c *Client) GetMovieTranslation(ctx context.Context, id int64, language string) (*models.Translation, error) {
	var response struct {
		Status string             `json:"status"`
		Data   models.Translation `json:"data"`
	}

	path := fmt.Sprintf("/movies/%d/translations/%s", id, url.PathEscape(language))
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// MovieFilterOptions are options for filtering movies.
type MovieFilterOptions struct {
	Company       *int   // production company ID
	ContentRating *int   // content rating ID
	Country       string // required - country of origin (e.g., "usa")
	Genre         *int   // genre ID
	Lang          string // required - original language (e.g., "eng")
	Sort          string // "score", "firstAired", or "name"
	Status        *int   // status ID (1, 2, or 3)
	Year          *int   // release year
}

// GetMoviesFilter searches movies based on filter parameters.
func (c *Client) GetMoviesFilter(ctx context.Context, opts MovieFilterOptions) ([]models.MovieBaseRecord, error) {
	if opts.Country == "" {
		return nil, fmt.Errorf("tvdb: Country is required for movies filter")
	}
	if opts.Lang == "" {
		return nil, fmt.Errorf("tvdb: Lang is required for movies filter")
	}

	var response struct {
		Status string                   `json:"status"`
		Data   []models.MovieBaseRecord `json:"data"`
	}

	query := url.Values{}
	query.Set("country", opts.Country)
	query.Set("lang", opts.Lang)

	if opts.Company != nil {
		query.Set("company", strconv.Itoa(*opts.Company))
	}
	if opts.ContentRating != nil {
		query.Set("contentRating", strconv.Itoa(*opts.ContentRating))
	}
	if opts.Genre != nil {
		query.Set("genre", strconv.Itoa(*opts.Genre))
	}
	if opts.Sort != "" {
		query.Set("sort", opts.Sort)
	}
	if opts.Status != nil {
		query.Set("status", strconv.Itoa(*opts.Status))
	}
	if opts.Year != nil {
		query.Set("year", strconv.Itoa(*opts.Year))
	}

	if err := c.get(ctx, "/movies/filter", query, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetAllMovieStatuses returns a list of movie status records.
func (c *Client) GetAllMovieStatuses(ctx context.Context) ([]models.Status, error) {
	var response struct {
		Status string          `json:"status"`
		Data   []models.Status `json:"data"`
	}

	if err := c.get(ctx, "/movies/statuses", nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
