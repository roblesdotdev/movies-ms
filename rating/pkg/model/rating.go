package model

// RecordId defines a record id. Together with RecordType
// identifies unique records across all types.
type (
	RecordId   string
	RecordType string
)

// Existing record types
const (
	RecordTypeMovie = RecordType("movie")
)

type UserId string

type RatingValue int

type Rating struct {
	RecordId   string      `json:"recordId"`
	RecordType string      `json:"recordType"`
	UserId     UserId      `json:"userId"`
	Value      RatingValue `json:"value"`
}
