package tags

import (
	"archive/zip"
	. "github.com/phoobynet/sec-financial-statements/utils"
	"gorm.io/gorm"
)

func Load(db *gorm.DB, f *zip.File) {
	ProcessFile[Tag](db, f, func(line string) *Tag {
		return ProcessLine[Tag](line)
	})
}
