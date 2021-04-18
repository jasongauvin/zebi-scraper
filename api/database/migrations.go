package database

import (
	"github.com/jasongauvin/zebi-scraper/api/models"
)

func MakeMigrations() {
	DB.AutoMigrate(&models.Customer{})
}
