package models

// EpisodeBaseRecord represents a base episode record.
type EpisodeBaseRecord struct {
	AbsoluteNumber       int                `json:"absoluteNumber,omitempty"`
	Aired                string             `json:"aired,omitempty"`
	AirsAfterSeason      int                `json:"airsAfterSeason,omitempty"`
	AirsBeforeEpisode    int                `json:"airsBeforeEpisode,omitempty"`
	AirsBeforeSeason     int                `json:"airsBeforeSeason,omitempty"`
	FinaleType           string             `json:"finaleType,omitempty"`
	ID                   int64              `json:"id,omitempty"`
	Image                string             `json:"image,omitempty"`
	ImageType            *int               `json:"imageType,omitempty"`
	IsMovie              int64              `json:"isMovie,omitempty"`
	LastUpdated          string             `json:"lastUpdated,omitempty"`
	LinkedMovie          int                `json:"linkedMovie,omitempty"`
	Name                 string             `json:"name,omitempty"`
	NameTranslations     []string           `json:"nameTranslations,omitempty"`
	Number               int                `json:"number,omitempty"`
	Overview             string             `json:"overview,omitempty"`
	OverviewTranslations []string           `json:"overviewTranslations,omitempty"`
	Runtime              *int               `json:"runtime,omitempty"`
	SeasonNumber         int                `json:"seasonNumber,omitempty"`
	Seasons              []SeasonBaseRecord `json:"seasons,omitempty"`
	SeriesID             int64              `json:"seriesId,omitempty"`
	SeasonName           string             `json:"seasonName,omitempty"`
	Year                 string             `json:"year,omitempty"`
}

// EpisodeExtendedRecord represents an extended episode record.
type EpisodeExtendedRecord struct {
	Aired                string                   `json:"aired,omitempty"`
	AirsAfterSeason      int                      `json:"airsAfterSeason,omitempty"`
	AirsBeforeEpisode    int                      `json:"airsBeforeEpisode,omitempty"`
	AirsBeforeSeason     int                      `json:"airsBeforeSeason,omitempty"`
	Awards               []AwardBaseRecord        `json:"awards,omitempty"`
	Characters           []Character              `json:"characters,omitempty"`
	Companies            []Company                `json:"companies,omitempty"`
	ContentRatings       []ContentRating          `json:"contentRatings,omitempty"`
	FinaleType           string                   `json:"finaleType,omitempty"`
	ID                   int64                    `json:"id,omitempty"`
	Image                string                   `json:"image,omitempty"`
	ImageType            *int                     `json:"imageType,omitempty"`
	IsMovie              int64                    `json:"isMovie,omitempty"`
	LastUpdated          string                   `json:"lastUpdated,omitempty"`
	LinkedMovie          int                      `json:"linkedMovie,omitempty"`
	Name                 string                   `json:"name,omitempty"`
	NameTranslations     []string                 `json:"nameTranslations,omitempty"`
	Networks             []Company                `json:"networks,omitempty"`
	Nominations          []AwardNomineeBaseRecord `json:"nominations,omitempty"`
	Number               int                      `json:"number,omitempty"`
	Overview             string                   `json:"overview,omitempty"`
	OverviewTranslations []string                 `json:"overviewTranslations,omitempty"`
	ProductionCode       string                   `json:"productionCode,omitempty"`
	RemoteIDs            []RemoteID               `json:"remoteIds,omitempty"`
	Runtime              *int                     `json:"runtime,omitempty"`
	SeasonNumber         int                      `json:"seasonNumber,omitempty"`
	Seasons              []SeasonBaseRecord       `json:"seasons,omitempty"`
	SeriesID             int64                    `json:"seriesId,omitempty"`
	Studios              []Company                `json:"studios,omitempty"`
	TagOptions           []TagOption              `json:"tagOptions,omitempty"`
	Trailers             []Trailer                `json:"trailers,omitempty"`
	Translations         *TranslationExtended     `json:"translations,omitempty"`
	Year                 string                   `json:"year,omitempty"`
}
