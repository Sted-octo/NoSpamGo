package dataprovider

import (
	"NoSpamGo/domain"
	"NoSpamGo/usecases"
	"database/sql"
	"log"
)

type UserSaver struct{}

func (o *UserSaver) Save(
	user domain.User,
	dbConnector usecases.IDatabaseConnector[*sql.DB],
	userByMailLoader usecases.IUserByMailLoader[*sql.DB],
	cryptoHelper usecases.ICryptoHelper) bool {

	if dbConnector.Get() == nil {
		log.Println("Database access error in service UserSaver")
		return false
	}

	encryptedPassword, nonce, err := cryptoHelper.Encrypt(user.ImapPassword)
	if err != nil {
		log.Println("Erreur d'encryption du mot de passe")
		return false
	}

	userLoaded := userByMailLoader.Load(user.Mail, dbConnector)
	if userLoaded == nil {
		stmt, _ := dbConnector.Get().Prepare("INSERT INTO users(mail, secret, mailbox_username, mailbox_password, mailbox_password_nonce, mailbox_server, mailbox_port) VALUES(?, ?, ?, ?, ?, ?, ?)")
		_, err := stmt.Exec(user.Mail, []byte(user.Secret), user.ImapUsername, encryptedPassword, nonce, user.ImapServerUrl, user.ImapServerPort)
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	}

	stmt, _ := dbConnector.Get().Prepare("UPDATE users set mailbox_username = ?, mailbox_password = ?, mailbox_password_nonce = ?, mailbox_server = ?, mailbox_port = ? where mail = ?")
	_, err = stmt.Exec(user.ImapUsername, encryptedPassword, nonce, user.ImapServerUrl, user.ImapServerPort, user.Mail)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
