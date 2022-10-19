package subs

import (
	"archive/zip"
	. "github.com/phoobynet/sec-financial-statements/utils"
	"gorm.io/gorm"
)

func Load(db *gorm.DB, f *zip.File) {
	tx := db.Begin()

	ProcessFile[Sub](tx, f, func(line string) *Sub {
		return ProcessLine[Sub](line)
	})

	tx.Commit()
}
