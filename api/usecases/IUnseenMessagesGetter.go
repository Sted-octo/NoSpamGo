package usecases

import "NoSpamGo/domain"

type IUnseenMessagesGetter[T any] interface {
	Get(clientConnector IClientConnector[T]) []domain.Message
}
