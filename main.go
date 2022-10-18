package main

import (
	"archive/zip"
	"flag"
	"github.com/phoobynet/sec-financial-statements/database"
	//. "github.com/phoobynet/sec-financial-statements/nums"
	//. "github.com/phoobynet/sec-financial-statements/pres"
	. "github.com/phoobynet/sec-financial-statements/submissions"
	//. "github.com/phoobynet/sec-financial-statements/tags"
	. "github.com/phoobynet/sec-financial-statements/utils"
	"log"
	"os"
)

func main() {
	sourceZip := flag.String("s", "/Volumes/raid/Downloads/browser/2022q3.zip", "source zip file")
	databasePath := flag.String("d", "", "Path to the database file")
	flag.Parse()

	if _, exists := os.Stat(*sourceZip); os.IsNotExist(exists) {
		flag.PrintDefaults()
		log.Fatalln("Source zip file does not exist")
	}

	if *databasePath == "" {
		flag.PrintDefaults()
		return
	}

	db := database.Init(*databasePath)

	finStatementZip, migrateErr := zip.OpenReader(*sourceZip)

	if migrateErr != nil {
		panic(migrateErr)
	}

	for _, file := range finStatementZip.File {
		switch file.Name {
		case "sub.txt":
			ProcessFile[Submission](db, file, func(line string) *Submission {
				return ProcessLine[Submission](line)
			})
			//case "num.txt":
			//	ProcessFile[Num](db, file, func(line string) *Num {
			//		return ProcessLine[Num](line)
			//	})
			//case "pre.txt":
			//	ProcessFile[Pre](db, file, func(line string) *Pre {
			//		return ProcessLine[Pre](line)
			//	})
			//case "tag.txt":
			//	ProcessFile[Tag](db, file, func(line string) *Tag {
			//		return ProcessLine[Tag](line)
			//	})
			//}
		}
	}
}
