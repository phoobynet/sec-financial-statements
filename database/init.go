package database

import (
	"github.com/phoobynet/sec-financial-statements/companies"
	"github.com/phoobynet/sec-financial-statements/nums"
	"github.com/phoobynet/sec-financial-statements/pres"
	"github.com/phoobynet/sec-financial-statements/sics"
	"github.com/phoobynet/sec-financial-statements/subs"
	"github.com/phoobynet/sec-financial-statements/tags"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
)

func Init(databasePath string) *gorm.DB {
	if db != nil {
		return db
	}

	d, openErr := gorm.Open(sqlite.Open(databasePath), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Silent),
		PrepareStmt:            true,
	})

	if openErr != nil {
		log.Fatalln(openErr)
	}

	migrateErr := d.AutoMigrate(&subs.Sub{}, &pres.Pre{}, &nums.Num{}, &tags.Tag{}, &companies.Company{}, &sics.SIC{})

	if migrateErr != nil {
		log.Fatalln(migrateErr)
	}

	d.Exec("PRAGMA cache_size = 400000; PRAGMA synchronous = OFF; PRAGMA journal_mode = OFF; PRAGMA locking_mode = EXCLUSIVE; PRAGMA count_changes = OFF; PRAGMA temp_store = MEMORY; PRAGMA auto_vacuum = NONE;")

	db = d

	return db
}
