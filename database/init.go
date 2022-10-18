package database

import (
	"github.com/phoobynet/sec-financial-statements/submissions"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func Init(databasePath string) *gorm.DB {
	if db != nil {
		return db
	}

	d, openErr := gorm.Open(sqlite.Open(databasePath), &gorm.Config{
		SkipDefaultTransaction: true,
	})

	if openErr != nil {
		log.Fatalln(openErr)
	}

	migrateErr := d.AutoMigrate(&submissions.Submission{})

	if migrateErr != nil {
		log.Fatalln(migrateErr)
	}

	db = d

	return db
}
