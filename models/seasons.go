package models

// SeasonBaseRecord represents a base season record.
type SeasonBaseRecord struct {
	ID                   int         `json:"id,omitempty"`
	Image                string      `json:"image,omitempty"`
	ImageType            int         `json:"imageType,omitempty"`
	LastUpdated          string      `json:"lastUpdated,omitempty"`
	Name                 string      `json:"name,omitempty"`
	NameTranslations     []string    `json:"nameTranslations,omitempty"`
	Number               int64       `json:"number,omitempty"`
	OverviewTranslations []string    `json:"overviewTranslations,omitempty"`
	Companies            *Companies  `json:"companies,omitempty"`
	SeriesID             int64       `json:"seriesId,omitempty"`
	Type                 *SeasonType `json:"type,omitempty"`
	Year                 string      `json:"year,omitempty"`
}

// SeasonExtendedRecord represents an extended season record.
type SeasonExtendedRecord struct {
	Artwork              []ArtworkBaseRecord `json:"artwork,omitempty"`
	Companies            *Companies          `json:"companies,omitempty"`
	Episodes             []EpisodeBaseRecord `json:"episodes,omitempty"`
	ID                   int                 `json:"id,omitempty"`
	Image                string              `json:"image,omitempty"`
	ImageType            int                 `json:"imageType,omitempty"`
	LastUpdated          string              `json:"lastUpdated,omitempty"`
	Name                 string              `json:"name,omitempty"`
	NameTranslations     []string            `json:"nameTranslations,omitempty"`
	Number               int64               `json:"number,omitempty"`
	OverviewTranslations []string            `json:"overviewTranslations,omitempty"`
	SeriesID             int64               `json:"seriesId,omitempty"`
	Trailers             []Trailer           `json:"trailers,omitempty"`
	Type                 *SeasonType         `json:"type,omitempty"`
	TagOptions           []TagOption         `json:"tagOptions,omitempty"`
	Translations         []Translation       `json:"translations,omitempty"`
	Year                 string              `json:"year,omitempty"`
}

// SeasonType represents a season type record.
type SeasonType struct {
	AlternateName string `json:"alternateName,omitempty"`
	ID            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Type          string `json:"type,omitempty"`
}
