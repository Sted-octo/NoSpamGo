package usecases

type IClientConnector[T any] interface {
	Connect(imapUrl string, port int, userName string, password string) error
	Close()
	Get() T
}
