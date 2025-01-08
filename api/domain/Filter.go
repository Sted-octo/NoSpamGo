package domain

type Filter struct {
	Name                 string `json:"Name"`
	NumberOfSpamDetected int    `json:"NumberOfSpamDetected,omitempty"`
}
