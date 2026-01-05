package models

// Character represents a character record.
type Character struct {
	Aliases              []Alias     `json:"aliases,omitempty"`
	Episode              *RecordInfo `json:"episode,omitempty"`
	EpisodeId            *int        `json:"episodeId,omitempty"`
	ID                   int64       `json:"id,omitempty"`
	Image                string      `json:"image,omitempty"`
	IsFeatured           bool        `json:"isFeatured,omitempty"`
	MovieId              *int        `json:"movieId,omitempty"`
	Movie                *RecordInfo `json:"movie,omitempty"`
	Name                 string      `json:"name,omitempty"`
	NameTranslations     []string    `json:"nameTranslations,omitempty"`
	OverviewTranslations []string    `json:"overviewTranslations,omitempty"`
	PeopleId             int         `json:"peopleId,omitempty"`
	PersonImgURL         string      `json:"personImgURL,omitempty"`
	PeopleType           string      `json:"peopleType,omitempty"`
	SeriesId             *int        `json:"seriesId,omitempty"`
	Series               *RecordInfo `json:"series,omitempty"`
	Sort                 int64       `json:"sort,omitempty"`
	TagOptions           []TagOption `json:"tagOptions,omitempty"`
	Type                 int64       `json:"type,omitempty"`
	URL                  string      `json:"url,omitempty"`
	PersonName           string      `json:"personName,omitempty"`
}
