package presenter

type EmailResult struct {
	Mail              string `json:"Mail"`
	CountSpamDetected int    `json:"CountSpamDetected"`
}
