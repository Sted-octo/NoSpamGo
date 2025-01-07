package usecases

type ICryptoHelper interface {
	Encrypt(plaintext string) ([]byte, []byte, error)
	Decrypt(ciphertext []byte, nonce []byte) (string, error)
}
