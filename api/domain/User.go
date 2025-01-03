package domain

type User struct {
	Mail         string `json:"Mail"`
	Secret       string `json:"Secret,omitempty"`
	IsEnabled2FA bool   `json:"IsEnabled2FA"`
}
