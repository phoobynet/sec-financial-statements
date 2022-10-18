package database

import (
	"github.com/phoobynet/sec-financial-statements/nums"
	"github.com/phoobynet/sec-financial-statements/pres"
	"github.com/phoobynet/sec-financial-statements/submissions"
	"github.com/phoobynet/sec-financial-statements/tags"
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

	migrateErr := d.AutoMigrate(&submissions.Submission{}, &pres.Pre{}, &nums.Num{}, &tags.Tag{})

	if migrateErr != nil {
		log.Fatalln(migrateErr)
	}

	db = d

	return db
}
