package domain

type User struct {
	Username     string `json:"Username"`
	Secret       string `json:"Secret,omitempty"`
	IsEnabled2FA bool   `json:"IsEnabled2FA"`
}
