package models

// MovieBaseRecord represents a base movie record.
type MovieBaseRecord struct {
	Aliases              []Alias  `json:"aliases,omitempty"`
	ID                   int64    `json:"id,omitempty"`
	Image                string   `json:"image,omitempty"`
	LastUpdated          string   `json:"lastUpdated,omitempty"`
	Name                 string   `json:"name,omitempty"`
	NameTranslations     []string `json:"nameTranslations,omitempty"`
	OverviewTranslations []string `json:"overviewTranslations,omitempty"`
	Score                float64  `json:"score,omitempty"`
	Slug                 string   `json:"slug,omitempty"`
	Status               *Status  `json:"status,omitempty"`
	Runtime              *int     `json:"runtime,omitempty"`
	Year                 string   `json:"year,omitempty"`
}

// MovieExtendedRecord represents an extended movie record.
type MovieExtendedRecord struct {
	Aliases              []Alias              `json:"aliases,omitempty"`
	Artworks             []ArtworkBaseRecord  `json:"artworks,omitempty"`
	AudioLanguages       []string             `json:"audioLanguages,omitempty"`
	Awards               []AwardBaseRecord    `json:"awards,omitempty"`
	BoxOffice            string               `json:"boxOffice,omitempty"`
	BoxOfficeUS          string               `json:"boxOfficeUS,omitempty"`
	Budget               string               `json:"budget,omitempty"`
	Characters           []Character          `json:"characters,omitempty"`
	Companies            *Companies           `json:"companies,omitempty"`
	ContentRatings       []ContentRating      `json:"contentRatings,omitempty"`
	FirstRelease         *Release             `json:"first_release,omitempty"`
	Genres               []GenreBaseRecord    `json:"genres,omitempty"`
	ID                   int64                `json:"id,omitempty"`
	Image                string               `json:"image,omitempty"`
	Inspirations         []Inspiration        `json:"inspirations,omitempty"`
	LastUpdated          string               `json:"lastUpdated,omitempty"`
	Lists                []ListBaseRecord     `json:"lists,omitempty"`
	Name                 string               `json:"name,omitempty"`
	NameTranslations     []string             `json:"nameTranslations,omitempty"`
	OriginalCountry      string               `json:"originalCountry,omitempty"`
	OriginalLanguage     string               `json:"originalLanguage,omitempty"`
	OverviewTranslations []string             `json:"overviewTranslations,omitempty"`
	ProductionCountries  []ProductionCountry  `json:"production_countries,omitempty"`
	Releases             []Release            `json:"releases,omitempty"`
	RemoteIDs            []RemoteID           `json:"remoteIds,omitempty"`
	Runtime              *int                 `json:"runtime,omitempty"`
	Score                float64              `json:"score,omitempty"`
	Slug                 string               `json:"slug,omitempty"`
	SpokenLanguages      []string             `json:"spoken_languages,omitempty"`
	Status               *Status              `json:"status,omitempty"`
	Studios              []StudioBaseRecord   `json:"studios,omitempty"`
	SubtitleLanguages    []string             `json:"subtitleLanguages,omitempty"`
	TagOptions           []TagOption          `json:"tagOptions,omitempty"`
	Trailers             []Trailer            `json:"trailers,omitempty"`
	Translations         *TranslationExtended `json:"translations,omitempty"`
	Year                 string               `json:"year,omitempty"`
}

// Release represents a release record.
type Release struct {
	Country string `json:"country,omitempty"`
	Date    string `json:"date,omitempty"`
	Detail  string `json:"detail,omitempty"`
}

// ProductionCountry represents a production country record.
type ProductionCountry struct {
	ID      int64  `json:"id,omitempty"`
	Country string `json:"country,omitempty"`
	Name    string `json:"name,omitempty"`
}

// Inspiration represents a movie inspiration record.
type Inspiration struct {
	ID       int64  `json:"id,omitempty"`
	Type     string `json:"type,omitempty"`
	TypeName string `json:"type_name,omitempty"`
	URL      string `json:"url,omitempty"`
}
