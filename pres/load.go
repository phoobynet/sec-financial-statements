package pres

import (
	"archive/zip"
	. "github.com/phoobynet/sec-financial-statements/utils"
	"gorm.io/gorm"
)

func Load(db *gorm.DB, f *zip.File) {
	ProcessFile[Pre](db, f, func(line string) *Pre {
		return ProcessLine[Pre](line)
	})
}
