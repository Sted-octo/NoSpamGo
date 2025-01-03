package dataprovider

import (
	"NoSpamGo/domain"
	"NoSpamGo/usecases"
	"database/sql"
	"log"
)

type UserByMailLoader struct{}

func (o *UserByMailLoader) Load(mail string, dbConnector usecases.IDatabaseConnector[*sql.DB]) *domain.User {

	if dbConnector.Get() == nil {
		log.Fatal("Database access error in service UserSaver")
	}

	var user domain.User
	row := dbConnector.Get().QueryRow("SELECT mail, secret FROM users WHERE mail = ?", mail)
	row.Scan(&user.Mail, &user.Secret)
	return &user
}
