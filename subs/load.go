package subs

import (
	"archive/zip"
	. "github.com/phoobynet/sec-financial-statements/utils"
	"gorm.io/gorm"
)

func Load(db *gorm.DB, f *zip.File) {
	ProcessFile[Sub](db, f, func(line string) *Sub {
		return ProcessLine[Sub](line)
	})
}
