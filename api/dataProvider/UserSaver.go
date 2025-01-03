package dataprovider

import (
	"NoSpamGo/domain"
	"NoSpamGo/usecases"
	"database/sql"
	"log"
)

type UserSaver struct{}

func (o *UserSaver) Save(user domain.User, dbConnector usecases.IDatabaseConnector[*sql.DB]) {

	if dbConnector.Get() == nil {
		log.Fatal("Database access error in service UserSaver")
	}
	stmt, _ := dbConnector.Get().Prepare("INSERT INTO users(mail, secret) VALUES(?, ?)")
	_, err := stmt.Exec(user.Mail, []byte(user.Secret))
	if err != nil {
		log.Fatal(err)
	}
}
