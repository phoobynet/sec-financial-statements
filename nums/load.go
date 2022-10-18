package nums

import (
	"archive/zip"
	. "github.com/phoobynet/sec-financial-statements/utils"
	"gorm.io/gorm"
)

func Load(db *gorm.DB, f *zip.File) {
	ProcessFile[Num](db, f, func(line string) *Num {
		return ProcessLine[Num](line)
	})
}
