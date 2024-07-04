package db

import (
    "spyingCats/models"

    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
    "log"
)

var DB *gorm.DB

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("cats.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

    DB.AutoMigrate(
        &models.Cat{},
        &models.Mission{},
        &models.Target{},
        &models.Note{},
    )
}
