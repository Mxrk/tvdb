package models

// ListBaseRecord represents a base list record.
type ListBaseRecord struct {
	Aliases              []Alias     `json:"aliases,omitempty"`
	ID                   int64       `json:"id,omitempty"`
	Image                string      `json:"image,omitempty"`
	ImageIsFallback      bool        `json:"imageIsFallback,omitempty"`
	IsOfficial           bool        `json:"isOfficial,omitempty"`
	Name                 string      `json:"name,omitempty"`
	NameTranslations     []string    `json:"nameTranslations,omitempty"`
	Overview             string      `json:"overview,omitempty"`
	OverviewTranslations []string    `json:"overviewTranslations,omitempty"`
	RemoteIDs            []RemoteID  `json:"remoteIds,omitempty"`
	Tags                 []TagOption `json:"tags,omitempty"`
	Score                int         `json:"score,omitempty"`
	URL                  string      `json:"url,omitempty"`
}

// ListExtendedRecord represents an extended list record.
type ListExtendedRecord struct {
	Aliases              []Alias  `json:"aliases,omitempty"`
	Entities             []Entity `json:"entities,omitempty"`
	ID                   int64    `json:"id,omitempty"`
	Image                string   `json:"image,omitempty"`
	ImageIsFallback      bool     `json:"imageIsFallback,omitempty"`
	IsOfficial           bool     `json:"isOfficial,omitempty"`
	Name                 string   `json:"name,omitempty"`
	NameTranslations     []string `json:"nameTranslations,omitempty"`
	Overview             string   `json:"overview,omitempty"`
	OverviewTranslations []string `json:"overviewTranslations,omitempty"`
	Score                int64    `json:"score,omitempty"`
	URL                  string   `json:"url,omitempty"`
}

// Entity represents an entity record in a list.
type Entity struct {
	MovieId  int   `json:"movieId,omitempty"`
	Order    int64 `json:"order,omitempty"`
	SeriesId int   `json:"seriesId,omitempty"`
}
