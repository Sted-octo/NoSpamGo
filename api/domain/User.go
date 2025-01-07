package domain

type User struct {
	Mail           string `json:"Mail"`
	Secret         string `json:"Secret,omitempty"`
	ImapUsername   string `json:"ImapUsername,omitempty"`
	ImapPassword   string `json:"ImapPassword,omitempty"`
	ImapServerUrl  string `json:"ImapServerUrl,omitempty"`
	ImapServerPort int    `json:"ImapServerPort,omitempty"`
}
