// Package models contains all data types for the TVDB API.
package models

// Alias is a name alias with language.
type Alias struct {
	Language string `json:"language,omitempty"`
	Name     string `json:"name,omitempty"`
}

// Status represents a status record for series/movies.
type Status struct {
	ID          *int64 `json:"id,omitempty"`
	KeepUpdated bool   `json:"keepUpdated,omitempty"`
	Name        string `json:"name,omitempty"`
	RecordType  string `json:"recordType,omitempty"`
}

// Links represents pagination links.
type Links struct {
	Prev       *string `json:"prev,omitempty"`
	Self       *string `json:"self,omitempty"`
	Next       string  `json:"next,omitempty"`
	TotalItems int     `json:"total_items,omitempty"`
	PageSize   int     `json:"page_size,omitempty"`
}

// RemoteID represents a remote identifier (IMDB, EIDR, etc.).
type RemoteID struct {
	ID         string `json:"id,omitempty"`
	Type       int64  `json:"type,omitempty"`
	SourceName string `json:"sourceName,omitempty"`
}

// TagOption represents a tag option record.
type TagOption struct {
	HelpText string `json:"helpText,omitempty"`
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	Tag      int64  `json:"tag,omitempty"`
	TagName  string `json:"tagName,omitempty"`
}

// Tag represents a tag record.
type Tag struct {
	AllowsMultiple bool        `json:"allowsMultiple,omitempty"`
	HelpText       string      `json:"helpText,omitempty"`
	ID             int64       `json:"id,omitempty"`
	Name           string      `json:"name,omitempty"`
	Options        []TagOption `json:"options,omitempty"`
}

// Translation represents a translation record.
type Translation struct {
	Aliases   []string `json:"aliases,omitempty"`
	IsAlias   bool     `json:"isAlias,omitempty"`
	IsPrimary bool     `json:"isPrimary,omitempty"`
	Language  string   `json:"language,omitempty"`
	Name      string   `json:"name,omitempty"`
	Overview  string   `json:"overview,omitempty"`
	Tagline   string   `json:"tagline,omitempty"`
}

// TranslationExtended represents extended translation data.
type TranslationExtended struct {
	NameTranslations     []Translation `json:"nameTranslations,omitempty"`
	OverviewTranslations []Translation `json:"overviewTranslations,omitempty"`
	Alias                []string      `json:"alias,omitempty"`
}

// TranslationSimple is a map of language code to translated text.
type TranslationSimple map[string]string

// Gender represents a gender record.
type Gender struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Language represents a language record.
type Language struct {
	ID         string `json:"id,omitempty"`
	Name       string `json:"name,omitempty"`
	NativeName string `json:"nativeName,omitempty"`
	ShortCode  string `json:"shortCode,omitempty"`
}

// Country represents a country record.
type Country struct {
	ID        string `json:"id,omitempty"`
	Name      string `json:"name,omitempty"`
	ShortCode string `json:"shortCode,omitempty"`
}

// ContentRating represents a content rating record.
type ContentRating struct {
	ID          int64  `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	Country     string `json:"country,omitempty"`
	ContentType string `json:"contentType,omitempty"`
	Order       int    `json:"order,omitempty"`
	FullName    string `json:"fullName,omitempty"`
}

// GenreBaseRecord represents a genre record.
type GenreBaseRecord struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Slug string `json:"slug,omitempty"`
}

// Trailer represents a trailer record.
type Trailer struct {
	ID       int64  `json:"id,omitempty"`
	Language string `json:"language,omitempty"`
	Name     string `json:"name,omitempty"`
	URL      string `json:"url,omitempty"`
	Runtime  int    `json:"runtime,omitempty"`
}

// Biography represents a biography record.
type Biography struct {
	Biography string `json:"biography,omitempty"`
	Language  string `json:"language,omitempty"`
}

// RecordInfo represents base record info.
type RecordInfo struct {
	Image string `json:"image,omitempty"`
	Name  string `json:"name,omitempty"`
	Year  string `json:"year,omitempty"`
}

// EntityType represents an entity type record.
type EntityType struct {
	ID          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	HasSpecials bool   `json:"hasSpecials,omitempty"`
}

// SourceType represents a source type record.
type SourceType struct {
	ID      int64  `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Postfix string `json:"postfix,omitempty"`
	Prefix  string `json:"prefix,omitempty"`
	Slug    string `json:"slug,omitempty"`
	Sort    int64  `json:"sort,omitempty"`
}

// InspirationType represents a movie inspiration type.
type InspirationType struct {
	ID            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	Description   string `json:"description,omitempty"`
	ReferenceName string `json:"reference_name,omitempty"`
	URL           string `json:"url,omitempty"`
}
