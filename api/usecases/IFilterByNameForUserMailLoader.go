package usecases

import "NoSpamGo/domain"

type IFilterByNameForUserMailLoader[T any] interface {
	Load(mail string, name string, dbConnector IDatabaseConnector[T]) *domain.Filter
}
