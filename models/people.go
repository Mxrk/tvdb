package models

// PeopleBaseRecord represents a base people record.
type PeopleBaseRecord struct {
	Aliases              []Alias  `json:"aliases,omitempty"`
	ID                   int64    `json:"id,omitempty"`
	Image                string   `json:"image,omitempty"`
	LastUpdated          string   `json:"lastUpdated,omitempty"`
	Name                 string   `json:"name,omitempty"`
	NameTranslations     []string `json:"nameTranslations,omitempty"`
	OverviewTranslations []string `json:"overviewTranslations,omitempty"`
	Score                int64    `json:"score,omitempty"`
}

// PeopleExtendedRecord represents an extended people record.
type PeopleExtendedRecord struct {
	Aliases              []Alias              `json:"aliases,omitempty"`
	Awards               []AwardBaseRecord    `json:"awards,omitempty"`
	Biographies          []Biography          `json:"biographies,omitempty"`
	Birth                string               `json:"birth,omitempty"`
	BirthPlace           string               `json:"birthPlace,omitempty"`
	Characters           []Character          `json:"characters,omitempty"`
	Death                string               `json:"death,omitempty"`
	Gender               int                  `json:"gender,omitempty"`
	ID                   int64                `json:"id,omitempty"`
	Image                string               `json:"image,omitempty"`
	LastUpdated          string               `json:"lastUpdated,omitempty"`
	Name                 string               `json:"name,omitempty"`
	NameTranslations     []string             `json:"nameTranslations,omitempty"`
	OverviewTranslations []string             `json:"overviewTranslations,omitempty"`
	Races                []Race               `json:"races,omitempty"`
	RemoteIDs            []RemoteID           `json:"remoteIds,omitempty"`
	Score                int64                `json:"score,omitempty"`
	Slug                 string               `json:"slug,omitempty"`
	TagOptions           []TagOption          `json:"tagOptions,omitempty"`
	Translations         *TranslationExtended `json:"translations,omitempty"`
}

// PeopleType represents a people type record.
type PeopleType struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// Race represents a race record (currently empty in the API spec).
type Race struct{}
