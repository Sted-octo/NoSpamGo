package tools

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
	"log"
)

type CryptoHelper struct {
	encryptionKey []byte
}

func NewCryptoHelper(key []byte) *CryptoHelper {
	if len(key) != 32 {
		log.Fatalf("la clé de chiffrement doit faire 32 bytes, elle fait %d", len(key))
	}
	return &CryptoHelper{
		encryptionKey: key,
	}
}

func (s *CryptoHelper) Encrypt(plaintext string) ([]byte, []byte, error) {

	block, err := aes.NewCipher(s.encryptionKey)
	if err != nil {
		return nil, nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	ciphertext := gcm.Seal(nil, nonce, []byte(plaintext), nil)
	return ciphertext, nonce, nil
}

func (s *CryptoHelper) Decrypt(ciphertext []byte, nonce []byte) (string, error) {
	block, err := aes.NewCipher(s.encryptionKey)
	if err != nil {
		return "", err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}

	return string(plaintext), nil
}
