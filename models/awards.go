package models

// AwardBaseRecord represents a base award record.
type AwardBaseRecord struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// AwardExtendedRecord represents an extended award record.
type AwardExtendedRecord struct {
	Categories []AwardCategoryBaseRecord `json:"categories,omitempty"`
	ID         int                       `json:"id,omitempty"`
	Name       string                    `json:"name,omitempty"`
	Score      int64                     `json:"score,omitempty"`
}

// AwardCategoryBaseRecord represents a base award category record.
type AwardCategoryBaseRecord struct {
	AllowCoNominees bool             `json:"allowCoNominees,omitempty"`
	Award           *AwardBaseRecord `json:"award,omitempty"`
	ForMovies       bool             `json:"forMovies,omitempty"`
	ForSeries       bool             `json:"forSeries,omitempty"`
	ID              int64            `json:"id,omitempty"`
	Name            string           `json:"name,omitempty"`
}

// AwardCategoryExtendedRecord represents an extended award category record.
type AwardCategoryExtendedRecord struct {
	AllowCoNominees bool                     `json:"allowCoNominees,omitempty"`
	Award           *AwardBaseRecord         `json:"award,omitempty"`
	ForMovies       bool                     `json:"forMovies,omitempty"`
	ForSeries       bool                     `json:"forSeries,omitempty"`
	ID              int64                    `json:"id,omitempty"`
	Name            string                   `json:"name,omitempty"`
	Nominees        []AwardNomineeBaseRecord `json:"nominees,omitempty"`
}

// AwardNomineeBaseRecord represents a base award nominee record.
type AwardNomineeBaseRecord struct {
	Character *Character         `json:"character,omitempty"`
	Details   string             `json:"details,omitempty"`
	Episode   *EpisodeBaseRecord `json:"episode,omitempty"`
	ID        int64              `json:"id,omitempty"`
	IsWinner  bool               `json:"isWinner,omitempty"`
	Movie     *MovieBaseRecord   `json:"movie,omitempty"`
	Series    *SeriesBaseRecord  `json:"series,omitempty"`
	Year      string             `json:"year,omitempty"`
	Category  string             `json:"category,omitempty"`
	Name      string             `json:"name,omitempty"`
}
