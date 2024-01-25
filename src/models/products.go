package models

import (
	"dasar_backend_go/src/config"

	"github.com/jinzhu/gorm"
)

type Product struct {
	gorm.Model
	Name  string
	Price int
	Stock int
}

func SelectAll() *gorm.DB {
	items := []Product{}
	return config.DB.Find(&items)
}

func Select(id string) *gorm.DB {
	var item Product
	return config.DB.First(&item, "id = ?", id)
}

func Post(item *Product) *gorm.DB {
	return config.DB.Create(&item)
}

func Updates(id string, newProduct *Product) *gorm.DB {
	var item Product
	return config.DB.Model(&item).Where("id = ?", id).Updates(&newProduct)
}

func Deletes(id string) *gorm.DB {
	var item Product
	return config.DB.Delete(&item, "id = ?", id)
}
