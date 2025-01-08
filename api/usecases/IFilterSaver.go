package usecases

import "NoSpamGo/domain"

type IFilterSaver[T any] interface {
	Save(mail string,
		filter domain.Filter,
		dbConnector IDatabaseConnector[T],
		filterByNameForUserMailLoader IFilterByNameForUserMailLoader[T]) bool
}
