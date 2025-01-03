package usecases

type IDatabaseConnector[T any] interface {
	Connect() error
	Close()
	Get() T
}
