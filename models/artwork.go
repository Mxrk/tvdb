package models

// ArtworkBaseRecord represents a base artwork record.
type ArtworkBaseRecord struct {
	Height       int64   `json:"height,omitempty"`
	ID           int     `json:"id,omitempty"`
	Image        string  `json:"image,omitempty"`
	IncludesText bool    `json:"includesText,omitempty"`
	Language     string  `json:"language,omitempty"`
	Score        float64 `json:"score,omitempty"`
	Thumbnail    string  `json:"thumbnail,omitempty"`
	Type         int64   `json:"type,omitempty"`
	Width        int64   `json:"width,omitempty"`
}

// ArtworkExtendedRecord represents an extended artwork record.
type ArtworkExtendedRecord struct {
	EpisodeId       int            `json:"episodeId,omitempty"`
	Height          int64          `json:"height,omitempty"`
	ID              int64          `json:"id,omitempty"`
	Image           string         `json:"image,omitempty"`
	IncludesText    bool           `json:"includesText,omitempty"`
	Language        string         `json:"language,omitempty"`
	MovieId         int            `json:"movieId,omitempty"`
	NetworkId       int            `json:"networkId,omitempty"`
	PeopleId        int            `json:"peopleId,omitempty"`
	Score           float64        `json:"score,omitempty"`
	SeasonId        int            `json:"seasonId,omitempty"`
	SeriesId        int            `json:"seriesId,omitempty"`
	SeriesPeopleId  int            `json:"seriesPeopleId,omitempty"`
	Status          *ArtworkStatus `json:"status,omitempty"`
	TagOptions      []TagOption    `json:"tagOptions,omitempty"`
	Thumbnail       string         `json:"thumbnail,omitempty"`
	ThumbnailHeight int64          `json:"thumbnailHeight,omitempty"`
	ThumbnailWidth  int64          `json:"thumbnailWidth,omitempty"`
	Type            int64          `json:"type,omitempty"`
	UpdatedAt       int64          `json:"updatedAt,omitempty"`
	Width           int64          `json:"width,omitempty"`
}

// ArtworkStatus represents an artwork status record.
type ArtworkStatus struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// ArtworkType represents an artwork type record.
type ArtworkType struct {
	Height      int64  `json:"height,omitempty"`
	ID          int64  `json:"id,omitempty"`
	ImageFormat string `json:"imageFormat,omitempty"`
	Name        string `json:"name,omitempty"`
	RecordType  string `json:"recordType,omitempty"`
	Slug        string `json:"slug,omitempty"`
	ThumbHeight int64  `json:"thumbHeight,omitempty"`
	ThumbWidth  int64  `json:"thumbWidth,omitempty"`
	Width       int64  `json:"width,omitempty"`
}
