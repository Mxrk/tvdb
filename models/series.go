package models

// SeriesBaseRecord represents a base series record.
type SeriesBaseRecord struct {
	Aliases              []Alias             `json:"aliases,omitempty"`
	AverageRuntime       *int                `json:"averageRuntime,omitempty"`
	Country              string              `json:"country,omitempty"`
	DefaultSeasonType    int64               `json:"defaultSeasonType,omitempty"`
	Episodes             []EpisodeBaseRecord `json:"episodes,omitempty"`
	FirstAired           string              `json:"firstAired,omitempty"`
	ID                   int                 `json:"id,omitempty"`
	Image                string              `json:"image,omitempty"`
	IsOrderRandomized    bool                `json:"isOrderRandomized,omitempty"`
	LastAired            string              `json:"lastAired,omitempty"`
	LastUpdated          string              `json:"lastUpdated,omitempty"`
	Name                 string              `json:"name,omitempty"`
	NameTranslations     []string            `json:"nameTranslations,omitempty"`
	NextAired            string              `json:"nextAired,omitempty"`
	OriginalCountry      string              `json:"originalCountry,omitempty"`
	OriginalLanguage     string              `json:"originalLanguage,omitempty"`
	OverviewTranslations []string            `json:"overviewTranslations,omitempty"`
	Score                float64             `json:"score,omitempty"`
	Slug                 string              `json:"slug,omitempty"`
	Status               *Status             `json:"status,omitempty"`
	Year                 string              `json:"year,omitempty"`
}

// SeriesExtendedRecord represents an extended series record.
type SeriesExtendedRecord struct {
	Abbreviation         string                  `json:"abbreviation,omitempty"`
	AirsDays             *SeriesAirsDays         `json:"airsDays,omitempty"`
	AirsTime             string                  `json:"airsTime,omitempty"`
	Aliases              []Alias                 `json:"aliases,omitempty"`
	Artworks             []ArtworkExtendedRecord `json:"artworks,omitempty"`
	AverageRuntime       *int                    `json:"averageRuntime,omitempty"`
	Characters           []Character             `json:"characters,omitempty"`
	ContentRatings       []ContentRating         `json:"contentRatings,omitempty"`
	Country              string                  `json:"country,omitempty"`
	DefaultSeasonType    int64                   `json:"defaultSeasonType,omitempty"`
	Episodes             []EpisodeBaseRecord     `json:"episodes,omitempty"`
	FirstAired           string                  `json:"firstAired,omitempty"`
	Lists                []ListBaseRecord        `json:"lists,omitempty"`
	Genres               []GenreBaseRecord       `json:"genres,omitempty"`
	ID                   int                     `json:"id,omitempty"`
	Image                string                  `json:"image,omitempty"`
	IsOrderRandomized    bool                    `json:"isOrderRandomized,omitempty"`
	LastAired            string                  `json:"lastAired,omitempty"`
	LastUpdated          string                  `json:"lastUpdated,omitempty"`
	Name                 string                  `json:"name,omitempty"`
	NameTranslations     []string                `json:"nameTranslations,omitempty"`
	Companies            []Company               `json:"companies,omitempty"`
	NextAired            string                  `json:"nextAired,omitempty"`
	OriginalCountry      string                  `json:"originalCountry,omitempty"`
	OriginalLanguage     string                  `json:"originalLanguage,omitempty"`
	OriginalNetwork      *Company                `json:"originalNetwork,omitempty"`
	Overview             string                  `json:"overview,omitempty"`
	LatestNetwork        *Company                `json:"latestNetwork,omitempty"`
	OverviewTranslations []string                `json:"overviewTranslations,omitempty"`
	RemoteIDs            []RemoteID              `json:"remoteIds,omitempty"`
	Score                float64                 `json:"score,omitempty"`
	Seasons              []SeasonBaseRecord      `json:"seasons,omitempty"`
	SeasonTypes          []SeasonType            `json:"seasonTypes,omitempty"`
	Slug                 string                  `json:"slug,omitempty"`
	Status               *Status                 `json:"status,omitempty"`
	Tags                 []TagOption             `json:"tags,omitempty"`
	Trailers             []Trailer               `json:"trailers,omitempty"`
	Translations         *TranslationExtended    `json:"translations,omitempty"`
	Year                 string                  `json:"year,omitempty"`
}

// SeriesAirsDays represents which days a series airs.
type SeriesAirsDays struct {
	Friday    bool `json:"friday,omitempty"`
	Monday    bool `json:"monday,omitempty"`
	Saturday  bool `json:"saturday,omitempty"`
	Sunday    bool `json:"sunday,omitempty"`
	Thursday  bool `json:"thursday,omitempty"`
	Tuesday   bool `json:"tuesday,omitempty"`
	Wednesday bool `json:"wednesday,omitempty"`
}

// SeriesEpisodesResponse represents the response for series episodes.
type SeriesEpisodesResponse struct {
	Series   *SeriesBaseRecord   `json:"series,omitempty"`
	Episodes []EpisodeBaseRecord `json:"episodes,omitempty"`
}
