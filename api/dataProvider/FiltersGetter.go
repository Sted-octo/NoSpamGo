package dataprovider

import (
	"NoSpamGo/domain"
	"NoSpamGo/usecases"
	"database/sql"
	"log"
)

type FiltersGetter struct{}

func (o *FiltersGetter) Get(mail string, dbConnector usecases.IDatabaseConnector[*sql.DB]) []domain.Filter {

	if dbConnector.Get() == nil {
		log.Println("Database access error in service UserSaver")
		return nil
	}

	rows, err := dbConnector.Get().Query(`
		SELECT mail, filter_name, filter_number_of_spam_detected 
		FROM filters 
		WHERE mail = ?
		ORDER by filter_number_of_spam_detected DESC
		`, mail)
	if err == sql.ErrNoRows {
		return nil
	}
	if err != nil {
		log.Println(err)
		return nil
	}
	defer rows.Close()

	filters := make([]domain.Filter, 0)

	for rows.Next() {
		var filterDb = new(Filter)
		if err := rows.Scan(&filterDb.Mail, &filterDb.FilterName, &filterDb.FilterNumberOfSpamDetected); err != nil {
			log.Printf("Erreur lors de la lecture d'un filter: %v", err)
			continue
		}
		filters = append(filters, *filterDb.ToDomain())
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil
	}

	return filters
}
