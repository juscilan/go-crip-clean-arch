package usecase

import "github.com/juscilan/go-crip-clean-arch/internal/domain"

type EncryptorUsecase struct {
	encryptor domain.Encryptor
}

func NewEncryptorUsecase(encryptor domain.Encryptor) *EncryptorUsecase {
	return &EncryptorUsecase{encryptor}
}

func (uc *EncryptorUsecase) Encrypt(data []byte, key []byte) (string, error) {
	return uc.encryptor.Encrypt(data, key)
}

func (uc *EncryptorUsecase) Decrypt(encryptedString string, key []byte) ([]byte, error) {
	return uc.encryptor.Decrypt(encryptedString, key)
}
