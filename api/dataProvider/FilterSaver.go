package dataprovider

import (
	"NoSpamGo/domain"
	"NoSpamGo/usecases"
	"database/sql"
	"log"
)

type FilterSaver struct{}

func (o *FilterSaver) Save(
	mail string,
	filter domain.Filter,
	dbConnector usecases.IDatabaseConnector[*sql.DB],
	filterByNameForUserMailLoader usecases.IFilterByNameForUserMailLoader[*sql.DB]) bool {

	if dbConnector.Get() == nil {
		log.Println("Database access error in service UserSaver")
		return false
	}

	filterLoaded := filterByNameForUserMailLoader.Load(mail, filter.Name, dbConnector)
	if filterLoaded == nil {
		stmt, _ := dbConnector.Get().Prepare("INSERT INTO filters(mail, filter_name, filter_number_of_spam_detected ) VALUES(?, ?, ?)")
		_, err := stmt.Exec(mail, filter.Name, filter.NumberOfSpamDetected)
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	}

	stmt, _ := dbConnector.Get().Prepare("UPDATE filter set filter_number_of_spam_detected = ? where mail = ? and filter_name = ?")
	_, err := stmt.Exec(filter.NumberOfSpamDetected, mail, filter.Name)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
