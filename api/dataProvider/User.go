package dataprovider

import (
	"NoSpamGo/domain"
	"database/sql"
)

type User struct {
	Mail           string
	Secret         string
	ImapUsername   sql.NullString
	ImapPassword   sql.NullString
	ImapServerUrl  sql.NullString
	ImapServerPort sql.NullInt64
}

func (u *User) ToDomain() *domain.User {
	return &domain.User{
		Mail:           u.Mail,
		Secret:         u.Secret,
		ImapUsername:   valueOrEmpty(u.ImapUsername),
		ImapPassword:   valueOrEmpty(u.ImapPassword),
		ImapServerUrl:  valueOrEmpty(u.ImapServerUrl),
		ImapServerPort: int(valueOrZero(u.ImapServerPort)),
	}
}

func valueOrEmpty(n sql.NullString) string {
	if n.Valid {
		return n.String
	}
	return ""
}

func valueOrZero(n sql.NullInt64) int64 {
	if n.Valid {
		return n.Int64
	}
	return 0
}
