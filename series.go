package tvdb

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/mxrkw/tvdb/models"
)

// SeriesResponse represents a paginated series response.
type SeriesResponse struct {
	Data  []models.SeriesBaseRecord `json:"data"`
	Links *models.Links             `json:"links,omitempty"`
}

// GetAllSeries returns a paginated list of series base records.
func (c *Client) GetAllSeries(ctx context.Context, page *int) (SeriesResponse, error) {
	var response struct {
		Status string                    `json:"status"`
		Data   []models.SeriesBaseRecord `json:"data"`
		Links  *models.Links             `json:"links,omitempty"`
	}

	query := url.Values{}
	if page != nil {
		query.Set("page", strconv.Itoa(*page))
	}

	if err := c.get(ctx, "/series", query, &response); err != nil {
		return SeriesResponse{}, err
	}

	return SeriesResponse{
		Data:  response.Data,
		Links: response.Links,
	}, nil
}

// GetSeriesBase returns a series base record.
func (c *Client) GetSeriesBase(ctx context.Context, id int64) (*models.SeriesBaseRecord, error) {
	var response struct {
		Status string                  `json:"status"`
		Data   models.SeriesBaseRecord `json:"data"`
	}

	path := fmt.Sprintf("/series/%d", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetSeriesBaseBySlug returns a series base record by slug.
func (c *Client) GetSeriesBaseBySlug(ctx context.Context, slug string) (*models.SeriesBaseRecord, error) {
	var response struct {
		Status string                  `json:"status"`
		Data   models.SeriesBaseRecord `json:"data"`
	}

	path := fmt.Sprintf("/series/slug/%s", url.PathEscape(slug))
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// SeriesExtendedOptions are options for getting an extended series record.
type SeriesExtendedOptions struct {
	Meta  string // "translations" or "episodes" to include
	Short bool   // reduce payload without characters and artworks
}

// GetSeriesExtended returns a series extended record.
func (c *Client) GetSeriesExtended(ctx context.Context, id int64, opts *SeriesExtendedOptions) (*models.SeriesExtendedRecord, error) {
	var response struct {
		Status string                      `json:"status"`
		Data   models.SeriesExtendedRecord `json:"data"`
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

	path := fmt.Sprintf("/series/%d/extended", id)
	if err := c.get(ctx, path, query, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// SeriesArtworksOptions are options for getting series artworks.
type SeriesArtworksOptions struct {
	Lang string // language code (e.g., "eng", "spa")
	Type *int   // artwork type ID
}

// GetSeriesArtworks returns series artworks based on language and type.
func (c *Client) GetSeriesArtworks(ctx context.Context, id int64, opts *SeriesArtworksOptions) (*models.SeriesExtendedRecord, error) {
	var response struct {
		Status string                      `json:"status"`
		Data   models.SeriesExtendedRecord `json:"data"`
	}

	query := url.Values{}
	if opts != nil {
		if opts.Lang != "" {
			query.Set("lang", opts.Lang)
		}
		if opts.Type != nil {
			query.Set("type", strconv.Itoa(*opts.Type))
		}
	}

	path := fmt.Sprintf("/series/%d/artworks", id)
	if err := c.get(ctx, path, query, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetSeriesNextAired returns series base record including the nextAired field.
func (c *Client) GetSeriesNextAired(ctx context.Context, id int64) (*models.SeriesBaseRecord, error) {
	var response struct {
		Status string                  `json:"status"`
		Data   models.SeriesBaseRecord `json:"data"`
	}

	path := fmt.Sprintf("/series/%d/nextAired", id)
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// SeriesEpisodesOptions are options for getting series episodes.
type SeriesEpisodesOptions struct {
	Page          int    // required - page number (0-indexed)
	Season        *int   // filter by season number
	EpisodeNumber *int   // filter by episode number (requires Season)
	AirDate       string // filter by air date (yyyy-mm-dd)
}

// GetSeriesEpisodes returns series episodes from the specified season type.
// seasonType can be: "default", "official", "dvd", "absolute", "alternate", "regional"
func (c *Client) GetSeriesEpisodes(ctx context.Context, id int64, seasonType string, opts SeriesEpisodesOptions) (*models.SeriesEpisodesResponse, error) {
	var response struct {
		Status string                        `json:"status"`
		Data   models.SeriesEpisodesResponse `json:"data"`
	}

	query := url.Values{}
	query.Set("page", strconv.Itoa(opts.Page))
	if opts.Season != nil {
		query.Set("season", strconv.Itoa(*opts.Season))
	}
	if opts.EpisodeNumber != nil {
		query.Set("episodeNumber", strconv.Itoa(*opts.EpisodeNumber))
	}
	if opts.AirDate != "" {
		query.Set("airDate", opts.AirDate)
	}

	path := fmt.Sprintf("/series/%d/episodes/%s", id, url.PathEscape(seasonType))
	if err := c.get(ctx, path, query, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// GetSeriesSeasonEpisodesTranslated returns series episodes with translations.
func (c *Client) GetSeriesSeasonEpisodesTranslated(ctx context.Context, id int64, seasonType, lang string, page int) (*models.SeriesBaseRecord, error) {
	var response struct {
		Status string `json:"status"`
		Data   struct {
			Series *models.SeriesBaseRecord `json:"series"`
		} `json:"data"`
	}

	query := url.Values{}
	query.Set("page", strconv.Itoa(page))

	path := fmt.Sprintf("/series/%d/episodes/%s/%s", id, url.PathEscape(seasonType), url.PathEscape(lang))
	if err := c.get(ctx, path, query, &response); err != nil {
		return nil, err
	}

	return response.Data.Series, nil
}

// GetSeriesTranslation returns a series translation record.
func (c *Client) GetSeriesTranslation(ctx context.Context, id int64, language string) (*models.Translation, error) {
	var response struct {
		Status string             `json:"status"`
		Data   models.Translation `json:"data"`
	}

	path := fmt.Sprintf("/series/%d/translations/%s", id, url.PathEscape(language))
	if err := c.get(ctx, path, nil, &response); err != nil {
		return nil, err
	}

	return &response.Data, nil
}

// SeriesFilterOptions are options for filtering series.
type SeriesFilterOptions struct {
	Company       *int   // production company ID
	ContentRating *int   // content rating ID
	Country       string // required - country of origin (e.g., "usa")
	Genre         *int   // genre ID
	Lang          string // required - original language (e.g., "eng")
	Sort          string // "score", "firstAired", "lastAired", or "name"
	SortType      string // "asc" or "desc"
	Status        *int   // status ID (1, 2, or 3)
	Year          *int   // release year
}

// GetSeriesFilter searches series based on filter parameters.
func (c *Client) GetSeriesFilter(ctx context.Context, opts SeriesFilterOptions) ([]models.SeriesBaseRecord, error) {
	if opts.Country == "" {
		return nil, fmt.Errorf("tvdb: Country is required for series filter")
	}
	if opts.Lang == "" {
		return nil, fmt.Errorf("tvdb: Lang is required for series filter")
	}

	var response struct {
		Status string                    `json:"status"`
		Data   []models.SeriesBaseRecord `json:"data"`
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
	if opts.SortType != "" {
		query.Set("sortType", opts.SortType)
	}
	if opts.Status != nil {
		query.Set("status", strconv.Itoa(*opts.Status))
	}
	if opts.Year != nil {
		query.Set("year", strconv.Itoa(*opts.Year))
	}

	if err := c.get(ctx, "/series/filter", query, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}

// GetAllSeriesStatuses returns a list of series status records.
func (c *Client) GetAllSeriesStatuses(ctx context.Context) ([]models.Status, error) {
	var response struct {
		Status string          `json:"status"`
		Data   []models.Status `json:"data"`
	}

	if err := c.get(ctx, "/series/statuses", nil, &response); err != nil {
		return nil, err
	}

	return response.Data, nil
}
