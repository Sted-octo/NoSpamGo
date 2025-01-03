package usecases

import "NoSpamGo/domain"

type IUserByMailLoader[T any] interface {
	Load(mail string, dbConnector IDatabaseConnector[T]) *domain.User
}
