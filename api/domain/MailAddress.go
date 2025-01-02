package domain

type MailAddress struct {
	PersonalName string `json:"PersonalName"`
	MailboxName  string `json:"MailboxName"`
	HostName     string `json:"HostName"`
}
