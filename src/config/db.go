package config

import (
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func InitDB() {
	url := os.Getenv("URL")
	var err error
	DB, err = gorm.Open("postgres", url)
	if err != nil {
		panic("Failed to connect database")
	}
}
