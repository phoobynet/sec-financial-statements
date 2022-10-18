package quarterly

import (
	"archive/zip"
	"github.com/phoobynet/sec-financial-statements/nums"
	"github.com/phoobynet/sec-financial-statements/pres"
	"github.com/phoobynet/sec-financial-statements/subs"
	"github.com/phoobynet/sec-financial-statements/tags"
	"gorm.io/gorm"
)

func Load(db *gorm.DB, sourceZip string) {
	finStatementZip, migrateErr := zip.OpenReader(sourceZip)

	if migrateErr != nil {
		panic(migrateErr)
	}

	for _, file := range finStatementZip.File {
		switch file.Name {
		case "sub.txt":
			subs.Load(db, file)
		case "num.txt":
			nums.Load(db, file)
		case "pre.txt":
			pres.Load(db, file)
		case "tag.txt":
			tags.Load(db, file)
		}
	}
}
