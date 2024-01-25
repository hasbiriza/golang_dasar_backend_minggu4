package helper

import (
	"dasar_backend_go/src/config"
	"dasar_backend_go/src/models"
)

func Migration() {
	config.DB.AutoMigrate(&models.Product{})
}
