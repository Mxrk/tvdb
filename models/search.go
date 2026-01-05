package models

// SearchResult represents a search result.
type SearchResult struct {
	Aliases              []string          `json:"aliases,omitempty"`
	Companies            []string          `json:"companies,omitempty"`
	CompanyType          string            `json:"companyType,omitempty"`
	Country              string            `json:"country,omitempty"`
	Director             string            `json:"director,omitempty"`
	FirstAirTime         string            `json:"first_air_time,omitempty"`
	Genres               []string          `json:"genres,omitempty"`
	ID                   string            `json:"id,omitempty"`
	ImageURL             string            `json:"image_url,omitempty"`
	Name                 string            `json:"name,omitempty"`
	IsOfficial           bool              `json:"is_official,omitempty"`
	NameTranslated       string            `json:"name_translated,omitempty"`
	Network              string            `json:"network,omitempty"`
	ObjectID             string            `json:"objectID,omitempty"`
	OfficialList         string            `json:"officialList,omitempty"`
	Overview             string            `json:"overview,omitempty"`
	Overviews            TranslationSimple `json:"overviews,omitempty"`
	OverviewTranslated   []string          `json:"overview_translated,omitempty"`
	Poster               string            `json:"poster,omitempty"`
	Posters              []string          `json:"posters,omitempty"`
	PrimaryLanguage      string            `json:"primary_language,omitempty"`
	RemoteIDs            []RemoteID        `json:"remote_ids,omitempty"`
	Status               string            `json:"status,omitempty"`
	Slug                 string            `json:"slug,omitempty"`
	Studios              []string          `json:"studios,omitempty"`
	Title                string            `json:"title,omitempty"`
	Thumbnail            string            `json:"thumbnail,omitempty"`
	Translations         TranslationSimple `json:"translations,omitempty"`
	TranslationsWithLang []string          `json:"translationsWithLang,omitempty"`
	TvdbID               string            `json:"tvdb_id,omitempty"`
	Type                 string            `json:"type,omitempty"`
	Year                 string            `json:"year,omitempty"`
}

// SearchByRemoteIdResult represents a search by remote ID result.
type SearchByRemoteIdResult struct {
	Series  *SeriesBaseRecord  `json:"series,omitempty"`
	People  *PeopleBaseRecord  `json:"people,omitempty"`
	Movie   *MovieBaseRecord   `json:"movie,omitempty"`
	Episode *EpisodeBaseRecord `json:"episode,omitempty"`
	Company *Company           `json:"company,omitempty"`
}
