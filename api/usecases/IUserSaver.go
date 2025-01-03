package usecases

import "NoSpamGo/domain"

type IUserSaver[T any] interface {
	Save(user domain.User, dbConnector IDatabaseConnector[T])
}
