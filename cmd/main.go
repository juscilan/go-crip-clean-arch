// main.go

package main

import (
	"github.com/gin-gonic/gin"
	"github.com/juscilan/go-crip-clean-arch/internal/delivery/http"
	"github.com/juscilan/go-crip-clean-arch/internal/domain"
	"github.com/juscilan/go-crip-clean-arch/internal/usecase"
)

func main() {
	r := gin.Default()

	// Initialize dependencies
	encryptor := domain.NewEncryptor()
	encryptorUsecase := usecase.NewEncryptorUsecase(encryptor)
	encryptorHandler := http.NewEncryptorHandler(encryptorUsecase)

	// Routes
	api := r.Group("/api")
	{
		api.POST("/encrypt", encryptorHandler.Encrypt)
		api.POST("/decrypt", encryptorHandler.Decrypt)
	}

	r.Run(":8080")
}
