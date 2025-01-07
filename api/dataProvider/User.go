package dataprovider

import (
	"NoSpamGo/domain"
	"NoSpamGo/tools"
	"NoSpamGo/usecases"
	"database/sql"
	"os"
)

type User struct {
	Mail              string
	Secret            string
	ImapUsername      sql.NullString
	ImapPassword      sql.NullString
	ImapServerUrl     sql.NullString
	ImapServerPort    sql.NullInt64
	ImapPasswordNonce sql.NullString
}

func (u *User) ToDomain() *domain.User {
	var cryptoHelper usecases.ICryptoHelper = tools.NewCryptoHelper([]byte(os.Getenv("CRYPTO_KEY")))

	decryptedPassword, err := cryptoHelper.Decrypt([]byte(valueOrEmpty(u.ImapPassword)), []byte(valueOrEmpty(u.ImapPasswordNonce)))

	if err != nil {
		return &domain.User{}
	}
	return &domain.User{
		Mail:           u.Mail,
		Secret:         u.Secret,
		ImapUsername:   valueOrEmpty(u.ImapUsername),
		ImapPassword:   decryptedPassword,
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
