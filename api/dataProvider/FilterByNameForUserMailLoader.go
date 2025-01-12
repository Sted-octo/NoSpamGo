package dataprovider

import (
	"NoSpamGo/domain"
	"NoSpamGo/usecases"
	"database/sql"
	"log"
)

type FilterByNameForUserMailLoader struct{}

func (o *FilterByNameForUserMailLoader) Load(mail string, name string, dbConnector usecases.IDatabaseConnector[*sql.DB]) *domain.Filter {

	if dbConnector.Get() == nil {
		log.Println("Database access error in service UserSaver")
		return nil
	}

	var filterDb = new(Filter)
	row := dbConnector.Get().QueryRow("SELECT mail, filter_name, filter_number_of_spam_detected FROM filters WHERE mail = ? AND filter_name = ?", mail, name)
	err := row.Scan(&filterDb.Mail, &filterDb.FilterName, &filterDb.FilterNumberOfSpamDetected)

	if err == sql.ErrNoRows {
		return nil
	}

	if err != nil {
		log.Printf("erreur lors de la requête : %v", err)
		return nil
	}

	return filterDb.ToDomain()
}
