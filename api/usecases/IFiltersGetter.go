package usecases

import "NoSpamGo/domain"

type IFiltersGetter[T any] interface {
	Get(mail string, dbConnector IDatabaseConnector[T]) []domain.Filter
}
