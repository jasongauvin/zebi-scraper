package database

import (
	"github.com/jasongauvin/gogoleplate/api/models"
)

func MakeMigrations() {
	DB.AutoMigrate(&models.Customer{})
}
