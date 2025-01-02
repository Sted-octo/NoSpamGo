package domain

type Message struct {
	Subject string        `json:"Subject"`
	Id      uint32        `json:"ID"`
	Mails   []MailAddress `json:"Mails"`
}
