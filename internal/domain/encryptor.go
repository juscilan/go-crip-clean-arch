package domain

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

type Encryptor interface {
	Encrypt(data []byte, key []byte) (string, error)
	Decrypt(encryptedString string, key []byte) ([]byte, error)
}

type encryptor struct{}

func NewEncryptor() Encryptor {
	return &encryptor{}
}

func (e *encryptor) Encrypt(data []byte, key []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Encrypt the data
	ciphertext := make([]byte, aes.BlockSize+len(data))
	iv := ciphertext[:aes.BlockSize]
	if _, err := rand.Read(iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], data)

	// Convert the result to a hexadecimal string
	return hex.EncodeToString(ciphertext), nil
}

func (e *encryptor) Decrypt(encryptedString string, key []byte) ([]byte, error) {
	// Decode the hexadecimal string
	ciphertext, err := hex.DecodeString(encryptedString)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Check if the ciphertext is valid
	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("ciphertext too short")
	}

	// Decrypt the data
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}
