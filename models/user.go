package models

// UserInfo represents a user info record.
type UserInfo struct {
	ID       int    `json:"id,omitempty"`
	Language string `json:"language,omitempty"`
	Name     string `json:"name,omitempty"`
	Type     string `json:"type,omitempty"`
}

// Favorites represents user favorites.
type Favorites struct {
	Series   []int `json:"series,omitempty"`
	Movies   []int `json:"movies,omitempty"`
	Episodes []int `json:"episodes,omitempty"`
	Artwork  []int `json:"artwork,omitempty"`
	People   []int `json:"people,omitempty"`
	Lists    []int `json:"lists,omitempty"`
}

// FavoriteRecord represents a single favorite record to add.
type FavoriteRecord struct {
	Series  int `json:"series,omitempty"`
	Movie   int `json:"movie,omitempty"`
	Episode int `json:"episode,omitempty"`
	Artwork int `json:"artwork,omitempty"`
	People  int `json:"people,omitempty"`
	List    int `json:"list,omitempty"`
}
