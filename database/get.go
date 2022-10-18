package database

import "gorm.io/gorm"

func Get() *gorm.DB {
	if db == nil {
		panic("database not initialized, have you called database.Init()?")
	}

	return db
}
