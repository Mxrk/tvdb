package models

// EntityUpdate represents an entity update record.
type EntityUpdate struct {
	EntityType        string `json:"entityType,omitempty"`
	MethodInt         int    `json:"methodInt,omitempty"`
	Method            string `json:"method,omitempty"`
	ExtraInfo         string `json:"extraInfo,omitempty"`
	UserID            int    `json:"userId,omitempty"`
	RecordType        string `json:"recordType,omitempty"`
	RecordID          int64  `json:"recordId,omitempty"`
	TimeStamp         int64  `json:"timeStamp,omitempty"`
	SeriesID          int64  `json:"seriesId,omitempty"`
	MergeToId         int64  `json:"mergeToId,omitempty"`
	MergeToEntityType string `json:"mergeToEntityType,omitempty"`
}
