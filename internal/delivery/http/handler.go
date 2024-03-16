// internal/delivery/http/handler.go

package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/juscilan/go-crip-clean-arch/internal/usecase"
)

type EncryptorHandler struct {
	ucase *usecase.EncryptorUsecase
}

func NewEncryptorHandler(ucase *usecase.EncryptorUsecase) *EncryptorHandler {
	return &EncryptorHandler{ucase}
}

func (h *EncryptorHandler) Encrypt(c *gin.Context) {
	data := []byte(c.PostForm("data"))
	key := []byte(c.PostForm("key"))

	encryptedString, err := h.ucase.Encrypt(data, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"encrypted_data": encryptedString})
}

func (h *EncryptorHandler) Decrypt(c *gin.Context) {
	encryptedString := c.PostForm("encrypted_string")
	key := []byte(c.PostForm("key"))

	decryptedData, err := h.ucase.Decrypt(encryptedString, key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"decrypted_data": string(decryptedData)})
}
