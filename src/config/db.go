package config

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func InitDB()
	url := os.Getenv("URL")
	var err error
	DB, err = gorm.Open("postgres", url)
	if err != nil {
		panic ("Failed to connect database")
	}