package usecases

type ISpamMover[T any] interface {
	Move(clientConnector IClientConnector[T], ids []uint32)
}
