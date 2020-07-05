package structs

import "github.com/jinzhu/gorm"

type DonorMin struct {
	gorm.Model
	Name           string
	Email          string
	DonationAmount string
	UniqueCode     string
	Status         int `gorm:"default:'0'"`
	ArticleID      int
}
