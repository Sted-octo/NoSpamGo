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

	var userDb = new(User)
	row := dbConnector.Get().QueryRow("SELECT mail, secret, mailbox_username, mailbox_password, mailbox_password_nonce, mailbox_server, mailbox_port FROM users WHERE mail = ?", mail)
	err := row.Scan(&userDb.Mail, &userDb.Secret, &userDb.ImapUsername, &userDb.ImapPassword, &userDb.ImapPasswordNonce, &userDb.ImapServerUrl, &userDb.ImapServerPort)

	if err == sql.ErrNoRows {
		return nil
	}

	if err != nil {
		log.Printf("erreur lors de la requête : %v", err)
		return nil
	}

	return userDb.ToDomain()
}
